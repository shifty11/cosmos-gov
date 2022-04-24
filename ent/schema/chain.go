package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("name").
			Unique(),
		field.String("display_name").
			Unique(),
		field.Bool("is_enabled").
			Default(true),
	}
}

// Edges of the Chain.
func (Chain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("chains"),
		edge.To("proposals", Proposal.Type),
		edge.To("rpc_endpoints", RpcEndpoint.Type),
	}
}

func (Chain) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique(),
	}
}
