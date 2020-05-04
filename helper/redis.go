// Copyright 2019 tree xie
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
	"encoding/json"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/vicanso/cybertect/config"
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/hes"
	"go.uber.org/zap"
)

var (
	redisClient *redis.Client
	redisNoop   = func() error {
		return nil
	}
	errRedisNil               = hes.New("key is not exists or expired")
	redisSrv                  = new(Redis)
	rh                        *redisHook
	ErrRedisTooManyProcessing = &hes.Error{
		Message:    "too many processing",
		StatusCode: http.StatusInternalServerError,
		Category:   "redis",
	}
)

type (

	// redisHook redis hook
	redisHook struct {
		maxProcessing  uint32
		slow           time.Duration
		processing     uint32
		pipeProcessing uint32
		total          uint64
	}
)

type (
	// RedisDone redis done function
	RedisDone func() error
	// Redis redis service
	Redis struct{}

	// RedisSessionStore redis session store
	RedisSessionStore struct {
		Prefix string
	}
)

// 对于慢或出错请求输出日志并写入influxdb
func (rh *redisHook) logSlowOrError(ctx context.Context, cmd, err string) {
	t := ctx.Value(startedAtKey).(*time.Time)
	d := time.Since(*t)
	if d > rh.slow || err != "" {
		logger.Info("redis process slow or error",
			zap.String("cmd", cmd),
			zap.String("use", d.String()),
			zap.String("error", err),
		)
		tags := map[string]string{
			"cmd": cmd,
		}
		fields := map[string]interface{}{
			"use":   d.Milliseconds(),
			"error": err,
		}
		GetInfluxSrv().Write(cs.MeasurementRedis, fields, tags)
	}
}

// BeforeProcess before process
func (rh *redisHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	t := time.Now()
	ctx = context.WithValue(ctx, startedAtKey, &t)
	atomic.AddUint32(&rh.processing, 1)
	atomic.AddUint64(&rh.total, 1)
	return ctx, nil
}

// AfterProcess after process
func (rh *redisHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	message := ""
	err := cmd.Err()
	if err != nil {
		message = err.Error()
	}
	rh.logSlowOrError(ctx, cmd.Name(), message)
	atomic.AddUint32(&rh.processing, ^uint32(0))
	return nil
}

// BeforeProcessPipeline before process pipeline
func (rh *redisHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	t := time.Now()
	ctx = context.WithValue(ctx, startedAtKey, &t)
	atomic.AddUint32(&rh.pipeProcessing, 1)
	atomic.AddUint64(&rh.total, 1)
	return ctx, nil
}

// AfterProcessPipeline after process pipeline
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
	atomic.AddUint32(&rh.pipeProcessing, ^uint32(0))
	return nil
}

// getProcessingAndTotal get processing and total
func (rh *redisHook) getProcessingAndTotal() (uint32, uint32, uint64) {
	processing := atomic.LoadUint32(&rh.processing)
	pipeProcessing := atomic.LoadUint32(&rh.pipeProcessing)
	total := atomic.LoadUint64(&rh.total)
	return processing, pipeProcessing, total
}

func (rh *redisHook) Allow() error {
	// 如果处理请求量超出，则不允许继续请求
	if atomic.LoadUint32(&rh.processing) > rh.maxProcessing {
		return ErrRedisTooManyProcessing
	}
	return nil
}

func (*redisHook) ReportResult(result error) {
	if result != nil {
		logger.Error("redis process fail",
			zap.Error(result),
		)
	}
}

func init() {
	options, err := config.GetRedisConfig()
	if err != nil {
		panic(err)
	}
	logger.Info("connect to redis",
		zap.String("addr", options.Addr),
		zap.Int("db", options.DB),
	)
	rh = &redisHook{
		slow:          options.Slow,
		maxProcessing: options.MaxProcessing,
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     options.Addr,
		Password: options.Password,
		DB:       options.DB,
		Limiter:  rh,
	})
	redisClient.AddHook(rh)
}

// RedisGetClient get redis client
func RedisGetClient() *redis.Client {
	return redisClient
}

// IsRedisNilError is redis nil errror
func IsRedisNilError(err error) bool {
	return err == errRedisNil || err == redis.Nil
}

// RedisStats get redis stats
func RedisStats() map[string]interface{} {
	stats := redisClient.PoolStats()
	processing, pipeProcessing, total := rh.getProcessingAndTotal()
	return map[string]interface{}{
		"hits":           stats.Hits,
		"missed":         stats.Misses,
		"timeouts":       stats.Timeouts,
		"totalConns":     stats.TotalConns,
		"idleConns":      stats.IdleConns,
		"staleConns":     stats.StaleConns,
		"processing":     processing,
		"pipeProcessing": pipeProcessing,
		"total":          total,
	}
}

// RedisPing redis ping
func RedisPing() (err error) {
	_, err = redisClient.Ping().Result()
	return
}

// Lock lock the key for ttl seconds
func (srv *Redis) Lock(key string, ttl time.Duration) (bool, error) {
	return redisClient.SetNX(key, true, ttl).Result()
}

// Del del the key of redis
func (srv *Redis) Del(key string) (err error) {
	_, err = redisClient.Del(key).Result()
	return
}

// LockWithDone lock the key for ttl, and with done function
func (srv *Redis) LockWithDone(key string, ttl time.Duration) (bool, RedisDone, error) {
	success, err := srv.Lock(key, ttl)
	// 如果lock失败，则返回no op 的done function
	if err != nil || !success {
		return false, redisNoop, err
	}
	done := func() error {
		err := srv.Del(key)
		return err
	}
	return true, done, nil
}

// IncWithTTL inc value with ttl
func (srv *Redis) IncWithTTL(key string, ttl time.Duration) (count int64, err error) {
	pipe := redisClient.TxPipeline()
	// 保证只有首次会设置ttl
	pipe.SetNX(key, 0, ttl)
	incr := pipe.Incr(key)
	_, err = pipe.Exec()
	if err != nil {
		return
	}
	count = incr.Val()
	return
}

// Get get value
func (srv *Redis) Get(key string) (result string, err error) {
	result, err = redisClient.Get(key).Result()
	if err == redis.Nil {
		err = errRedisNil
	}
	return
}

// GetIgnoreNilErr get value ignore nil error
func (srv *Redis) GetIgnoreNilErr(key string) (result string, err error) {
	result, err = srv.Get(key)
	if IsRedisNilError(err) {
		err = nil
	}
	return
}

// GetAndDel get value and del
func (srv *Redis) GetAndDel(key string) (result string, err error) {
	pipe := redisClient.TxPipeline()
	cmd := pipe.Get(key)
	pipe.Del(key)
	_, err = pipe.Exec()
	if err != nil {
		if err == redis.Nil {
			err = errRedisNil
		}
		return
	}
	result = cmd.Val()
	return
}

// Set redis set with ttl
func (srv *Redis) Set(key string, value interface{}, ttl time.Duration) (err error) {
	redisClient.Set(key, value, ttl)
	return
}

// GetStruct get struct
func (srv *Redis) GetStruct(key string, value interface{}) (err error) {
	result, err := srv.Get(key)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(result), value)
	return
}

// SetStruct redis set struct with ttl
func (srv *Redis) SetStruct(key string, value interface{}, ttl time.Duration) (err error) {
	buf, err := json.Marshal(value)
	if err != nil {
		return
	}
	return srv.Set(key, string(buf), ttl)
}

func (rs *RedisSessionStore) getKey(key string) string {
	return rs.Prefix + key
}

// Get get the session from redis
func (rs *RedisSessionStore) Get(key string) ([]byte, error) {
	result, err := redisSrv.Get(rs.getKey(key))
	if IsRedisNilError(err) {
		return nil, nil
	}
	return []byte(result), err
}

// Set set the session to redis
func (rs *RedisSessionStore) Set(key string, data []byte, ttl time.Duration) error {
	return redisSrv.Set(rs.getKey(key), data, ttl)
}

// Destroy remove the session from redis
func (rs *RedisSessionStore) Destroy(key string) error {
	return redisSrv.Del(rs.getKey(key))
}
