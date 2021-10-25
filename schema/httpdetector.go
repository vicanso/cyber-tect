package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// HTTPDetector holds the schema definition for the HTTP entity.
type HTTPDetector struct {
	ent.Schema
}

// Fields of the HTTP.
func (HTTPDetector) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("ips").
			Comment("IP列表"),
		field.String("url").
			NotEmpty().
			Comment("测试URL"),
		field.String("script").
			Comment("检测脚本"),
	}
}

// Mixin http mixin
func (HTTPDetector) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		StatusMixin{},
		DetectorMixin{},
	}
}

// Edges of the HTTP.
func (HTTPDetector) Edges() []ent.Edge {
	return nil
}
