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

package validate

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func init() {
	AddAlias("xLimit", "number,min=0,max=10")
	AddAlias("xOffset", "number,min=0,max=1000")
	AddAlias("xFields", "min=0,max=100")

	durationRegexp := regexp.MustCompile("^[1-9][0-9]*(ms|[smh])$")
	Add("xDuration", func(fl validator.FieldLevel) bool {
		value, ok := toString(fl)
		if !ok {
			return false
		}
		return durationRegexp.MatchString(value)
	})
}
