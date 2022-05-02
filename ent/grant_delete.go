// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-gov/ent/grant"
	"github.com/shifty11/cosmos-gov/ent/predicate"
)

// GrantDelete is the builder for deleting a Grant entity.
type GrantDelete struct {
	config
	hooks    []Hook
	mutation *GrantMutation
}

// Where appends a list predicates to the GrantDelete builder.
func (gd *GrantDelete) Where(ps ...predicate.Grant) *GrantDelete {
	gd.mutation.Where(ps...)
	return gd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gd *GrantDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gd.hooks) == 0 {
		affected, err = gd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GrantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gd.mutation = mutation
			affected, err = gd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gd.hooks) - 1; i >= 0; i-- {
			if gd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gd *GrantDelete) ExecX(ctx context.Context) int {
	n, err := gd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gd *GrantDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: grant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: grant.FieldID,
			},
		},
	}
	if ps := gd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, gd.driver, _spec)
}

// GrantDeleteOne is the builder for deleting a single Grant entity.
type GrantDeleteOne struct {
	gd *GrantDelete
}

// Exec executes the deletion query.
func (gdo *GrantDeleteOne) Exec(ctx context.Context) error {
	n, err := gdo.gd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{grant.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gdo *GrantDeleteOne) ExecX(ctx context.Context) {
	gdo.gd.ExecX(ctx)
}