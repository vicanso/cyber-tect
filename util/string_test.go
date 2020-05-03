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

package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(8, len(RandomString(8)))
	assert.Equal(4, len(RandomDigit(4)))
	assert.Equal(26, len(GenUlid()))
}

func TestContainsString(t *testing.T) {
	assert := assert.New(t)
	assert.True(ContainsString([]string{
		"a",
		"b",
	}, "a"))
	assert.False(ContainsString([]string{
		"a",
		"b",
	}, "c"))
}

func TestUserRoleIsValid(t *testing.T) {
	assert := assert.New(t)
	assert.True(UserRoleIsValid([]string{
		"admin",
	}, []string{
		"admin",
		"su",
	}))
	assert.False(UserRoleIsValid([]string{
		"staff",
	}, []string{
		"admin",
		"su",
	}))
}

func TestEncrypt(t *testing.T) {
	assert := assert.New(t)
	key := []byte("01234567890123456789012345678901")
	data := []byte("abcd")

	result, err := Encrypt(key, data)
	assert.Nil(err)
	result, err = Decrypt(key, result)
	assert.Nil(err)
	assert.Equal(data, result)
}
