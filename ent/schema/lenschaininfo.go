package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// LensChainInfo holds the schema definition for the LensChainInfo entity.
type LensChainInfo struct {
	ent.Schema
}

func (LensChainInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the LensChainInfo.
func (LensChainInfo) Fields() []ent.Field {
	return []ent.Field{
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
