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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vicanso/elton"
)

func TestRouterLimiter(t *testing.T) {
	assert := assert.New(t)

	InitRouterConcurrencyLimiter([]elton.RouterInfo{
		{
			Method: "GET",
			Route:  "/",
		},
		{
			Method: "GET",
			Route:  "/users/me",
		},
	})
	rc := GetRouterConcurrencyLimiter()
	key := "GET /"
	count, max := rc.IncConcurrency(key)
	assert.Equal(uint32(1), count)
	assert.Equal(uint32(0), max)

	assert.Equal(uint32(1), rc.GetConcurrency(key))

	rc.DecConcurrency(key)
	assert.Equal(uint32(0), rc.GetConcurrency(key))

	// 重置路由并发配置
	ResetRouterConcurrency([]string{
		`{
			"route": "/",
			"method": "GET",
			"max": 10,
			"rateLimit": "100/s"
		}`,
	})
	count, max = rc.IncConcurrency(key)
	assert.Equal(uint32(1), count)
	assert.Equal(uint32(10), max)
}

func TestRouterConfig(t *testing.T) {
	assert := assert.New(t)
	updateRouterMockConfigs([]string{
		`{
			"route": "/",
			"method": "GET",
			"status": 400
		}`,
	})
	routeConfig := RouterGetConfig("GET", "/")
	assert.Equal(400, routeConfig.Status)
}
