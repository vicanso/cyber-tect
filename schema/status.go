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
	"entgo.io/ent/schema/mixin"
)

type Status int8

// StatusInfo 状态信息
type StatusInfo struct {
	Name  string `json:"name"`
	Value Status `json:"value"`
}

const (
	// 状态启用
	StatusEnabled Status = iota + 1
	// 状态禁用
	StatusDisabled
)

// ToInt8 转换为int8
func (status Status) Int8() int8 {
	return int8(status)
}

// String 转换为string
func (status Status) String() string {
	switch status {
	case StatusEnabled:
		return "启用"
	case StatusDisabled:
		return "禁用"
	default:
		return "未知"
	}
}

// GetSchemaStatusList 获取schema的状态列表
func GetStatusList() []*StatusInfo {
	values := []Status{
		StatusEnabled,
		StatusDisabled,
	}
	list := make([]*StatusInfo, len(values))
	for index, value := range values {
		list[index] = &StatusInfo{
			Name:  value.String(),
			Value: value,
		}
	}
	return list
}

// StatusMixin 状态的schema
type StatusMixin struct {
	mixin.Schema
}

// Fields 公共的status的字段
func (StatusMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("status").
			Range(StatusEnabled.Int8(), StatusDisabled.Int8()).
			Default(StatusEnabled.Int8()).
			GoType(Status(StatusEnabled)).
			Comment("状态，默认为启用状态"),
	}
}
