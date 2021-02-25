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
	"errors"
	"fmt"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxdbAPI "github.com/influxdata/influxdb-client-go/v2/api"
	influxdbDomain "github.com/influxdata/influxdb-client-go/v2/domain"
	"github.com/vicanso/cybertect/config"
	"github.com/vicanso/cybertect/log"
	"go.uber.org/zap"
)

type (
	InfluxSrv struct {
		client influxdb2.Client
		writer influxdbAPI.WriteAPI
		config config.InfluxdbConfig
	}
)

var hostname, _ = os.Hostname()
var defaultInfluxSrv = mustNewInfluxSrv()

// mustNewInfluxSrv 创建新的influx服务
func mustNewInfluxSrv() *InfluxSrv {
	influxdbConfig := config.GetInfluxdbConfig()
	if influxdbConfig.Disabled {

		return new(InfluxSrv)
	}
	opts := influxdb2.DefaultOptions()
	// 设置批量提交的大小
	opts.SetBatchSize(influxdbConfig.BatchSize)
	// 如果定时提交间隔大于1秒，则设定定时提交间隔
	if influxdbConfig.FlushInterval > time.Second {
		v := influxdbConfig.FlushInterval / time.Millisecond
		opts.SetFlushInterval(uint(v))
	}
	opts.SetUseGZip(influxdbConfig.Gzip)
	log.Default().Info("new influxdb client",
		zap.String("uri", influxdbConfig.URI),
		zap.String("org", influxdbConfig.Org),
		zap.String("bucket", influxdbConfig.Bucket),
		zap.Uint("batchSize", influxdbConfig.BatchSize),
		zap.String("token", influxdbConfig.Token[:5]+"..."),
		zap.Duration("interval", influxdbConfig.FlushInterval),
	)
	c := influxdb2.NewClientWithOptions(influxdbConfig.URI, influxdbConfig.Token, opts)
	writer := c.WriteAPI(influxdbConfig.Org, influxdbConfig.Bucket)
	go newInfluxdbErrorLogger(writer)

	return &InfluxSrv{
		client: c,
		writer: writer,
		config: influxdbConfig,
	}
}

// newInfluxdbErrorLogger 创建读取出错日志处理，需要注意此功能需要启用新的goroutine
func newInfluxdbErrorLogger(writer influxdbAPI.WriteAPI) {
	for err := range writer.Errors() {
		log.Default().Error("influxdb write fail",
			zap.Error(err),
		)
	}
}

// GetInfluxSrv 获取默认的influxdb服务
func GetInfluxSrv() *InfluxSrv {
	return defaultInfluxSrv
}

// Health check influxdb health
func (srv *InfluxSrv) Health() (err error) {
	if srv.client == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := srv.client.Health(ctx)
	if err != nil {
		return
	}
	if result.Status != influxdbDomain.HealthCheckStatusPass {
		err = errors.New(string(result.Status))
		return
	}
	return
}

func (srv *InfluxSrv) query(ctx context.Context, query string) (items []map[string]interface{}, err error) {
	if srv.client == nil {
		return
	}
	result, err := srv.client.QueryAPI(srv.config.Org).Query(ctx, query)
	if err != nil {
		return
	}
	items = make([]map[string]interface{}, 0)
	for result.Next() {
		items = append(items, result.Record().Values())
	}
	err = result.Err()
	if err != nil {
		return
	}

	return
}

// Query query records
func (srv *InfluxSrv) Query(ctx context.Context, query string) (items []map[string]interface{}, err error) {
	query = fmt.Sprintf(`from(bucket: "%s")`, srv.config.Bucket) + query
	return srv.query(ctx, query)
}

// ListTagValue list value of tag
func (srv *InfluxSrv) ListTagValue(ctx context.Context, measurement, tag string) (values []string, err error) {
	query := fmt.Sprintf(`import "influxdata/influxdb/schema"

	schema.measurementTagValues(
	  bucket: "%s",
	  measurement: "%s",
	  tag: "%s"
	)`, srv.config.Bucket, measurement, tag)
	items, err := srv.query(ctx, query)
	if err != nil {
		return
	}
	for _, item := range items {
		v, ok := item["_value"]
		if !ok {
			continue
		}
		value, ok := v.(string)
		if !ok {
			continue
		}
		values = append(values, value)
	}
	return
}

// Write 写入数据
func (srv *InfluxSrv) Write(measurement string, tags map[string]string, fields map[string]interface{}, ts ...time.Time) {
	if srv.writer == nil {
		return
	}
	var now time.Time
	if len(ts) != 0 {
		now = ts[0]
	} else {
		now = time.Now()
	}
	if fields == nil {
		fields = make(map[string]interface{})
	}
	if hostname != "" && fields["hostname"] == nil {
		fields["hostname"] = hostname
	}
	srv.writer.WritePoint(influxdb2.NewPoint(measurement, tags, fields, now))
}

// Close 关闭当前client
func (srv *InfluxSrv) Close() {
	if srv.client == nil {
		return
	}
	srv.client.Close()
}
