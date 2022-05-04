package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Chain holds the schema definition for the Chain entity.
type Chain struct {
	ent.Schema
}

func (Chain) Mixin() []ent.Mixin {
	return []ent.Mixin{
		//mixin.Time{},
	}
}

// Fields of the Chain.
func (Chain) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").Optional(),
		field.Time("update_time").Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),

		field.String("chain_id").
			Unique().
			Optional(), // TODO: remove optional
		field.String("account_prefix").
			Optional(), // TODO: remove optional
		field.String("name").
			Unique(),
		field.String("display_name").
			Unique(),
		field.Bool("is_enabled").
			Default(true),
	}
}

// Edges of the Chain.
func (Chain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type). // TODO: has to be removed
						Ref("chains"),
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
