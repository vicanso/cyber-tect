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
	"testing"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
	"github.com/vicanso/cybertect/util"
)

func TestIsRedisNilError(t *testing.T) {
	assert := assert.New(t)
	assert.True(IsRedisNilError(errRedisNil))
	assert.True(IsRedisNilError(redis.Nil))
}

func TestRedisSrv(t *testing.T) {
	assert := assert.New(t)
	srv := new(Redis)
	t.Run("lock", func(t *testing.T) {
		key := util.RandomString(8)
		ttl := 2 * time.Millisecond
		ok, err := srv.Lock(key, ttl)
		assert.Nil(err)
		assert.True(ok)

		ok, err = srv.Lock(key, ttl)
		assert.Nil(err)
		assert.False(ok)

		time.Sleep(2 * ttl)
		ok, err = srv.Lock(key, ttl)
		assert.Nil(err)
		assert.True(ok)
	})

	t.Run("lock with done", func(t *testing.T) {
		key := util.RandomString(8)
		ttl := 2 * time.Millisecond
		ok, done, err := srv.LockWithDone(key, ttl)
		assert.Nil(err)
		assert.True(ok)

		ok, _, err = srv.LockWithDone(key, ttl)
		assert.Nil(err)
		assert.False(ok)

		err = done()
		assert.Nil(err)
		ok, _, err = srv.LockWithDone(key, ttl)
		assert.Nil(err)
		assert.True(ok)
	})

	t.Run("inc with ttl", func(t *testing.T) {
		key := util.RandomString(8)
		ttl := 2 * time.Millisecond
		count, err := srv.IncWithTTL(key, ttl)
		assert.Nil(err)
		assert.Equal(int64(1), count)

		count, err = srv.IncWithTTL(key, ttl)
		assert.Nil(err)
		assert.Equal(int64(2), count)

		time.Sleep(2 * ttl)
		count, err = srv.IncWithTTL(key, ttl)
		assert.Nil(err)
		assert.Equal(int64(1), count)
	})

	t.Run("get/set", func(t *testing.T) {
		key := util.RandomString(8)
		_, err := srv.Get(key)
		assert.Equal(errRedisNil, err)

		_, err = srv.GetIgnoreNilErr(key)
		assert.Nil(err)

		err = srv.Set(key, "1", time.Second)
		assert.Nil(err)
		value, err := srv.Get(key)
		assert.Nil(err)
		assert.Equal("1", value)

		value, err = srv.GetAndDel(key)
		assert.Nil(err)
		assert.Equal("1", value)

		_, err = srv.Get(key)
		assert.Equal(errRedisNil, err)
	})

	t.Run("get/set struct", func(t *testing.T) {
		type tmpStruct struct {
			Name string
		}
		key := util.RandomString(8)
		s1 := tmpStruct{
			Name: "name",
		}
		err := srv.SetStruct(key, &s1, time.Second)
		assert.Nil(err)

		s2 := tmpStruct{}
		err = srv.GetStruct(key, &s2)
		assert.Nil(err)
		assert.Equal(s1.Name, s2.Name)
	})
}

func TestRedisSessionStore(t *testing.T) {
	assert := assert.New(t)
	rs := &RedisSessionStore{
		Prefix: "sess:",
	}
	data := []byte("abcd")
	key := util.RandomString(8)
	err := rs.Set(key, data, time.Second)
	assert.Nil(err)

	result, err := rs.Get(key)
	assert.Nil(err)
	assert.Equal(data, result)

	err = rs.Destroy(key)
	assert.Nil(err)

	result, err = rs.Get(key)
	assert.Nil(err)
	assert.Empty(result)
}
