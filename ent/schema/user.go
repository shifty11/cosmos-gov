package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id"),
		field.String("name"),
		//field.Int64("chat_id"). // TODO: has to be removed
		//			Immutable(),
		field.Enum("type").
			Values("telegram", "discord").
			Immutable(),
		field.String("login_token").
			Default(""),

		//TODO: remove this fields (or rename)
		//field.Int64("user_id").
		//	Default(0),
		//field.String("user_name").
		//	Default("<not set>"),
		//field.String("chat_name").
		//	Default("<not set>"),
		//field.Bool("is_group").
		//	Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("chains", Chain.Type), // TODO: has to be removed
		edge.From("telegram_chats", TelegramChat.Type).
			Ref("user").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.From("discord_channels", DiscordChannel.Type).
			Ref("user").
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("wallets", Wallet.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		//index.Fields("chat_id", "type"). // TODO: has to be removed
		//					Unique(),
		index.Fields("user_id", "type"). // TODO: disable for migration
							Unique(),
	}
}
