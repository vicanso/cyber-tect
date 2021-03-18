package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// DNSDetector holds the schema definition for the DNSDetector entity.
type DNSDetector struct {
	ent.Schema
}

// Fields of the DNSDetector.
func (DNSDetector) Fields() []ent.Field {
	return []ent.Field{
		field.String("host").
			NotEmpty().
			Comment("域名地址"),
		field.Strings("ips").
			Comment("域名配置的IP列表"),
		field.Strings("servers").
			Comment("DNS服务器列表"),
	}
}

// Mixin of the DNSDetector.
func (DNSDetector) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		StatusMixin{},
		DetectorMixin{},
	}
}

// Edges of the DNSDetector.
func (DNSDetector) Edges() []ent.Edge {
	return nil
}
