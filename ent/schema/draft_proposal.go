package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// DraftProposal holds the schema definition for the DraftProposal entity.
type DraftProposal struct {
	ent.Schema
}

func (DraftProposal) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the DraftProposal.
func (DraftProposal) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("draft_proposal_id"),
		field.String("title"),
		field.String("url"),
	}
}

// Edges of the DraftProposal.
func (DraftProposal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chain", Chain.Type).
			Ref("draft_proposals").
			Unique(),
	}
}

func (DraftProposal) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("draft_proposal_id").
			Edges("chain").
			Unique(),
	}
}
