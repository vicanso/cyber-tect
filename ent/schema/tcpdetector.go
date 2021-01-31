package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// TCPDetector holds the schema definition for the TCPDetector entity.
type TCPDetector struct {
	ent.Schema
}

// Fields of the TCPDetector.
func (TCPDetector) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("addrs").
			Comment("检测地址列表"),
	}
}

// Mixin of the TCPDetector
func (TCPDetector) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		StatusMixin{},
		DetectorMixin{},
	}
}

// Edges of the TCPDetector.
func (TCPDetector) Edges() []ent.Edge {
	return nil
}
