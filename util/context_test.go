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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vicanso/elton"
	"github.com/vicanso/cyber-tect/config"

	"github.com/stretchr/testify/assert"
)

func TestGetTrackID(t *testing.T) {
	assert := assert.New(t)
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	c := elton.NewContext(resp, req)
	assert.Empty(GetTrackID(c))
	cookie := http.Cookie{
		Name:  config.GetTrackKey(),
		Value: "abcd",
	}
	req.AddCookie(&cookie)
	assert.Equal(cookie.Value, GetTrackID(c))
}

func TestGetSessionID(t *testing.T) {
	assert := assert.New(t)
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	c := elton.NewContext(resp, req)
	assert.Empty(GetSessionID(c))
	cookie := http.Cookie{
		Name:  config.GetSessionConfig().Key,
		Value: "abcd",
	}
	req.AddCookie(&cookie)
	assert.Equal(cookie.Value, GetSessionID(c))
}
