// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/cosmos-gov/ent/migrationinfo"
)

// MigrationInfo is the model entity for the MigrationInfo schema.
type MigrationInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// IsMigrated holds the value of the "is_migrated" field.
	IsMigrated bool `json:"is_migrated,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MigrationInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case migrationinfo.FieldIsMigrated:
			values[i] = new(sql.NullBool)
		case migrationinfo.FieldID:
			values[i] = new(sql.NullInt64)
		case migrationinfo.FieldCreatedAt, migrationinfo.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MigrationInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MigrationInfo fields.
func (mi *MigrationInfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case migrationinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mi.ID = int(value.Int64)
		case migrationinfo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mi.CreatedAt = value.Time
			}
		case migrationinfo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mi.UpdatedAt = value.Time
			}
		case migrationinfo.FieldIsMigrated:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_migrated", values[i])
			} else if value.Valid {
				mi.IsMigrated = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this MigrationInfo.
// Note that you need to call MigrationInfo.Unwrap() before calling this method if this MigrationInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (mi *MigrationInfo) Update() *MigrationInfoUpdateOne {
	return (&MigrationInfoClient{config: mi.config}).UpdateOne(mi)
}

// Unwrap unwraps the MigrationInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mi *MigrationInfo) Unwrap() *MigrationInfo {
	tx, ok := mi.config.driver.(*txDriver)
	if !ok {
		panic("ent: MigrationInfo is not a transactional entity")
	}
	mi.config.driver = tx.drv
	return mi
}

// String implements the fmt.Stringer.
func (mi *MigrationInfo) String() string {
	var builder strings.Builder
	builder.WriteString("MigrationInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", mi.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(mi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(mi.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", is_migrated=")
	builder.WriteString(fmt.Sprintf("%v", mi.IsMigrated))
	builder.WriteByte(')')
	return builder.String()
}

// MigrationInfos is a parsable slice of MigrationInfo.
type MigrationInfos []*MigrationInfo

func (mi MigrationInfos) config(cfg config) {
	for _i := range mi {
		mi[_i].config = cfg
	}
}
