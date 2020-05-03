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
	"encoding/json"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/vicanso/hes"
)

var (
	defaultValidator = validator.New()

	errCategory          = "validate"
	errJSONParseCategory = "json-parse"
)

func toString(fl validator.FieldLevel) (string, bool) {
	value := fl.Field()
	if value.Kind() != reflect.String {
		return "", false
	}
	return value.String(), true
}
func toInt(fl validator.FieldLevel) (int, bool) {
	value := fl.Field()
	if value.Kind() != reflect.Int {
		return 0, false
	}
	return int(value.Int()), true
}
func isInInt(fl validator.FieldLevel, values []int) bool {
	value, ok := toInt(fl)
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
func isZero(fl validator.FieldLevel) bool {
	return fl.Field().IsZero()
}

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

// Do do validate
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

// Add add validate
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
