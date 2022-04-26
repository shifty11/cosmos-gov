package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Grant holds the schema definition for the Grant entity.
type Grant struct {
	ent.Schema
}

func (Grant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Grant.
func (Grant) Fields() []ent.Field {
	return []ent.Field{
		field.String("grantee"),
		field.String("type"),
		field.Time("expires_at"),
	}
}

// Edges of the Grant.
func (Grant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("granter", Wallet.Type).
			Ref("grants").
			Unique(),
	}
}
