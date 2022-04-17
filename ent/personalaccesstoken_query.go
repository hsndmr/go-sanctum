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
	"github.com/hsndmr/go-sanctum/ent/personalaccesstoken"
	"github.com/hsndmr/go-sanctum/ent/predicate"
	"github.com/hsndmr/go-sanctum/ent/user"
)

// PersonalAccessTokenQuery is the builder for querying PersonalAccessToken entities.
type PersonalAccessTokenQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PersonalAccessToken
	// eager-loading edges.
	withUser *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PersonalAccessTokenQuery builder.
func (patq *PersonalAccessTokenQuery) Where(ps ...predicate.PersonalAccessToken) *PersonalAccessTokenQuery {
	patq.predicates = append(patq.predicates, ps...)
	return patq
}

// Limit adds a limit step to the query.
func (patq *PersonalAccessTokenQuery) Limit(limit int) *PersonalAccessTokenQuery {
	patq.limit = &limit
	return patq
}

// Offset adds an offset step to the query.
func (patq *PersonalAccessTokenQuery) Offset(offset int) *PersonalAccessTokenQuery {
	patq.offset = &offset
	return patq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (patq *PersonalAccessTokenQuery) Unique(unique bool) *PersonalAccessTokenQuery {
	patq.unique = &unique
	return patq
}

// Order adds an order step to the query.
func (patq *PersonalAccessTokenQuery) Order(o ...OrderFunc) *PersonalAccessTokenQuery {
	patq.order = append(patq.order, o...)
	return patq
}

// QueryUser chains the current query on the "user" edge.
func (patq *PersonalAccessTokenQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: patq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := patq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := patq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(personalaccesstoken.Table, personalaccesstoken.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, personalaccesstoken.UserTable, personalaccesstoken.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(patq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PersonalAccessToken entity from the query.
// Returns a *NotFoundError when no PersonalAccessToken was found.
func (patq *PersonalAccessTokenQuery) First(ctx context.Context) (*PersonalAccessToken, error) {
	nodes, err := patq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{personalaccesstoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (patq *PersonalAccessTokenQuery) FirstX(ctx context.Context) *PersonalAccessToken {
	node, err := patq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PersonalAccessToken ID from the query.
// Returns a *NotFoundError when no PersonalAccessToken ID was found.
func (patq *PersonalAccessTokenQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = patq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{personalaccesstoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (patq *PersonalAccessTokenQuery) FirstIDX(ctx context.Context) int {
	id, err := patq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PersonalAccessToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PersonalAccessToken entity is found.
// Returns a *NotFoundError when no PersonalAccessToken entities are found.
func (patq *PersonalAccessTokenQuery) Only(ctx context.Context) (*PersonalAccessToken, error) {
	nodes, err := patq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{personalaccesstoken.Label}
	default:
		return nil, &NotSingularError{personalaccesstoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (patq *PersonalAccessTokenQuery) OnlyX(ctx context.Context) *PersonalAccessToken {
	node, err := patq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PersonalAccessToken ID in the query.
// Returns a *NotSingularError when more than one PersonalAccessToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (patq *PersonalAccessTokenQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = patq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{personalaccesstoken.Label}
	default:
		err = &NotSingularError{personalaccesstoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (patq *PersonalAccessTokenQuery) OnlyIDX(ctx context.Context) int {
	id, err := patq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PersonalAccessTokens.
func (patq *PersonalAccessTokenQuery) All(ctx context.Context) ([]*PersonalAccessToken, error) {
	if err := patq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return patq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (patq *PersonalAccessTokenQuery) AllX(ctx context.Context) []*PersonalAccessToken {
	nodes, err := patq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PersonalAccessToken IDs.
func (patq *PersonalAccessTokenQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := patq.Select(personalaccesstoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (patq *PersonalAccessTokenQuery) IDsX(ctx context.Context) []int {
	ids, err := patq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (patq *PersonalAccessTokenQuery) Count(ctx context.Context) (int, error) {
	if err := patq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return patq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (patq *PersonalAccessTokenQuery) CountX(ctx context.Context) int {
	count, err := patq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (patq *PersonalAccessTokenQuery) Exist(ctx context.Context) (bool, error) {
	if err := patq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return patq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (patq *PersonalAccessTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := patq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PersonalAccessTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (patq *PersonalAccessTokenQuery) Clone() *PersonalAccessTokenQuery {
	if patq == nil {
		return nil
	}
	return &PersonalAccessTokenQuery{
		config:     patq.config,
		limit:      patq.limit,
		offset:     patq.offset,
		order:      append([]OrderFunc{}, patq.order...),
		predicates: append([]predicate.PersonalAccessToken{}, patq.predicates...),
		withUser:   patq.withUser.Clone(),
		// clone intermediate query.
		sql:    patq.sql.Clone(),
		path:   patq.path,
		unique: patq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (patq *PersonalAccessTokenQuery) WithUser(opts ...func(*UserQuery)) *PersonalAccessTokenQuery {
	query := &UserQuery{config: patq.config}
	for _, opt := range opts {
		opt(query)
	}
	patq.withUser = query
	return patq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PersonalAccessToken.Query().
//		GroupBy(personalaccesstoken.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (patq *PersonalAccessTokenQuery) GroupBy(field string, fields ...string) *PersonalAccessTokenGroupBy {
	group := &PersonalAccessTokenGroupBy{config: patq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := patq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return patq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.PersonalAccessToken.Query().
//		Select(personalaccesstoken.FieldName).
//		Scan(ctx, &v)
//
func (patq *PersonalAccessTokenQuery) Select(fields ...string) *PersonalAccessTokenSelect {
	patq.fields = append(patq.fields, fields...)
	return &PersonalAccessTokenSelect{PersonalAccessTokenQuery: patq}
}

func (patq *PersonalAccessTokenQuery) prepareQuery(ctx context.Context) error {
	for _, f := range patq.fields {
		if !personalaccesstoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if patq.path != nil {
		prev, err := patq.path(ctx)
		if err != nil {
			return err
		}
		patq.sql = prev
	}
	return nil
}

func (patq *PersonalAccessTokenQuery) sqlAll(ctx context.Context) ([]*PersonalAccessToken, error) {
	var (
		nodes       = []*PersonalAccessToken{}
		_spec       = patq.querySpec()
		loadedTypes = [1]bool{
			patq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PersonalAccessToken{config: patq.config}
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
	if err := sqlgraph.QueryNodes(ctx, patq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := patq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*PersonalAccessToken)
		for i := range nodes {
			fk := nodes[i].UserID
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
				return nil, fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (patq *PersonalAccessTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := patq.querySpec()
	_spec.Node.Columns = patq.fields
	if len(patq.fields) > 0 {
		_spec.Unique = patq.unique != nil && *patq.unique
	}
	return sqlgraph.CountNodes(ctx, patq.driver, _spec)
}

func (patq *PersonalAccessTokenQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := patq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (patq *PersonalAccessTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   personalaccesstoken.Table,
			Columns: personalaccesstoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: personalaccesstoken.FieldID,
			},
		},
		From:   patq.sql,
		Unique: true,
	}
	if unique := patq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := patq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, personalaccesstoken.FieldID)
		for i := range fields {
			if fields[i] != personalaccesstoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := patq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := patq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := patq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := patq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (patq *PersonalAccessTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(patq.driver.Dialect())
	t1 := builder.Table(personalaccesstoken.Table)
	columns := patq.fields
	if len(columns) == 0 {
		columns = personalaccesstoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if patq.sql != nil {
		selector = patq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if patq.unique != nil && *patq.unique {
		selector.Distinct()
	}
	for _, p := range patq.predicates {
		p(selector)
	}
	for _, p := range patq.order {
		p(selector)
	}
	if offset := patq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := patq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PersonalAccessTokenGroupBy is the group-by builder for PersonalAccessToken entities.
type PersonalAccessTokenGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (patgb *PersonalAccessTokenGroupBy) Aggregate(fns ...AggregateFunc) *PersonalAccessTokenGroupBy {
	patgb.fns = append(patgb.fns, fns...)
	return patgb
}

// Scan applies the group-by query and scans the result into the given value.
func (patgb *PersonalAccessTokenGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := patgb.path(ctx)
	if err != nil {
		return err
	}
	patgb.sql = query
	return patgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (patgb *PersonalAccessTokenGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := patgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (patgb *PersonalAccessTokenGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(patgb.fields) > 1 {
		return nil, errors.New("ent: PersonalAccessTokenGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := patgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (patgb *PersonalAccessTokenGroupBy) StringsX(ctx context.Context) []string {
	v, err := patgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (patgb *PersonalAccessTokenGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = patgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{personalaccesstoken.Label}
	default:
		err = fmt.Errorf("ent: PersonalAccessTokenGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (patgb *PersonalAccessTokenGroupBy) StringX(ctx context.Context) string {
	v, err := patgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (patgb *PersonalAccessTokenGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(patgb.fields) > 1 {
		return nil, errors.New("ent: PersonalAccessTokenGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := patgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (patgb *PersonalAccessTokenGroupBy) IntsX(ctx context.Context) []int {
	v, err := patgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (patgb *PersonalAccessTokenGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = patgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{personalaccesstoken.Label}
	default:
		err = fmt.Errorf("ent: PersonalAccessTokenGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (patgb *PersonalAccessTokenGroupBy) IntX(ctx context.Context) int {
	v, err := patgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (patgb *PersonalAccessTokenGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(patgb.fields) > 1 {
		return nil, errors.New("ent: PersonalAccessTokenGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := patgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (patgb *PersonalAccessTokenGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := patgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (patgb *PersonalAccessTokenGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = patgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{personalaccesstoken.Label}
	default:
		err = fmt.Errorf("ent: PersonalAccessTokenGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (patgb *PersonalAccessTokenGroupBy) Float64X(ctx context.Context) float64 {
	v, err := patgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (patgb *PersonalAccessTokenGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(patgb.fields) > 1 {
		return nil, errors.New("ent: PersonalAccessTokenGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := patgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (patgb *PersonalAccessTokenGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := patgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (patgb *PersonalAccessTokenGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = patgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{personalaccesstoken.Label}
	default:
		err = fmt.Errorf("ent: PersonalAccessTokenGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (patgb *PersonalAccessTokenGroupBy) BoolX(ctx context.Context) bool {
	v, err := patgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (patgb *PersonalAccessTokenGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range patgb.fields {
		if !personalaccesstoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := patgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := patgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (patgb *PersonalAccessTokenGroupBy) sqlQuery() *sql.Selector {
	selector := patgb.sql.Select()
	aggregation := make([]string, 0, len(patgb.fns))
	for _, fn := range patgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(patgb.fields)+len(patgb.fns))
		for _, f := range patgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(patgb.fields...)...)
}

// PersonalAccessTokenSelect is the builder for selecting fields of PersonalAccessToken entities.
type PersonalAccessTokenSelect struct {
	*PersonalAccessTokenQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pats *PersonalAccessTokenSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pats.prepareQuery(ctx); err != nil {
		return err
	}
	pats.sql = pats.PersonalAccessTokenQuery.sqlQuery(ctx)
	return pats.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pats *PersonalAccessTokenSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pats.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pats *PersonalAccessTokenSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pats.fields) > 1 {
		return nil, errors.New("ent: PersonalAccessTokenSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pats *PersonalAccessTokenSelect) StringsX(ctx context.Context) []string {
	v, err := pats.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pats *PersonalAccessTokenSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pats.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{personalaccesstoken.Label}
	default:
		err = fmt.Errorf("ent: PersonalAccessTokenSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pats *PersonalAccessTokenSelect) StringX(ctx context.Context) string {
	v, err := pats.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pats *PersonalAccessTokenSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pats.fields) > 1 {
		return nil, errors.New("ent: PersonalAccessTokenSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pats *PersonalAccessTokenSelect) IntsX(ctx context.Context) []int {
	v, err := pats.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pats *PersonalAccessTokenSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pats.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{personalaccesstoken.Label}
	default:
		err = fmt.Errorf("ent: PersonalAccessTokenSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pats *PersonalAccessTokenSelect) IntX(ctx context.Context) int {
	v, err := pats.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pats *PersonalAccessTokenSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pats.fields) > 1 {
		return nil, errors.New("ent: PersonalAccessTokenSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pats *PersonalAccessTokenSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pats.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pats *PersonalAccessTokenSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pats.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{personalaccesstoken.Label}
	default:
		err = fmt.Errorf("ent: PersonalAccessTokenSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pats *PersonalAccessTokenSelect) Float64X(ctx context.Context) float64 {
	v, err := pats.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pats *PersonalAccessTokenSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pats.fields) > 1 {
		return nil, errors.New("ent: PersonalAccessTokenSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pats *PersonalAccessTokenSelect) BoolsX(ctx context.Context) []bool {
	v, err := pats.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pats *PersonalAccessTokenSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pats.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{personalaccesstoken.Label}
	default:
		err = fmt.Errorf("ent: PersonalAccessTokenSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pats *PersonalAccessTokenSelect) BoolX(ctx context.Context) bool {
	v, err := pats.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pats *PersonalAccessTokenSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pats.sql.Query()
	if err := pats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
