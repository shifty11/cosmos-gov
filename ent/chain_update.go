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
	"github.com/shifty11/cosmos-gov/ent/proposal"
	"github.com/shifty11/cosmos-gov/ent/user"
)

// ChainUpdate is the builder for updating Chain entities.
type ChainUpdate struct {
	config
	hooks    []Hook
	mutation *ChainMutation
}

// Where appends a list predicates to the ChainUpdate builder.
func (cu *ChainUpdate) Where(ps ...predicate.Chain) *ChainUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ChainUpdate) SetUpdatedAt(t time.Time) *ChainUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetName sets the "name" field.
func (cu *ChainUpdate) SetName(s string) *ChainUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetDisplayName sets the "display_name" field.
func (cu *ChainUpdate) SetDisplayName(s string) *ChainUpdate {
	cu.mutation.SetDisplayName(s)
	return cu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (cu *ChainUpdate) AddUserIDs(ids ...int) *ChainUpdate {
	cu.mutation.AddUserIDs(ids...)
	return cu
}

// AddUsers adds the "users" edges to the User entity.
func (cu *ChainUpdate) AddUsers(u ...*User) *ChainUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.AddUserIDs(ids...)
}

// AddProposalIDs adds the "proposals" edge to the Proposal entity by IDs.
func (cu *ChainUpdate) AddProposalIDs(ids ...int) *ChainUpdate {
	cu.mutation.AddProposalIDs(ids...)
	return cu
}

// AddProposals adds the "proposals" edges to the Proposal entity.
func (cu *ChainUpdate) AddProposals(p ...*Proposal) *ChainUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddProposalIDs(ids...)
}

// Mutation returns the ChainMutation object of the builder.
func (cu *ChainUpdate) Mutation() *ChainMutation {
	return cu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (cu *ChainUpdate) ClearUsers() *ChainUpdate {
	cu.mutation.ClearUsers()
	return cu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (cu *ChainUpdate) RemoveUserIDs(ids ...int) *ChainUpdate {
	cu.mutation.RemoveUserIDs(ids...)
	return cu
}

// RemoveUsers removes "users" edges to User entities.
func (cu *ChainUpdate) RemoveUsers(u ...*User) *ChainUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.RemoveUserIDs(ids...)
}

// ClearProposals clears all "proposals" edges to the Proposal entity.
func (cu *ChainUpdate) ClearProposals() *ChainUpdate {
	cu.mutation.ClearProposals()
	return cu
}

// RemoveProposalIDs removes the "proposals" edge to Proposal entities by IDs.
func (cu *ChainUpdate) RemoveProposalIDs(ids ...int) *ChainUpdate {
	cu.mutation.RemoveProposalIDs(ids...)
	return cu
}

// RemoveProposals removes "proposals" edges to Proposal entities.
func (cu *ChainUpdate) RemoveProposals(p ...*Proposal) *ChainUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemoveProposalIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChainUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChainMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChainUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChainUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChainUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ChainUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := chain.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *ChainUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chain.Table,
			Columns: chain.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: chain.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chain.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chain.FieldName,
		})
	}
	if value, ok := cu.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chain.FieldDisplayName,
		})
	}
	if cu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.UsersTable,
			Columns: chain.UsersPrimaryKey,
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
	if nodes := cu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !cu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.UsersTable,
			Columns: chain.UsersPrimaryKey,
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
	if nodes := cu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.UsersTable,
			Columns: chain.UsersPrimaryKey,
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
	if cu.mutation.ProposalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ProposalsTable,
			Columns: []string{chain.ProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: proposal.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedProposalsIDs(); len(nodes) > 0 && !cu.mutation.ProposalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ProposalsTable,
			Columns: []string{chain.ProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: proposal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ProposalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ProposalsTable,
			Columns: []string{chain.ProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: proposal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chain.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ChainUpdateOne is the builder for updating a single Chain entity.
type ChainUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChainMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ChainUpdateOne) SetUpdatedAt(t time.Time) *ChainUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetName sets the "name" field.
func (cuo *ChainUpdateOne) SetName(s string) *ChainUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetDisplayName sets the "display_name" field.
func (cuo *ChainUpdateOne) SetDisplayName(s string) *ChainUpdateOne {
	cuo.mutation.SetDisplayName(s)
	return cuo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (cuo *ChainUpdateOne) AddUserIDs(ids ...int) *ChainUpdateOne {
	cuo.mutation.AddUserIDs(ids...)
	return cuo
}

// AddUsers adds the "users" edges to the User entity.
func (cuo *ChainUpdateOne) AddUsers(u ...*User) *ChainUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.AddUserIDs(ids...)
}

// AddProposalIDs adds the "proposals" edge to the Proposal entity by IDs.
func (cuo *ChainUpdateOne) AddProposalIDs(ids ...int) *ChainUpdateOne {
	cuo.mutation.AddProposalIDs(ids...)
	return cuo
}

// AddProposals adds the "proposals" edges to the Proposal entity.
func (cuo *ChainUpdateOne) AddProposals(p ...*Proposal) *ChainUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddProposalIDs(ids...)
}

// Mutation returns the ChainMutation object of the builder.
func (cuo *ChainUpdateOne) Mutation() *ChainMutation {
	return cuo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (cuo *ChainUpdateOne) ClearUsers() *ChainUpdateOne {
	cuo.mutation.ClearUsers()
	return cuo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (cuo *ChainUpdateOne) RemoveUserIDs(ids ...int) *ChainUpdateOne {
	cuo.mutation.RemoveUserIDs(ids...)
	return cuo
}

// RemoveUsers removes "users" edges to User entities.
func (cuo *ChainUpdateOne) RemoveUsers(u ...*User) *ChainUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.RemoveUserIDs(ids...)
}

// ClearProposals clears all "proposals" edges to the Proposal entity.
func (cuo *ChainUpdateOne) ClearProposals() *ChainUpdateOne {
	cuo.mutation.ClearProposals()
	return cuo
}

// RemoveProposalIDs removes the "proposals" edge to Proposal entities by IDs.
func (cuo *ChainUpdateOne) RemoveProposalIDs(ids ...int) *ChainUpdateOne {
	cuo.mutation.RemoveProposalIDs(ids...)
	return cuo
}

// RemoveProposals removes "proposals" edges to Proposal entities.
func (cuo *ChainUpdateOne) RemoveProposals(p ...*Proposal) *ChainUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemoveProposalIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChainUpdateOne) Select(field string, fields ...string) *ChainUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Chain entity.
func (cuo *ChainUpdateOne) Save(ctx context.Context) (*Chain, error) {
	var (
		err  error
		node *Chain
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChainMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChainUpdateOne) SaveX(ctx context.Context) *Chain {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChainUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChainUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ChainUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := chain.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *ChainUpdateOne) sqlSave(ctx context.Context) (_node *Chain, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chain.Table,
			Columns: chain.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: chain.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Chain.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chain.FieldID)
		for _, f := range fields {
			if !chain.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chain.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chain.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chain.FieldName,
		})
	}
	if value, ok := cuo.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chain.FieldDisplayName,
		})
	}
	if cuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.UsersTable,
			Columns: chain.UsersPrimaryKey,
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
	if nodes := cuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !cuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.UsersTable,
			Columns: chain.UsersPrimaryKey,
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
	if nodes := cuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.UsersTable,
			Columns: chain.UsersPrimaryKey,
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
	if cuo.mutation.ProposalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ProposalsTable,
			Columns: []string{chain.ProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: proposal.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedProposalsIDs(); len(nodes) > 0 && !cuo.mutation.ProposalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ProposalsTable,
			Columns: []string{chain.ProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: proposal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ProposalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ProposalsTable,
			Columns: []string{chain.ProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: proposal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Chain{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chain.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
