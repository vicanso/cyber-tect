package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// PingDetectorResult holds the schema definition for the PingDetectorResult entity.
type PingDetectorResult struct {
	ent.Schema
}

type PingDetectorSubResult struct {
	Result   DetectorResult `json:"result,omitempty"`
	IP       string         `json:"ip,omitempty"`
	Duration int            `json:"duration,omitempty"`
	Message  string         `json:"message,omitempty"`
}

type PingDetectorSubResults []*PingDetectorSubResult

// Fields of the PingDetectorResult.
func (PingDetectorResult) Fields() []ent.Field {
	return []ent.Field{
		field.String("ips").
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
