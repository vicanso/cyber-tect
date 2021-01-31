package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// TCPDetectorResult holds the schema definition for the TCPDetectorResult entity.
type TCPDetectorResult struct {
	ent.Schema
}

type TCPDetectorSubResult struct {
	Result   DetectorResult `json:"result,omitempty"`
	Addr     string         `json:"addr,omitempty"`
	Duration int            `json:"duration,omitempty"`
	Message  string         `json:"message,omitempty"`
}

type TCPDetectorSubResults []*TCPDetectorSubResult

// Fields of the TCPDetectorResult.
func (TCPDetectorResult) Fields() []ent.Field {
	return []ent.Field{
		field.String("addrs").
			Comment("检测地址"),
		field.JSON("results", TCPDetectorSubResults{}).
			Comment("检测结果列表"),
	}
}

// Mixin of the TCPDetectorResult.
func (TCPDetectorResult) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		StatusMixin{},
		DetectorResultMixin{},
	}
}

// Edges of the TCPDetectorResult.
func (TCPDetectorResult) Edges() []ent.Edge {
	return nil
}
