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

package cache

import (
	"time"

	"github.com/vicanso/cybertect/config"
	"github.com/vicanso/cybertect/helper"
	goCache "github.com/vicanso/go-cache"
	lruttl "github.com/vicanso/lru-ttl"
)

var redisCache = newRedisCache()
var redisSession = newRedisSession()
var redisConfig = config.GetRedisConfig()

func newRedisCache() *goCache.RedisCache {
	c := goCache.NewRedisCache(helper.RedisGetClient())
	// 设置前缀
	c.SetPrefix(redisConfig.Prefix)
	return c
}

func newRedisSession() *goCache.RedisSession {
	ss := goCache.NewRedisSession(helper.RedisGetClient())
	// 设置前缀
	ss.SetPrefix(redisConfig.Prefix + "ss:")
	return ss
}

// GetRedisCache get redis cache
func GetRedisCache() *goCache.RedisCache {
	return redisCache
}

// GetRedisSession get redis session
func GetRedisSession() *goCache.RedisSession {
	return redisSession
}

// NewMultilevelCache create a new multilevel cache
func NewMultilevelCache(lruSize int, ttl time.Duration, prefix string) *lruttl.L2Cache {
	return goCache.NewMultilevelCache(goCache.MultilevelCacheOptions{
		Cache:   redisCache,
		LRUSize: lruSize,
		TTL:     ttl,
		Prefix:  prefix,
	})
}

// NewLRUCache new lru cache with ttl
func NewLRUCache(maxEntries int, defaultTTL time.Duration) *lruttl.Cache {
	return lruttl.New(maxEntries, defaultTTL)
}
