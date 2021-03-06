// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/wallet"
)

// Wallet is the model entity for the Wallet schema.
type Wallet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WalletQuery when eager-loading is set.
	Edges         WalletEdges `json:"edges"`
	chain_wallets *int
}

// WalletEdges holds the relations/edges for other nodes in the graph.
type WalletEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Chain holds the value of the chain edge.
	Chain *Chain `json:"chain,omitempty"`
	// Grants holds the value of the grants edge.
	Grants []*Grant `json:"grants,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e WalletEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// ChainOrErr returns the Chain value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WalletEdges) ChainOrErr() (*Chain, error) {
	if e.loadedTypes[1] {
		if e.Chain == nil {
			// The edge chain was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: chain.Label}
		}
		return e.Chain, nil
	}
	return nil, &NotLoadedError{edge: "chain"}
}

// GrantsOrErr returns the Grants value or an error if the edge
// was not loaded in eager-loading.
func (e WalletEdges) GrantsOrErr() ([]*Grant, error) {
	if e.loadedTypes[2] {
		return e.Grants, nil
	}
	return nil, &NotLoadedError{edge: "grants"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Wallet) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case wallet.FieldID:
			values[i] = new(sql.NullInt64)
		case wallet.FieldAddress:
			values[i] = new(sql.NullString)
		case wallet.FieldCreateTime, wallet.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case wallet.ForeignKeys[0]: // chain_wallets
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Wallet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Wallet fields.
func (w *Wallet) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wallet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case wallet.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				w.CreateTime = value.Time
			}
		case wallet.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				w.UpdateTime = value.Time
			}
		case wallet.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				w.Address = value.String
			}
		case wallet.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field chain_wallets", value)
			} else if value.Valid {
				w.chain_wallets = new(int)
				*w.chain_wallets = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Wallet entity.
func (w *Wallet) QueryUsers() *UserQuery {
	return (&WalletClient{config: w.config}).QueryUsers(w)
}

// QueryChain queries the "chain" edge of the Wallet entity.
func (w *Wallet) QueryChain() *ChainQuery {
	return (&WalletClient{config: w.config}).QueryChain(w)
}

// QueryGrants queries the "grants" edge of the Wallet entity.
func (w *Wallet) QueryGrants() *GrantQuery {
	return (&WalletClient{config: w.config}).QueryGrants(w)
}

// Update returns a builder for updating this Wallet.
// Note that you need to call Wallet.Unwrap() before calling this method if this Wallet
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Wallet) Update() *WalletUpdateOne {
	return (&WalletClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the Wallet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Wallet) Unwrap() *Wallet {
	tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Wallet is not a transactional entity")
	}
	w.config.driver = tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Wallet) String() string {
	var builder strings.Builder
	builder.WriteString("Wallet(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(w.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(w.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", address=")
	builder.WriteString(w.Address)
	builder.WriteByte(')')
	return builder.String()
}

// Wallets is a parsable slice of Wallet.
type Wallets []*Wallet

func (w Wallets) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}
