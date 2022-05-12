package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Chain holds the schema definition for the Chain entity.
type Chain struct {
	ent.Schema
}

func (Chain) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Chain.
func (Chain) Fields() []ent.Field {
	return []ent.Field{
		field.String("chain_id").
			Unique(),
		field.String("account_prefix"),
		field.String("name").
			Unique(),
		field.String("display_name").
			Unique(),
		field.Bool("is_enabled").
			Default(false),
		field.Bool("is_voting_enabled").
			Default(false),
		field.Bool("is_feegrant_used").
			Default(false),
	}
}

// Edges of the Chain.
func (Chain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("proposals", Proposal.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.From("telegram_chats", TelegramChat.Type).
			Ref("chains").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.From("discord_channels", DiscordChannel.Type).
			Ref("chains").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("rpc_endpoints", RpcEndpoint.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("wallets", Wallet.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

func (Chain) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique(),
	}
}
