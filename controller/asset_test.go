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

package controller

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vicanso/elton"
)

func TestStaticFile(t *testing.T) {
	assert := assert.New(t)
	assert.True(assetFS.Exists("index.html"))
	assert.False(assetFS.Exists("test.html"))

	buf, err := assetFS.Get("index.html")
	assert.Nil(err)
	assert.NotEmpty(buf)

	assert.Nil(assetFS.Stat("index.html"))

	r, err := assetFS.NewReader("index.html")
	assert.Nil(err)
	assert.NotNil(r)
	buf, err = io.ReadAll(r)
	assert.Nil(err)
	assert.NotEmpty(buf)
}

func TestAassetCtrl(t *testing.T) {
	assert := assert.New(t)
	ctrl := assetCtrl{}
	t.Run("get index", func(t *testing.T) {
		c := elton.NewContext(httptest.NewRecorder(), nil)
		err := ctrl.getIndex(c)
		assert.Nil(err)
		assert.NotEmpty(c.BodyBuffer.String())
		assert.Equal("text/html; charset=utf-8", c.GetHeader(elton.HeaderContentType))
		assert.Equal("public, max-age=10", c.GetHeader(elton.HeaderCacheControl))
	})
	t.Run("get favicon", func(t *testing.T) {
		c := elton.NewContext(httptest.NewRecorder(), nil)
		err := ctrl.getFavIcon(c)
		assert.Nil(err)
		assert.NotEqual(0, c.BodyBuffer.Len())
		assert.Equal("public, max-age=3600, s-maxage=600", c.GetHeader(elton.HeaderCacheControl))
	})
}
