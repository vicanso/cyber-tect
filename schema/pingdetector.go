package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PingDetector holds the schema definition for the PingDetector entity.
type PingDetector struct {
	ent.Schema
}

// Fields of the PingDetector.
func (PingDetector) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("ips").
			Comment("检测IP列表"),
	}
}

// Mixin of the TCPDetector
func (PingDetector) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		StatusMixin{},
		DetectorMixin{},
	}
}

// Edges of the PingDetector.
func (PingDetector) Edges() []ent.Edge {
	return nil
}
