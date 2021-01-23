package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// HTTP holds the schema definition for the HTTP entity.
type HTTP struct {
	ent.Schema
}

// Fields of the HTTP.
func (HTTP) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("配置名称"),
		field.String("owner").
			NotEmpty().
			Comment("配置拥有者"),
		field.String("description").
			Comment("配置描述"),
		field.Strings("receivers").
			Comment("接收者列表"),
		field.Strings("ips").
			Comment("IP列表"),
		field.String("url").
			NotEmpty().
			Comment("测试URL"),
		field.String("timeout").
			NotEmpty().
			Comment("超时设置"),
	}
}

// Mixin http mixin
func (HTTP) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		StatusMixin{},
	}
}

// Edges of the HTTP.
func (HTTP) Edges() []ent.Edge {
	return nil
}

// Indexes http 索引
func (HTTP) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner"),
	}
}

// Annotations of the User.
func (HTTP) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "https"},
	}
}
