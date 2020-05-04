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

func TestTime(t *testing.T) {
	assert := assert.New(t)
	mockTime := "2020-04-26T20:34:33+08:00"
	SetMockTime(mockTime)
	defer SetMockTime("")

	assert.Equal(int64(1587904473000000000), Now().UnixNano())
	// assert.Equal(mockTime, NowString())
	assert.Equal("2020-04-26 12:34:33 +0000 UTC", UTCNow().String())

	value, err := ParseTime(mockTime)
	assert.Nil(err)
	assert.Equal("2020-04-26T20:34:33+08:00", FormatTime(value))

	assert.Equal("2020-04-26T20:34:33+08:00", FormatTime(ChinaNow()))
}
