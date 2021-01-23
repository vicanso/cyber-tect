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

package helper

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vicanso/elton"
	"github.com/vicanso/go-axios"
	"github.com/vicanso/hes"
)

func TestHTTPStats(t *testing.T) {
	assert := assert.New(t)

	data := []byte("abcd")
	ins := NewHTTPInstance("test", "https://test.com", time.Second)
	done := ins.Mock(&axios.Response{
		Status: 200,
		Data:   data,
	})
	defer done()
	resp, err := ins.Get("/")
	assert.Nil(err)
	assert.Equal(200, resp.Status)
	assert.Equal(data, resp.Data)
	assert.NotNil(resp.Config.HTTPTrace)
}

func TestConvertResponseToError(t *testing.T) {
	assert := assert.New(t)
	fn := newHTTPConvertResponseToError("test")
	data := []byte(`{
		"message": "error message"
	}`)
	err := fn(&axios.Response{
		Status: 400,
		Data:   data,
	})
	assert.Equal("message=error message", err.Error())

	ins := NewHTTPInstance("test", "https://test.com", time.Second)
	done := ins.Mock(&axios.Response{
		Status: 400,
		Data:   data,
	})
	defer done()
	resp, err := ins.Get("/")
	assert.Equal("category=test, message=error message", err.Error())
	assert.Equal(400, resp.Status)
	assert.Equal(data, resp.Data)
}

func TestOnError(t *testing.T) {
	assert := assert.New(t)
	data := []byte(`{
		"message": "error message"
	}`)
	ins := NewHTTPInstance("test", "https://test.com", time.Second)
	done := ins.Mock(&axios.Response{
		Status: 400,
		Data:   data,
	})
	resp, err := ins.Request(&axios.Config{
		Route: "/",
	})
	done()
	he := hes.Wrap(err)
	assert.Equal(`{"statusCode":400,"category":"test","message":"error message","extra":{"requestCURL":"curl -XGET 'https://test.com'","requestRoute":"/","requestService":"test"}}`, string(he.ToJSON()))
	assert.Equal("/", resp.Config.Route)

	data = []byte("abc")
	done = ins.Mock(&axios.Response{
		Status: 400,
		Data:   data,
	})
	resp, err = ins.Request(&axios.Config{
		Route: "/",
	})
	done()
	he = hes.Wrap(err)
	assert.Equal(`{"statusCode":400,"category":"test","message":"abc","extra":{"requestCURL":"curl -XGET 'https://test.com'","requestRoute":"/","requestService":"test"}}`, string(he.ToJSON()))
	assert.Equal("/", resp.Config.Route)
}

func TestAttachWithContext(t *testing.T) {
	assert := assert.New(t)
	config := &axios.Config{}
	req := httptest.NewRequest("GET", "/", nil)
	c := elton.NewContext(nil, req)
	c.ID = "abcd"
	AttachWithContext(config, c)
	assert.Equal(c.Context(), config.Context)
}
