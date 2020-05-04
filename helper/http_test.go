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
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vicanso/elton"
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/go-axios"
	"github.com/vicanso/hes"
)

func TestGetHTTPStats(t *testing.T) {
	assert := assert.New(t)
	service := "test"
	cid := "cid"
	route := "/users/:id"
	method := "get"
	url := "/users/1"
	status := 200
	conf := &axios.Config{
		Route:  route,
		URL:    url,
		Method: method,
	}
	conf.Set(cs.CID, cid)
	resp := &axios.Response{
		Status: status,
		Config: conf,
	}
	tags, fields := getHTTPStats(service, resp)
	assert.Equal(tags["route"], route)
	assert.Equal(tags["service"], service)
	assert.Equal(tags["method"], method)
	assert.Equal(fields["cid"], cid)
	assert.Equal(fields["url"], url)
	assert.Equal(fields["status"], status)
}

func TestConvertResponseToError(t *testing.T) {
	assert := assert.New(t)
	service := "test"
	fn := newConvertResponseToError(service)

	err := fn(&axios.Response{
		Status: 200,
	})
	assert.Nil(err)

	resp := &axios.Response{
		Status: 400,
		Data: []byte(`{
			"message": "出错了"
		}`),
	}
	err = fn(resp)
	assert.NotNil(err)
	assert.Equal("出错了", err.Error())
}

func TestOnError(t *testing.T) {
	assert := assert.New(t)
	service := "test"
	message := "error message"
	conf := &axios.Config{}
	fn := newOnError(service)
	err := fn(errors.New(message), conf)
	he, ok := err.(*hes.Error)
	assert.True(ok)
	assert.Equal(message, he.Message)
}

func TestNewInstance(t *testing.T) {
	assert := assert.New(t)
	timeout := 10 * time.Second
	baseURL := "https://example.com"
	ins := NewInstance("test", baseURL, timeout)
	assert.Equal(timeout, ins.Config.Timeout)
	assert.Equal(baseURL, ins.Config.BaseURL)
}

func TestAttachWithContext(t *testing.T) {
	assert := assert.New(t)
	c := elton.NewContext(nil, nil)
	c.ID = "context id"
	conf := &axios.Config{}
	AttachWithContext(conf, c)
	assert.Equal(c.ID, conf.GetString(cs.CID))
}
