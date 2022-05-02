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
	"github.com/shifty11/cosmos-gov/ent/grant"
	"github.com/shifty11/cosmos-gov/ent/predicate"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/ent/wallet"
)

// WalletUpdate is the builder for updating Wallet entities.
type WalletUpdate struct {
	config
	hooks    []Hook
	mutation *WalletMutation
}

// Where appends a list predicates to the WalletUpdate builder.
func (wu *WalletUpdate) Where(ps ...predicate.Wallet) *WalletUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetUpdateTime sets the "update_time" field.
func (wu *WalletUpdate) SetUpdateTime(t time.Time) *WalletUpdate {
	wu.mutation.SetUpdateTime(t)
	return wu
}

// SetAddress sets the "address" field.
func (wu *WalletUpdate) SetAddress(s string) *WalletUpdate {
	wu.mutation.SetAddress(s)
	return wu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (wu *WalletUpdate) AddUserIDs(ids ...int) *WalletUpdate {
	wu.mutation.AddUserIDs(ids...)
	return wu
}

// AddUsers adds the "users" edges to the User entity.
func (wu *WalletUpdate) AddUsers(u ...*User) *WalletUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return wu.AddUserIDs(ids...)
}

// SetChainID sets the "chain" edge to the Chain entity by ID.
func (wu *WalletUpdate) SetChainID(id int) *WalletUpdate {
	wu.mutation.SetChainID(id)
	return wu
}

// SetNillableChainID sets the "chain" edge to the Chain entity by ID if the given value is not nil.
func (wu *WalletUpdate) SetNillableChainID(id *int) *WalletUpdate {
	if id != nil {
		wu = wu.SetChainID(*id)
	}
	return wu
}

// SetChain sets the "chain" edge to the Chain entity.
func (wu *WalletUpdate) SetChain(c *Chain) *WalletUpdate {
	return wu.SetChainID(c.ID)
}

// AddGrantIDs adds the "grants" edge to the Grant entity by IDs.
func (wu *WalletUpdate) AddGrantIDs(ids ...int) *WalletUpdate {
	wu.mutation.AddGrantIDs(ids...)
	return wu
}

// AddGrants adds the "grants" edges to the Grant entity.
func (wu *WalletUpdate) AddGrants(g ...*Grant) *WalletUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return wu.AddGrantIDs(ids...)
}

// Mutation returns the WalletMutation object of the builder.
func (wu *WalletUpdate) Mutation() *WalletMutation {
	return wu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (wu *WalletUpdate) ClearUsers() *WalletUpdate {
	wu.mutation.ClearUsers()
	return wu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (wu *WalletUpdate) RemoveUserIDs(ids ...int) *WalletUpdate {
	wu.mutation.RemoveUserIDs(ids...)
	return wu
}

// RemoveUsers removes "users" edges to User entities.
func (wu *WalletUpdate) RemoveUsers(u ...*User) *WalletUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return wu.RemoveUserIDs(ids...)
}

// ClearChain clears the "chain" edge to the Chain entity.
func (wu *WalletUpdate) ClearChain() *WalletUpdate {
	wu.mutation.ClearChain()
	return wu
}

// ClearGrants clears all "grants" edges to the Grant entity.
func (wu *WalletUpdate) ClearGrants() *WalletUpdate {
	wu.mutation.ClearGrants()
	return wu
}

// RemoveGrantIDs removes the "grants" edge to Grant entities by IDs.
func (wu *WalletUpdate) RemoveGrantIDs(ids ...int) *WalletUpdate {
	wu.mutation.RemoveGrantIDs(ids...)
	return wu
}

// RemoveGrants removes "grants" edges to Grant entities.
func (wu *WalletUpdate) RemoveGrants(g ...*Grant) *WalletUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return wu.RemoveGrantIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WalletUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	wu.defaults()
	if len(wu.hooks) == 0 {
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WalletUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WalletUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WalletUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wu *WalletUpdate) defaults() {
	if _, ok := wu.mutation.UpdateTime(); !ok {
		v := wallet.UpdateDefaultUpdateTime()
		wu.mutation.SetUpdateTime(v)
	}
}

func (wu *WalletUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wallet.Table,
			Columns: wallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: wallet.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: wallet.FieldUpdateTime,
		})
	}
	if value, ok := wu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wallet.FieldAddress,
		})
	}
	if wu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   wallet.UsersTable,
			Columns: wallet.UsersPrimaryKey,
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
	if nodes := wu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !wu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   wallet.UsersTable,
			Columns: wallet.UsersPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   wallet.UsersTable,
			Columns: wallet.UsersPrimaryKey,
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
	if wu.mutation.ChainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wallet.ChainTable,
			Columns: []string{wallet.ChainColumn},
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
	if nodes := wu.mutation.ChainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wallet.ChainTable,
			Columns: []string{wallet.ChainColumn},
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
	if wu.mutation.GrantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.GrantsTable,
			Columns: []string{wallet.GrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: grant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedGrantsIDs(); len(nodes) > 0 && !wu.mutation.GrantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.GrantsTable,
			Columns: []string{wallet.GrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: grant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.GrantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.GrantsTable,
			Columns: []string{wallet.GrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: grant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wallet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WalletUpdateOne is the builder for updating a single Wallet entity.
type WalletUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WalletMutation
}

// SetUpdateTime sets the "update_time" field.
func (wuo *WalletUpdateOne) SetUpdateTime(t time.Time) *WalletUpdateOne {
	wuo.mutation.SetUpdateTime(t)
	return wuo
}

// SetAddress sets the "address" field.
func (wuo *WalletUpdateOne) SetAddress(s string) *WalletUpdateOne {
	wuo.mutation.SetAddress(s)
	return wuo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (wuo *WalletUpdateOne) AddUserIDs(ids ...int) *WalletUpdateOne {
	wuo.mutation.AddUserIDs(ids...)
	return wuo
}

// AddUsers adds the "users" edges to the User entity.
func (wuo *WalletUpdateOne) AddUsers(u ...*User) *WalletUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return wuo.AddUserIDs(ids...)
}

// SetChainID sets the "chain" edge to the Chain entity by ID.
func (wuo *WalletUpdateOne) SetChainID(id int) *WalletUpdateOne {
	wuo.mutation.SetChainID(id)
	return wuo
}

// SetNillableChainID sets the "chain" edge to the Chain entity by ID if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableChainID(id *int) *WalletUpdateOne {
	if id != nil {
		wuo = wuo.SetChainID(*id)
	}
	return wuo
}

// SetChain sets the "chain" edge to the Chain entity.
func (wuo *WalletUpdateOne) SetChain(c *Chain) *WalletUpdateOne {
	return wuo.SetChainID(c.ID)
}

// AddGrantIDs adds the "grants" edge to the Grant entity by IDs.
func (wuo *WalletUpdateOne) AddGrantIDs(ids ...int) *WalletUpdateOne {
	wuo.mutation.AddGrantIDs(ids...)
	return wuo
}

// AddGrants adds the "grants" edges to the Grant entity.
func (wuo *WalletUpdateOne) AddGrants(g ...*Grant) *WalletUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return wuo.AddGrantIDs(ids...)
}

// Mutation returns the WalletMutation object of the builder.
func (wuo *WalletUpdateOne) Mutation() *WalletMutation {
	return wuo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (wuo *WalletUpdateOne) ClearUsers() *WalletUpdateOne {
	wuo.mutation.ClearUsers()
	return wuo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (wuo *WalletUpdateOne) RemoveUserIDs(ids ...int) *WalletUpdateOne {
	wuo.mutation.RemoveUserIDs(ids...)
	return wuo
}

// RemoveUsers removes "users" edges to User entities.
func (wuo *WalletUpdateOne) RemoveUsers(u ...*User) *WalletUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return wuo.RemoveUserIDs(ids...)
}

// ClearChain clears the "chain" edge to the Chain entity.
func (wuo *WalletUpdateOne) ClearChain() *WalletUpdateOne {
	wuo.mutation.ClearChain()
	return wuo
}

// ClearGrants clears all "grants" edges to the Grant entity.
func (wuo *WalletUpdateOne) ClearGrants() *WalletUpdateOne {
	wuo.mutation.ClearGrants()
	return wuo
}

// RemoveGrantIDs removes the "grants" edge to Grant entities by IDs.
func (wuo *WalletUpdateOne) RemoveGrantIDs(ids ...int) *WalletUpdateOne {
	wuo.mutation.RemoveGrantIDs(ids...)
	return wuo
}

// RemoveGrants removes "grants" edges to Grant entities.
func (wuo *WalletUpdateOne) RemoveGrants(g ...*Grant) *WalletUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return wuo.RemoveGrantIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WalletUpdateOne) Select(field string, fields ...string) *WalletUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Wallet entity.
func (wuo *WalletUpdateOne) Save(ctx context.Context) (*Wallet, error) {
	var (
		err  error
		node *Wallet
	)
	wuo.defaults()
	if len(wuo.hooks) == 0 {
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WalletUpdateOne) SaveX(ctx context.Context) *Wallet {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WalletUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WalletUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuo *WalletUpdateOne) defaults() {
	if _, ok := wuo.mutation.UpdateTime(); !ok {
		v := wallet.UpdateDefaultUpdateTime()
		wuo.mutation.SetUpdateTime(v)
	}
}

func (wuo *WalletUpdateOne) sqlSave(ctx context.Context) (_node *Wallet, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wallet.Table,
			Columns: wallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: wallet.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Wallet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wallet.FieldID)
		for _, f := range fields {
			if !wallet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != wallet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: wallet.FieldUpdateTime,
		})
	}
	if value, ok := wuo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wallet.FieldAddress,
		})
	}
	if wuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   wallet.UsersTable,
			Columns: wallet.UsersPrimaryKey,
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
	if nodes := wuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !wuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   wallet.UsersTable,
			Columns: wallet.UsersPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   wallet.UsersTable,
			Columns: wallet.UsersPrimaryKey,
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
	if wuo.mutation.ChainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wallet.ChainTable,
			Columns: []string{wallet.ChainColumn},
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
	if nodes := wuo.mutation.ChainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wallet.ChainTable,
			Columns: []string{wallet.ChainColumn},
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
	if wuo.mutation.GrantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.GrantsTable,
			Columns: []string{wallet.GrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: grant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedGrantsIDs(); len(nodes) > 0 && !wuo.mutation.GrantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.GrantsTable,
			Columns: []string{wallet.GrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: grant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.GrantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.GrantsTable,
			Columns: []string{wallet.GrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: grant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Wallet{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wallet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}