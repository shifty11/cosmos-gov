// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldChatID holds the string denoting the chat_id field in the database.
	FieldChatID = "chat_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldLogingToken holds the string denoting the loging_token field in the database.
	FieldLogingToken = "loging_token"
	// EdgeChains holds the string denoting the chains edge name in mutations.
	EdgeChains = "chains"
	// EdgeTelegramChats holds the string denoting the telegram_chats edge name in mutations.
	EdgeTelegramChats = "telegram_chats"
	// EdgeDiscordChannels holds the string denoting the discord_channels edge name in mutations.
	EdgeDiscordChannels = "discord_channels"
	// EdgeWallets holds the string denoting the wallets edge name in mutations.
	EdgeWallets = "wallets"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ChainsTable is the table that holds the chains relation/edge. The primary key declared below.
	ChainsTable = "user_chains"
	// ChainsInverseTable is the table name for the Chain entity.
	// It exists in this package in order to avoid circular dependency with the "chain" package.
	ChainsInverseTable = "chains"
	// TelegramChatsTable is the table that holds the telegram_chats relation/edge.
	TelegramChatsTable = "telegram_chats"
	// TelegramChatsInverseTable is the table name for the TelegramChat entity.
	// It exists in this package in order to avoid circular dependency with the "telegramchat" package.
	TelegramChatsInverseTable = "telegram_chats"
	// TelegramChatsColumn is the table column denoting the telegram_chats relation/edge.
	TelegramChatsColumn = "telegram_chat_user"
	// DiscordChannelsTable is the table that holds the discord_channels relation/edge.
	DiscordChannelsTable = "discord_channels"
	// DiscordChannelsInverseTable is the table name for the DiscordChannel entity.
	// It exists in this package in order to avoid circular dependency with the "discordchannel" package.
	DiscordChannelsInverseTable = "discord_channels"
	// DiscordChannelsColumn is the table column denoting the discord_channels relation/edge.
	DiscordChannelsColumn = "discord_channel_user"
	// WalletsTable is the table that holds the wallets relation/edge. The primary key declared below.
	WalletsTable = "user_wallets"
	// WalletsInverseTable is the table name for the Wallet entity.
	// It exists in this package in order to avoid circular dependency with the "wallet" package.
	WalletsInverseTable = "wallets"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldChatID,
	FieldType,
	FieldLogingToken,
}

var (
	// ChainsPrimaryKey and ChainsColumn2 are the table columns denoting the
	// primary key for the chains relation (M2M).
	ChainsPrimaryKey = []string{"user_id", "chain_id"}
	// WalletsPrimaryKey and WalletsColumn2 are the table columns denoting the
	// primary key for the wallets relation (M2M).
	WalletsPrimaryKey = []string{"user_id", "wallet_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultLogingToken holds the default value on creation for the "loging_token" field.
	DefaultLogingToken string
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeTelegram Type = "telegram"
	TypeDiscord  Type = "discord"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeTelegram, TypeDiscord:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for type field: %q", _type)
	}
}
