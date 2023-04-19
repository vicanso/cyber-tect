// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vicanso/cybertect/ent/configuration"
	"github.com/vicanso/cybertect/ent/predicate"
)

// ConfigurationQuery is the builder for querying Configuration entities.
type ConfigurationQuery struct {
	config
	ctx        *QueryContext
	order      []configuration.OrderOption
	inters     []Interceptor
	predicates []predicate.Configuration
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ConfigurationQuery builder.
func (cq *ConfigurationQuery) Where(ps ...predicate.Configuration) *ConfigurationQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ConfigurationQuery) Limit(limit int) *ConfigurationQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ConfigurationQuery) Offset(offset int) *ConfigurationQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ConfigurationQuery) Unique(unique bool) *ConfigurationQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ConfigurationQuery) Order(o ...configuration.OrderOption) *ConfigurationQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// First returns the first Configuration entity from the query.
// Returns a *NotFoundError when no Configuration was found.
func (cq *ConfigurationQuery) First(ctx context.Context) (*Configuration, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{configuration.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ConfigurationQuery) FirstX(ctx context.Context) *Configuration {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Configuration ID from the query.
// Returns a *NotFoundError when no Configuration ID was found.
func (cq *ConfigurationQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{configuration.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ConfigurationQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Configuration entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Configuration entity is found.
// Returns a *NotFoundError when no Configuration entities are found.
func (cq *ConfigurationQuery) Only(ctx context.Context) (*Configuration, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{configuration.Label}
	default:
		return nil, &NotSingularError{configuration.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ConfigurationQuery) OnlyX(ctx context.Context) *Configuration {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Configuration ID in the query.
// Returns a *NotSingularError when more than one Configuration ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ConfigurationQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{configuration.Label}
	default:
		err = &NotSingularError{configuration.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ConfigurationQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Configurations.
func (cq *ConfigurationQuery) All(ctx context.Context) ([]*Configuration, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Configuration, *ConfigurationQuery]()
	return withInterceptors[[]*Configuration](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ConfigurationQuery) AllX(ctx context.Context) []*Configuration {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Configuration IDs.
func (cq *ConfigurationQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(configuration.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ConfigurationQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ConfigurationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ConfigurationQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ConfigurationQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ConfigurationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ConfigurationQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ConfigurationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ConfigurationQuery) Clone() *ConfigurationQuery {
	if cq == nil {
		return nil
	}
	return &ConfigurationQuery{
		config:     cq.config,
		ctx:        cq.ctx.Clone(),
		order:      append([]configuration.OrderOption{}, cq.order...),
		inters:     append([]Interceptor{}, cq.inters...),
		predicates: append([]predicate.Configuration{}, cq.predicates...),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
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
//	client.Configuration.Query().
//		GroupBy(configuration.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ConfigurationQuery) GroupBy(field string, fields ...string) *ConfigurationGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ConfigurationGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = configuration.Label
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
//	client.Configuration.Query().
//		Select(configuration.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *ConfigurationQuery) Select(fields ...string) *ConfigurationSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ConfigurationSelect{ConfigurationQuery: cq}
	sbuild.label = configuration.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ConfigurationSelect configured with the given aggregations.
func (cq *ConfigurationQuery) Aggregate(fns ...AggregateFunc) *ConfigurationSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ConfigurationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !configuration.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ConfigurationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Configuration, error) {
	var (
		nodes = []*Configuration{}
		_spec = cq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Configuration).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Configuration{config: cq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cq *ConfigurationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ConfigurationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(configuration.Table, configuration.Columns, sqlgraph.NewFieldSpec(configuration.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, configuration.FieldID)
		for i := range fields {
			if fields[i] != configuration.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ConfigurationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(configuration.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = configuration.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cq.modifiers {
		m(selector)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cq *ConfigurationQuery) Modify(modifiers ...func(s *sql.Selector)) *ConfigurationSelect {
	cq.modifiers = append(cq.modifiers, modifiers...)
	return cq.Select()
}

// ConfigurationGroupBy is the group-by builder for Configuration entities.
type ConfigurationGroupBy struct {
	selector
	build *ConfigurationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ConfigurationGroupBy) Aggregate(fns ...AggregateFunc) *ConfigurationGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ConfigurationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConfigurationQuery, *ConfigurationGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ConfigurationGroupBy) sqlScan(ctx context.Context, root *ConfigurationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ConfigurationSelect is the builder for selecting fields of Configuration entities.
type ConfigurationSelect struct {
	*ConfigurationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ConfigurationSelect) Aggregate(fns ...AggregateFunc) *ConfigurationSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ConfigurationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ConfigurationQuery, *ConfigurationSelect](ctx, cs.ConfigurationQuery, cs, cs.inters, v)
}

func (cs *ConfigurationSelect) sqlScan(ctx context.Context, root *ConfigurationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cs *ConfigurationSelect) Modify(modifiers ...func(s *sql.Selector)) *ConfigurationSelect {
	cs.modifiers = append(cs.modifiers, modifiers...)
	return cs
}
