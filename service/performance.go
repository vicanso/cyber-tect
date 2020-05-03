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

package service

import (
	"runtime"
	"sync/atomic"
)

type (
	// Performance performance
	Performance struct {
		GoMaxProcs   int    `json:"goMaxProcs,omitempty"`
		Concurrency  uint32 `json:"concurrency,omitempty"`
		Sys          int    `json:"sys,omitempty"`
		HeapSys      int    `json:"heapSys,omitempty"`
		HeapInuse    int    `json:"heapInuse,omitempty"`
		RoutineCount int    `json:"routineCount,omitempty"`
	}
)

var (
	concurrency uint32
)

// GetPerformance get performance
func GetPerformance() *Performance {
	var mb uint64 = 1024 * 1024
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)
	return &Performance{
		GoMaxProcs:   runtime.GOMAXPROCS(0),
		Concurrency:  GetConcurrency(),
		Sys:          int(m.Sys / mb),
		HeapSys:      int(m.HeapSys / mb),
		HeapInuse:    int(m.HeapInuse / mb),
		RoutineCount: runtime.NumGoroutine(),
	}
}

// IncreaseConcurrency increase concurrency count
func IncreaseConcurrency() uint32 {
	return atomic.AddUint32(&concurrency, 1)
}

// DecreaseConcurrency decrease concurrency count
func DecreaseConcurrency() uint32 {
	return atomic.AddUint32(&concurrency, ^uint32(0))
}

// GetConcurrency get concurrency
func GetConcurrency() uint32 {
	return atomic.LoadUint32(&concurrency)
}
