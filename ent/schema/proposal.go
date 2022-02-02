package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Proposal holds the schema definition for the Proposal entity.
type Proposal struct {
	ent.Schema
}

// Fields of the Proposal.
func (Proposal) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int("proposal_id"),
		field.String("title"),
		field.String("description"),
		field.Time("voting_start_time"),
		field.Time("voting_end_time"),
		field.String("status"),
	}
}

// Edges of the Proposal.
func (Proposal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chain", Chain.Type).
			Ref("proposals").
			Unique(),
	}
}
