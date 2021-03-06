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
	"github.com/shifty11/cosmos-gov/ent/predicate"
	"github.com/shifty11/cosmos-gov/ent/telegramchat"
	"github.com/shifty11/cosmos-gov/ent/user"
)

// TelegramChatUpdate is the builder for updating TelegramChat entities.
type TelegramChatUpdate struct {
	config
	hooks    []Hook
	mutation *TelegramChatMutation
}

// Where appends a list predicates to the TelegramChatUpdate builder.
func (tcu *TelegramChatUpdate) Where(ps ...predicate.TelegramChat) *TelegramChatUpdate {
	tcu.mutation.Where(ps...)
	return tcu
}

// SetUpdateTime sets the "update_time" field.
func (tcu *TelegramChatUpdate) SetUpdateTime(t time.Time) *TelegramChatUpdate {
	tcu.mutation.SetUpdateTime(t)
	return tcu
}

// SetChatID sets the "chat_id" field.
func (tcu *TelegramChatUpdate) SetChatID(i int64) *TelegramChatUpdate {
	tcu.mutation.ResetChatID()
	tcu.mutation.SetChatID(i)
	return tcu
}

// AddChatID adds i to the "chat_id" field.
func (tcu *TelegramChatUpdate) AddChatID(i int64) *TelegramChatUpdate {
	tcu.mutation.AddChatID(i)
	return tcu
}

// SetName sets the "name" field.
func (tcu *TelegramChatUpdate) SetName(s string) *TelegramChatUpdate {
	tcu.mutation.SetName(s)
	return tcu
}

// SetWantsDraftProposals sets the "wants_draft_proposals" field.
func (tcu *TelegramChatUpdate) SetWantsDraftProposals(b bool) *TelegramChatUpdate {
	tcu.mutation.SetWantsDraftProposals(b)
	return tcu
}

// SetNillableWantsDraftProposals sets the "wants_draft_proposals" field if the given value is not nil.
func (tcu *TelegramChatUpdate) SetNillableWantsDraftProposals(b *bool) *TelegramChatUpdate {
	if b != nil {
		tcu.SetWantsDraftProposals(*b)
	}
	return tcu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tcu *TelegramChatUpdate) SetUserID(id int) *TelegramChatUpdate {
	tcu.mutation.SetUserID(id)
	return tcu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (tcu *TelegramChatUpdate) SetNillableUserID(id *int) *TelegramChatUpdate {
	if id != nil {
		tcu = tcu.SetUserID(*id)
	}
	return tcu
}

// SetUser sets the "user" edge to the User entity.
func (tcu *TelegramChatUpdate) SetUser(u *User) *TelegramChatUpdate {
	return tcu.SetUserID(u.ID)
}

// AddChainIDs adds the "chains" edge to the Chain entity by IDs.
func (tcu *TelegramChatUpdate) AddChainIDs(ids ...int) *TelegramChatUpdate {
	tcu.mutation.AddChainIDs(ids...)
	return tcu
}

// AddChains adds the "chains" edges to the Chain entity.
func (tcu *TelegramChatUpdate) AddChains(c ...*Chain) *TelegramChatUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tcu.AddChainIDs(ids...)
}

// Mutation returns the TelegramChatMutation object of the builder.
func (tcu *TelegramChatUpdate) Mutation() *TelegramChatMutation {
	return tcu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (tcu *TelegramChatUpdate) ClearUser() *TelegramChatUpdate {
	tcu.mutation.ClearUser()
	return tcu
}

// ClearChains clears all "chains" edges to the Chain entity.
func (tcu *TelegramChatUpdate) ClearChains() *TelegramChatUpdate {
	tcu.mutation.ClearChains()
	return tcu
}

// RemoveChainIDs removes the "chains" edge to Chain entities by IDs.
func (tcu *TelegramChatUpdate) RemoveChainIDs(ids ...int) *TelegramChatUpdate {
	tcu.mutation.RemoveChainIDs(ids...)
	return tcu
}

// RemoveChains removes "chains" edges to Chain entities.
func (tcu *TelegramChatUpdate) RemoveChains(c ...*Chain) *TelegramChatUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tcu.RemoveChainIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tcu *TelegramChatUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tcu.defaults()
	if len(tcu.hooks) == 0 {
		affected, err = tcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TelegramChatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tcu.mutation = mutation
			affected, err = tcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tcu.hooks) - 1; i >= 0; i-- {
			if tcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tcu *TelegramChatUpdate) SaveX(ctx context.Context) int {
	affected, err := tcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tcu *TelegramChatUpdate) Exec(ctx context.Context) error {
	_, err := tcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcu *TelegramChatUpdate) ExecX(ctx context.Context) {
	if err := tcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcu *TelegramChatUpdate) defaults() {
	if _, ok := tcu.mutation.UpdateTime(); !ok {
		v := telegramchat.UpdateDefaultUpdateTime()
		tcu.mutation.SetUpdateTime(v)
	}
}

func (tcu *TelegramChatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   telegramchat.Table,
			Columns: telegramchat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: telegramchat.FieldID,
			},
		},
	}
	if ps := tcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: telegramchat.FieldUpdateTime,
		})
	}
	if value, ok := tcu.mutation.ChatID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: telegramchat.FieldChatID,
		})
	}
	if value, ok := tcu.mutation.AddedChatID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: telegramchat.FieldChatID,
		})
	}
	if value, ok := tcu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telegramchat.FieldName,
		})
	}
	if value, ok := tcu.mutation.WantsDraftProposals(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: telegramchat.FieldWantsDraftProposals,
		})
	}
	if tcu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   telegramchat.UserTable,
			Columns: []string{telegramchat.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   telegramchat.UserTable,
			Columns: []string{telegramchat.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tcu.mutation.ChainsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.RemovedChainsIDs(); len(nodes) > 0 && !tcu.mutation.ChainsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.ChainsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{telegramchat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TelegramChatUpdateOne is the builder for updating a single TelegramChat entity.
type TelegramChatUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TelegramChatMutation
}

// SetUpdateTime sets the "update_time" field.
func (tcuo *TelegramChatUpdateOne) SetUpdateTime(t time.Time) *TelegramChatUpdateOne {
	tcuo.mutation.SetUpdateTime(t)
	return tcuo
}

// SetChatID sets the "chat_id" field.
func (tcuo *TelegramChatUpdateOne) SetChatID(i int64) *TelegramChatUpdateOne {
	tcuo.mutation.ResetChatID()
	tcuo.mutation.SetChatID(i)
	return tcuo
}

// AddChatID adds i to the "chat_id" field.
func (tcuo *TelegramChatUpdateOne) AddChatID(i int64) *TelegramChatUpdateOne {
	tcuo.mutation.AddChatID(i)
	return tcuo
}

// SetName sets the "name" field.
func (tcuo *TelegramChatUpdateOne) SetName(s string) *TelegramChatUpdateOne {
	tcuo.mutation.SetName(s)
	return tcuo
}

// SetWantsDraftProposals sets the "wants_draft_proposals" field.
func (tcuo *TelegramChatUpdateOne) SetWantsDraftProposals(b bool) *TelegramChatUpdateOne {
	tcuo.mutation.SetWantsDraftProposals(b)
	return tcuo
}

// SetNillableWantsDraftProposals sets the "wants_draft_proposals" field if the given value is not nil.
func (tcuo *TelegramChatUpdateOne) SetNillableWantsDraftProposals(b *bool) *TelegramChatUpdateOne {
	if b != nil {
		tcuo.SetWantsDraftProposals(*b)
	}
	return tcuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tcuo *TelegramChatUpdateOne) SetUserID(id int) *TelegramChatUpdateOne {
	tcuo.mutation.SetUserID(id)
	return tcuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (tcuo *TelegramChatUpdateOne) SetNillableUserID(id *int) *TelegramChatUpdateOne {
	if id != nil {
		tcuo = tcuo.SetUserID(*id)
	}
	return tcuo
}

// SetUser sets the "user" edge to the User entity.
func (tcuo *TelegramChatUpdateOne) SetUser(u *User) *TelegramChatUpdateOne {
	return tcuo.SetUserID(u.ID)
}

// AddChainIDs adds the "chains" edge to the Chain entity by IDs.
func (tcuo *TelegramChatUpdateOne) AddChainIDs(ids ...int) *TelegramChatUpdateOne {
	tcuo.mutation.AddChainIDs(ids...)
	return tcuo
}

// AddChains adds the "chains" edges to the Chain entity.
func (tcuo *TelegramChatUpdateOne) AddChains(c ...*Chain) *TelegramChatUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tcuo.AddChainIDs(ids...)
}

// Mutation returns the TelegramChatMutation object of the builder.
func (tcuo *TelegramChatUpdateOne) Mutation() *TelegramChatMutation {
	return tcuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (tcuo *TelegramChatUpdateOne) ClearUser() *TelegramChatUpdateOne {
	tcuo.mutation.ClearUser()
	return tcuo
}

// ClearChains clears all "chains" edges to the Chain entity.
func (tcuo *TelegramChatUpdateOne) ClearChains() *TelegramChatUpdateOne {
	tcuo.mutation.ClearChains()
	return tcuo
}

// RemoveChainIDs removes the "chains" edge to Chain entities by IDs.
func (tcuo *TelegramChatUpdateOne) RemoveChainIDs(ids ...int) *TelegramChatUpdateOne {
	tcuo.mutation.RemoveChainIDs(ids...)
	return tcuo
}

// RemoveChains removes "chains" edges to Chain entities.
func (tcuo *TelegramChatUpdateOne) RemoveChains(c ...*Chain) *TelegramChatUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tcuo.RemoveChainIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tcuo *TelegramChatUpdateOne) Select(field string, fields ...string) *TelegramChatUpdateOne {
	tcuo.fields = append([]string{field}, fields...)
	return tcuo
}

// Save executes the query and returns the updated TelegramChat entity.
func (tcuo *TelegramChatUpdateOne) Save(ctx context.Context) (*TelegramChat, error) {
	var (
		err  error
		node *TelegramChat
	)
	tcuo.defaults()
	if len(tcuo.hooks) == 0 {
		node, err = tcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TelegramChatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tcuo.mutation = mutation
			node, err = tcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tcuo.hooks) - 1; i >= 0; i-- {
			if tcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tcuo *TelegramChatUpdateOne) SaveX(ctx context.Context) *TelegramChat {
	node, err := tcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tcuo *TelegramChatUpdateOne) Exec(ctx context.Context) error {
	_, err := tcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcuo *TelegramChatUpdateOne) ExecX(ctx context.Context) {
	if err := tcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcuo *TelegramChatUpdateOne) defaults() {
	if _, ok := tcuo.mutation.UpdateTime(); !ok {
		v := telegramchat.UpdateDefaultUpdateTime()
		tcuo.mutation.SetUpdateTime(v)
	}
}

func (tcuo *TelegramChatUpdateOne) sqlSave(ctx context.Context) (_node *TelegramChat, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   telegramchat.Table,
			Columns: telegramchat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: telegramchat.FieldID,
			},
		},
	}
	id, ok := tcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TelegramChat.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, telegramchat.FieldID)
		for _, f := range fields {
			if !telegramchat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != telegramchat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: telegramchat.FieldUpdateTime,
		})
	}
	if value, ok := tcuo.mutation.ChatID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: telegramchat.FieldChatID,
		})
	}
	if value, ok := tcuo.mutation.AddedChatID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: telegramchat.FieldChatID,
		})
	}
	if value, ok := tcuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: telegramchat.FieldName,
		})
	}
	if value, ok := tcuo.mutation.WantsDraftProposals(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: telegramchat.FieldWantsDraftProposals,
		})
	}
	if tcuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   telegramchat.UserTable,
			Columns: []string{telegramchat.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   telegramchat.UserTable,
			Columns: []string{telegramchat.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tcuo.mutation.ChainsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.RemovedChainsIDs(); len(nodes) > 0 && !tcuo.mutation.ChainsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.ChainsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TelegramChat{config: tcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{telegramchat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
