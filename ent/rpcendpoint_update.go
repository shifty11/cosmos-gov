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
	"github.com/shifty11/cosmos-gov/ent/rpcendpoint"
)

// RpcEndpointUpdate is the builder for updating RpcEndpoint entities.
type RpcEndpointUpdate struct {
	config
	hooks    []Hook
	mutation *RpcEndpointMutation
}

// Where appends a list predicates to the RpcEndpointUpdate builder.
func (reu *RpcEndpointUpdate) Where(ps ...predicate.RpcEndpoint) *RpcEndpointUpdate {
	reu.mutation.Where(ps...)
	return reu
}

// SetUpdateTime sets the "update_time" field.
func (reu *RpcEndpointUpdate) SetUpdateTime(t time.Time) *RpcEndpointUpdate {
	reu.mutation.SetUpdateTime(t)
	return reu
}

// SetEndpoint sets the "endpoint" field.
func (reu *RpcEndpointUpdate) SetEndpoint(s string) *RpcEndpointUpdate {
	reu.mutation.SetEndpoint(s)
	return reu
}

// SetChainID sets the "chain" edge to the Chain entity by ID.
func (reu *RpcEndpointUpdate) SetChainID(id int) *RpcEndpointUpdate {
	reu.mutation.SetChainID(id)
	return reu
}

// SetNillableChainID sets the "chain" edge to the Chain entity by ID if the given value is not nil.
func (reu *RpcEndpointUpdate) SetNillableChainID(id *int) *RpcEndpointUpdate {
	if id != nil {
		reu = reu.SetChainID(*id)
	}
	return reu
}

// SetChain sets the "chain" edge to the Chain entity.
func (reu *RpcEndpointUpdate) SetChain(c *Chain) *RpcEndpointUpdate {
	return reu.SetChainID(c.ID)
}

// Mutation returns the RpcEndpointMutation object of the builder.
func (reu *RpcEndpointUpdate) Mutation() *RpcEndpointMutation {
	return reu.mutation
}

// ClearChain clears the "chain" edge to the Chain entity.
func (reu *RpcEndpointUpdate) ClearChain() *RpcEndpointUpdate {
	reu.mutation.ClearChain()
	return reu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (reu *RpcEndpointUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	reu.defaults()
	if len(reu.hooks) == 0 {
		affected, err = reu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RpcEndpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			reu.mutation = mutation
			affected, err = reu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(reu.hooks) - 1; i >= 0; i-- {
			if reu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = reu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, reu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (reu *RpcEndpointUpdate) SaveX(ctx context.Context) int {
	affected, err := reu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (reu *RpcEndpointUpdate) Exec(ctx context.Context) error {
	_, err := reu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (reu *RpcEndpointUpdate) ExecX(ctx context.Context) {
	if err := reu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (reu *RpcEndpointUpdate) defaults() {
	if _, ok := reu.mutation.UpdateTime(); !ok {
		v := rpcendpoint.UpdateDefaultUpdateTime()
		reu.mutation.SetUpdateTime(v)
	}
}

func (reu *RpcEndpointUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rpcendpoint.Table,
			Columns: rpcendpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rpcendpoint.FieldID,
			},
		},
	}
	if ps := reu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := reu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rpcendpoint.FieldUpdateTime,
		})
	}
	if value, ok := reu.mutation.Endpoint(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rpcendpoint.FieldEndpoint,
		})
	}
	if reu.mutation.ChainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rpcendpoint.ChainTable,
			Columns: []string{rpcendpoint.ChainColumn},
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
	if nodes := reu.mutation.ChainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rpcendpoint.ChainTable,
			Columns: []string{rpcendpoint.ChainColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, reu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rpcendpoint.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RpcEndpointUpdateOne is the builder for updating a single RpcEndpoint entity.
type RpcEndpointUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RpcEndpointMutation
}

// SetUpdateTime sets the "update_time" field.
func (reuo *RpcEndpointUpdateOne) SetUpdateTime(t time.Time) *RpcEndpointUpdateOne {
	reuo.mutation.SetUpdateTime(t)
	return reuo
}

// SetEndpoint sets the "endpoint" field.
func (reuo *RpcEndpointUpdateOne) SetEndpoint(s string) *RpcEndpointUpdateOne {
	reuo.mutation.SetEndpoint(s)
	return reuo
}

// SetChainID sets the "chain" edge to the Chain entity by ID.
func (reuo *RpcEndpointUpdateOne) SetChainID(id int) *RpcEndpointUpdateOne {
	reuo.mutation.SetChainID(id)
	return reuo
}

// SetNillableChainID sets the "chain" edge to the Chain entity by ID if the given value is not nil.
func (reuo *RpcEndpointUpdateOne) SetNillableChainID(id *int) *RpcEndpointUpdateOne {
	if id != nil {
		reuo = reuo.SetChainID(*id)
	}
	return reuo
}

// SetChain sets the "chain" edge to the Chain entity.
func (reuo *RpcEndpointUpdateOne) SetChain(c *Chain) *RpcEndpointUpdateOne {
	return reuo.SetChainID(c.ID)
}

// Mutation returns the RpcEndpointMutation object of the builder.
func (reuo *RpcEndpointUpdateOne) Mutation() *RpcEndpointMutation {
	return reuo.mutation
}

// ClearChain clears the "chain" edge to the Chain entity.
func (reuo *RpcEndpointUpdateOne) ClearChain() *RpcEndpointUpdateOne {
	reuo.mutation.ClearChain()
	return reuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (reuo *RpcEndpointUpdateOne) Select(field string, fields ...string) *RpcEndpointUpdateOne {
	reuo.fields = append([]string{field}, fields...)
	return reuo
}

// Save executes the query and returns the updated RpcEndpoint entity.
func (reuo *RpcEndpointUpdateOne) Save(ctx context.Context) (*RpcEndpoint, error) {
	var (
		err  error
		node *RpcEndpoint
	)
	reuo.defaults()
	if len(reuo.hooks) == 0 {
		node, err = reuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RpcEndpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			reuo.mutation = mutation
			node, err = reuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(reuo.hooks) - 1; i >= 0; i-- {
			if reuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = reuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, reuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (reuo *RpcEndpointUpdateOne) SaveX(ctx context.Context) *RpcEndpoint {
	node, err := reuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (reuo *RpcEndpointUpdateOne) Exec(ctx context.Context) error {
	_, err := reuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (reuo *RpcEndpointUpdateOne) ExecX(ctx context.Context) {
	if err := reuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (reuo *RpcEndpointUpdateOne) defaults() {
	if _, ok := reuo.mutation.UpdateTime(); !ok {
		v := rpcendpoint.UpdateDefaultUpdateTime()
		reuo.mutation.SetUpdateTime(v)
	}
}

func (reuo *RpcEndpointUpdateOne) sqlSave(ctx context.Context) (_node *RpcEndpoint, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rpcendpoint.Table,
			Columns: rpcendpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rpcendpoint.FieldID,
			},
		},
	}
	id, ok := reuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RpcEndpoint.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := reuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rpcendpoint.FieldID)
		for _, f := range fields {
			if !rpcendpoint.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rpcendpoint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := reuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := reuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rpcendpoint.FieldUpdateTime,
		})
	}
	if value, ok := reuo.mutation.Endpoint(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rpcendpoint.FieldEndpoint,
		})
	}
	if reuo.mutation.ChainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rpcendpoint.ChainTable,
			Columns: []string{rpcendpoint.ChainColumn},
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
	if nodes := reuo.mutation.ChainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rpcendpoint.ChainTable,
			Columns: []string{rpcendpoint.ChainColumn},
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
	_node = &RpcEndpoint{config: reuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, reuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rpcendpoint.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
