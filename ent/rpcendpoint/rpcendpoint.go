// Code generated by entc, DO NOT EDIT.

package rpcendpoint

import (
	"time"
)

const (
	// Label holds the string label denoting the rpcendpoint type in the database.
	Label = "rpc_endpoint"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdatedTime holds the string denoting the updated_time field in the database.
	FieldUpdatedTime = "updated_time"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldEndpoint holds the string denoting the endpoint field in the database.
	FieldEndpoint = "endpoint"
	// EdgeChain holds the string denoting the chain edge name in mutations.
	EdgeChain = "chain"
	// Table holds the table name of the rpcendpoint in the database.
	Table = "rpc_endpoints"
	// ChainTable is the table that holds the chain relation/edge.
	ChainTable = "rpc_endpoints"
	// ChainInverseTable is the table name for the Chain entity.
	// It exists in this package in order to avoid circular dependency with the "chain" package.
	ChainInverseTable = "chains"
	// ChainColumn is the table column denoting the chain relation/edge.
	ChainColumn = "chain_rpc_endpoints"
)

// Columns holds all SQL columns for rpcendpoint fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdatedTime,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldEndpoint,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "rpc_endpoints"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"chain_rpc_endpoints",
}

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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
