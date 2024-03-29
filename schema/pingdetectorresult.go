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

// PingDetectorResult holds the schema definition for the PingDetectorResult entity.
type PingDetectorResult struct {
	ent.Schema
}

type PingDetectorSubResult struct {
	Result   DetectorResult `json:"result"`
	IP       string         `json:"ip"`
	Duration int            `json:"duration"`
	Message  string         `json:"message"`
}

type PingDetectorSubResults []*PingDetectorSubResult

// Fields of the PingDetectorResult.
func (PingDetectorResult) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("ips").
			Comment("检测IP"),
		field.JSON("results", PingDetectorSubResults{}).
			Comment("检测结果列表"),
	}
}

func (PingDetectorResult) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		DetectorResultMixin{},
	}
}

// Edges of the PingDetectorResult.
func (PingDetectorResult) Edges() []ent.Edge {
	return nil
}
