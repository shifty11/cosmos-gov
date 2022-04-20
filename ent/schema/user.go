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
		field.Int64("id").
			Immutable(),
		field.String("name"),
		field.Int64("chat_id"). // TODO: has to be removed
					Immutable(),
		field.Enum("type").
			Values("telegram", "discord").
			Immutable(),
		field.String("loging_token").
			Default(""),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chains", Chain.Type), // TODO: has to be removed
		edge.From("telegram_chats", TelegramChat.Type).
			Ref("user"),
		edge.From("discord_channels", DiscordChannel.Type).
			Ref("user"),
		edge.To("wallets", Wallet.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("chat_id", "type"). // TODO: has to be removed
							Unique(),
	}
}
