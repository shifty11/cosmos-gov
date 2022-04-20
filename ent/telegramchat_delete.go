// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-gov/ent/predicate"
	"github.com/shifty11/cosmos-gov/ent/telegramchat"
)

// TelegramChatDelete is the builder for deleting a TelegramChat entity.
type TelegramChatDelete struct {
	config
	hooks    []Hook
	mutation *TelegramChatMutation
}

// Where appends a list predicates to the TelegramChatDelete builder.
func (tcd *TelegramChatDelete) Where(ps ...predicate.TelegramChat) *TelegramChatDelete {
	tcd.mutation.Where(ps...)
	return tcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tcd *TelegramChatDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tcd.hooks) == 0 {
		affected, err = tcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TelegramChatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tcd.mutation = mutation
			affected, err = tcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tcd.hooks) - 1; i >= 0; i-- {
			if tcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcd *TelegramChatDelete) ExecX(ctx context.Context) int {
	n, err := tcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tcd *TelegramChatDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: telegramchat.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: telegramchat.FieldID,
			},
		},
	}
	if ps := tcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tcd.driver, _spec)
}

// TelegramChatDeleteOne is the builder for deleting a single TelegramChat entity.
type TelegramChatDeleteOne struct {
	tcd *TelegramChatDelete
}

// Exec executes the deletion query.
func (tcdo *TelegramChatDeleteOne) Exec(ctx context.Context) error {
	n, err := tcdo.tcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{telegramchat.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tcdo *TelegramChatDeleteOne) ExecX(ctx context.Context) {
	tcdo.tcd.ExecX(ctx)
}
