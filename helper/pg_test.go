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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPGFormatOrder(t *testing.T) {
	assert := assert.New(t)
	order := PGFormatOrder("-id,date,-price")
	assert.Equal("id desc,date,price desc", order)
}

func TestPGFormatSelect(t *testing.T) {
	assert := assert.New(t)
	fields := PGFormatSelect("id,updatedAt")
	assert.Equal("id,updated_at", fields)
}

func TestPGQuery(t *testing.T) {
	assert := assert.New(t)
	db := PGQuery(PGQueryParams{
		Limit:  10,
		Offset: 1,
		Fields: "id,title",
		Order:  "-id,date,-price",
	})
	assert.NotNil(db)
}
