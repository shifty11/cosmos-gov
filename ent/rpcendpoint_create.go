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
	"github.com/shifty11/cosmos-gov/ent/rpcendpoint"
)

// RpcEndpointCreate is the builder for creating a RpcEndpoint entity.
type RpcEndpointCreate struct {
	config
	mutation *RpcEndpointMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rec *RpcEndpointCreate) SetCreatedAt(t time.Time) *RpcEndpointCreate {
	rec.mutation.SetCreatedAt(t)
	return rec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rec *RpcEndpointCreate) SetNillableCreatedAt(t *time.Time) *RpcEndpointCreate {
	if t != nil {
		rec.SetCreatedAt(*t)
	}
	return rec
}

// SetUpdatedAt sets the "updated_at" field.
func (rec *RpcEndpointCreate) SetUpdatedAt(t time.Time) *RpcEndpointCreate {
	rec.mutation.SetUpdatedAt(t)
	return rec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rec *RpcEndpointCreate) SetNillableUpdatedAt(t *time.Time) *RpcEndpointCreate {
	if t != nil {
		rec.SetUpdatedAt(*t)
	}
	return rec
}

// SetEndpoint sets the "endpoint" field.
func (rec *RpcEndpointCreate) SetEndpoint(s string) *RpcEndpointCreate {
	rec.mutation.SetEndpoint(s)
	return rec
}

// SetChainID sets the "chain" edge to the Chain entity by ID.
func (rec *RpcEndpointCreate) SetChainID(id int) *RpcEndpointCreate {
	rec.mutation.SetChainID(id)
	return rec
}

// SetNillableChainID sets the "chain" edge to the Chain entity by ID if the given value is not nil.
func (rec *RpcEndpointCreate) SetNillableChainID(id *int) *RpcEndpointCreate {
	if id != nil {
		rec = rec.SetChainID(*id)
	}
	return rec
}

// SetChain sets the "chain" edge to the Chain entity.
func (rec *RpcEndpointCreate) SetChain(c *Chain) *RpcEndpointCreate {
	return rec.SetChainID(c.ID)
}

// Mutation returns the RpcEndpointMutation object of the builder.
func (rec *RpcEndpointCreate) Mutation() *RpcEndpointMutation {
	return rec.mutation
}

// Save creates the RpcEndpoint in the database.
func (rec *RpcEndpointCreate) Save(ctx context.Context) (*RpcEndpoint, error) {
	var (
		err  error
		node *RpcEndpoint
	)
	rec.defaults()
	if len(rec.hooks) == 0 {
		if err = rec.check(); err != nil {
			return nil, err
		}
		node, err = rec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RpcEndpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rec.check(); err != nil {
				return nil, err
			}
			rec.mutation = mutation
			if node, err = rec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rec.hooks) - 1; i >= 0; i-- {
			if rec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rec *RpcEndpointCreate) SaveX(ctx context.Context) *RpcEndpoint {
	v, err := rec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rec *RpcEndpointCreate) Exec(ctx context.Context) error {
	_, err := rec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rec *RpcEndpointCreate) ExecX(ctx context.Context) {
	if err := rec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rec *RpcEndpointCreate) defaults() {
	if _, ok := rec.mutation.CreatedAt(); !ok {
		v := rpcendpoint.DefaultCreatedAt()
		rec.mutation.SetCreatedAt(v)
	}
	if _, ok := rec.mutation.UpdatedAt(); !ok {
		v := rpcendpoint.DefaultUpdatedAt()
		rec.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rec *RpcEndpointCreate) check() error {
	if _, ok := rec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RpcEndpoint.created_at"`)}
	}
	if _, ok := rec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "RpcEndpoint.updated_at"`)}
	}
	if _, ok := rec.mutation.Endpoint(); !ok {
		return &ValidationError{Name: "endpoint", err: errors.New(`ent: missing required field "RpcEndpoint.endpoint"`)}
	}
	return nil
}

func (rec *RpcEndpointCreate) sqlSave(ctx context.Context) (*RpcEndpoint, error) {
	_node, _spec := rec.createSpec()
	if err := sqlgraph.CreateNode(ctx, rec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rec *RpcEndpointCreate) createSpec() (*RpcEndpoint, *sqlgraph.CreateSpec) {
	var (
		_node = &RpcEndpoint{config: rec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rpcendpoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rpcendpoint.FieldID,
			},
		}
	)
	if value, ok := rec.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rpcendpoint.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := rec.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rpcendpoint.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := rec.mutation.Endpoint(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rpcendpoint.FieldEndpoint,
		})
		_node.Endpoint = value
	}
	if nodes := rec.mutation.ChainIDs(); len(nodes) > 0 {
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
		_node.chain_rpc_endpoints = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RpcEndpointCreateBulk is the builder for creating many RpcEndpoint entities in bulk.
type RpcEndpointCreateBulk struct {
	config
	builders []*RpcEndpointCreate
}

// Save creates the RpcEndpoint entities in the database.
func (recb *RpcEndpointCreateBulk) Save(ctx context.Context) ([]*RpcEndpoint, error) {
	specs := make([]*sqlgraph.CreateSpec, len(recb.builders))
	nodes := make([]*RpcEndpoint, len(recb.builders))
	mutators := make([]Mutator, len(recb.builders))
	for i := range recb.builders {
		func(i int, root context.Context) {
			builder := recb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RpcEndpointMutation)
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
					_, err = mutators[i+1].Mutate(root, recb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, recb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, recb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (recb *RpcEndpointCreateBulk) SaveX(ctx context.Context) []*RpcEndpoint {
	v, err := recb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (recb *RpcEndpointCreateBulk) Exec(ctx context.Context) error {
	_, err := recb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (recb *RpcEndpointCreateBulk) ExecX(ctx context.Context) {
	if err := recb.Exec(ctx); err != nil {
		panic(err)
	}
}