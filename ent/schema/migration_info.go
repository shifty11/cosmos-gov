package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// MigrationInfo holds the schema definition for the MigrationInfo entity.
type MigrationInfo struct {
	ent.Schema
}

// Fields of the MigrationInfo.
func (MigrationInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
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
