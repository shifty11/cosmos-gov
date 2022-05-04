// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-gov/ent/migrationinfo"
)

// MigrationInfoCreate is the builder for creating a MigrationInfo entity.
type MigrationInfoCreate struct {
	config
	mutation *MigrationInfoMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (mic *MigrationInfoCreate) SetCreateTime(t time.Time) *MigrationInfoCreate {
	mic.mutation.SetCreateTime(t)
	return mic
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (mic *MigrationInfoCreate) SetNillableCreateTime(t *time.Time) *MigrationInfoCreate {
	if t != nil {
		mic.SetCreateTime(*t)
	}
	return mic
}

// SetUpdateTime sets the "update_time" field.
func (mic *MigrationInfoCreate) SetUpdateTime(t time.Time) *MigrationInfoCreate {
	mic.mutation.SetUpdateTime(t)
	return mic
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (mic *MigrationInfoCreate) SetNillableUpdateTime(t *time.Time) *MigrationInfoCreate {
	if t != nil {
		mic.SetUpdateTime(*t)
	}
	return mic
}

// SetIsMigrated sets the "is_migrated" field.
func (mic *MigrationInfoCreate) SetIsMigrated(b bool) *MigrationInfoCreate {
	mic.mutation.SetIsMigrated(b)
	return mic
}

// SetNillableIsMigrated sets the "is_migrated" field if the given value is not nil.
func (mic *MigrationInfoCreate) SetNillableIsMigrated(b *bool) *MigrationInfoCreate {
	if b != nil {
		mic.SetIsMigrated(*b)
	}
	return mic
}

// Mutation returns the MigrationInfoMutation object of the builder.
func (mic *MigrationInfoCreate) Mutation() *MigrationInfoMutation {
	return mic.mutation
}

// Save creates the MigrationInfo in the database.
func (mic *MigrationInfoCreate) Save(ctx context.Context) (*MigrationInfo, error) {
	var (
		err  error
		node *MigrationInfo
	)
	mic.defaults()
	if len(mic.hooks) == 0 {
		if err = mic.check(); err != nil {
			return nil, err
		}
		node, err = mic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MigrationInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mic.check(); err != nil {
				return nil, err
			}
			mic.mutation = mutation
			if node, err = mic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mic.hooks) - 1; i >= 0; i-- {
			if mic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mic *MigrationInfoCreate) SaveX(ctx context.Context) *MigrationInfo {
	v, err := mic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mic *MigrationInfoCreate) Exec(ctx context.Context) error {
	_, err := mic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mic *MigrationInfoCreate) ExecX(ctx context.Context) {
	if err := mic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mic *MigrationInfoCreate) defaults() {
	if _, ok := mic.mutation.CreateTime(); !ok {
		v := migrationinfo.DefaultCreateTime()
		mic.mutation.SetCreateTime(v)
	}
	if _, ok := mic.mutation.UpdateTime(); !ok {
		v := migrationinfo.DefaultUpdateTime()
		mic.mutation.SetUpdateTime(v)
	}
	if _, ok := mic.mutation.IsMigrated(); !ok {
		v := migrationinfo.DefaultIsMigrated
		mic.mutation.SetIsMigrated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mic *MigrationInfoCreate) check() error {
	if _, ok := mic.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "MigrationInfo.create_time"`)}
	}
	if _, ok := mic.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "MigrationInfo.update_time"`)}
	}
	if _, ok := mic.mutation.IsMigrated(); !ok {
		return &ValidationError{Name: "is_migrated", err: errors.New(`ent: missing required field "MigrationInfo.is_migrated"`)}
	}
	return nil
}

func (mic *MigrationInfoCreate) sqlSave(ctx context.Context) (*MigrationInfo, error) {
	_node, _spec := mic.createSpec()
	if err := sqlgraph.CreateNode(ctx, mic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mic *MigrationInfoCreate) createSpec() (*MigrationInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &MigrationInfo{config: mic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: migrationinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: migrationinfo.FieldID,
			},
		}
	)
	if value, ok := mic.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: migrationinfo.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := mic.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: migrationinfo.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := mic.mutation.IsMigrated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: migrationinfo.FieldIsMigrated,
		})
		_node.IsMigrated = value
	}
	return _node, _spec
}

// MigrationInfoCreateBulk is the builder for creating many MigrationInfo entities in bulk.
type MigrationInfoCreateBulk struct {
	config
	builders []*MigrationInfoCreate
}

// Save creates the MigrationInfo entities in the database.
func (micb *MigrationInfoCreateBulk) Save(ctx context.Context) ([]*MigrationInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(micb.builders))
	nodes := make([]*MigrationInfo, len(micb.builders))
	mutators := make([]Mutator, len(micb.builders))
	for i := range micb.builders {
		func(i int, root context.Context) {
			builder := micb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MigrationInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, micb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, micb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, micb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (micb *MigrationInfoCreateBulk) SaveX(ctx context.Context) []*MigrationInfo {
	v, err := micb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (micb *MigrationInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := micb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (micb *MigrationInfoCreateBulk) ExecX(ctx context.Context) {
	if err := micb.Exec(ctx); err != nil {
		panic(err)
	}
}
