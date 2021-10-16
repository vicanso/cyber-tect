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
	"crypto/tls"
	"database/sql"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v4"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/databasedetector"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/go-parallel"
	"github.com/vicanso/hes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseSrv struct{}

type databaseCheckParams struct {
	// 连接串
	ConnectionURI string
	// 超时
	Timeout time.Duration
	// tls配置
	TLSConfig *tls.Config
}

type DatabaseChecker func(ctx context.Context, params databaseCheckParams) (maskConnectionURI string, err error)

var databaseCheckers = map[string]DatabaseChecker{}

const (
	databaseCheckerRedis    = "redis"
	databaseCheckerPostgres = "postgres"
	databaseCheckerMysql    = "mysql"
	databaseCheckerMongodb  = "mongodb"
)

func init() {
	databaseCheckers[databaseCheckerRedis] = redisCheck
	databaseCheckers[databaseCheckerPostgres] = postgresCheck
	databaseCheckers[databaseCheckerMysql] = mysqlCheck
	databaseCheckers[databaseCheckerMongodb] = mongodbCheck
}

// redis检测
func redisCheck(ctx context.Context, params databaseCheckParams) (string, error) {
	connectionURI := params.ConnectionURI
	options, err := parseRedisConnectionURI(connectionURI)
	if err != nil {
		return "", err
	}

	maskURI := connectionURI
	if options.Password != "" {
		maskURI = strings.ReplaceAll(maskURI, options.Password, "***")
	}
	if options.SentinelPassword != "" {
		maskURI = strings.ReplaceAll(maskURI, options.SentinelPassword, "***")
	}
	options.DialTimeout = params.Timeout
	options.TLSConfig = params.TLSConfig
	c := redis.NewUniversalClient(options)
	defer c.Close()
	err = c.Ping(ctx).Err()

	return maskURI, err
}
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

// postgres检测
func postgresCheck(ctx context.Context, params databaseCheckParams) (string, error) {
	connectionURI := params.ConnectionURI
	maskURI := connectionURI
	info, err := url.Parse(maskURI)
	if err != nil {
		return "", err
	}
	password, _ := info.User.Password()
	if password != "" {
		maskURI = strings.ReplaceAll(maskURI, password, "***")
	}
	conf, err := pgx.ParseConfig(connectionURI)
	if err != nil {
		return maskURI, err
	}
	conf.TLSConfig = params.TLSConfig
	ctx, cancel := context.WithTimeout(ctx, params.Timeout)
	defer cancel()
	conn, err := pgx.ConnectConfig(ctx, conf)
	if err != nil {
		return maskURI, err
	}
	defer conn.Close(ctx)
	err = conn.Ping(ctx)
	return maskURI, err
}

// mysql检测
func mysqlCheck(ctx context.Context, params databaseCheckParams) (string, error) {
	connectionURI := params.ConnectionURI
	maskURI := connectionURI
	reg := regexp.MustCompile(`://\S+:(\S+?)@`)
	values := reg.FindStringSubmatch(maskURI)
	if len(values) == 2 {
		maskURI = strings.ReplaceAll(maskURI, values[1], "***")
	}
	connectionURI = strings.ReplaceAll(connectionURI, databaseCheckerMysql+"://", "")
	// mysql 支持tls
	if params.TLSConfig != nil {
		// 生成随机串
		name := util.RandomString(10)
		// 添加tls配置
		err := mysql.RegisterTLSConfig(name, params.TLSConfig)
		if err != nil {
			return maskURI, err
		}
		// 删除tls配置
		defer mysql.DeregisterTLSConfig(name)
		joinStr := "?"
		// 连接串增加tls参数
		if strings.Contains(connectionURI, "?") {
			joinStr = "&"
		}
		connectionURI += (joinStr + "tls=" + name)
	}

	db, err := sql.Open(databaseCheckerMysql, connectionURI)
	if err != nil {
		return maskURI, err
	}
	defer db.Close()
	err = db.PingContext(ctx)
	return maskURI, err
}

func mongodbCheck(ctx context.Context, params databaseCheckParams) (string, error) {
	connectionURI := params.ConnectionURI
	clientOpts := options.Client().ApplyURI(connectionURI)
	maskURI := connectionURI
	if clientOpts.Auth != nil && clientOpts.Auth.Password != "" {
		maskURI = strings.ReplaceAll(connectionURI, clientOpts.Auth.Password, "")
	}
	clientOpts.TLSConfig = params.TLSConfig
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		return maskURI, err
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, nil)
	return maskURI, err
}

func (srv *DatabaseSrv) check(ctx context.Context, params databaseCheckParams) (string, error) {
	connectionURI := params.ConnectionURI
	protocol := "://"
	if !strings.Contains(connectionURI, protocol) {
		return "", hes.New("数据库连接串有误")
	}
	arr := strings.Split(connectionURI, protocol)
	schema := arr[0]
	fn, ok := databaseCheckers[schema]
	if !ok {
		return "", hes.New("暂未支持(" + schema + ")检测")
	}
	return fn(ctx, params)
}

func (srv *DatabaseSrv) detect(ctx context.Context, config *ent.DatabaseDetector) (*ent.DatabaseDetectorResult, error) {
	timeout, _ := time.ParseDuration(config.Timeout)
	if timeout == 0 {
		timeout = defaultTimeout
	}
	result := schema.DetectorResultSuccess
	subResults := make(schema.DatabaseDetectorSubResults, 0)
	maxDuration := 0
	messages := make([]string, 0)
	uris := make([]string, len(config.Uris))
	var tlsConfig *tls.Config
	if config.CertPem != "" && config.KeyPem != "" {
		cert, err := tls.X509KeyPair([]byte(config.CertPem), []byte(config.KeyPem))
		if err != nil {
			return nil, err
		}
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
			Certificates: []tls.Certificate{
				cert,
			},
		}
	}
	for index, uri := range config.Uris {
		startedAt := time.Now()
		maskURI, err := srv.check(ctx, databaseCheckParams{
			ConnectionURI: uri,
			Timeout:       timeout,
			TLSConfig:     tlsConfig,
		})
		subResult := schema.DatabaseDetectorSubResult{
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
	return getEntClient().DatabaseDetectorResult.Create().
		SetTask(config.ID).
		SetResult(schema.DetectorResult(result)).
		SetResults(subResults).
		SetMaxDuration(maxDuration).
		SetMessages(messages).
		SetUris(uris).
		Save(ctx)
}

func (srv *DatabaseSrv) doAlarm(ctx context.Context, name string, receivers []string, result *ent.DatabaseDetectorResult) {
	if result == nil {
		return
	}
	doAlarm(ctx, alarmDetail{
		Name:      name,
		Receivers: receivers,
		Task:      fmt.Sprintf("database-%d", result.Task),
		IsSuccess: result.Result == schema.DetectorResultSuccess,
		Messages:  result.Messages,
	})
}

func (srv *DatabaseSrv) Detect(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	result, err := getEntClient().DatabaseDetector.Query().
		Where(databasedetector.StatusEQ(schema.StatusEnabled)).
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
		err = convertParallelError(pErr, "database detect fail")
	}
	if err != nil {
		return err
	}
	return nil
}
