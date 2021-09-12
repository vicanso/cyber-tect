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
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/cast"
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
func doValidate(s interface{}, data interface{}) error {
	// statusCode := http.StatusBadRequest
	if data != nil {
		switch data := data.(type) {
		case []byte:
			if len(data) == 0 {
				he := hes.New("data is empty")
				he.Category = errJSONParseCategory
				return he
			}
			err := json.Unmarshal(data, s)
			if err != nil {
				he := hes.Wrap(err)
				he.Category = errJSONParseCategory
				return he
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
	defaults.SetDefaults(s)
	return defaultValidator.Struct(s)
}

func wrapError(err error) error {

	he := hes.Wrap(err)
	if he.Category == "" {
		he.Category = errCategory
	}
	he.StatusCode = http.StatusBadRequest
	return he
}

// Query 转换数据后执行校验，用于将query转换为struct时使用
func Query(s interface{}, data map[string]string) (err error) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr {
		err = wrapError(errors.New("only support pointer"))
		return
	}
	v = v.Elem()
	t := v.Type()
	len := t.NumField()
	for i := 0; i < len; i++ {
		field := t.Field(i)
		// we can't access the value of unexported fields
		if field.PkgPath != "" {
			continue
		}
		value := v.FieldByIndex(field.Index)
		tag := field.Tag.Get("json")
		tagValue := data[tag]
		// 如果值为空，则不做赋值处理
		if tagValue == "" {
			continue
		}
		switch value.Kind() {
		case reflect.Int:
			fallthrough
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			v, e := cast.ToInt64E(tagValue)
			if e != nil {
				err = wrapError(e)
				return
			}
			value.SetInt(v)
		case reflect.Float64:
			fallthrough
		case reflect.Float32:
			v, e := cast.ToFloat64E(tagValue)
			if e != nil {
				err = wrapError(e)
				return
			}
			value.SetFloat(v)
		case reflect.Bool:
			v, e := cast.ToBoolE(tagValue)
			if e != nil {
				err = wrapError(e)
				return
			}
			value.SetBool(v)
		case reflect.String:
			value.SetString(tagValue)
		default:
			err = wrapError(fmt.Errorf("not support the type:%s", tag))
			return
		}
	}
	defaults.SetDefaults(s)
	err = defaultValidator.Struct(s)
	return
}

// Do 执行校验
func Do(s interface{}, data interface{}) error {
	err := doValidate(s, data)
	if err != nil {
		return wrapError(err)
	}
	return nil
}

// 对struct校验
func Struct(s interface{}) error {
	defaults.SetDefaults(s)
	err := defaultValidator.Struct(s)
	if err != nil {
		return wrapError(err)
	}
	return nil
}

// 任何参数均返回true，不校验。用于临时将某个校验禁用
func notValidate(fl validator.FieldLevel) bool {
	return true
}

func getCustomDefine(tag string) string {
	return os.Getenv("VALIDATE_" + tag)
}

// Add 添加一个校验函数
func Add(tag string, fn validator.Func, args ...bool) {
	custom := getCustomDefine(tag)
	if custom == "*" {
		_ = defaultValidator.RegisterValidation(tag, notValidate)
		return
	}
	if custom != "" {
		defaultValidator.RegisterAlias(tag, custom)
		return
	}
	err := defaultValidator.RegisterValidation(tag, fn, args...)
	if err != nil {
		panic(err)
	}
}

// AddAlias add alias
func AddAlias(alias, tags string) {
	custom := getCustomDefine(alias)
	if custom == "*" {
		_ = defaultValidator.RegisterValidation(alias, notValidate)
		return
	}
	if custom != "" {
		tags = custom
	}
	defaultValidator.RegisterAlias(alias, tags)
}
