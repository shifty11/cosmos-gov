// Code generated by entc, DO NOT EDIT.

package wallet

import (
	"time"
)

const (
	// Label holds the string label denoting the wallet type in the database.
	Label = "wallet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeChains holds the string denoting the chains edge name in mutations.
	EdgeChains = "chains"
	// Table holds the table name of the wallet in the database.
	Table = "wallets"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_wallets"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// ChainsTable is the table that holds the chains relation/edge.
	ChainsTable = "chains"
	// ChainsInverseTable is the table name for the Chain entity.
	// It exists in this package in order to avoid circular dependency with the "chain" package.
	ChainsInverseTable = "chains"
	// ChainsColumn is the table column denoting the chains relation/edge.
	ChainsColumn = "wallet_chains"
)

// Columns holds all SQL columns for wallet fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAddress,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "wallet_id"}
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
)
