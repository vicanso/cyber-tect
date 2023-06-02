// Copyright 2020 tree xie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helper

import (
	"context"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vicanso/cybertect/config"
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/log"
	"github.com/vicanso/go-gauge"
	"github.com/vicanso/hes"
	"go.uber.org/atomic"
)

var (
	defaultRedisClient, defaultRedisHook = mustNewRedisClient()

	// ErrRedisTooManyProcessing 处理请求太多时的出错
	ErrRedisTooManyProcessing = &hes.Error{
		Message:    "too many processing",
		StatusCode: http.StatusInternalServerError,
		Category:   "redis",
	}
)

type (

	// redisHook redis的hook配置
	redisHook struct {
		// 连接池大小
		poolSize int
		// 最大正在处理数量
		maxProcessing uint32
		// 慢请求阀值
		slow time.Duration
		// 正在处理数
		processing atomic.Uint32
		// pipe的正在处理数
		pipeProcessing atomic.Uint32
		// 总的处理请求数
		total     atomic.Uint64
		gauge     *gauge.Gauge
		pipeGauge *gauge.Gauge
	}
)

func init() {
	redis.SetLogger(log.NewRedisLogger())
}
func mustNewRedisClient() (redis.UniversalClient, *redisHook) {
	redisConfig := config.MustGetRedisConfig()
	log.Info(context.Background()).
		Strs("addr", redisConfig.Addrs).
		Msg("connect to redis")
	slow := redisConfig.Slow
	if slow < time.Millisecond {
		slow = 100 * time.Millisecond
	}
	hook := &redisHook{
		slow:          redisConfig.Slow,
		maxProcessing: redisConfig.MaxProcessing,
		// 记录每分钟最大并发数
		gauge: gauge.New(gauge.PeriodOption(time.Minute)),
		// 记录每分钟pipe最大并发数
		pipeGauge: gauge.New(gauge.PeriodOption(time.Minute)),
	}
	opts := &redis.UniversalOptions{
		Addrs:            redisConfig.Addrs,
		Username:         redisConfig.Username,
		Password:         redisConfig.Password,
		SentinelPassword: redisConfig.Password,
		MasterName:       redisConfig.Master,
		PoolSize:         redisConfig.PoolSize,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			log.Info(ctx).Msg("redis new connection is established")
			GetInfluxDB().Write(cs.MeasurementRedisConn, nil, map[string]any{
				cs.FieldCount: 1,
				cs.FieldAddr:  cn.String(),
			})
			return nil
		},
		MinIdleConns: 2,
	}
	var c redis.UniversalClient
	// 需要对增加limiter，因此单独判断处理
	if opts.MasterName != "" {
		// TODO 确认有无可能增加limiter
		failoverOpts := opts.Failover()
		c = redis.NewFailoverClient(failoverOpts)
		hook.poolSize = failoverOpts.PoolSize
	} else if len(opts.Addrs) > 1 {
		clusterOpts := opts.Cluster()
		clusterOpts.NewClient = func(opt *redis.Options) *redis.Client {
			// 对每个client的增加limiter
			opt.Limiter = hook
			return redis.NewClient(opt)
		}
		c = redis.NewClusterClient(clusterOpts)
		hook.poolSize = clusterOpts.PoolSize
	} else {
		simpleOpts := opts.Simple()
		simpleOpts.Limiter = hook
		c = redis.NewClient(simpleOpts)
		hook.poolSize = simpleOpts.PoolSize
	}
	// https://github.com/redis/go-redis/issues/2453
	// race condition
	// c.AddHook(hook)
	return c, hook
}

// DialHook redis连接时的hook
func (rh *redisHook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}

// ProcessHook redis 命令执行时的hook
func (rh *redisHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		t := time.Now()
		v := rh.processing.Inc()
		rh.gauge.SetMax(int64(v))
		rh.total.Inc()
		err := next(ctx, cmd)
		d := time.Since(t)
		message := ""
		if err != nil {
			message = err.Error()
		}
		rh.addStats(ctx, cmd.FullName(), message, d)
		rh.processing.Dec()
		if log.DebugEnabled() {
			// 由于redis是较频繁的操作
			// 由于cmd string的执行也有耗时，因此判断是否启用debug再输出
			log.Debug(ctx).
				Str("category", "redisProcess").
				Str("message", message).
				Msg(cmd.String())
		}
		return err
	}
}

// ProcessPipelineHook redis pipeline 执行时的hook
func (rh *redisHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		t := time.Now()
		v := rh.pipeProcessing.Inc()
		rh.pipeGauge.SetMax(int64(v))
		rh.total.Inc()
		fullNameList := new(strings.Builder)
		for index, cmd := range cmds {
			if index != 0 {
				fullNameList.WriteString(",")
			}
			fullNameList.WriteString(cmd.FullName())
		}

		err := next(ctx, cmds)
		d := time.Since(t)
		message := ""

		if err != nil {
			message = err.Error()
		}
		rh.addStats(ctx, fullNameList.String(), message, d)
		rh.pipeProcessing.Dec()
		if log.DebugEnabled() {
			// 由于redis是较频繁的操作
			// 由于cmd string的执行也有耗时，因此判断是否启用debug再输出
			cmdList := new(strings.Builder)
			for index, cmd := range cmds {
				if index != 0 {
					cmdList.WriteString(",")
				}
				cmdList.WriteString(cmd.String())
			}
			log.Debug(ctx).
				Str("category", "redisProcess").
				Str("message", message).
				Msg(cmdList.String())
		}
		return err
	}
}

// 添加统计至influxdb
func (rh *redisHook) addStats(ctx context.Context, cmd, err string, d time.Duration) {
	if d > rh.slow || err != "" {
		log.Info(ctx).
			Str("category", "redisSlowOrErr").
			Str("cmd", cmd).
			Str("use", d.String()).
			Str("error", err).
			Msg("")
	}

	tags := map[string]string{
		cs.TagOP: cmd,
	}
	fields := map[string]any{
		cs.FieldLatency: int(d.Milliseconds()),
	}
	if len(err) != 0 {
		fields[cs.FieldError] = err
	}
	GetInfluxDB().Write(cs.MeasurementRedisOP, tags, fields)
}

// getProcessingAndTotal 获取正在处理中的请求与总请求量
func (rh *redisHook) getProcessingAndTotal() (uint32, uint32, uint64) {
	processing := uint32(rh.gauge.Count())
	pipeProcessing := uint32(rh.pipeGauge.Count())
	total := rh.total.Load()
	return processing, pipeProcessing, total
}

// Allow 是否允许继续执行redis
func (rh *redisHook) Allow() error {
	// 如果处理请求量超出，则不允许继续请求
	if rh.processing.Load()+rh.pipeProcessing.Load() > rh.maxProcessing {
		return ErrRedisTooManyProcessing
	}
	return nil
}

// ReportResult 记录结果
func (*redisHook) ReportResult(result error) {
	// 仅是调用redis完成时触发
	// 需要注意，只有allow通过后才会触发
	// 仅仅nil error则忽略
	if result != nil && !RedisIsNilError(result) {
		log.Error(context.Background()).
			Str("category", "redisProcessFail").
			Err(result).
			Msg("")
		GetInfluxDB().Write(cs.MeasurementRedisError, nil, map[string]any{
			cs.FieldError: result.Error(),
		})
	}
}

// RedisGetClient 获取redis client
func RedisGetClient() redis.UniversalClient {
	return defaultRedisClient
}

// RedisIsNilError 判断是否redis的nil error
func RedisIsNilError(err error) bool {
	return err == redis.Nil
}

// RedisStats 获取redis的性能统计
func RedisStats() map[string]any {
	stats := RedisGetClient().PoolStats()
	processing, pipeProcessing, total := defaultRedisHook.getProcessingAndTotal()
	return map[string]any{
		cs.FieldHits:          int(stats.Hits),
		cs.FieldMisses:        int(stats.Misses),
		cs.FieldTimeouts:      int(stats.Timeouts),
		cs.FieldTotalConns:    int(stats.TotalConns),
		cs.FieldIdleConns:     int(stats.IdleConns),
		cs.FieldStaleConns:    int(stats.StaleConns),
		cs.FieldProcessing:    int(processing),
		cs.FilePipeProcessing: int(pipeProcessing),
		cs.FieldTotal:         int(total),
		cs.FieldPoolSize:      defaultRedisHook.poolSize,
	}
}

// RedisPing ping操作
func RedisPing() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := RedisGetClient().Ping(ctx).Result()
	return err
}
