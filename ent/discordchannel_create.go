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
	"github.com/shifty11/cosmos-gov/ent/discordchannel"
	"github.com/shifty11/cosmos-gov/ent/user"
)

// DiscordChannelCreate is the builder for creating a DiscordChannel entity.
type DiscordChannelCreate struct {
	config
	mutation *DiscordChannelMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (dcc *DiscordChannelCreate) SetCreateTime(t time.Time) *DiscordChannelCreate {
	dcc.mutation.SetCreateTime(t)
	return dcc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (dcc *DiscordChannelCreate) SetNillableCreateTime(t *time.Time) *DiscordChannelCreate {
	if t != nil {
		dcc.SetCreateTime(*t)
	}
	return dcc
}

// SetUpdateTime sets the "update_time" field.
func (dcc *DiscordChannelCreate) SetUpdateTime(t time.Time) *DiscordChannelCreate {
	dcc.mutation.SetUpdateTime(t)
	return dcc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (dcc *DiscordChannelCreate) SetNillableUpdateTime(t *time.Time) *DiscordChannelCreate {
	if t != nil {
		dcc.SetUpdateTime(*t)
	}
	return dcc
}

// SetChannelID sets the "channel_id" field.
func (dcc *DiscordChannelCreate) SetChannelID(i int64) *DiscordChannelCreate {
	dcc.mutation.SetChannelID(i)
	return dcc
}

// SetName sets the "name" field.
func (dcc *DiscordChannelCreate) SetName(s string) *DiscordChannelCreate {
	dcc.mutation.SetName(s)
	return dcc
}

// SetIsGroup sets the "is_group" field.
func (dcc *DiscordChannelCreate) SetIsGroup(b bool) *DiscordChannelCreate {
	dcc.mutation.SetIsGroup(b)
	return dcc
}

// SetRoles sets the "roles" field.
func (dcc *DiscordChannelCreate) SetRoles(s string) *DiscordChannelCreate {
	dcc.mutation.SetRoles(s)
	return dcc
}

// SetNillableRoles sets the "roles" field if the given value is not nil.
func (dcc *DiscordChannelCreate) SetNillableRoles(s *string) *DiscordChannelCreate {
	if s != nil {
		dcc.SetRoles(*s)
	}
	return dcc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (dcc *DiscordChannelCreate) SetUserID(id int) *DiscordChannelCreate {
	dcc.mutation.SetUserID(id)
	return dcc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (dcc *DiscordChannelCreate) SetNillableUserID(id *int) *DiscordChannelCreate {
	if id != nil {
		dcc = dcc.SetUserID(*id)
	}
	return dcc
}

// SetUser sets the "user" edge to the User entity.
func (dcc *DiscordChannelCreate) SetUser(u *User) *DiscordChannelCreate {
	return dcc.SetUserID(u.ID)
}

// AddChainIDs adds the "chains" edge to the Chain entity by IDs.
func (dcc *DiscordChannelCreate) AddChainIDs(ids ...int) *DiscordChannelCreate {
	dcc.mutation.AddChainIDs(ids...)
	return dcc
}

// AddChains adds the "chains" edges to the Chain entity.
func (dcc *DiscordChannelCreate) AddChains(c ...*Chain) *DiscordChannelCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dcc.AddChainIDs(ids...)
}

// Mutation returns the DiscordChannelMutation object of the builder.
func (dcc *DiscordChannelCreate) Mutation() *DiscordChannelMutation {
	return dcc.mutation
}

// Save creates the DiscordChannel in the database.
func (dcc *DiscordChannelCreate) Save(ctx context.Context) (*DiscordChannel, error) {
	var (
		err  error
		node *DiscordChannel
	)
	dcc.defaults()
	if len(dcc.hooks) == 0 {
		if err = dcc.check(); err != nil {
			return nil, err
		}
		node, err = dcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiscordChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dcc.check(); err != nil {
				return nil, err
			}
			dcc.mutation = mutation
			if node, err = dcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dcc.hooks) - 1; i >= 0; i-- {
			if dcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dcc *DiscordChannelCreate) SaveX(ctx context.Context) *DiscordChannel {
	v, err := dcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcc *DiscordChannelCreate) Exec(ctx context.Context) error {
	_, err := dcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcc *DiscordChannelCreate) ExecX(ctx context.Context) {
	if err := dcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dcc *DiscordChannelCreate) defaults() {
	if _, ok := dcc.mutation.CreateTime(); !ok {
		v := discordchannel.DefaultCreateTime()
		dcc.mutation.SetCreateTime(v)
	}
	if _, ok := dcc.mutation.UpdateTime(); !ok {
		v := discordchannel.DefaultUpdateTime()
		dcc.mutation.SetUpdateTime(v)
	}
	if _, ok := dcc.mutation.Roles(); !ok {
		v := discordchannel.DefaultRoles
		dcc.mutation.SetRoles(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dcc *DiscordChannelCreate) check() error {
	if _, ok := dcc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "DiscordChannel.create_time"`)}
	}
	if _, ok := dcc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "DiscordChannel.update_time"`)}
	}
	if _, ok := dcc.mutation.ChannelID(); !ok {
		return &ValidationError{Name: "channel_id", err: errors.New(`ent: missing required field "DiscordChannel.channel_id"`)}
	}
	if _, ok := dcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "DiscordChannel.name"`)}
	}
	if _, ok := dcc.mutation.IsGroup(); !ok {
		return &ValidationError{Name: "is_group", err: errors.New(`ent: missing required field "DiscordChannel.is_group"`)}
	}
	if _, ok := dcc.mutation.Roles(); !ok {
		return &ValidationError{Name: "roles", err: errors.New(`ent: missing required field "DiscordChannel.roles"`)}
	}
	return nil
}

func (dcc *DiscordChannelCreate) sqlSave(ctx context.Context) (*DiscordChannel, error) {
	_node, _spec := dcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dcc *DiscordChannelCreate) createSpec() (*DiscordChannel, *sqlgraph.CreateSpec) {
	var (
		_node = &DiscordChannel{config: dcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: discordchannel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: discordchannel.FieldID,
			},
		}
	)
	if value, ok := dcc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: discordchannel.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := dcc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: discordchannel.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := dcc.mutation.ChannelID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: discordchannel.FieldChannelID,
		})
		_node.ChannelID = value
	}
	if value, ok := dcc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discordchannel.FieldName,
		})
		_node.Name = value
	}
	if value, ok := dcc.mutation.IsGroup(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: discordchannel.FieldIsGroup,
		})
		_node.IsGroup = value
	}
	if value, ok := dcc.mutation.Roles(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discordchannel.FieldRoles,
		})
		_node.Roles = value
	}
	if nodes := dcc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   discordchannel.UserTable,
			Columns: []string{discordchannel.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.discord_channel_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dcc.mutation.ChainsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   discordchannel.ChainsTable,
			Columns: discordchannel.ChainsPrimaryKey,
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

// DiscordChannelCreateBulk is the builder for creating many DiscordChannel entities in bulk.
type DiscordChannelCreateBulk struct {
	config
	builders []*DiscordChannelCreate
}

// Save creates the DiscordChannel entities in the database.
func (dccb *DiscordChannelCreateBulk) Save(ctx context.Context) ([]*DiscordChannel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dccb.builders))
	nodes := make([]*DiscordChannel, len(dccb.builders))
	mutators := make([]Mutator, len(dccb.builders))
	for i := range dccb.builders {
		func(i int, root context.Context) {
			builder := dccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DiscordChannelMutation)
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
					_, err = mutators[i+1].Mutate(root, dccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dccb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, dccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dccb *DiscordChannelCreateBulk) SaveX(ctx context.Context) []*DiscordChannel {
	v, err := dccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dccb *DiscordChannelCreateBulk) Exec(ctx context.Context) error {
	_, err := dccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dccb *DiscordChannelCreateBulk) ExecX(ctx context.Context) {
	if err := dccb.Exec(ctx); err != nil {
		panic(err)
	}
}
