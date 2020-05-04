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
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vicanso/elton"
	"github.com/vicanso/cybertect/service"
)

func TestNoQuery(t *testing.T) {
	assert := assert.New(t)
	req := httptest.NewRequest("GET", "/users/me?type=1", nil)
	resp := httptest.NewRecorder()
	c := elton.NewContext(resp, req)
	c.Next = func() error {
		return nil
	}
	err := NoQuery(c)
	assert.Equal(errQueryNotAllow, err)

	c.Request = httptest.NewRequest("GET", "/users/me", nil)
	err = NoQuery(c)
	assert.Nil(err)
}

func TestWaitFor(t *testing.T) {
	assert := assert.New(t)
	req := httptest.NewRequest("GET", "/users/me", nil)
	resp := httptest.NewRecorder()
	c := elton.NewContext(resp, req)
	c.Next = func() error {
		return nil
	}
	start := time.Now()
	fn := WaitFor(2 * time.Millisecond)
	err := fn(c)
	assert.Nil(err)
	assert.True(time.Since(start) > time.Millisecond)
}

func TestValidateCaptcha(t *testing.T) {
	assert := assert.New(t)
	info, err := service.GetCaptcha("255,255,255", "102,102,102")
	assert.Nil(err)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set(xCaptchHeader, fmt.Sprintf("%s:%s", info.ID, info.Value))
	resp := httptest.NewRecorder()
	c := elton.NewContext(resp, req)
	c.Next = func() error {
		return nil
	}
	fn := ValidateCaptcha("")
	err = fn(c)
	assert.Nil(err)
}
