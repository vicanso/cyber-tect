// Copyright 2021 tree xie
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

package detector

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/redisdetector"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/go-parallel"
)

type RedisSrv struct{}

func parseRedisConnectionURI(connectionURI string) (*redis.UniversalOptions, error) {
	info, err := url.Parse(connectionURI)
	if err != nil {
		return nil, err
	}
	addrs := strings.Split(info.Host, ",")
	username := info.User.Username()
	password, _ := info.User.Password()
	master := info.Query().Get("master")
	sentinelPassword := info.Query().Get("sentinelPassword")
	if sentinelPassword == "" {
		sentinelPassword = password
	}
	return &redis.UniversalOptions{
		Addrs:            addrs,
		Username:         username,
		Password:         password,
		SentinelPassword: sentinelPassword,
		MasterName:       master,
	}, nil
}

func (srv *RedisSrv) check(ctx context.Context, connectionURI string, timeout time.Duration) (string, error) {
	options, err := parseRedisConnectionURI(connectionURI)
	if err != nil {
		return "", err
	}
	options.DialTimeout = timeout
	c := redis.NewUniversalClient(options)
	defer c.Close()
	err = c.Ping(ctx).Err()
	maskURI := connectionURI
	if options.Password != "" {
		maskURI = strings.ReplaceAll(maskURI, options.Password, "***")
	}
	if options.SentinelPassword != "" {
		maskURI = strings.ReplaceAll(maskURI, options.SentinelPassword, "***")
	}
	return maskURI, err
}

func (srv *RedisSrv) detect(ctx context.Context, config *ent.RedisDetector) (*ent.RedisDetectorResult, error) {
	timeout, _ := time.ParseDuration(config.Timeout)
	if timeout == 0 {
		timeout = defaultTimeout
	}
	result := schema.DetectorResultSuccess
	subResults := make(schema.RedisDetectorSubResults, 0)
	maxDuration := 0
	messages := make([]string, 0)
	uris := make([]string, len(config.Uris))
	for index, uri := range config.Uris {
		startedAt := time.Now()
		maskURI, err := srv.check(ctx, uri, timeout)
		subResult := schema.RedisDetectorSubResult{
			URI:      maskURI,
			Duration: ceilToMs(time.Since(startedAt)),
		}
		uris[index] = maskURI
		if err != nil {
			subResult.Result = schema.DetectorResultFail
			subResult.Message = err.Error()
			result = schema.DetectorResultFail
			messages = append(messages, subResult.Message)
		} else {
			subResult.Result = schema.DetectorResultSuccess
		}
		if subResult.Duration > maxDuration {
			maxDuration = subResult.Duration
		}
		subResults = append(subResults, &subResult)
	}
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	return getEntClient().RedisDetectorResult.Create().
		SetTask(config.ID).
		SetResult(schema.DetectorResult(result)).
		SetResults(subResults).
		SetMaxDuration(maxDuration).
		SetMessages(messages).
		SetUris(uris).
		Save(ctx)
}

func (srv *RedisSrv) doAlarm(ctx context.Context, name string, receivers []string, result *ent.RedisDetectorResult) {
	if result == nil {
		return
	}
	doAlarm(ctx, alarmDetail{
		Name:      name,
		Receivers: receivers,
		Task:      fmt.Sprintf("redis-%d", result.Task),
		IsSuccess: result.Result == schema.DetectorResultSuccess,
		Messages:  result.Messages,
	})
}

func (srv *RedisSrv) Detect(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	result, err := getEntClient().RedisDetector.Query().
		Where(redisdetector.StatusEQ(schema.StatusEnabled)).
		All(ctx)
	if err != nil {
		return err
	}
	pErr := parallel.Parallel(func(index int) error {
		item := result[index]
		detectResult, err := srv.detect(ctx, item)
		srv.doAlarm(ctx, item.Name, item.Receivers, detectResult)
		return err
	}, len(result), detectorConfig.Concurrency)
	// 如果parallel检测失败，则转换
	if pErr != nil {
		err = convertParallelError(pErr, "redis detect fail")
	}
	if err != nil {
		return err
	}
	return nil
}
