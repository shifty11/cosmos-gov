// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/discordchannel"
	"github.com/shifty11/cosmos-gov/ent/predicate"
	"github.com/shifty11/cosmos-gov/ent/user"
)

// DiscordChannelUpdate is the builder for updating DiscordChannel entities.
type DiscordChannelUpdate struct {
	config
	hooks    []Hook
	mutation *DiscordChannelMutation
}

// Where appends a list predicates to the DiscordChannelUpdate builder.
func (dcu *DiscordChannelUpdate) Where(ps ...predicate.DiscordChannel) *DiscordChannelUpdate {
	dcu.mutation.Where(ps...)
	return dcu
}

// SetUpdateTime sets the "update_time" field.
func (dcu *DiscordChannelUpdate) SetUpdateTime(t time.Time) *DiscordChannelUpdate {
	dcu.mutation.SetUpdateTime(t)
	return dcu
}

// SetName sets the "name" field.
func (dcu *DiscordChannelUpdate) SetName(s string) *DiscordChannelUpdate {
	dcu.mutation.SetName(s)
	return dcu
}

// SetRoles sets the "roles" field.
func (dcu *DiscordChannelUpdate) SetRoles(s string) *DiscordChannelUpdate {
	dcu.mutation.SetRoles(s)
	return dcu
}

// SetNillableRoles sets the "roles" field if the given value is not nil.
func (dcu *DiscordChannelUpdate) SetNillableRoles(s *string) *DiscordChannelUpdate {
	if s != nil {
		dcu.SetRoles(*s)
	}
	return dcu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (dcu *DiscordChannelUpdate) SetUserID(id int64) *DiscordChannelUpdate {
	dcu.mutation.SetUserID(id)
	return dcu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (dcu *DiscordChannelUpdate) SetNillableUserID(id *int64) *DiscordChannelUpdate {
	if id != nil {
		dcu = dcu.SetUserID(*id)
	}
	return dcu
}

// SetUser sets the "user" edge to the User entity.
func (dcu *DiscordChannelUpdate) SetUser(u *User) *DiscordChannelUpdate {
	return dcu.SetUserID(u.ID)
}

// AddChainIDs adds the "chains" edge to the Chain entity by IDs.
func (dcu *DiscordChannelUpdate) AddChainIDs(ids ...int) *DiscordChannelUpdate {
	dcu.mutation.AddChainIDs(ids...)
	return dcu
}

// AddChains adds the "chains" edges to the Chain entity.
func (dcu *DiscordChannelUpdate) AddChains(c ...*Chain) *DiscordChannelUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dcu.AddChainIDs(ids...)
}

// Mutation returns the DiscordChannelMutation object of the builder.
func (dcu *DiscordChannelUpdate) Mutation() *DiscordChannelMutation {
	return dcu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (dcu *DiscordChannelUpdate) ClearUser() *DiscordChannelUpdate {
	dcu.mutation.ClearUser()
	return dcu
}

// ClearChains clears all "chains" edges to the Chain entity.
func (dcu *DiscordChannelUpdate) ClearChains() *DiscordChannelUpdate {
	dcu.mutation.ClearChains()
	return dcu
}

// RemoveChainIDs removes the "chains" edge to Chain entities by IDs.
func (dcu *DiscordChannelUpdate) RemoveChainIDs(ids ...int) *DiscordChannelUpdate {
	dcu.mutation.RemoveChainIDs(ids...)
	return dcu
}

// RemoveChains removes "chains" edges to Chain entities.
func (dcu *DiscordChannelUpdate) RemoveChains(c ...*Chain) *DiscordChannelUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dcu.RemoveChainIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dcu *DiscordChannelUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	dcu.defaults()
	if len(dcu.hooks) == 0 {
		affected, err = dcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiscordChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dcu.mutation = mutation
			affected, err = dcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dcu.hooks) - 1; i >= 0; i-- {
			if dcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dcu *DiscordChannelUpdate) SaveX(ctx context.Context) int {
	affected, err := dcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dcu *DiscordChannelUpdate) Exec(ctx context.Context) error {
	_, err := dcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcu *DiscordChannelUpdate) ExecX(ctx context.Context) {
	if err := dcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dcu *DiscordChannelUpdate) defaults() {
	if _, ok := dcu.mutation.UpdateTime(); !ok {
		v := discordchannel.UpdateDefaultUpdateTime()
		dcu.mutation.SetUpdateTime(v)
	}
}

func (dcu *DiscordChannelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   discordchannel.Table,
			Columns: discordchannel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: discordchannel.FieldID,
			},
		},
	}
	if ps := dcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dcu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: discordchannel.FieldUpdateTime,
		})
	}
	if value, ok := dcu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discordchannel.FieldName,
		})
	}
	if value, ok := dcu.mutation.Roles(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discordchannel.FieldRoles,
		})
	}
	if dcu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   discordchannel.UserTable,
			Columns: []string{discordchannel.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dcu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   discordchannel.UserTable,
			Columns: []string{discordchannel.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dcu.mutation.ChainsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dcu.mutation.RemovedChainsIDs(); len(nodes) > 0 && !dcu.mutation.ChainsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dcu.mutation.ChainsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{discordchannel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DiscordChannelUpdateOne is the builder for updating a single DiscordChannel entity.
type DiscordChannelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DiscordChannelMutation
}

// SetUpdateTime sets the "update_time" field.
func (dcuo *DiscordChannelUpdateOne) SetUpdateTime(t time.Time) *DiscordChannelUpdateOne {
	dcuo.mutation.SetUpdateTime(t)
	return dcuo
}

// SetName sets the "name" field.
func (dcuo *DiscordChannelUpdateOne) SetName(s string) *DiscordChannelUpdateOne {
	dcuo.mutation.SetName(s)
	return dcuo
}

// SetRoles sets the "roles" field.
func (dcuo *DiscordChannelUpdateOne) SetRoles(s string) *DiscordChannelUpdateOne {
	dcuo.mutation.SetRoles(s)
	return dcuo
}

// SetNillableRoles sets the "roles" field if the given value is not nil.
func (dcuo *DiscordChannelUpdateOne) SetNillableRoles(s *string) *DiscordChannelUpdateOne {
	if s != nil {
		dcuo.SetRoles(*s)
	}
	return dcuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (dcuo *DiscordChannelUpdateOne) SetUserID(id int64) *DiscordChannelUpdateOne {
	dcuo.mutation.SetUserID(id)
	return dcuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (dcuo *DiscordChannelUpdateOne) SetNillableUserID(id *int64) *DiscordChannelUpdateOne {
	if id != nil {
		dcuo = dcuo.SetUserID(*id)
	}
	return dcuo
}

// SetUser sets the "user" edge to the User entity.
func (dcuo *DiscordChannelUpdateOne) SetUser(u *User) *DiscordChannelUpdateOne {
	return dcuo.SetUserID(u.ID)
}

// AddChainIDs adds the "chains" edge to the Chain entity by IDs.
func (dcuo *DiscordChannelUpdateOne) AddChainIDs(ids ...int) *DiscordChannelUpdateOne {
	dcuo.mutation.AddChainIDs(ids...)
	return dcuo
}

// AddChains adds the "chains" edges to the Chain entity.
func (dcuo *DiscordChannelUpdateOne) AddChains(c ...*Chain) *DiscordChannelUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dcuo.AddChainIDs(ids...)
}

// Mutation returns the DiscordChannelMutation object of the builder.
func (dcuo *DiscordChannelUpdateOne) Mutation() *DiscordChannelMutation {
	return dcuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (dcuo *DiscordChannelUpdateOne) ClearUser() *DiscordChannelUpdateOne {
	dcuo.mutation.ClearUser()
	return dcuo
}

// ClearChains clears all "chains" edges to the Chain entity.
func (dcuo *DiscordChannelUpdateOne) ClearChains() *DiscordChannelUpdateOne {
	dcuo.mutation.ClearChains()
	return dcuo
}

// RemoveChainIDs removes the "chains" edge to Chain entities by IDs.
func (dcuo *DiscordChannelUpdateOne) RemoveChainIDs(ids ...int) *DiscordChannelUpdateOne {
	dcuo.mutation.RemoveChainIDs(ids...)
	return dcuo
}

// RemoveChains removes "chains" edges to Chain entities.
func (dcuo *DiscordChannelUpdateOne) RemoveChains(c ...*Chain) *DiscordChannelUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dcuo.RemoveChainIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dcuo *DiscordChannelUpdateOne) Select(field string, fields ...string) *DiscordChannelUpdateOne {
	dcuo.fields = append([]string{field}, fields...)
	return dcuo
}

// Save executes the query and returns the updated DiscordChannel entity.
func (dcuo *DiscordChannelUpdateOne) Save(ctx context.Context) (*DiscordChannel, error) {
	var (
		err  error
		node *DiscordChannel
	)
	dcuo.defaults()
	if len(dcuo.hooks) == 0 {
		node, err = dcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiscordChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dcuo.mutation = mutation
			node, err = dcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dcuo.hooks) - 1; i >= 0; i-- {
			if dcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dcuo *DiscordChannelUpdateOne) SaveX(ctx context.Context) *DiscordChannel {
	node, err := dcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dcuo *DiscordChannelUpdateOne) Exec(ctx context.Context) error {
	_, err := dcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcuo *DiscordChannelUpdateOne) ExecX(ctx context.Context) {
	if err := dcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dcuo *DiscordChannelUpdateOne) defaults() {
	if _, ok := dcuo.mutation.UpdateTime(); !ok {
		v := discordchannel.UpdateDefaultUpdateTime()
		dcuo.mutation.SetUpdateTime(v)
	}
}

func (dcuo *DiscordChannelUpdateOne) sqlSave(ctx context.Context) (_node *DiscordChannel, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   discordchannel.Table,
			Columns: discordchannel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: discordchannel.FieldID,
			},
		},
	}
	id, ok := dcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DiscordChannel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, discordchannel.FieldID)
		for _, f := range fields {
			if !discordchannel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != discordchannel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dcuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: discordchannel.FieldUpdateTime,
		})
	}
	if value, ok := dcuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discordchannel.FieldName,
		})
	}
	if value, ok := dcuo.mutation.Roles(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discordchannel.FieldRoles,
		})
	}
	if dcuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   discordchannel.UserTable,
			Columns: []string{discordchannel.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dcuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   discordchannel.UserTable,
			Columns: []string{discordchannel.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dcuo.mutation.ChainsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dcuo.mutation.RemovedChainsIDs(); len(nodes) > 0 && !dcuo.mutation.ChainsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dcuo.mutation.ChainsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DiscordChannel{config: dcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{discordchannel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
