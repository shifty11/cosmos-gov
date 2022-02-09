// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/proposal"
)

// Proposal is the model entity for the Proposal schema.
type Proposal struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ProposalID holds the value of the "proposal_id" field.
	ProposalID uint64 `json:"proposal_id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// VotingStartTime holds the value of the "voting_start_time" field.
	VotingStartTime time.Time `json:"voting_start_time,omitempty"`
	// VotingEndTime holds the value of the "voting_end_time" field.
	VotingEndTime time.Time `json:"voting_end_time,omitempty"`
	// Status holds the value of the "status" field.
	Status proposal.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProposalQuery when eager-loading is set.
	Edges           ProposalEdges `json:"edges"`
	chain_proposals *int
}

// ProposalEdges holds the relations/edges for other nodes in the graph.
type ProposalEdges struct {
	// Chain holds the value of the chain edge.
	Chain *Chain `json:"chain,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ChainOrErr returns the Chain value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProposalEdges) ChainOrErr() (*Chain, error) {
	if e.loadedTypes[0] {
		if e.Chain == nil {
			// The edge chain was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: chain.Label}
		}
		return e.Chain, nil
	}
	return nil, &NotLoadedError{edge: "chain"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Proposal) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case proposal.FieldID, proposal.FieldProposalID:
			values[i] = new(sql.NullInt64)
		case proposal.FieldTitle, proposal.FieldDescription, proposal.FieldStatus:
			values[i] = new(sql.NullString)
		case proposal.FieldCreatedAt, proposal.FieldUpdatedAt, proposal.FieldVotingStartTime, proposal.FieldVotingEndTime:
			values[i] = new(sql.NullTime)
		case proposal.ForeignKeys[0]: // chain_proposals
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Proposal", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Proposal fields.
func (pr *Proposal) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case proposal.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case proposal.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case proposal.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case proposal.FieldProposalID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field proposal_id", values[i])
			} else if value.Valid {
				pr.ProposalID = uint64(value.Int64)
			}
		case proposal.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pr.Title = value.String
			}
		case proposal.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case proposal.FieldVotingStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field voting_start_time", values[i])
			} else if value.Valid {
				pr.VotingStartTime = value.Time
			}
		case proposal.FieldVotingEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field voting_end_time", values[i])
			} else if value.Valid {
				pr.VotingEndTime = value.Time
			}
		case proposal.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pr.Status = proposal.Status(value.String)
			}
		case proposal.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field chain_proposals", value)
			} else if value.Valid {
				pr.chain_proposals = new(int)
				*pr.chain_proposals = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryChain queries the "chain" edge of the Proposal entity.
func (pr *Proposal) QueryChain() *ChainQuery {
	return (&ProposalClient{config: pr.config}).QueryChain(pr)
}

// Update returns a builder for updating this Proposal.
// Note that you need to call Proposal.Unwrap() before calling this method if this Proposal
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Proposal) Update() *ProposalUpdateOne {
	return (&ProposalClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Proposal entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Proposal) Unwrap() *Proposal {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Proposal is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Proposal) String() string {
	var builder strings.Builder
	builder.WriteString("Proposal(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", proposal_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.ProposalID))
	builder.WriteString(", title=")
	builder.WriteString(pr.Title)
	builder.WriteString(", description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", voting_start_time=")
	builder.WriteString(pr.VotingStartTime.Format(time.ANSIC))
	builder.WriteString(", voting_end_time=")
	builder.WriteString(pr.VotingEndTime.Format(time.ANSIC))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", pr.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Proposals is a parsable slice of Proposal.
type Proposals []*Proposal

func (pr Proposals) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}