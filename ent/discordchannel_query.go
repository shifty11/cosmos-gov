// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/discordchannel"
	"github.com/shifty11/cosmos-gov/ent/predicate"
	"github.com/shifty11/cosmos-gov/ent/user"
)

// DiscordChannelQuery is the builder for querying DiscordChannel entities.
type DiscordChannelQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DiscordChannel
	// eager-loading edges.
	withUser   *UserQuery
	withChains *ChainQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DiscordChannelQuery builder.
func (dcq *DiscordChannelQuery) Where(ps ...predicate.DiscordChannel) *DiscordChannelQuery {
	dcq.predicates = append(dcq.predicates, ps...)
	return dcq
}

// Limit adds a limit step to the query.
func (dcq *DiscordChannelQuery) Limit(limit int) *DiscordChannelQuery {
	dcq.limit = &limit
	return dcq
}

// Offset adds an offset step to the query.
func (dcq *DiscordChannelQuery) Offset(offset int) *DiscordChannelQuery {
	dcq.offset = &offset
	return dcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dcq *DiscordChannelQuery) Unique(unique bool) *DiscordChannelQuery {
	dcq.unique = &unique
	return dcq
}

// Order adds an order step to the query.
func (dcq *DiscordChannelQuery) Order(o ...OrderFunc) *DiscordChannelQuery {
	dcq.order = append(dcq.order, o...)
	return dcq
}

// QueryUser chains the current query on the "user" edge.
func (dcq *DiscordChannelQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: dcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(discordchannel.Table, discordchannel.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, discordchannel.UserTable, discordchannel.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(dcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChains chains the current query on the "chains" edge.
func (dcq *DiscordChannelQuery) QueryChains() *ChainQuery {
	query := &ChainQuery{config: dcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(discordchannel.Table, discordchannel.FieldID, selector),
			sqlgraph.To(chain.Table, chain.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, discordchannel.ChainsTable, discordchannel.ChainsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DiscordChannel entity from the query.
// Returns a *NotFoundError when no DiscordChannel was found.
func (dcq *DiscordChannelQuery) First(ctx context.Context) (*DiscordChannel, error) {
	nodes, err := dcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{discordchannel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dcq *DiscordChannelQuery) FirstX(ctx context.Context) *DiscordChannel {
	node, err := dcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DiscordChannel ID from the query.
// Returns a *NotFoundError when no DiscordChannel ID was found.
func (dcq *DiscordChannelQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = dcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{discordchannel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dcq *DiscordChannelQuery) FirstIDX(ctx context.Context) int64 {
	id, err := dcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DiscordChannel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DiscordChannel entity is found.
// Returns a *NotFoundError when no DiscordChannel entities are found.
func (dcq *DiscordChannelQuery) Only(ctx context.Context) (*DiscordChannel, error) {
	nodes, err := dcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{discordchannel.Label}
	default:
		return nil, &NotSingularError{discordchannel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dcq *DiscordChannelQuery) OnlyX(ctx context.Context) *DiscordChannel {
	node, err := dcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DiscordChannel ID in the query.
// Returns a *NotSingularError when more than one DiscordChannel ID is found.
// Returns a *NotFoundError when no entities are found.
func (dcq *DiscordChannelQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = dcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{discordchannel.Label}
	default:
		err = &NotSingularError{discordchannel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dcq *DiscordChannelQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := dcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DiscordChannels.
func (dcq *DiscordChannelQuery) All(ctx context.Context) ([]*DiscordChannel, error) {
	if err := dcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dcq *DiscordChannelQuery) AllX(ctx context.Context) []*DiscordChannel {
	nodes, err := dcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DiscordChannel IDs.
func (dcq *DiscordChannelQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := dcq.Select(discordchannel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dcq *DiscordChannelQuery) IDsX(ctx context.Context) []int64 {
	ids, err := dcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dcq *DiscordChannelQuery) Count(ctx context.Context) (int, error) {
	if err := dcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dcq *DiscordChannelQuery) CountX(ctx context.Context) int {
	count, err := dcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dcq *DiscordChannelQuery) Exist(ctx context.Context) (bool, error) {
	if err := dcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dcq *DiscordChannelQuery) ExistX(ctx context.Context) bool {
	exist, err := dcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DiscordChannelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dcq *DiscordChannelQuery) Clone() *DiscordChannelQuery {
	if dcq == nil {
		return nil
	}
	return &DiscordChannelQuery{
		config:     dcq.config,
		limit:      dcq.limit,
		offset:     dcq.offset,
		order:      append([]OrderFunc{}, dcq.order...),
		predicates: append([]predicate.DiscordChannel{}, dcq.predicates...),
		withUser:   dcq.withUser.Clone(),
		withChains: dcq.withChains.Clone(),
		// clone intermediate query.
		sql:    dcq.sql.Clone(),
		path:   dcq.path,
		unique: dcq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (dcq *DiscordChannelQuery) WithUser(opts ...func(*UserQuery)) *DiscordChannelQuery {
	query := &UserQuery{config: dcq.config}
	for _, opt := range opts {
		opt(query)
	}
	dcq.withUser = query
	return dcq
}

// WithChains tells the query-builder to eager-load the nodes that are connected to
// the "chains" edge. The optional arguments are used to configure the query builder of the edge.
func (dcq *DiscordChannelQuery) WithChains(opts ...func(*ChainQuery)) *DiscordChannelQuery {
	query := &ChainQuery{config: dcq.config}
	for _, opt := range opts {
		opt(query)
	}
	dcq.withChains = query
	return dcq
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
//	client.DiscordChannel.Query().
//		GroupBy(discordchannel.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dcq *DiscordChannelQuery) GroupBy(field string, fields ...string) *DiscordChannelGroupBy {
	group := &DiscordChannelGroupBy{config: dcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dcq.sqlQuery(ctx), nil
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
//	client.DiscordChannel.Query().
//		Select(discordchannel.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (dcq *DiscordChannelQuery) Select(fields ...string) *DiscordChannelSelect {
	dcq.fields = append(dcq.fields, fields...)
	return &DiscordChannelSelect{DiscordChannelQuery: dcq}
}

func (dcq *DiscordChannelQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dcq.fields {
		if !discordchannel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dcq.path != nil {
		prev, err := dcq.path(ctx)
		if err != nil {
			return err
		}
		dcq.sql = prev
	}
	return nil
}

func (dcq *DiscordChannelQuery) sqlAll(ctx context.Context) ([]*DiscordChannel, error) {
	var (
		nodes       = []*DiscordChannel{}
		withFKs     = dcq.withFKs
		_spec       = dcq.querySpec()
		loadedTypes = [2]bool{
			dcq.withUser != nil,
			dcq.withChains != nil,
		}
	)
	if dcq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, discordchannel.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DiscordChannel{config: dcq.config}
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
	if err := sqlgraph.QueryNodes(ctx, dcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dcq.withUser; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*DiscordChannel)
		for i := range nodes {
			if nodes[i].discord_channel_user == nil {
				continue
			}
			fk := *nodes[i].discord_channel_user
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "discord_channel_user" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := dcq.withChains; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int64]*DiscordChannel, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Chains = []*Chain{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*DiscordChannel)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   discordchannel.ChainsTable,
				Columns: discordchannel.ChainsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(discordchannel.ChainsPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := eout.Int64
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, dcq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "chains": %w`, err)
		}
		query.Where(chain.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "chains" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Chains = append(nodes[i].Edges.Chains, n)
			}
		}
	}

	return nodes, nil
}

func (dcq *DiscordChannelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dcq.querySpec()
	_spec.Node.Columns = dcq.fields
	if len(dcq.fields) > 0 {
		_spec.Unique = dcq.unique != nil && *dcq.unique
	}
	return sqlgraph.CountNodes(ctx, dcq.driver, _spec)
}

func (dcq *DiscordChannelQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dcq *DiscordChannelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   discordchannel.Table,
			Columns: discordchannel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: discordchannel.FieldID,
			},
		},
		From:   dcq.sql,
		Unique: true,
	}
	if unique := dcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, discordchannel.FieldID)
		for i := range fields {
			if fields[i] != discordchannel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dcq *DiscordChannelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dcq.driver.Dialect())
	t1 := builder.Table(discordchannel.Table)
	columns := dcq.fields
	if len(columns) == 0 {
		columns = discordchannel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dcq.sql != nil {
		selector = dcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dcq.unique != nil && *dcq.unique {
		selector.Distinct()
	}
	for _, p := range dcq.predicates {
		p(selector)
	}
	for _, p := range dcq.order {
		p(selector)
	}
	if offset := dcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DiscordChannelGroupBy is the group-by builder for DiscordChannel entities.
type DiscordChannelGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dcgb *DiscordChannelGroupBy) Aggregate(fns ...AggregateFunc) *DiscordChannelGroupBy {
	dcgb.fns = append(dcgb.fns, fns...)
	return dcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dcgb *DiscordChannelGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dcgb.path(ctx)
	if err != nil {
		return err
	}
	dcgb.sql = query
	return dcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dcgb *DiscordChannelGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dcgb *DiscordChannelGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dcgb.fields) > 1 {
		return nil, errors.New("ent: DiscordChannelGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dcgb *DiscordChannelGroupBy) StringsX(ctx context.Context) []string {
	v, err := dcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dcgb *DiscordChannelGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordchannel.Label}
	default:
		err = fmt.Errorf("ent: DiscordChannelGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dcgb *DiscordChannelGroupBy) StringX(ctx context.Context) string {
	v, err := dcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dcgb *DiscordChannelGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dcgb.fields) > 1 {
		return nil, errors.New("ent: DiscordChannelGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dcgb *DiscordChannelGroupBy) IntsX(ctx context.Context) []int {
	v, err := dcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dcgb *DiscordChannelGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordchannel.Label}
	default:
		err = fmt.Errorf("ent: DiscordChannelGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dcgb *DiscordChannelGroupBy) IntX(ctx context.Context) int {
	v, err := dcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dcgb *DiscordChannelGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dcgb.fields) > 1 {
		return nil, errors.New("ent: DiscordChannelGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dcgb *DiscordChannelGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dcgb *DiscordChannelGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordchannel.Label}
	default:
		err = fmt.Errorf("ent: DiscordChannelGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dcgb *DiscordChannelGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dcgb *DiscordChannelGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dcgb.fields) > 1 {
		return nil, errors.New("ent: DiscordChannelGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dcgb *DiscordChannelGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dcgb *DiscordChannelGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordchannel.Label}
	default:
		err = fmt.Errorf("ent: DiscordChannelGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dcgb *DiscordChannelGroupBy) BoolX(ctx context.Context) bool {
	v, err := dcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dcgb *DiscordChannelGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dcgb.fields {
		if !discordchannel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dcgb *DiscordChannelGroupBy) sqlQuery() *sql.Selector {
	selector := dcgb.sql.Select()
	aggregation := make([]string, 0, len(dcgb.fns))
	for _, fn := range dcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dcgb.fields)+len(dcgb.fns))
		for _, f := range dcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dcgb.fields...)...)
}

// DiscordChannelSelect is the builder for selecting fields of DiscordChannel entities.
type DiscordChannelSelect struct {
	*DiscordChannelQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dcs *DiscordChannelSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dcs.prepareQuery(ctx); err != nil {
		return err
	}
	dcs.sql = dcs.DiscordChannelQuery.sqlQuery(ctx)
	return dcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dcs *DiscordChannelSelect) ScanX(ctx context.Context, v interface{}) {
	if err := dcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (dcs *DiscordChannelSelect) Strings(ctx context.Context) ([]string, error) {
	if len(dcs.fields) > 1 {
		return nil, errors.New("ent: DiscordChannelSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := dcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dcs *DiscordChannelSelect) StringsX(ctx context.Context) []string {
	v, err := dcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (dcs *DiscordChannelSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordchannel.Label}
	default:
		err = fmt.Errorf("ent: DiscordChannelSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dcs *DiscordChannelSelect) StringX(ctx context.Context) string {
	v, err := dcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (dcs *DiscordChannelSelect) Ints(ctx context.Context) ([]int, error) {
	if len(dcs.fields) > 1 {
		return nil, errors.New("ent: DiscordChannelSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := dcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dcs *DiscordChannelSelect) IntsX(ctx context.Context) []int {
	v, err := dcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (dcs *DiscordChannelSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordchannel.Label}
	default:
		err = fmt.Errorf("ent: DiscordChannelSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dcs *DiscordChannelSelect) IntX(ctx context.Context) int {
	v, err := dcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (dcs *DiscordChannelSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(dcs.fields) > 1 {
		return nil, errors.New("ent: DiscordChannelSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := dcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dcs *DiscordChannelSelect) Float64sX(ctx context.Context) []float64 {
	v, err := dcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (dcs *DiscordChannelSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordchannel.Label}
	default:
		err = fmt.Errorf("ent: DiscordChannelSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dcs *DiscordChannelSelect) Float64X(ctx context.Context) float64 {
	v, err := dcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (dcs *DiscordChannelSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(dcs.fields) > 1 {
		return nil, errors.New("ent: DiscordChannelSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := dcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dcs *DiscordChannelSelect) BoolsX(ctx context.Context) []bool {
	v, err := dcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (dcs *DiscordChannelSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordchannel.Label}
	default:
		err = fmt.Errorf("ent: DiscordChannelSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dcs *DiscordChannelSelect) BoolX(ctx context.Context) bool {
	v, err := dcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dcs *DiscordChannelSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dcs.sql.Query()
	if err := dcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
