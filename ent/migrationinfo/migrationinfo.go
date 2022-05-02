// Code generated by entc, DO NOT EDIT.

package migrationinfo

import (
	"time"
)

const (
	// Label holds the string label denoting the migrationinfo type in the database.
	Label = "migration_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldIsMigrated holds the string denoting the is_migrated field in the database.
	FieldIsMigrated = "is_migrated"
	// Table holds the table name of the migrationinfo in the database.
	Table = "migration_infos"
)

// Columns holds all SQL columns for migrationinfo fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldIsMigrated,
}

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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultIsMigrated holds the default value on creation for the "is_migrated" field.
	DefaultIsMigrated bool
)