package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int64("chat_id"),
		field.Enum("type").
			Values("telegram", "discord"),

		field.Int64("user_id").
			Default(0),
		field.String("user_name").
			Default("<not set>"),
		field.String("chat_name").
			Default("<not set>"),
		field.Bool("is_group").
			Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chains", Chain.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("chat_id", "type").
			Unique(),
	}
}
