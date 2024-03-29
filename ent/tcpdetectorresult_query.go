// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/ent/tcpdetectorresult"
)

// TCPDetectorResultQuery is the builder for querying TCPDetectorResult entities.
type TCPDetectorResultQuery struct {
	config
	ctx        *QueryContext
	order      []tcpdetectorresult.OrderOption
	inters     []Interceptor
	predicates []predicate.TCPDetectorResult
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TCPDetectorResultQuery builder.
func (tdrq *TCPDetectorResultQuery) Where(ps ...predicate.TCPDetectorResult) *TCPDetectorResultQuery {
	tdrq.predicates = append(tdrq.predicates, ps...)
	return tdrq
}

// Limit the number of records to be returned by this query.
func (tdrq *TCPDetectorResultQuery) Limit(limit int) *TCPDetectorResultQuery {
	tdrq.ctx.Limit = &limit
	return tdrq
}

// Offset to start from.
func (tdrq *TCPDetectorResultQuery) Offset(offset int) *TCPDetectorResultQuery {
	tdrq.ctx.Offset = &offset
	return tdrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tdrq *TCPDetectorResultQuery) Unique(unique bool) *TCPDetectorResultQuery {
	tdrq.ctx.Unique = &unique
	return tdrq
}

// Order specifies how the records should be ordered.
func (tdrq *TCPDetectorResultQuery) Order(o ...tcpdetectorresult.OrderOption) *TCPDetectorResultQuery {
	tdrq.order = append(tdrq.order, o...)
	return tdrq
}

// First returns the first TCPDetectorResult entity from the query.
// Returns a *NotFoundError when no TCPDetectorResult was found.
func (tdrq *TCPDetectorResultQuery) First(ctx context.Context) (*TCPDetectorResult, error) {
	nodes, err := tdrq.Limit(1).All(setContextOp(ctx, tdrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tcpdetectorresult.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tdrq *TCPDetectorResultQuery) FirstX(ctx context.Context) *TCPDetectorResult {
	node, err := tdrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TCPDetectorResult ID from the query.
// Returns a *NotFoundError when no TCPDetectorResult ID was found.
func (tdrq *TCPDetectorResultQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tdrq.Limit(1).IDs(setContextOp(ctx, tdrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tcpdetectorresult.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tdrq *TCPDetectorResultQuery) FirstIDX(ctx context.Context) int {
	id, err := tdrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TCPDetectorResult entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TCPDetectorResult entity is found.
// Returns a *NotFoundError when no TCPDetectorResult entities are found.
func (tdrq *TCPDetectorResultQuery) Only(ctx context.Context) (*TCPDetectorResult, error) {
	nodes, err := tdrq.Limit(2).All(setContextOp(ctx, tdrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tcpdetectorresult.Label}
	default:
		return nil, &NotSingularError{tcpdetectorresult.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tdrq *TCPDetectorResultQuery) OnlyX(ctx context.Context) *TCPDetectorResult {
	node, err := tdrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TCPDetectorResult ID in the query.
// Returns a *NotSingularError when more than one TCPDetectorResult ID is found.
// Returns a *NotFoundError when no entities are found.
func (tdrq *TCPDetectorResultQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tdrq.Limit(2).IDs(setContextOp(ctx, tdrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tcpdetectorresult.Label}
	default:
		err = &NotSingularError{tcpdetectorresult.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tdrq *TCPDetectorResultQuery) OnlyIDX(ctx context.Context) int {
	id, err := tdrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TCPDetectorResults.
func (tdrq *TCPDetectorResultQuery) All(ctx context.Context) ([]*TCPDetectorResult, error) {
	ctx = setContextOp(ctx, tdrq.ctx, "All")
	if err := tdrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TCPDetectorResult, *TCPDetectorResultQuery]()
	return withInterceptors[[]*TCPDetectorResult](ctx, tdrq, qr, tdrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tdrq *TCPDetectorResultQuery) AllX(ctx context.Context) []*TCPDetectorResult {
	nodes, err := tdrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TCPDetectorResult IDs.
func (tdrq *TCPDetectorResultQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tdrq.ctx.Unique == nil && tdrq.path != nil {
		tdrq.Unique(true)
	}
	ctx = setContextOp(ctx, tdrq.ctx, "IDs")
	if err = tdrq.Select(tcpdetectorresult.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tdrq *TCPDetectorResultQuery) IDsX(ctx context.Context) []int {
	ids, err := tdrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tdrq *TCPDetectorResultQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tdrq.ctx, "Count")
	if err := tdrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tdrq, querierCount[*TCPDetectorResultQuery](), tdrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tdrq *TCPDetectorResultQuery) CountX(ctx context.Context) int {
	count, err := tdrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tdrq *TCPDetectorResultQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tdrq.ctx, "Exist")
	switch _, err := tdrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tdrq *TCPDetectorResultQuery) ExistX(ctx context.Context) bool {
	exist, err := tdrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TCPDetectorResultQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tdrq *TCPDetectorResultQuery) Clone() *TCPDetectorResultQuery {
	if tdrq == nil {
		return nil
	}
	return &TCPDetectorResultQuery{
		config:     tdrq.config,
		ctx:        tdrq.ctx.Clone(),
		order:      append([]tcpdetectorresult.OrderOption{}, tdrq.order...),
		inters:     append([]Interceptor{}, tdrq.inters...),
		predicates: append([]predicate.TCPDetectorResult{}, tdrq.predicates...),
		// clone intermediate query.
		sql:  tdrq.sql.Clone(),
		path: tdrq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt" sql:"created_at"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TCPDetectorResult.Query().
//		GroupBy(tcpdetectorresult.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tdrq *TCPDetectorResultQuery) GroupBy(field string, fields ...string) *TCPDetectorResultGroupBy {
	tdrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TCPDetectorResultGroupBy{build: tdrq}
	grbuild.flds = &tdrq.ctx.Fields
	grbuild.label = tcpdetectorresult.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt" sql:"created_at"`
//	}
//
//	client.TCPDetectorResult.Query().
//		Select(tcpdetectorresult.FieldCreatedAt).
//		Scan(ctx, &v)
func (tdrq *TCPDetectorResultQuery) Select(fields ...string) *TCPDetectorResultSelect {
	tdrq.ctx.Fields = append(tdrq.ctx.Fields, fields...)
	sbuild := &TCPDetectorResultSelect{TCPDetectorResultQuery: tdrq}
	sbuild.label = tcpdetectorresult.Label
	sbuild.flds, sbuild.scan = &tdrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TCPDetectorResultSelect configured with the given aggregations.
func (tdrq *TCPDetectorResultQuery) Aggregate(fns ...AggregateFunc) *TCPDetectorResultSelect {
	return tdrq.Select().Aggregate(fns...)
}

func (tdrq *TCPDetectorResultQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tdrq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tdrq); err != nil {
				return err
			}
		}
	}
	for _, f := range tdrq.ctx.Fields {
		if !tcpdetectorresult.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tdrq.path != nil {
		prev, err := tdrq.path(ctx)
		if err != nil {
			return err
		}
		tdrq.sql = prev
	}
	return nil
}

func (tdrq *TCPDetectorResultQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TCPDetectorResult, error) {
	var (
		nodes = []*TCPDetectorResult{}
		_spec = tdrq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TCPDetectorResult).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TCPDetectorResult{config: tdrq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(tdrq.modifiers) > 0 {
		_spec.Modifiers = tdrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tdrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tdrq *TCPDetectorResultQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tdrq.querySpec()
	if len(tdrq.modifiers) > 0 {
		_spec.Modifiers = tdrq.modifiers
	}
	_spec.Node.Columns = tdrq.ctx.Fields
	if len(tdrq.ctx.Fields) > 0 {
		_spec.Unique = tdrq.ctx.Unique != nil && *tdrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tdrq.driver, _spec)
}

func (tdrq *TCPDetectorResultQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tcpdetectorresult.Table, tcpdetectorresult.Columns, sqlgraph.NewFieldSpec(tcpdetectorresult.FieldID, field.TypeInt))
	_spec.From = tdrq.sql
	if unique := tdrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tdrq.path != nil {
		_spec.Unique = true
	}
	if fields := tdrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tcpdetectorresult.FieldID)
		for i := range fields {
			if fields[i] != tcpdetectorresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tdrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tdrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tdrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tdrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tdrq *TCPDetectorResultQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tdrq.driver.Dialect())
	t1 := builder.Table(tcpdetectorresult.Table)
	columns := tdrq.ctx.Fields
	if len(columns) == 0 {
		columns = tcpdetectorresult.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tdrq.sql != nil {
		selector = tdrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tdrq.ctx.Unique != nil && *tdrq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range tdrq.modifiers {
		m(selector)
	}
	for _, p := range tdrq.predicates {
		p(selector)
	}
	for _, p := range tdrq.order {
		p(selector)
	}
	if offset := tdrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tdrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tdrq *TCPDetectorResultQuery) Modify(modifiers ...func(s *sql.Selector)) *TCPDetectorResultSelect {
	tdrq.modifiers = append(tdrq.modifiers, modifiers...)
	return tdrq.Select()
}

// TCPDetectorResultGroupBy is the group-by builder for TCPDetectorResult entities.
type TCPDetectorResultGroupBy struct {
	selector
	build *TCPDetectorResultQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tdrgb *TCPDetectorResultGroupBy) Aggregate(fns ...AggregateFunc) *TCPDetectorResultGroupBy {
	tdrgb.fns = append(tdrgb.fns, fns...)
	return tdrgb
}

// Scan applies the selector query and scans the result into the given value.
func (tdrgb *TCPDetectorResultGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tdrgb.build.ctx, "GroupBy")
	if err := tdrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TCPDetectorResultQuery, *TCPDetectorResultGroupBy](ctx, tdrgb.build, tdrgb, tdrgb.build.inters, v)
}

func (tdrgb *TCPDetectorResultGroupBy) sqlScan(ctx context.Context, root *TCPDetectorResultQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tdrgb.fns))
	for _, fn := range tdrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tdrgb.flds)+len(tdrgb.fns))
		for _, f := range *tdrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tdrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tdrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TCPDetectorResultSelect is the builder for selecting fields of TCPDetectorResult entities.
type TCPDetectorResultSelect struct {
	*TCPDetectorResultQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tdrs *TCPDetectorResultSelect) Aggregate(fns ...AggregateFunc) *TCPDetectorResultSelect {
	tdrs.fns = append(tdrs.fns, fns...)
	return tdrs
}

// Scan applies the selector query and scans the result into the given value.
func (tdrs *TCPDetectorResultSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tdrs.ctx, "Select")
	if err := tdrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TCPDetectorResultQuery, *TCPDetectorResultSelect](ctx, tdrs.TCPDetectorResultQuery, tdrs, tdrs.inters, v)
}

func (tdrs *TCPDetectorResultSelect) sqlScan(ctx context.Context, root *TCPDetectorResultQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tdrs.fns))
	for _, fn := range tdrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tdrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tdrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tdrs *TCPDetectorResultSelect) Modify(modifiers ...func(s *sql.Selector)) *TCPDetectorResultSelect {
	tdrs.modifiers = append(tdrs.modifiers, modifiers...)
	return tdrs
}
