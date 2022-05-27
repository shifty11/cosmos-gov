// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-gov/ent/draftproposal"
	"github.com/shifty11/cosmos-gov/ent/predicate"
)

// DraftProposalDelete is the builder for deleting a DraftProposal entity.
type DraftProposalDelete struct {
	config
	hooks    []Hook
	mutation *DraftProposalMutation
}

// Where appends a list predicates to the DraftProposalDelete builder.
func (dpd *DraftProposalDelete) Where(ps ...predicate.DraftProposal) *DraftProposalDelete {
	dpd.mutation.Where(ps...)
	return dpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dpd *DraftProposalDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dpd.hooks) == 0 {
		affected, err = dpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DraftProposalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dpd.mutation = mutation
			affected, err = dpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dpd.hooks) - 1; i >= 0; i-- {
			if dpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpd *DraftProposalDelete) ExecX(ctx context.Context) int {
	n, err := dpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dpd *DraftProposalDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: draftproposal.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: draftproposal.FieldID,
			},
		},
	}
	if ps := dpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, dpd.driver, _spec)
}

// DraftProposalDeleteOne is the builder for deleting a single DraftProposal entity.
type DraftProposalDeleteOne struct {
	dpd *DraftProposalDelete
}

// Exec executes the deletion query.
func (dpdo *DraftProposalDeleteOne) Exec(ctx context.Context) error {
	n, err := dpdo.dpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{draftproposal.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dpdo *DraftProposalDeleteOne) ExecX(ctx context.Context) {
	dpdo.dpd.ExecX(ctx)
}
