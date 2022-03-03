package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// LensChainInfo holds the schema definition for the LensChainInfo entity.
type LensChainInfo struct {
	ent.Schema
}

// Fields of the LensChainInfo.
func (LensChainInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("name").
			Unique(),
		field.Int("cnt_errors"),
	}
}

// Edges of the LensChainInfo.
func (LensChainInfo) Edges() []ent.Edge {
	return nil
}

func (LensChainInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique(),
	}
}
