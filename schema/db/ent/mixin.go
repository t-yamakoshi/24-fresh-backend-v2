package ent

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Optional(),
		field.Time("updated_at").Optional(),
	}
}

type TimeMixinAdminOnly struct {
	mixin.Schema
}

func (TimeMixinAdminOnly) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Optional(),
		field.Time("updated_at").Optional(),
	}
}

type IdMixin struct {
	mixin.Schema
}

func (IdMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive().Immutable().SchemaType(
			map[string]string{
				dialect.MySQL: "bigint unsigned",
			},
		),
	}
}
