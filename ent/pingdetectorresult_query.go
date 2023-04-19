// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vicanso/cybertect/ent/pingdetectorresult"
	"github.com/vicanso/cybertect/ent/predicate"
)

// PingDetectorResultQuery is the builder for querying PingDetectorResult entities.
type PingDetectorResultQuery struct {
	config
	ctx        *QueryContext
	order      []pingdetectorresult.OrderOption
	inters     []Interceptor
	predicates []predicate.PingDetectorResult
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PingDetectorResultQuery builder.
func (pdrq *PingDetectorResultQuery) Where(ps ...predicate.PingDetectorResult) *PingDetectorResultQuery {
	pdrq.predicates = append(pdrq.predicates, ps...)
	return pdrq
}

// Limit the number of records to be returned by this query.
func (pdrq *PingDetectorResultQuery) Limit(limit int) *PingDetectorResultQuery {
	pdrq.ctx.Limit = &limit
	return pdrq
}

// Offset to start from.
func (pdrq *PingDetectorResultQuery) Offset(offset int) *PingDetectorResultQuery {
	pdrq.ctx.Offset = &offset
	return pdrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pdrq *PingDetectorResultQuery) Unique(unique bool) *PingDetectorResultQuery {
	pdrq.ctx.Unique = &unique
	return pdrq
}

// Order specifies how the records should be ordered.
func (pdrq *PingDetectorResultQuery) Order(o ...pingdetectorresult.OrderOption) *PingDetectorResultQuery {
	pdrq.order = append(pdrq.order, o...)
	return pdrq
}

// First returns the first PingDetectorResult entity from the query.
// Returns a *NotFoundError when no PingDetectorResult was found.
func (pdrq *PingDetectorResultQuery) First(ctx context.Context) (*PingDetectorResult, error) {
	nodes, err := pdrq.Limit(1).All(setContextOp(ctx, pdrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pingdetectorresult.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pdrq *PingDetectorResultQuery) FirstX(ctx context.Context) *PingDetectorResult {
	node, err := pdrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PingDetectorResult ID from the query.
// Returns a *NotFoundError when no PingDetectorResult ID was found.
func (pdrq *PingDetectorResultQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pdrq.Limit(1).IDs(setContextOp(ctx, pdrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pingdetectorresult.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pdrq *PingDetectorResultQuery) FirstIDX(ctx context.Context) int {
	id, err := pdrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PingDetectorResult entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PingDetectorResult entity is found.
// Returns a *NotFoundError when no PingDetectorResult entities are found.
func (pdrq *PingDetectorResultQuery) Only(ctx context.Context) (*PingDetectorResult, error) {
	nodes, err := pdrq.Limit(2).All(setContextOp(ctx, pdrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pingdetectorresult.Label}
	default:
		return nil, &NotSingularError{pingdetectorresult.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pdrq *PingDetectorResultQuery) OnlyX(ctx context.Context) *PingDetectorResult {
	node, err := pdrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PingDetectorResult ID in the query.
// Returns a *NotSingularError when more than one PingDetectorResult ID is found.
// Returns a *NotFoundError when no entities are found.
func (pdrq *PingDetectorResultQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pdrq.Limit(2).IDs(setContextOp(ctx, pdrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pingdetectorresult.Label}
	default:
		err = &NotSingularError{pingdetectorresult.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pdrq *PingDetectorResultQuery) OnlyIDX(ctx context.Context) int {
	id, err := pdrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PingDetectorResults.
func (pdrq *PingDetectorResultQuery) All(ctx context.Context) ([]*PingDetectorResult, error) {
	ctx = setContextOp(ctx, pdrq.ctx, "All")
	if err := pdrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PingDetectorResult, *PingDetectorResultQuery]()
	return withInterceptors[[]*PingDetectorResult](ctx, pdrq, qr, pdrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pdrq *PingDetectorResultQuery) AllX(ctx context.Context) []*PingDetectorResult {
	nodes, err := pdrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PingDetectorResult IDs.
func (pdrq *PingDetectorResultQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pdrq.ctx.Unique == nil && pdrq.path != nil {
		pdrq.Unique(true)
	}
	ctx = setContextOp(ctx, pdrq.ctx, "IDs")
	if err = pdrq.Select(pingdetectorresult.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pdrq *PingDetectorResultQuery) IDsX(ctx context.Context) []int {
	ids, err := pdrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pdrq *PingDetectorResultQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pdrq.ctx, "Count")
	if err := pdrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pdrq, querierCount[*PingDetectorResultQuery](), pdrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pdrq *PingDetectorResultQuery) CountX(ctx context.Context) int {
	count, err := pdrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pdrq *PingDetectorResultQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pdrq.ctx, "Exist")
	switch _, err := pdrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pdrq *PingDetectorResultQuery) ExistX(ctx context.Context) bool {
	exist, err := pdrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PingDetectorResultQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pdrq *PingDetectorResultQuery) Clone() *PingDetectorResultQuery {
	if pdrq == nil {
		return nil
	}
	return &PingDetectorResultQuery{
		config:     pdrq.config,
		ctx:        pdrq.ctx.Clone(),
		order:      append([]pingdetectorresult.OrderOption{}, pdrq.order...),
		inters:     append([]Interceptor{}, pdrq.inters...),
		predicates: append([]predicate.PingDetectorResult{}, pdrq.predicates...),
		// clone intermediate query.
		sql:  pdrq.sql.Clone(),
		path: pdrq.path,
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
//	client.PingDetectorResult.Query().
//		GroupBy(pingdetectorresult.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pdrq *PingDetectorResultQuery) GroupBy(field string, fields ...string) *PingDetectorResultGroupBy {
	pdrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PingDetectorResultGroupBy{build: pdrq}
	grbuild.flds = &pdrq.ctx.Fields
	grbuild.label = pingdetectorresult.Label
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
//	client.PingDetectorResult.Query().
//		Select(pingdetectorresult.FieldCreatedAt).
//		Scan(ctx, &v)
func (pdrq *PingDetectorResultQuery) Select(fields ...string) *PingDetectorResultSelect {
	pdrq.ctx.Fields = append(pdrq.ctx.Fields, fields...)
	sbuild := &PingDetectorResultSelect{PingDetectorResultQuery: pdrq}
	sbuild.label = pingdetectorresult.Label
	sbuild.flds, sbuild.scan = &pdrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PingDetectorResultSelect configured with the given aggregations.
func (pdrq *PingDetectorResultQuery) Aggregate(fns ...AggregateFunc) *PingDetectorResultSelect {
	return pdrq.Select().Aggregate(fns...)
}

func (pdrq *PingDetectorResultQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pdrq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pdrq); err != nil {
				return err
			}
		}
	}
	for _, f := range pdrq.ctx.Fields {
		if !pingdetectorresult.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pdrq.path != nil {
		prev, err := pdrq.path(ctx)
		if err != nil {
			return err
		}
		pdrq.sql = prev
	}
	return nil
}

func (pdrq *PingDetectorResultQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PingDetectorResult, error) {
	var (
		nodes = []*PingDetectorResult{}
		_spec = pdrq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PingDetectorResult).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PingDetectorResult{config: pdrq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(pdrq.modifiers) > 0 {
		_spec.Modifiers = pdrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pdrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pdrq *PingDetectorResultQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pdrq.querySpec()
	if len(pdrq.modifiers) > 0 {
		_spec.Modifiers = pdrq.modifiers
	}
	_spec.Node.Columns = pdrq.ctx.Fields
	if len(pdrq.ctx.Fields) > 0 {
		_spec.Unique = pdrq.ctx.Unique != nil && *pdrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pdrq.driver, _spec)
}

func (pdrq *PingDetectorResultQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(pingdetectorresult.Table, pingdetectorresult.Columns, sqlgraph.NewFieldSpec(pingdetectorresult.FieldID, field.TypeInt))
	_spec.From = pdrq.sql
	if unique := pdrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pdrq.path != nil {
		_spec.Unique = true
	}
	if fields := pdrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pingdetectorresult.FieldID)
		for i := range fields {
			if fields[i] != pingdetectorresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pdrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pdrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pdrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pdrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pdrq *PingDetectorResultQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pdrq.driver.Dialect())
	t1 := builder.Table(pingdetectorresult.Table)
	columns := pdrq.ctx.Fields
	if len(columns) == 0 {
		columns = pingdetectorresult.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pdrq.sql != nil {
		selector = pdrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pdrq.ctx.Unique != nil && *pdrq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range pdrq.modifiers {
		m(selector)
	}
	for _, p := range pdrq.predicates {
		p(selector)
	}
	for _, p := range pdrq.order {
		p(selector)
	}
	if offset := pdrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pdrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pdrq *PingDetectorResultQuery) Modify(modifiers ...func(s *sql.Selector)) *PingDetectorResultSelect {
	pdrq.modifiers = append(pdrq.modifiers, modifiers...)
	return pdrq.Select()
}

// PingDetectorResultGroupBy is the group-by builder for PingDetectorResult entities.
type PingDetectorResultGroupBy struct {
	selector
	build *PingDetectorResultQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pdrgb *PingDetectorResultGroupBy) Aggregate(fns ...AggregateFunc) *PingDetectorResultGroupBy {
	pdrgb.fns = append(pdrgb.fns, fns...)
	return pdrgb
}

// Scan applies the selector query and scans the result into the given value.
func (pdrgb *PingDetectorResultGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pdrgb.build.ctx, "GroupBy")
	if err := pdrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PingDetectorResultQuery, *PingDetectorResultGroupBy](ctx, pdrgb.build, pdrgb, pdrgb.build.inters, v)
}

func (pdrgb *PingDetectorResultGroupBy) sqlScan(ctx context.Context, root *PingDetectorResultQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pdrgb.fns))
	for _, fn := range pdrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pdrgb.flds)+len(pdrgb.fns))
		for _, f := range *pdrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pdrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pdrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PingDetectorResultSelect is the builder for selecting fields of PingDetectorResult entities.
type PingDetectorResultSelect struct {
	*PingDetectorResultQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pdrs *PingDetectorResultSelect) Aggregate(fns ...AggregateFunc) *PingDetectorResultSelect {
	pdrs.fns = append(pdrs.fns, fns...)
	return pdrs
}

// Scan applies the selector query and scans the result into the given value.
func (pdrs *PingDetectorResultSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pdrs.ctx, "Select")
	if err := pdrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PingDetectorResultQuery, *PingDetectorResultSelect](ctx, pdrs.PingDetectorResultQuery, pdrs, pdrs.inters, v)
}

func (pdrs *PingDetectorResultSelect) sqlScan(ctx context.Context, root *PingDetectorResultQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pdrs.fns))
	for _, fn := range pdrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pdrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pdrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pdrs *PingDetectorResultSelect) Modify(modifiers ...func(s *sql.Selector)) *PingDetectorResultSelect {
	pdrs.modifiers = append(pdrs.modifiers, modifiers...)
	return pdrs
}