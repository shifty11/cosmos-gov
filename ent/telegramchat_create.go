// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/telegramchat"
	"github.com/shifty11/cosmos-gov/ent/user"
)

// TelegramChatCreate is the builder for creating a TelegramChat entity.
type TelegramChatCreate struct {
	config
	mutation *TelegramChatMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tcc *TelegramChatCreate) SetCreatedAt(t time.Time) *TelegramChatCreate {
	tcc.mutation.SetCreatedAt(t)
	return tcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tcc *TelegramChatCreate) SetNillableCreatedAt(t *time.Time) *TelegramChatCreate {
	if t != nil {
		tcc.SetCreatedAt(*t)
	}
	return tcc
}

// SetUpdatedAt sets the "updated_at" field.
func (tcc *TelegramChatCreate) SetUpdatedAt(t time.Time) *TelegramChatCreate {
	tcc.mutation.SetUpdatedAt(t)
	return tcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tcc *TelegramChatCreate) SetNillableUpdatedAt(t *time.Time) *TelegramChatCreate {
	if t != nil {
		tcc.SetUpdatedAt(*t)
	}
	return tcc
}

// SetName sets the "name" field.
func (tcc *TelegramChatCreate) SetName(s string) *TelegramChatCreate {
	tcc.mutation.SetName(s)
	return tcc
}

// SetIsGroup sets the "is_group" field.
func (tcc *TelegramChatCreate) SetIsGroup(b bool) *TelegramChatCreate {
	tcc.mutation.SetIsGroup(b)
	return tcc
}

// SetID sets the "id" field.
func (tcc *TelegramChatCreate) SetID(i int64) *TelegramChatCreate {
	tcc.mutation.SetID(i)
	return tcc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tcc *TelegramChatCreate) SetUserID(id int64) *TelegramChatCreate {
	tcc.mutation.SetUserID(id)
	return tcc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (tcc *TelegramChatCreate) SetNillableUserID(id *int64) *TelegramChatCreate {
	if id != nil {
		tcc = tcc.SetUserID(*id)
	}
	return tcc
}

// SetUser sets the "user" edge to the User entity.
func (tcc *TelegramChatCreate) SetUser(u *User) *TelegramChatCreate {
	return tcc.SetUserID(u.ID)
}

// AddChainIDs adds the "chains" edge to the Chain entity by IDs.
func (tcc *TelegramChatCreate) AddChainIDs(ids ...int) *TelegramChatCreate {
	tcc.mutation.AddChainIDs(ids...)
	return tcc
}

// AddChains adds the "chains" edges to the Chain entity.
func (tcc *TelegramChatCreate) AddChains(c ...*Chain) *TelegramChatCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tcc.AddChainIDs(ids...)
}

// Mutation returns the TelegramChatMutation object of the builder.
func (tcc *TelegramChatCreate) Mutation() *TelegramChatMutation {
	return tcc.mutation
}

// Save creates the TelegramChat in the database.
func (tcc *TelegramChatCreate) Save(ctx context.Context) (*TelegramChat, error) {
	var (
		err  error
		node *TelegramChat
	)
	tcc.defaults()
	if len(tcc.hooks) == 0 {
		if err = tcc.check(); err != nil {
			return nil, err
		}
		node, err = tcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TelegramChatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tcc.check(); err != nil {
				return nil, err
			}
			tcc.mutation = mutation
			if node, err = tcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tcc.hooks) - 1; i >= 0; i-- {
			if tcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tcc *TelegramChatCreate) SaveX(ctx context.Context) *TelegramChat {
	v, err := tcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcc *TelegramChatCreate) Exec(ctx context.Context) error {
	_, err := tcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcc *TelegramChatCreate) ExecX(ctx context.Context) {
	if err := tcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcc *TelegramChatCreate) defaults() {
	if _, ok := tcc.mutation.CreatedAt(); !ok {
		v := telegramchat.DefaultCreatedAt()
		tcc.mutation.SetCreatedAt(v)
	}
	if _, ok := tcc.mutation.UpdatedAt(); !ok {
		v := telegramchat.DefaultUpdatedAt()
		tcc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tcc *TelegramChatCreate) check() error {
	if _, ok := tcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TelegramChat.created_at"`)}
	}
	if _, ok := tcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TelegramChat.updated_at"`)}
	}
	if _, ok := tcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "TelegramChat.name"`)}
	}
	if _, ok := tcc.mutation.IsGroup(); !ok {
		return &ValidationError{Name: "is_group", err: errors.New(`ent: missing required field "TelegramChat.is_group"`)}
	}
	return nil
}

func (tcc *TelegramChatCreate) sqlSave(ctx context.Context) (*TelegramChat, error) {
	_node, _spec := tcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (tcc *TelegramChatCreate) createSpec() (*TelegramChat, *sqlgraph.CreateSpec) {
	var (
		_node = &TelegramChat{config: tcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: telegramchat.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: telegramchat.FieldID,
			},
		}
	)
	if id, ok := tcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tcc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: telegramchat.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tcc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: telegramchat.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := tcc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telegramchat.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tcc.mutation.IsGroup(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: telegramchat.FieldIsGroup,
		})
		_node.IsGroup = value
	}
	if nodes := tcc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   telegramchat.UserTable,
			Columns: []string{telegramchat.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.telegram_chat_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tcc.mutation.ChainsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   telegramchat.ChainsTable,
			Columns: telegramchat.ChainsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: chain.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TelegramChatCreateBulk is the builder for creating many TelegramChat entities in bulk.
type TelegramChatCreateBulk struct {
	config
	builders []*TelegramChatCreate
}

// Save creates the TelegramChat entities in the database.
func (tccb *TelegramChatCreateBulk) Save(ctx context.Context) ([]*TelegramChat, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tccb.builders))
	nodes := make([]*TelegramChat, len(tccb.builders))
	mutators := make([]Mutator, len(tccb.builders))
	for i := range tccb.builders {
		func(i int, root context.Context) {
			builder := tccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TelegramChatMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tccb *TelegramChatCreateBulk) SaveX(ctx context.Context) []*TelegramChat {
	v, err := tccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tccb *TelegramChatCreateBulk) Exec(ctx context.Context) error {
	_, err := tccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tccb *TelegramChatCreateBulk) ExecX(ctx context.Context) {
	if err := tccb.Exec(ctx); err != nil {
		panic(err)
	}
}