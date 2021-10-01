package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TCPDetectorResult holds the schema definition for the TCPDetectorResult entity.
type TCPDetectorResult struct {
	ent.Schema
}

type TCPDetectorSubResult struct {
	Result   DetectorResult `json:"result"`
	Addr     string         `json:"addr"`
	Duration int            `json:"duration"`
	Message  string         `json:"message"`
}

type TCPDetectorSubResults []*TCPDetectorSubResult

// Fields of the TCPDetectorResult.
func (TCPDetectorResult) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("addrs").
			Comment("检测地址"),
		field.JSON("results", TCPDetectorSubResults{}).
			Comment("检测结果列表"),
	}
}

// Mixin of the TCPDetectorResult.
func (TCPDetectorResult) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		DetectorResultMixin{},
	}
}

// Edges of the TCPDetectorResult.
func (TCPDetectorResult) Edges() []ent.Edge {
	return nil
}
