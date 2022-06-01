// Code generated by entc, DO NOT EDIT.

package telegramchat

import (
	"time"
)

const (
	// Label holds the string label denoting the telegramchat type in the database.
	Label = "telegram_chat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldChatID holds the string denoting the chat_id field in the database.
	FieldChatID = "chat_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIsGroup holds the string denoting the is_group field in the database.
	FieldIsGroup = "is_group"
	// FieldWantsDraftProposals holds the string denoting the wants_draft_proposals field in the database.
	FieldWantsDraftProposals = "wants_draft_proposals"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeChains holds the string denoting the chains edge name in mutations.
	EdgeChains = "chains"
	// Table holds the table name of the telegramchat in the database.
	Table = "telegram_chats"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "telegram_chats"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "telegram_chat_user"
	// ChainsTable is the table that holds the chains relation/edge. The primary key declared below.
	ChainsTable = "telegram_chat_chains"
	// ChainsInverseTable is the table name for the Chain entity.
	// It exists in this package in order to avoid circular dependency with the "chain" package.
	ChainsInverseTable = "chains"
)

// Columns holds all SQL columns for telegramchat fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldChatID,
	FieldName,
	FieldIsGroup,
	FieldWantsDraftProposals,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "telegram_chats"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"telegram_chat_user",
}

var (
	// ChainsPrimaryKey and ChainsColumn2 are the table columns denoting the
	// primary key for the chains relation (M2M).
	ChainsPrimaryKey = []string{"telegram_chat_id", "chain_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultWantsDraftProposals holds the default value on creation for the "wants_draft_proposals" field.
	DefaultWantsDraftProposals bool
)
