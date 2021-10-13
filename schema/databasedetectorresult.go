// Copyright 2021 tree xie
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
)

type DatabaseDetectorResult struct {
	ent.Schema
}

type DatabaseDetectorSubResult struct {
	Result   DetectorResult `json:"result"`
	URI      string         `json:"uri"`
	Duration int            `json:"duration"`
	Message  string         `json:"message"`
}

type DatabaseDetectorSubResults []*DatabaseDetectorSubResult

// Fields of the DatabaseDetectorResult
func (DatabaseDetectorResult) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("uris").
			Comment("检测的redis连接地址"),
		field.JSON("results", DatabaseDetectorSubResults{}).
			Comment("检测结果列表"),
	}
}

func (DatabaseDetectorResult) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		DetectorResultMixin{},
	}
}
