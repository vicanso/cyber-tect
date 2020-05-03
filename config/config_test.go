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

package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfigGet(t *testing.T) {
	assert := assert.New(t)
	originEnv := env

	env = "test"
	assert.Equal(env, GetENV())
	env = originEnv

	key := "my-config-test"

	assert.Equal(0, GetInt(key))
	defaultViper.Set(key, 1)
	assert.Equal(1, GetInt(key))

	assert.Equal(uint(1), GetUint(key))
	assert.Equal(uint32(1), GetUint32(key))

	defaultViper.Set(key, nil)
	assert.Equal(2, GetIntDefault(key, 2))
	assert.Equal(uint32(2), GetUint32Default(key, 2))

	defaultViper.Set(key, "s")
	assert.Equal("s", GetString(key))
	defaultViper.Set(key, nil)
	assert.Equal("ss", GetStringDefault(key, "ss"))

	defaultViper.Set(key, time.Second)
	assert.Equal(time.Second, GetDuration(key))
	defaultViper.Set(key, nil)
	assert.Equal(time.Minute, GetDurationDefault(key, time.Minute))

	defaultViper.Set(key, []string{
		"a",
		"b",
	})
	assert.Equal([]string{
		"a",
		"b",
	}, GetStringSlice(key))

	defaultViper.Set(key, map[string]interface{}{
		"a": "b",
		"c": 1,
	})
	assert.Equal(map[string]interface{}{
		"a": "b",
		"c": 1,
	}, GetStringMap(key))
}
