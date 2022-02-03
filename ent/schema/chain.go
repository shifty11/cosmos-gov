package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Chain holds the schema definition for the Chain entity.
type Chain struct {
	ent.Schema
}

// Fields of the Chain.
func (Chain) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("name").
			Unique(),
		field.String("display_name").
			Unique(),
	}
}

// Edges of the Chain.
func (Chain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("chains"),
		edge.To("proposals", Proposal.Type),
	}
}
