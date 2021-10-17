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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRedisConnectionURI(t *testing.T) {
	assert := assert.New(t)

	options, err := parseRedisConnectionURI("redis://user:pass@127.0.0.1:6379/")
	assert.Nil(err)
	assert.Equal([]string{
		"127.0.0.1:6379",
	}, options.Addrs)
	assert.Equal("user", options.Username)
	assert.Equal("pass", options.Password)

	options, err = parseRedisConnectionURI("redis://user:pass@127.0.0.1:6379,127.0.0.1:6378/")
	assert.Nil(err)
	assert.Equal([]string{
		"127.0.0.1:6379",
		"127.0.0.1:6378",
	}, options.Addrs)
	assert.Equal("user", options.Username)
	assert.Equal("pass", options.Password)

	options, err = parseRedisConnectionURI("redis://user:pass@127.0.0.1:6379,127.0.0.1:6378/?master=abc&sentinelPassword=sentinelPass")
	assert.Nil(err)
	assert.Equal([]string{
		"127.0.0.1:6379",
		"127.0.0.1:6378",
	}, options.Addrs)
	assert.Equal("user", options.Username)
	assert.Equal("pass", options.Password)
	assert.Equal("abc", options.MasterName)
	assert.Equal("sentinelPass", options.SentinelPassword)

	options, err = parseRedisConnectionURI("redis://user:pass@127.0.0.1:6379,127.0.0.1:6378/?master=abc&sentinelPassword=123")
	assert.Nil(err)
	assert.Equal([]string{
		"127.0.0.1:6379",
		"127.0.0.1:6378",
	}, options.Addrs)
	assert.Equal("user", options.Username)
	assert.Equal("pass", options.Password)
	assert.Equal("abc", options.MasterName)
	assert.Equal("123", options.SentinelPassword)
}
