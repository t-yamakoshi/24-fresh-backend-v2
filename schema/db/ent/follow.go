package ent

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type FollowsModel struct {
	ent.Schema
}

func (FollowsModel) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("follower_user_id"),
		field.Int64("followee_user_id"),
	}
}

func (FollowsModel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("follower", UserModel.Type).
			Ref("followers").
			Required().
			Unique().
			Field("follower_user_id"),
		edge.From("followee", UserModel.Type).
			Ref("followees").
			Required().
			Unique().
			Field("followee_user_id"),
	}
}

func (FollowsModel) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("follower_user_id", "followee_user_id").Unique(),
	}
}

func (FollowsModel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
