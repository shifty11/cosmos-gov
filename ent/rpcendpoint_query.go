// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/predicate"
	"github.com/shifty11/cosmos-gov/ent/rpcendpoint"
)

// RpcEndpointQuery is the builder for querying RpcEndpoint entities.
type RpcEndpointQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.RpcEndpoint
	// eager-loading edges.
	withChain *ChainQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RpcEndpointQuery builder.
func (req *RpcEndpointQuery) Where(ps ...predicate.RpcEndpoint) *RpcEndpointQuery {
	req.predicates = append(req.predicates, ps...)
	return req
}

// Limit adds a limit step to the query.
func (req *RpcEndpointQuery) Limit(limit int) *RpcEndpointQuery {
	req.limit = &limit
	return req
}

// Offset adds an offset step to the query.
func (req *RpcEndpointQuery) Offset(offset int) *RpcEndpointQuery {
	req.offset = &offset
	return req
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (req *RpcEndpointQuery) Unique(unique bool) *RpcEndpointQuery {
	req.unique = &unique
	return req
}

// Order adds an order step to the query.
func (req *RpcEndpointQuery) Order(o ...OrderFunc) *RpcEndpointQuery {
	req.order = append(req.order, o...)
	return req
}

// QueryChain chains the current query on the "chain" edge.
func (req *RpcEndpointQuery) QueryChain() *ChainQuery {
	query := &ChainQuery{config: req.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := req.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := req.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rpcendpoint.Table, rpcendpoint.FieldID, selector),
			sqlgraph.To(chain.Table, chain.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, rpcendpoint.ChainTable, rpcendpoint.ChainColumn),
		)
		fromU = sqlgraph.SetNeighbors(req.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RpcEndpoint entity from the query.
// Returns a *NotFoundError when no RpcEndpoint was found.
func (req *RpcEndpointQuery) First(ctx context.Context) (*RpcEndpoint, error) {
	nodes, err := req.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rpcendpoint.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (req *RpcEndpointQuery) FirstX(ctx context.Context) *RpcEndpoint {
	node, err := req.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RpcEndpoint ID from the query.
// Returns a *NotFoundError when no RpcEndpoint ID was found.
func (req *RpcEndpointQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = req.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rpcendpoint.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (req *RpcEndpointQuery) FirstIDX(ctx context.Context) int {
	id, err := req.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RpcEndpoint entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RpcEndpoint entity is found.
// Returns a *NotFoundError when no RpcEndpoint entities are found.
func (req *RpcEndpointQuery) Only(ctx context.Context) (*RpcEndpoint, error) {
	nodes, err := req.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rpcendpoint.Label}
	default:
		return nil, &NotSingularError{rpcendpoint.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (req *RpcEndpointQuery) OnlyX(ctx context.Context) *RpcEndpoint {
	node, err := req.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RpcEndpoint ID in the query.
// Returns a *NotSingularError when more than one RpcEndpoint ID is found.
// Returns a *NotFoundError when no entities are found.
func (req *RpcEndpointQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = req.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rpcendpoint.Label}
	default:
		err = &NotSingularError{rpcendpoint.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (req *RpcEndpointQuery) OnlyIDX(ctx context.Context) int {
	id, err := req.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RpcEndpoints.
func (req *RpcEndpointQuery) All(ctx context.Context) ([]*RpcEndpoint, error) {
	if err := req.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return req.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (req *RpcEndpointQuery) AllX(ctx context.Context) []*RpcEndpoint {
	nodes, err := req.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RpcEndpoint IDs.
func (req *RpcEndpointQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := req.Select(rpcendpoint.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (req *RpcEndpointQuery) IDsX(ctx context.Context) []int {
	ids, err := req.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (req *RpcEndpointQuery) Count(ctx context.Context) (int, error) {
	if err := req.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return req.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (req *RpcEndpointQuery) CountX(ctx context.Context) int {
	count, err := req.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (req *RpcEndpointQuery) Exist(ctx context.Context) (bool, error) {
	if err := req.prepareQuery(ctx); err != nil {
		return false, err
	}
	return req.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (req *RpcEndpointQuery) ExistX(ctx context.Context) bool {
	exist, err := req.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RpcEndpointQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (req *RpcEndpointQuery) Clone() *RpcEndpointQuery {
	if req == nil {
		return nil
	}
	return &RpcEndpointQuery{
		config:     req.config,
		limit:      req.limit,
		offset:     req.offset,
		order:      append([]OrderFunc{}, req.order...),
		predicates: append([]predicate.RpcEndpoint{}, req.predicates...),
		withChain:  req.withChain.Clone(),
		// clone intermediate query.
		sql:    req.sql.Clone(),
		path:   req.path,
		unique: req.unique,
	}
}

// WithChain tells the query-builder to eager-load the nodes that are connected to
// the "chain" edge. The optional arguments are used to configure the query builder of the edge.
func (req *RpcEndpointQuery) WithChain(opts ...func(*ChainQuery)) *RpcEndpointQuery {
	query := &ChainQuery{config: req.config}
	for _, opt := range opts {
		opt(query)
	}
	req.withChain = query
	return req
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RpcEndpoint.Query().
//		GroupBy(rpcendpoint.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (req *RpcEndpointQuery) GroupBy(field string, fields ...string) *RpcEndpointGroupBy {
	group := &RpcEndpointGroupBy{config: req.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := req.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return req.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.RpcEndpoint.Query().
//		Select(rpcendpoint.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (req *RpcEndpointQuery) Select(fields ...string) *RpcEndpointSelect {
	req.fields = append(req.fields, fields...)
	return &RpcEndpointSelect{RpcEndpointQuery: req}
}

func (req *RpcEndpointQuery) prepareQuery(ctx context.Context) error {
	for _, f := range req.fields {
		if !rpcendpoint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if req.path != nil {
		prev, err := req.path(ctx)
		if err != nil {
			return err
		}
		req.sql = prev
	}
	return nil
}

func (req *RpcEndpointQuery) sqlAll(ctx context.Context) ([]*RpcEndpoint, error) {
	var (
		nodes       = []*RpcEndpoint{}
		withFKs     = req.withFKs
		_spec       = req.querySpec()
		loadedTypes = [1]bool{
			req.withChain != nil,
		}
	)
	if req.withChain != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, rpcendpoint.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &RpcEndpoint{config: req.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, req.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := req.withChain; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*RpcEndpoint)
		for i := range nodes {
			if nodes[i].chain_rpc_endpoints == nil {
				continue
			}
			fk := *nodes[i].chain_rpc_endpoints
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(chain.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "chain_rpc_endpoints" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Chain = n
			}
		}
	}

	return nodes, nil
}

func (req *RpcEndpointQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := req.querySpec()
	_spec.Node.Columns = req.fields
	if len(req.fields) > 0 {
		_spec.Unique = req.unique != nil && *req.unique
	}
	return sqlgraph.CountNodes(ctx, req.driver, _spec)
}

func (req *RpcEndpointQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := req.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (req *RpcEndpointQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rpcendpoint.Table,
			Columns: rpcendpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rpcendpoint.FieldID,
			},
		},
		From:   req.sql,
		Unique: true,
	}
	if unique := req.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := req.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rpcendpoint.FieldID)
		for i := range fields {
			if fields[i] != rpcendpoint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := req.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := req.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := req.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := req.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (req *RpcEndpointQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(req.driver.Dialect())
	t1 := builder.Table(rpcendpoint.Table)
	columns := req.fields
	if len(columns) == 0 {
		columns = rpcendpoint.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if req.sql != nil {
		selector = req.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if req.unique != nil && *req.unique {
		selector.Distinct()
	}
	for _, p := range req.predicates {
		p(selector)
	}
	for _, p := range req.order {
		p(selector)
	}
	if offset := req.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := req.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RpcEndpointGroupBy is the group-by builder for RpcEndpoint entities.
type RpcEndpointGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (regb *RpcEndpointGroupBy) Aggregate(fns ...AggregateFunc) *RpcEndpointGroupBy {
	regb.fns = append(regb.fns, fns...)
	return regb
}

// Scan applies the group-by query and scans the result into the given value.
func (regb *RpcEndpointGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := regb.path(ctx)
	if err != nil {
		return err
	}
	regb.sql = query
	return regb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (regb *RpcEndpointGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := regb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (regb *RpcEndpointGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(regb.fields) > 1 {
		return nil, errors.New("ent: RpcEndpointGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := regb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (regb *RpcEndpointGroupBy) StringsX(ctx context.Context) []string {
	v, err := regb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (regb *RpcEndpointGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = regb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rpcendpoint.Label}
	default:
		err = fmt.Errorf("ent: RpcEndpointGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (regb *RpcEndpointGroupBy) StringX(ctx context.Context) string {
	v, err := regb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (regb *RpcEndpointGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(regb.fields) > 1 {
		return nil, errors.New("ent: RpcEndpointGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := regb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (regb *RpcEndpointGroupBy) IntsX(ctx context.Context) []int {
	v, err := regb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (regb *RpcEndpointGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = regb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rpcendpoint.Label}
	default:
		err = fmt.Errorf("ent: RpcEndpointGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (regb *RpcEndpointGroupBy) IntX(ctx context.Context) int {
	v, err := regb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (regb *RpcEndpointGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(regb.fields) > 1 {
		return nil, errors.New("ent: RpcEndpointGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := regb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (regb *RpcEndpointGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := regb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (regb *RpcEndpointGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = regb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rpcendpoint.Label}
	default:
		err = fmt.Errorf("ent: RpcEndpointGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (regb *RpcEndpointGroupBy) Float64X(ctx context.Context) float64 {
	v, err := regb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (regb *RpcEndpointGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(regb.fields) > 1 {
		return nil, errors.New("ent: RpcEndpointGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := regb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (regb *RpcEndpointGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := regb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (regb *RpcEndpointGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = regb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rpcendpoint.Label}
	default:
		err = fmt.Errorf("ent: RpcEndpointGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (regb *RpcEndpointGroupBy) BoolX(ctx context.Context) bool {
	v, err := regb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (regb *RpcEndpointGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range regb.fields {
		if !rpcendpoint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := regb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := regb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (regb *RpcEndpointGroupBy) sqlQuery() *sql.Selector {
	selector := regb.sql.Select()
	aggregation := make([]string, 0, len(regb.fns))
	for _, fn := range regb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(regb.fields)+len(regb.fns))
		for _, f := range regb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(regb.fields...)...)
}

// RpcEndpointSelect is the builder for selecting fields of RpcEndpoint entities.
type RpcEndpointSelect struct {
	*RpcEndpointQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (res *RpcEndpointSelect) Scan(ctx context.Context, v interface{}) error {
	if err := res.prepareQuery(ctx); err != nil {
		return err
	}
	res.sql = res.RpcEndpointQuery.sqlQuery(ctx)
	return res.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (res *RpcEndpointSelect) ScanX(ctx context.Context, v interface{}) {
	if err := res.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (res *RpcEndpointSelect) Strings(ctx context.Context) ([]string, error) {
	if len(res.fields) > 1 {
		return nil, errors.New("ent: RpcEndpointSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := res.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (res *RpcEndpointSelect) StringsX(ctx context.Context) []string {
	v, err := res.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (res *RpcEndpointSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = res.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rpcendpoint.Label}
	default:
		err = fmt.Errorf("ent: RpcEndpointSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (res *RpcEndpointSelect) StringX(ctx context.Context) string {
	v, err := res.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (res *RpcEndpointSelect) Ints(ctx context.Context) ([]int, error) {
	if len(res.fields) > 1 {
		return nil, errors.New("ent: RpcEndpointSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := res.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (res *RpcEndpointSelect) IntsX(ctx context.Context) []int {
	v, err := res.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (res *RpcEndpointSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = res.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rpcendpoint.Label}
	default:
		err = fmt.Errorf("ent: RpcEndpointSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (res *RpcEndpointSelect) IntX(ctx context.Context) int {
	v, err := res.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (res *RpcEndpointSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(res.fields) > 1 {
		return nil, errors.New("ent: RpcEndpointSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := res.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (res *RpcEndpointSelect) Float64sX(ctx context.Context) []float64 {
	v, err := res.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (res *RpcEndpointSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = res.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rpcendpoint.Label}
	default:
		err = fmt.Errorf("ent: RpcEndpointSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (res *RpcEndpointSelect) Float64X(ctx context.Context) float64 {
	v, err := res.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (res *RpcEndpointSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(res.fields) > 1 {
		return nil, errors.New("ent: RpcEndpointSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := res.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (res *RpcEndpointSelect) BoolsX(ctx context.Context) []bool {
	v, err := res.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (res *RpcEndpointSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = res.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rpcendpoint.Label}
	default:
		err = fmt.Errorf("ent: RpcEndpointSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (res *RpcEndpointSelect) BoolX(ctx context.Context) bool {
	v, err := res.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (res *RpcEndpointSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := res.sql.Query()
	if err := res.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
