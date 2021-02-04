package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// DNSDetectorResult holds the schema definition for the DNSDetectorResult entity.
type DNSDetectorResult struct {
	ent.Schema
}

type DNSDetectorSubResult struct {
	Result   DetectorResult `json:"result,omitempty"`
	IPS      []string       `json:"ips,omitempty"`
	Server   string         `json:"server,omitempty"`
	Duration int            `json:"duration,omitempty"`
	Message  string         `json:"message,omitempty"`
}
type DNSDetectorSubResults []*DNSDetectorSubResult

// Fields of the DNSDetectorResult.
func (DNSDetectorResult) Fields() []ent.Field {
	return []ent.Field{
		field.String("host").
			Comment("检测Host"),
		field.JSON("results", DNSDetectorSubResults{}).
			Comment("检测结果列表"),
	}
}

// Mixin of the DNSDetectorResult
func (DNSDetectorResult) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		DetectorResultMixin{},
	}
}

// Edges of the DNSDetectorResult.
func (DNSDetectorResult) Edges() []ent.Edge {
	return nil
}
