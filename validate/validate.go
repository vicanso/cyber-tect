package validate

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	jsoniter "github.com/json-iterator/go"
	"github.com/vicanso/hes"
)

var (
	paramTagRegexMap = govalidator.ParamTagRegexMap
	paramTagMap      = govalidator.ParamTagMap
	customTypeTagMap = govalidator.CustomTypeTagMap
	json             = jsoniter.ConfigCompatibleWithStandardLibrary
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	AddRegex("xIntRange", "^xIntRange\\((\\d+)\\|(\\d+)\\)$", func(value string, params ...string) bool {
		return govalidator.InRangeInt(value, params[0], params[1])
	})

	AddRegex("xIntIn", `^xIntIn\((.*)\)$`, func(value string, params ...string) bool {
		if len(params) == 1 {
			rawParams := params[0]
			parsedParams := strings.Split(rawParams, "|")
			return govalidator.IsIn(value, parsedParams...)
		}
		return false
	})

	methods := []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"HEAD",
	}
	Add("xMethods", func(i interface{}, _ interface{}) bool {
		value, ok := i.(string)
		if !ok {
			return false
		}
		return govalidator.IsIn(value, methods...)
	})
}

// Do do validate
func Do(s interface{}, data interface{}) (err error) {
	statusCode := http.StatusBadRequest
	if data != nil {
		switch data.(type) {
		case []byte:
			e := json.Unmarshal(data.([]byte), s)
			if e != nil {
				err = hes.NewWithErrorStatusCode(e, statusCode)
				return
			}
		default:
			buf, e := json.Marshal(data)
			if e != nil {
				err = hes.NewWithErrorStatusCode(e, statusCode)
				return
			}
			e = json.Unmarshal(buf, s)
			if e != nil {
				err = hes.NewWithErrorStatusCode(e, statusCode)
				return
			}
		}
	}
	_, err = govalidator.ValidateStruct(s)
	if err != nil {
		err = hes.NewWithErrorStatusCode(err, statusCode)
	}
	return
}

// AddRegex add a regexp validate
func AddRegex(name, reg string, fn govalidator.ParamValidator) {
	if paramTagMap[name] != nil {
		panic(name + ", reg:" + reg + " is duplicated")
	}
	paramTagRegexMap[name] = regexp.MustCompile(reg)
	paramTagMap[name] = fn
}

// Add add validate
func Add(name string, fn govalidator.CustomTypeValidator) {
	_, exists := customTypeTagMap.Get(name)
	if exists {
		panic(name + " is duplicated")
	}
	customTypeTagMap.Set(name, fn)
}
