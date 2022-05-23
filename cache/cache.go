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
	goCache "github.com/vicanso/go-cache/v2"
)

var redisCache = newRedisCache()

var redisSession = newRedisSession()
var redisConfig = config.MustGetRedisConfig()

// 数据压缩的最小长度
var compressMinLength = 2 * 1024

func newRedisCache() *goCache.RedisCache {
	opts := []goCache.RedisCacheOption{
		goCache.RedisCachePrefixOption(redisConfig.Prefix),
	}
	c := goCache.NewRedisCache(helper.RedisGetClient(), opts...)
	return c
}

func MustNewSnappyCompressCache(ttl time.Duration) *goCache.Cache {
	return mustNewCompressCache(ttl, goCache.NewSnappyCompressor(compressMinLength))
}

func MustNewZSTDCompressCache(ttl time.Duration) *goCache.Cache {
	return mustNewCompressCache(ttl, goCache.NewZSTDCompressor(compressMinLength, 2))
}

func mustNewCompressCache(ttl time.Duration, compressor goCache.Compressor) *goCache.Cache {
	c, err := goCache.New(
		ttl,
		goCache.CacheCompressorOption(compressor),
		goCache.CacheKeyPrefixOption(redisConfig.Prefix),
		goCache.CacheStoreOption(
			goCache.NewRedisStore(helper.RedisGetClient()),
		),
	)
	if err != nil {
		panic(err)
	}
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

// MustNewMultilevelCache create a new multilevel cache
func MustNewMultilevelCache(ttl time.Duration, cacheSizeMB int, prefix string) *goCache.Cache {

	c, err := goCache.New(
		ttl,
		goCache.CacheHardMaxCacheSizeOption(cacheSizeMB),
		goCache.CacheKeyPrefixOption(prefix),
		goCache.CacheSecondaryStoreOption(
			goCache.NewRedisStore(helper.RedisGetClient()),
		),
	)
	if err != nil {
		panic(err)
	}
	return c
}

// MustNewLRUCache new lru cache with ttl
func MustNewLRUCache(ttl time.Duration, cacheSizeMB int) *goCache.Cache {
	c, err := goCache.New(
		ttl,
		goCache.CacheHardMaxCacheSizeOption(cacheSizeMB),
	)
	if err != nil {
		panic(err)
	}
	return c
}
