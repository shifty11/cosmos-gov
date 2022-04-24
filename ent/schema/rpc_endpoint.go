package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// RpcEndpoint holds the schema definition for the RpcEndpoint entity.
type RpcEndpoint struct {
	ent.Schema
}

// Fields of the RpcEndpoint.
func (RpcEndpoint) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
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
