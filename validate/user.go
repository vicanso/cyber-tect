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
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/vicanso/cybertect/cs"
)

func init() {
	// 账号
	AddAlias("xUserAccount", "ascii,len=0|min=4,max=10")
	AddAlias("xUserPassword", "ascii,len=0|len=44")
	AddAlias("xUserAccountKeyword", "ascii,min=0,max=10")
	Add("xUserRole", func(fl validator.FieldLevel) bool {
		return isZero(fl) || isInString(fl, []string{
			cs.UserRoleSu,
			cs.UserRoleAdmin,
		})
	})
	Add("xUserRoles", func(fl validator.FieldLevel) bool {
		if fl.Field().Kind() != reflect.Slice {
			return false
		}
		v := fl.Field().Interface()
		value, ok := v.([]string)
		if !ok {
			return false
		}
		valid := true
		for _, item := range value {
			exists := false
			for _, role := range []string{
				cs.UserRoleSu,
				cs.UserRoleAdmin,
			} {
				if item == role {
					exists = true
				}
			}
			if !exists {
				valid = false
			}
		}
		return valid
	})
}
