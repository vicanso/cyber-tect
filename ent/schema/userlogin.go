package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserLogin holds the schema definition for the UserLogin entity.
type UserLogin struct {
	ent.Schema
}

// Mixin 用户登录记录的mixin
func (UserLogin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields 用户登录表的相关字段
func (UserLogin) Fields() []ent.Field {
	return []ent.Field{
		field.String("account").
			NotEmpty().
			Immutable().
			Comment("登录账户"),
		field.String("user_agent").
			StructTag(`json:"userAgent,omitempty"`).
			Optional().
			Comment("用户浏览器的user-agent"),
		field.String("ip").
			Optional().
			Comment("用户IP"),
		field.String("track_id").
			StructTag(`json:"trackID,omitempty"`).
			Optional().
			Comment("用户的track id"),
		field.String("session_id").
			StructTag(`json:"sessionID,omitempty"`).
			Optional().
			Comment("用户的session id"),
		field.String("x_forwarded_for").
			StructTag(`json:"xForwardedFor,omitempty"`).
			Optional().
			Comment("用户登录时的x-forwarded-for"),
		field.String("country").
			Optional().
			Comment("用户登录IP定位的国家"),
		field.String("province").
			Optional().
			Comment("用户登录IP定位的省份"),
		field.String("city").
			Optional().
			Comment("用户登录IP定位的城市"),
		field.String("isp").
			Optional().
			Comment("用户登录IP的网络服务商"),
	}
}

// Edges of the UserLogin.
func (UserLogin) Edges() []ent.Edge {
	return nil
}

// Indexes 用户登录表索引
func (UserLogin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account"),
		index.Fields("created_at"),
	}
}
