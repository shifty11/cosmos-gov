package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// TelegramChat holds the schema definition for the TelegramChat entity.
type TelegramChat struct {
	ent.Schema
}

// Fields of the TelegramChat.
func (TelegramChat) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int64("id").
			Unique().
			Immutable(),
		field.String("name"),
		field.Bool("is_group").
			Immutable(),
	}
}

// Edges of the TelegramChat.
func (TelegramChat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique(),
		edge.To("chains", Chain.Type),
	}
}

func (TelegramChat) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").
			Unique(),
	}
}
