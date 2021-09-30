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

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// 检测结果
type DetectorResult int8

const (
	// DetectorResultSuccess 成功
	DetectorResultSuccess DetectorResult = iota + 1
	// DetectorResultFail 失败
	DetectorResultFail
)

func (result DetectorResult) String() string {
	switch result {
	case DetectorResultSuccess:
		return "成功"
	case DetectorResultFail:
		return "失败"
	default:
		return "未知"
	}
}

// DetectorMixin mixin of detector
type DetectorMixin struct {
	mixin.Schema
}

// Fields of the detector
func (DetectorMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("配置名称"),
		field.Strings("owners").
			Comment("配置拥有者"),
		field.Strings("receivers").
			Comment("接收者列表"),
		field.String("timeout").
			NotEmpty().
			Comment("超时设置"),
		field.String("description").
			Comment("配置描述"),
	}
}

// Indexes of the detector
func (DetectorMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owners"),
	}
}

type DetectorResultMixin struct {
	mixin.Schema
}

// Fields of the DetectorResultMixin
func (DetectorResultMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("task").
			Comment("任务ID"),
		field.Int8("result").
			Range(int8(DetectorResultSuccess), int8(DetectorResultFail)).
			GoType(DetectorResultFail).
			Comment("检测结果"),
		field.Int("maxDuration").
			StructTag(`json:"maxDuration,omitempty" sql:"max_duration"`).
			Comment("最长时长"),
		field.Strings("messages").
			Comment("出错信息汇总"),
	}
}

// Indexes of the DetectorResultMixin
func (DetectorResultMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("task"),
		index.Fields("result"),
	}
}
