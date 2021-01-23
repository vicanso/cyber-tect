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

package validate

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/vicanso/hes"
)

var (
	defaultValidator = validator.New()

	// validate默认的出错类别
	errCategory = "validate"
	// json parse失败时的出错类别
	errJSONParseCategory = "json-parse"
)

// toString 转换为string
func toString(fl validator.FieldLevel) (string, bool) {
	value := fl.Field()
	if value.Kind() != reflect.String {
		return "", false
	}
	return value.String(), true
}

// newNumberRange 校验number是否>=min <=max
func newNumberRange(min, max int) validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field()
		// 如果是int
		if value.Kind() == reflect.Int {
			number := int(value.Int())
			return number >= min && number <= max
		}
		// 如果是string
		if value.Kind() == reflect.String {
			number, err := strconv.Atoi(value.String())
			// 如果无法转换为int，则不符合
			if err != nil {
				return false
			}
			return number >= min && number <= max
		}
		return false
	}
}

// // isInt 判断是否int
// func isInt(fl validator.FieldLevel) bool {
// 	value := fl.Field()
// 	return value.Kind() == reflect.Int
// }

// // toInt 转换为int
// func toInt(fl validator.FieldLevel) (int, bool) {
// 	value := fl.Field()
// 	if value.Kind() != reflect.Int {
// 		return 0, false
// 	}
// 	return int(value.Int()), true
// }

// // isInInt 判断是否在int数组中
// func isInInt(fl validator.FieldLevel, values []int) bool {
// 	value, ok := toInt(fl)
// 	if !ok {
// 		return false
// 	}
// 	exists := false
// 	for _, v := range values {
// 		if v == value {
// 			exists = true
// 		}
// 	}
// 	return exists
// }

// newIsInString new is in string validator
func newIsInString(values []string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		return isInString(fl, values)
	}
}

// isInString 判断是否在string数组中
func isInString(fl validator.FieldLevel, values []string) bool {
	value, ok := toString(fl)
	if !ok {
		return false
	}
	exists := false
	for _, v := range values {
		if v == value {
			exists = true
		}
	}
	return exists
}

// // isAllInString 判断是否所有都在string数组中
// func isAllInString(fl validator.FieldLevel, values []string) bool {
// 	if fl.Field().Kind() != reflect.Slice {
// 		return false
// 	}
// 	v := fl.Field().Interface()
// 	value, ok := v.([]string)
// 	if !ok || len(value) == 0 {
// 		return false
// 	}
// 	valid := true
// 	for _, item := range value {
// 		exists := containsString(values, item)
// 		if !exists {
// 			valid = false
// 		}
// 	}
// 	return valid
// }

// // containsString 是否包含此string
// func containsString(arr []string, str string) (found bool) {
// 	for _, v := range arr {
// 		if found {
// 			break
// 		}
// 		if v == str {
// 			found = true
// 		}
// 	}
// 	return
// }

// doValidate 校验struct
func doValidate(s interface{}, data interface{}) (err error) {
	// statusCode := http.StatusBadRequest
	if data != nil {
		switch data := data.(type) {
		case []byte:
			err = json.Unmarshal(data, s)
			if err != nil {
				he := hes.Wrap(err)
				he.Category = errJSONParseCategory
				err = he
				return
			}
		default:
			buf, err := json.Marshal(data)
			if err != nil {
				return err
			}
			err = json.Unmarshal(buf, s)
			if err != nil {
				return err
			}
		}
	}
	err = defaultValidator.Struct(s)
	return
}

// Do 执行校验
func Do(s interface{}, data interface{}) (err error) {
	err = doValidate(s, data)
	if err != nil {
		he := hes.Wrap(err)
		if he.Category == "" {
			he.Category = errCategory
		}
		err = he
	}
	return
}

// Add 添加一个校验函数
func Add(tag string, fn validator.Func, args ...bool) {
	err := defaultValidator.RegisterValidation(tag, fn, args...)
	if err != nil {
		panic(err)
	}
}

// AddAlias add alias
func AddAlias(alias, tags string) {
	defaultValidator.RegisterAlias(alias, tags)
}
