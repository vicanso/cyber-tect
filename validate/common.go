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
	"net/url"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func init() {
	// 客户端使用的limit，最多只允许一次拉取100条
	Add("xLimit", newNumberRange(1, 200))
	// 更大的数量限制，一般管理后台接口使用
	Add("xLargerLimit", newNumberRange(1, 200))
	// offset设置已满足大部分场景需要，如果有更多的处理再调整
	Add("xOffset", newNumberRange(0, 100*1000))
	AddAlias("xOrder", "ascii,min=0,max=100")
	AddAlias("xFields", "ascii,min=0,max=100")
	AddAlias("xKeyword", "min=1,max=10")
	// 状态：启用、禁用
	AddAlias("xStatus", "numeric,min=1,max=2")
	// path校验
	AddAlias("xPath", "startswith=/")
	// boolean的字符串形式，0: false, 1:true
	AddAlias("xBoolean", "oneof=0 1")
	// http(s)校验
	Add("xHTTP", func(fl validator.FieldLevel) bool {
		v, ok := toString(fl)
		if !ok {
			return false
		}
		urlInfo, err := url.ParseRequestURI(v)
		if err != nil {
			return false
		}
		if urlInfo.Scheme != "http" &&
			urlInfo.Scheme != "https" {
			return false
		}
		return urlInfo.Host != "" && urlInfo.Path != ""
	})
	// duration配置
	durationReg := regexp.MustCompile(`^\d+(ms|s|m|h)$`)
	Add("xDuration", func(fl validator.FieldLevel) bool {
		v, _ := toString(fl)
		return durationReg.MatchString(v)
	})
}
