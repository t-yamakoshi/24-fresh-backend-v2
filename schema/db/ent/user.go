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
		field.String("name"),
		field.String("user_name"),
		field.String("email"),
	}
}
