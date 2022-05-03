package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"time"
)

// Proposal holds the schema definition for the Proposal entity.
type Proposal struct {
	ent.Schema
}

func (Proposal) Mixin() []ent.Mixin {
	return []ent.Mixin{
		//mixin.Time{},
	}
}

// Fields of the Proposal.
func (Proposal) Fields() []ent.Field {
	var statusValues []string
	for _, status := range types.ProposalStatus_name {
		statusValues = append(statusValues, status)
	}
	return []ent.Field{
		field.Time("create_time").Optional(),
		field.Time("updated_time").Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),

		field.Uint64("proposal_id"),
		field.String("title"),
		field.String("description"),
		field.Time("voting_start_time"),
		field.Time("voting_end_time"),
		field.Enum("status").
			Values(statusValues...),
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

func (Proposal) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("proposal_id").
			Edges("chain").
			Unique(),
	}
}
