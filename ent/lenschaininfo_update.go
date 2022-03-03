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
	"github.com/shifty11/cosmos-gov/ent/lenschaininfo"
	"github.com/shifty11/cosmos-gov/ent/predicate"
)

// LensChainInfoUpdate is the builder for updating LensChainInfo entities.
type LensChainInfoUpdate struct {
	config
	hooks    []Hook
	mutation *LensChainInfoMutation
}

// Where appends a list predicates to the LensChainInfoUpdate builder.
func (lciu *LensChainInfoUpdate) Where(ps ...predicate.LensChainInfo) *LensChainInfoUpdate {
	lciu.mutation.Where(ps...)
	return lciu
}

// SetUpdatedAt sets the "updated_at" field.
func (lciu *LensChainInfoUpdate) SetUpdatedAt(t time.Time) *LensChainInfoUpdate {
	lciu.mutation.SetUpdatedAt(t)
	return lciu
}

// SetName sets the "name" field.
func (lciu *LensChainInfoUpdate) SetName(s string) *LensChainInfoUpdate {
	lciu.mutation.SetName(s)
	return lciu
}

// SetCntErrors sets the "cnt_errors" field.
func (lciu *LensChainInfoUpdate) SetCntErrors(i int) *LensChainInfoUpdate {
	lciu.mutation.ResetCntErrors()
	lciu.mutation.SetCntErrors(i)
	return lciu
}

// AddCntErrors adds i to the "cnt_errors" field.
func (lciu *LensChainInfoUpdate) AddCntErrors(i int) *LensChainInfoUpdate {
	lciu.mutation.AddCntErrors(i)
	return lciu
}

// Mutation returns the LensChainInfoMutation object of the builder.
func (lciu *LensChainInfoUpdate) Mutation() *LensChainInfoMutation {
	return lciu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lciu *LensChainInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	lciu.defaults()
	if len(lciu.hooks) == 0 {
		affected, err = lciu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LensChainInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lciu.mutation = mutation
			affected, err = lciu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lciu.hooks) - 1; i >= 0; i-- {
			if lciu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lciu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lciu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lciu *LensChainInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := lciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lciu *LensChainInfoUpdate) Exec(ctx context.Context) error {
	_, err := lciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lciu *LensChainInfoUpdate) ExecX(ctx context.Context) {
	if err := lciu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lciu *LensChainInfoUpdate) defaults() {
	if _, ok := lciu.mutation.UpdatedAt(); !ok {
		v := lenschaininfo.UpdateDefaultUpdatedAt()
		lciu.mutation.SetUpdatedAt(v)
	}
}

func (lciu *LensChainInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lenschaininfo.Table,
			Columns: lenschaininfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lenschaininfo.FieldID,
			},
		},
	}
	if ps := lciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lciu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lenschaininfo.FieldUpdatedAt,
		})
	}
	if value, ok := lciu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lenschaininfo.FieldName,
		})
	}
	if value, ok := lciu.mutation.CntErrors(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: lenschaininfo.FieldCntErrors,
		})
	}
	if value, ok := lciu.mutation.AddedCntErrors(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: lenschaininfo.FieldCntErrors,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lenschaininfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// LensChainInfoUpdateOne is the builder for updating a single LensChainInfo entity.
type LensChainInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LensChainInfoMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (lciuo *LensChainInfoUpdateOne) SetUpdatedAt(t time.Time) *LensChainInfoUpdateOne {
	lciuo.mutation.SetUpdatedAt(t)
	return lciuo
}

// SetName sets the "name" field.
func (lciuo *LensChainInfoUpdateOne) SetName(s string) *LensChainInfoUpdateOne {
	lciuo.mutation.SetName(s)
	return lciuo
}

// SetCntErrors sets the "cnt_errors" field.
func (lciuo *LensChainInfoUpdateOne) SetCntErrors(i int) *LensChainInfoUpdateOne {
	lciuo.mutation.ResetCntErrors()
	lciuo.mutation.SetCntErrors(i)
	return lciuo
}

// AddCntErrors adds i to the "cnt_errors" field.
func (lciuo *LensChainInfoUpdateOne) AddCntErrors(i int) *LensChainInfoUpdateOne {
	lciuo.mutation.AddCntErrors(i)
	return lciuo
}

// Mutation returns the LensChainInfoMutation object of the builder.
func (lciuo *LensChainInfoUpdateOne) Mutation() *LensChainInfoMutation {
	return lciuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lciuo *LensChainInfoUpdateOne) Select(field string, fields ...string) *LensChainInfoUpdateOne {
	lciuo.fields = append([]string{field}, fields...)
	return lciuo
}

// Save executes the query and returns the updated LensChainInfo entity.
func (lciuo *LensChainInfoUpdateOne) Save(ctx context.Context) (*LensChainInfo, error) {
	var (
		err  error
		node *LensChainInfo
	)
	lciuo.defaults()
	if len(lciuo.hooks) == 0 {
		node, err = lciuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LensChainInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lciuo.mutation = mutation
			node, err = lciuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lciuo.hooks) - 1; i >= 0; i-- {
			if lciuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lciuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lciuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (lciuo *LensChainInfoUpdateOne) SaveX(ctx context.Context) *LensChainInfo {
	node, err := lciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lciuo *LensChainInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := lciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lciuo *LensChainInfoUpdateOne) ExecX(ctx context.Context) {
	if err := lciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lciuo *LensChainInfoUpdateOne) defaults() {
	if _, ok := lciuo.mutation.UpdatedAt(); !ok {
		v := lenschaininfo.UpdateDefaultUpdatedAt()
		lciuo.mutation.SetUpdatedAt(v)
	}
}

func (lciuo *LensChainInfoUpdateOne) sqlSave(ctx context.Context) (_node *LensChainInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lenschaininfo.Table,
			Columns: lenschaininfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lenschaininfo.FieldID,
			},
		},
	}
	id, ok := lciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LensChainInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, lenschaininfo.FieldID)
		for _, f := range fields {
			if !lenschaininfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != lenschaininfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lciuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: lenschaininfo.FieldUpdatedAt,
		})
	}
	if value, ok := lciuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lenschaininfo.FieldName,
		})
	}
	if value, ok := lciuo.mutation.CntErrors(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: lenschaininfo.FieldCntErrors,
		})
	}
	if value, ok := lciuo.mutation.AddedCntErrors(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: lenschaininfo.FieldCntErrors,
		})
	}
	_node = &LensChainInfo{config: lciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lenschaininfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}