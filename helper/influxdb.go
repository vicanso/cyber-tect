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
	"time"

	influxdb "github.com/influxdata/influxdb-client-go"
	"github.com/vicanso/cybertect/config"
)

var (
	defaultInfluxSrv *InfluxSrv
)

type (
	InfluxSrv struct {
		client influxdb.InfluxDBClient
		writer influxdb.WriteApi
	}
)

func init() {
	influxbConfig := config.GetInfluxdbConfig()
	if influxbConfig.Disabled {
		defaultInfluxSrv = new(InfluxSrv)
		return
	}
	opts := influxdb.DefaultOptions()
	opts.SetBatchSize(influxbConfig.BatchSize)
	if influxbConfig.FlushInterval > time.Millisecond {
		v := influxbConfig.FlushInterval / time.Millisecond
		opts.SetFlushInterval(uint(v))
	}
	c := influxdb.NewClientWithOptions(influxbConfig.URI, influxbConfig.Token, opts)
	writer := c.WriteApi(influxbConfig.Org, influxbConfig.Bucket)
	defaultInfluxSrv = &InfluxSrv{
		client: c,
		writer: writer,
	}
}

// GetInfluxSrv get default influx service
func GetInfluxSrv() *InfluxSrv {
	return defaultInfluxSrv
}

// Write write metric to influxdb
func (srv *InfluxSrv) Write(measurement string, fields map[string]interface{}, tags map[string]string) {
	if srv.writer == nil {
		return
	}
	srv.writer.WritePoint(influxdb.NewPoint(measurement, tags, fields, time.Now()))
}

// Flush flush metric list
func (srv *InfluxSrv) Flush() {
	if srv.writer == nil {
		return
	}
	srv.writer.Flush()
}

// Close flush the point to influxdb and close client
func (srv *InfluxSrv) Close() {
	if srv.client == nil {
		return
	}
	srv.client.Close()
}
