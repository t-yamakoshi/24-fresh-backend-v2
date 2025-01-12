package ent

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type UserModel struct {
	ent.Schema
}

func (UserModel) Fields() []ent.Field {
	return []ent.Field{
		field.String("cognito_id"),
		field.String("name"),
		field.String("user_name"),
		field.Time("birth_date"),
		field.Enum("gender").Optional().Values(
			"men", "women", "other",
		),
		field.String("self_introduction").Optional(),
		field.String("profile_image").Optional(),
		field.String("email"),
	}
}

func (UserModel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IdMixin{},
		TimeMixin{},
	}
}
