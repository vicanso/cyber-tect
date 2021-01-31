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
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/vicanso/cybertect/config"
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/log"
	"github.com/vicanso/hes"
	"go.uber.org/atomic"
	"go.uber.org/zap"
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
		maxProcessing  uint32
		slow           time.Duration
		processing     atomic.Uint32
		pipeProcessing atomic.Uint32
		total          atomic.Uint64
	}
)

func mustNewRedisClient() (*redis.Client, *redisHook) {
	redisConfig := config.GetRedisConfig()
	log.Default().Info("connect to redis",
		zap.String("addr", redisConfig.Addr),
		zap.Int("db", redisConfig.DB),
	)
	hook := &redisHook{
		slow:          redisConfig.Slow,
		maxProcessing: redisConfig.MaxProcessing,
	}
	redis.SetLogger(log.NewRedisLogger())
	c := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
		Limiter:  hook,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			log.Default().Info("redis new connection is established")
			GetInfluxSrv().Write(cs.MeasurementRedisConn, nil, map[string]interface{}{
				cs.FieldCount: 1,
			})
			return nil
		},
	})
	c.AddHook(hook)
	return c, hook
}

// 对于慢或出错请求输出日志并写入influxdb
func (rh *redisHook) logSlowOrError(ctx context.Context, cmd, err string) {
	t := ctx.Value(startedAtKey).(*time.Time)
	d := time.Since(*t)
	if d > rh.slow || err != "" {
		log.Default().Info("redis process slow or error",
			zap.String("cmd", cmd),
			zap.String("use", d.String()),
			zap.String("error", err),
		)
		tags := map[string]string{
			cs.TagOP: cmd,
		}
		fields := map[string]interface{}{
			cs.FieldUse:   int(d.Milliseconds()),
			cs.FieldError: err,
		}
		GetInfluxSrv().Write(cs.MeasurementRedisStats, tags, fields)
	}
}

// BeforeProcess redis处理命令前的hook函数
func (rh *redisHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	t := time.Now()
	ctx = context.WithValue(ctx, startedAtKey, &t)
	rh.processing.Inc()
	rh.total.Inc()
	return ctx, nil
}

// AfterProcess redis处理命令后的hook函数
func (rh *redisHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	message := ""
	err := cmd.Err()
	if err != nil {
		message = err.Error()
	}
	rh.logSlowOrError(ctx, cmd.Name(), message)
	rh.processing.Dec()
	return nil
}

// BeforeProcessPipeline redis pipeline命令前的hook函数
func (rh *redisHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	t := time.Now()
	ctx = context.WithValue(ctx, startedAtKey, &t)
	rh.pipeProcessing.Inc()
	rh.total.Inc()
	return ctx, nil
}

// AfterProcessPipeline redis pipeline命令后的hook函数
func (rh *redisHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	cmdSb := new(strings.Builder)
	message := ""
	for index, cmd := range cmds {
		if index != 0 {
			cmdSb.WriteString(",")
		}
		cmdSb.WriteString(cmd.Name())
		err := cmd.Err()
		if err != nil {
			message += err.Error()
		}
	}
	rh.logSlowOrError(ctx, cmdSb.String(), message)
	rh.pipeProcessing.Dec()
	return nil
}

// getProcessingAndTotal 获取正在处理中的请求与总请求量
func (rh *redisHook) getProcessingAndTotal() (uint32, uint32, uint64) {
	processing := rh.processing.Load()
	pipeProcessing := rh.pipeProcessing.Load()
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
	if result != nil && !RedisIsNilError(result) {
		log.Default().Error("redis process fail",
			zap.Error(result),
		)
		GetInfluxSrv().Write(cs.MeasurementRedisError, nil, map[string]interface{}{
			cs.FieldError: result.Error(),
		})
	}
}

// RedisGetClient 获取redis client
func RedisGetClient() *redis.Client {
	return defaultRedisClient
}

// RedisIsNilError 判断是否redis的nil error
func RedisIsNilError(err error) bool {
	return err == redis.Nil
}

// RedisStats 获取redis的性能统计
func RedisStats() map[string]interface{} {
	stats := RedisGetClient().PoolStats()
	processing, pipeProcessing, total := defaultRedisHook.getProcessingAndTotal()
	return map[string]interface{}{
		cs.FieldHits:          int(stats.Hits),
		cs.FieldMisses:        int(stats.Misses),
		cs.FieldTimeouts:      int(stats.Timeouts),
		cs.FieldTotalConns:    int(stats.TotalConns),
		cs.FieldIdleConns:     int(stats.IdleConns),
		cs.FieldStaleConns:    int(stats.StaleConns),
		cs.FieldProcessing:    int(processing),
		cs.FilePipeProcessing: int(pipeProcessing),
		cs.FieldTotal:         int(total),
	}
}

// RedisPing ping操作
func RedisPing() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = RedisGetClient().Ping(ctx).Result()
	return
}
