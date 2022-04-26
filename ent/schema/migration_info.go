package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// MigrationInfo holds the schema definition for the MigrationInfo entity.
type MigrationInfo struct {
	ent.Schema
}

func (MigrationInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the MigrationInfo.
func (MigrationInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_migrated").
			Default(false),
	}
}

// Edges of the MigrationInfo.
func (MigrationInfo) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (MigrationInfo) Indexes() []ent.Index {
	return []ent.Index{}
}
