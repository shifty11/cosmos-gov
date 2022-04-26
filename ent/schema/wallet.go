package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Wallet holds the schema definition for the Wallet entity.
type Wallet struct {
	ent.Schema
}

func (Wallet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").
			Unique(),
	}
}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("wallets"),
		edge.From("chain", Chain.Type).
			Ref("wallets").
			Unique(),
		edge.To("grants", Grant.Type),
	}
}

func (Wallet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("address").
			Unique(),
	}
}
