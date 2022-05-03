// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-gov/ent/lenschaininfo"
)

// LensChainInfoCreate is the builder for creating a LensChainInfo entity.
type LensChainInfoCreate struct {
	config
	mutation *LensChainInfoMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (lcic *LensChainInfoCreate) SetCreateTime(t time.Time) *LensChainInfoCreate {
	lcic.mutation.SetCreateTime(t)
	return lcic
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (lcic *LensChainInfoCreate) SetNillableCreateTime(t *time.Time) *LensChainInfoCreate {
	if t != nil {
		lcic.SetCreateTime(*t)
	}
	return lcic
}

// SetUpdatedTime sets the "updated_time" field.
func (lcic *LensChainInfoCreate) SetUpdatedTime(t time.Time) *LensChainInfoCreate {
	lcic.mutation.SetUpdatedTime(t)
	return lcic
}

// SetNillableUpdatedTime sets the "updated_time" field if the given value is not nil.
func (lcic *LensChainInfoCreate) SetNillableUpdatedTime(t *time.Time) *LensChainInfoCreate {
	if t != nil {
		lcic.SetUpdatedTime(*t)
	}
	return lcic
}

// SetCreatedAt sets the "created_at" field.
func (lcic *LensChainInfoCreate) SetCreatedAt(t time.Time) *LensChainInfoCreate {
	lcic.mutation.SetCreatedAt(t)
	return lcic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lcic *LensChainInfoCreate) SetNillableCreatedAt(t *time.Time) *LensChainInfoCreate {
	if t != nil {
		lcic.SetCreatedAt(*t)
	}
	return lcic
}

// SetUpdatedAt sets the "updated_at" field.
func (lcic *LensChainInfoCreate) SetUpdatedAt(t time.Time) *LensChainInfoCreate {
	lcic.mutation.SetUpdatedAt(t)
	return lcic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lcic *LensChainInfoCreate) SetNillableUpdatedAt(t *time.Time) *LensChainInfoCreate {
	if t != nil {
		lcic.SetUpdatedAt(*t)
	}
	return lcic
}

// SetName sets the "name" field.
func (lcic *LensChainInfoCreate) SetName(s string) *LensChainInfoCreate {
	lcic.mutation.SetName(s)
	return lcic
}

// SetCntErrors sets the "cnt_errors" field.
func (lcic *LensChainInfoCreate) SetCntErrors(i int) *LensChainInfoCreate {
	lcic.mutation.SetCntErrors(i)
	return lcic
}

// Mutation returns the LensChainInfoMutation object of the builder.
func (lcic *LensChainInfoCreate) Mutation() *LensChainInfoMutation {
	return lcic.mutation
}

// Save creates the LensChainInfo in the database.
func (lcic *LensChainInfoCreate) Save(ctx context.Context) (*LensChainInfo, error) {
	var (
		err  error
		node *LensChainInfo
	)
	lcic.defaults()
	if len(lcic.hooks) == 0 {
		if err = lcic.check(); err != nil {
			return nil, err
		}
		node, err = lcic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LensChainInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lcic.check(); err != nil {
				return nil, err
			}
			lcic.mutation = mutation
			if node, err = lcic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lcic.hooks) - 1; i >= 0; i-- {
			if lcic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lcic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lcic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lcic *LensChainInfoCreate) SaveX(ctx context.Context) *LensChainInfo {
	v, err := lcic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcic *LensChainInfoCreate) Exec(ctx context.Context) error {
	_, err := lcic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcic *LensChainInfoCreate) ExecX(ctx context.Context) {
	if err := lcic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lcic *LensChainInfoCreate) defaults() {
	if _, ok := lcic.mutation.CreatedAt(); !ok {
		v := lenschaininfo.DefaultCreatedAt()
		lcic.mutation.SetCreatedAt(v)
	}
	if _, ok := lcic.mutation.UpdatedAt(); !ok {
		v := lenschaininfo.DefaultUpdatedAt()
		lcic.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lcic *LensChainInfoCreate) check() error {
	if _, ok := lcic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "LensChainInfo.created_at"`)}
	}
	if _, ok := lcic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "LensChainInfo.updated_at"`)}
	}
	if _, ok := lcic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "LensChainInfo.name"`)}
	}
	if _, ok := lcic.mutation.CntErrors(); !ok {
		return &ValidationError{Name: "cnt_errors", err: errors.New(`ent: missing required field "LensChainInfo.cnt_errors"`)}
	}
	return nil
}

func (lcic *LensChainInfoCreate) sqlSave(ctx context.Context) (*LensChainInfo, error) {
	_node, _spec := lcic.createSpec()
	if err := sqlgraph.CreateNode(ctx, lcic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (lcic *LensChainInfoCreate) createSpec() (*LensChainInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &LensChainInfo{config: lcic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: lenschaininfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lenschaininfo.FieldID,
			},
		}
	)
	if value, ok := lcic.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lenschaininfo.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := lcic.mutation.UpdatedTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lenschaininfo.FieldUpdatedTime,
		})
		_node.UpdatedTime = value
	}
	if value, ok := lcic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lenschaininfo.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := lcic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lenschaininfo.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := lcic.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lenschaininfo.FieldName,
		})
		_node.Name = value
	}
	if value, ok := lcic.mutation.CntErrors(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: lenschaininfo.FieldCntErrors,
		})
		_node.CntErrors = value
	}
	return _node, _spec
}

// LensChainInfoCreateBulk is the builder for creating many LensChainInfo entities in bulk.
type LensChainInfoCreateBulk struct {
	config
	builders []*LensChainInfoCreate
}

// Save creates the LensChainInfo entities in the database.
func (lcicb *LensChainInfoCreateBulk) Save(ctx context.Context) ([]*LensChainInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcicb.builders))
	nodes := make([]*LensChainInfo, len(lcicb.builders))
	mutators := make([]Mutator, len(lcicb.builders))
	for i := range lcicb.builders {
		func(i int, root context.Context) {
			builder := lcicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LensChainInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, lcicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcicb *LensChainInfoCreateBulk) SaveX(ctx context.Context) []*LensChainInfo {
	v, err := lcicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcicb *LensChainInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := lcicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcicb *LensChainInfoCreateBulk) ExecX(ctx context.Context) {
	if err := lcicb.Exec(ctx); err != nil {
		panic(err)
	}
}
