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

package middleware

import (
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/vicanso/elton"
	"github.com/vicanso/hes"

	"github.com/stretchr/testify/assert"
)

func TestNewConcurrentLimit(t *testing.T) {
	assert := assert.New(t)

	ttl := 2 * time.Millisecond
	fn := NewConcurrentLimit([]string{
		"q:type",
	}, ttl, "lock-")
	req := httptest.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	c := elton.NewContext(resp, req)
	c.Next = func() error {
		return nil
	}
	err := fn(c)
	assert.Nil(err)

	err = fn(c)
	he, ok := err.(*hes.Error)
	assert.True(ok)
	assert.Equal("elton-concurrent-limiter", he.Category)

	time.Sleep(3 * time.Millisecond)
	err = fn(c)
	assert.Nil(err)
}

func TestNewIPLimit(t *testing.T) {
	assert := assert.New(t)

	ttl := 2 * time.Millisecond
	fn := NewIPLimit(1, ttl, "lock-")
	req := httptest.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	c := elton.NewContext(resp, req)
	c.Next = func() error {
		return nil
	}
	err := fn(c)
	assert.Nil(err)

	err = fn(c)
	assert.Equal(err, errTooFrequently)

	time.Sleep(3 * time.Millisecond)
	err = fn(c)
	assert.Nil(err)
}

func TestNewErrorLimit(t *testing.T) {
	assert := assert.New(t)

	ttl := 2 * time.Millisecond
	fn := NewErrorLimit(1, ttl, func(c *elton.Context) string {
		return c.Request.RequestURI
	})
	req := httptest.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	c := elton.NewContext(resp, req)
	customErr := errors.New("error")
	c.Next = func() error {
		return customErr
	}

	// 第一次可以通过，返回的是customErr
	err := fn(c)
	assert.Equal(customErr, err)

	// 第二次不通过
	err = fn(c)
	assert.Equal(errTooFrequently, err)

	// 过期之后，返回的是customErr
	time.Sleep(3 * time.Millisecond)
	err = fn(c)
	assert.Equal(customErr, err)
}
