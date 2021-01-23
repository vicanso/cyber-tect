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

package service

import (
	"net"
	"net/http"

	performance "github.com/vicanso/go-performance"
)

var httpServerConnStats = performance.NewHttpServerConnStats()
var requestProcessConcurrency = performance.NewConcurrency()

// GetHTTPServerConnState get http server on conn state function
func GetHTTPServerConnState() func(net.Conn, http.ConnState) {
	return httpServerConnStats.ConnState
}

type (
	// Performance 应用性能指标
	Performance struct {
		Concurrency           int32 `json:"concurrency,omitempty"`
		RequestProcessedTotal int64 `json:"requestProcessedTotal,omitempty"`
		performance.CPUMemory
		performance.ConnStats
	}
)

// GetPerformance 获取应用性能指标
func GetPerformance() Performance {
	return Performance{
		Concurrency:           GetConcurrency(),
		RequestProcessedTotal: requestProcessConcurrency.Total(),
		CPUMemory:             performance.CurrentCPUMemory(),
		ConnStats:             httpServerConnStats.Stats(),
	}

}

// IncreaseConcurrency 当前并发请求+1
func IncreaseConcurrency() int32 {
	return requestProcessConcurrency.Inc()
}

// DecreaseConcurrency 当前并发请求-1
func DecreaseConcurrency() int32 {
	return requestProcessConcurrency.Dec()
}

// GetConcurrency 获取当前并发请求
func GetConcurrency() int32 {
	return requestProcessConcurrency.Current()
}

// UpdateCPUUsage 更新cpu使用率
func UpdateCPUUsage() error {
	return performance.UpdateCPUUsage()
}
