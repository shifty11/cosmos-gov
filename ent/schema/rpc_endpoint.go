package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// RpcEndpoint holds the schema definition for the RpcEndpoint entity.
type RpcEndpoint struct {
	ent.Schema
}

func (RpcEndpoint) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the RpcEndpoint.
func (RpcEndpoint) Fields() []ent.Field {
	return []ent.Field{
		field.String("endpoint").
			Unique(),
	}
}

// Edges of the RpcEndpoint.
func (RpcEndpoint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chain", Chain.Type).
			Ref("rpc_endpoints").
			Unique(),
	}
}

func (RpcEndpoint) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("endpoint").
			Unique(),
	}
}
