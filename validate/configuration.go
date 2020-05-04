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
	"github.com/go-playground/validator/v10"
	"github.com/vicanso/cybertect/cs"
)

func init() {
	// 应用配置名称
	AddAlias("xConfigName", "alphanum,len=0|min=2,max=20")
	AddAlias("xConfigCategory", "alphanum,len=0|min=2,max=20")
	AddAlias("xConfigData", "min=0,max=500")
	AddAlias("xConfigNames", "min=0,max=100")

	Add("xConfigStatus", func(fl validator.FieldLevel) bool {
		// 公共配置的都允许为空
		return isZero(fl) || isInInt(fl, []int{
			cs.ConfigEnabled,
			cs.ConfigDiabled,
		})
	})
}
