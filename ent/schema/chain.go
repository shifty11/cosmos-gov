package schema

import (
	"entgo.io/ent"
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
		field.String("account_prefix").
			Unique(),
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
		//edge.From("users", User.Type). // TODO: has to be removed
		//				Ref("chains"),
		edge.To("proposals", Proposal.Type),
		edge.From("telegram_chats", TelegramChat.Type).
			Ref("chains"),
		edge.From("discord_channels", DiscordChannel.Type).
			Ref("chains"),
		edge.To("rpc_endpoints", RpcEndpoint.Type),
		edge.To("wallets", Wallet.Type),
	}
}

func (Chain) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique(),
	}
}
