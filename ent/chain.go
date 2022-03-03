// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/cosmos-gov/ent/chain"
)

// Chain is the model entity for the Chain schema.
type Chain struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// IsEnabled holds the value of the "is_enabled" field.
	IsEnabled bool `json:"is_enabled,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChainQuery when eager-loading is set.
	Edges ChainEdges `json:"edges"`
}

// ChainEdges holds the relations/edges for other nodes in the graph.
type ChainEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Proposals holds the value of the proposals edge.
	Proposals []*Proposal `json:"proposals,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e ChainEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// ProposalsOrErr returns the Proposals value or an error if the edge
// was not loaded in eager-loading.
func (e ChainEdges) ProposalsOrErr() ([]*Proposal, error) {
	if e.loadedTypes[1] {
		return e.Proposals, nil
	}
	return nil, &NotLoadedError{edge: "proposals"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Chain) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case chain.FieldIsEnabled:
			values[i] = new(sql.NullBool)
		case chain.FieldID:
			values[i] = new(sql.NullInt64)
		case chain.FieldName, chain.FieldDisplayName:
			values[i] = new(sql.NullString)
		case chain.FieldCreatedAt, chain.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Chain", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Chain fields.
func (c *Chain) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chain.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case chain.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case chain.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case chain.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case chain.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				c.DisplayName = value.String
			}
		case chain.FieldIsEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_enabled", values[i])
			} else if value.Valid {
				c.IsEnabled = value.Bool
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Chain entity.
func (c *Chain) QueryUsers() *UserQuery {
	return (&ChainClient{config: c.config}).QueryUsers(c)
}

// QueryProposals queries the "proposals" edge of the Chain entity.
func (c *Chain) QueryProposals() *ProposalQuery {
	return (&ChainClient{config: c.config}).QueryProposals(c)
}

// Update returns a builder for updating this Chain.
// Note that you need to call Chain.Unwrap() before calling this method if this Chain
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Chain) Update() *ChainUpdateOne {
	return (&ChainClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Chain entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Chain) Unwrap() *Chain {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Chain is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Chain) String() string {
	var builder strings.Builder
	builder.WriteString("Chain(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", display_name=")
	builder.WriteString(c.DisplayName)
	builder.WriteString(", is_enabled=")
	builder.WriteString(fmt.Sprintf("%v", c.IsEnabled))
	builder.WriteByte(')')
	return builder.String()
}

// Chains is a parsable slice of Chain.
type Chains []*Chain

func (c Chains) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
