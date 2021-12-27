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

func TestIsMatchAlarmCount(t *testing.T) {
	assert := assert.New(t)

	assert.True(isMatchAlarmCount(1))
	assert.True(isMatchAlarmCount(2))
	assert.True(isMatchAlarmCount(3))
	assert.False(isMatchAlarmCount(4))
	assert.True(isMatchAlarmCount(10))
	assert.False(isMatchAlarmCount(61))
	assert.True(isMatchAlarmCount(120))
}
