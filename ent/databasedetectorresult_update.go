// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/vicanso/cybertect/ent/databasedetectorresult"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/schema"
)

// DatabaseDetectorResultUpdate is the builder for updating DatabaseDetectorResult entities.
type DatabaseDetectorResultUpdate struct {
	config
	hooks     []Hook
	mutation  *DatabaseDetectorResultMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DatabaseDetectorResultUpdate builder.
func (ddru *DatabaseDetectorResultUpdate) Where(ps ...predicate.DatabaseDetectorResult) *DatabaseDetectorResultUpdate {
	ddru.mutation.Where(ps...)
	return ddru
}

// SetTask sets the "task" field.
func (ddru *DatabaseDetectorResultUpdate) SetTask(i int) *DatabaseDetectorResultUpdate {
	ddru.mutation.ResetTask()
	ddru.mutation.SetTask(i)
	return ddru
}

// AddTask adds i to the "task" field.
func (ddru *DatabaseDetectorResultUpdate) AddTask(i int) *DatabaseDetectorResultUpdate {
	ddru.mutation.AddTask(i)
	return ddru
}

// SetResult sets the "result" field.
func (ddru *DatabaseDetectorResultUpdate) SetResult(sr schema.DetectorResult) *DatabaseDetectorResultUpdate {
	ddru.mutation.ResetResult()
	ddru.mutation.SetResult(sr)
	return ddru
}

// AddResult adds sr to the "result" field.
func (ddru *DatabaseDetectorResultUpdate) AddResult(sr schema.DetectorResult) *DatabaseDetectorResultUpdate {
	ddru.mutation.AddResult(sr)
	return ddru
}

// SetMaxDuration sets the "maxDuration" field.
func (ddru *DatabaseDetectorResultUpdate) SetMaxDuration(i int) *DatabaseDetectorResultUpdate {
	ddru.mutation.ResetMaxDuration()
	ddru.mutation.SetMaxDuration(i)
	return ddru
}

// AddMaxDuration adds i to the "maxDuration" field.
func (ddru *DatabaseDetectorResultUpdate) AddMaxDuration(i int) *DatabaseDetectorResultUpdate {
	ddru.mutation.AddMaxDuration(i)
	return ddru
}

// SetMessages sets the "messages" field.
func (ddru *DatabaseDetectorResultUpdate) SetMessages(s []string) *DatabaseDetectorResultUpdate {
	ddru.mutation.SetMessages(s)
	return ddru
}

// AppendMessages appends s to the "messages" field.
func (ddru *DatabaseDetectorResultUpdate) AppendMessages(s []string) *DatabaseDetectorResultUpdate {
	ddru.mutation.AppendMessages(s)
	return ddru
}

// SetUris sets the "uris" field.
func (ddru *DatabaseDetectorResultUpdate) SetUris(s []string) *DatabaseDetectorResultUpdate {
	ddru.mutation.SetUris(s)
	return ddru
}

// AppendUris appends s to the "uris" field.
func (ddru *DatabaseDetectorResultUpdate) AppendUris(s []string) *DatabaseDetectorResultUpdate {
	ddru.mutation.AppendUris(s)
	return ddru
}

// SetResults sets the "results" field.
func (ddru *DatabaseDetectorResultUpdate) SetResults(sdsr schema.DatabaseDetectorSubResults) *DatabaseDetectorResultUpdate {
	ddru.mutation.SetResults(sdsr)
	return ddru
}

// AppendResults appends sdsr to the "results" field.
func (ddru *DatabaseDetectorResultUpdate) AppendResults(sdsr schema.DatabaseDetectorSubResults) *DatabaseDetectorResultUpdate {
	ddru.mutation.AppendResults(sdsr)
	return ddru
}

// Mutation returns the DatabaseDetectorResultMutation object of the builder.
func (ddru *DatabaseDetectorResultUpdate) Mutation() *DatabaseDetectorResultMutation {
	return ddru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ddru *DatabaseDetectorResultUpdate) Save(ctx context.Context) (int, error) {
	ddru.defaults()
	return withHooks[int, DatabaseDetectorResultMutation](ctx, ddru.sqlSave, ddru.mutation, ddru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ddru *DatabaseDetectorResultUpdate) SaveX(ctx context.Context) int {
	affected, err := ddru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ddru *DatabaseDetectorResultUpdate) Exec(ctx context.Context) error {
	_, err := ddru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddru *DatabaseDetectorResultUpdate) ExecX(ctx context.Context) {
	if err := ddru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ddru *DatabaseDetectorResultUpdate) defaults() {
	if _, ok := ddru.mutation.UpdatedAt(); !ok {
		v := databasedetectorresult.UpdateDefaultUpdatedAt()
		ddru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ddru *DatabaseDetectorResultUpdate) check() error {
	if v, ok := ddru.mutation.Result(); ok {
		if err := databasedetectorresult.ResultValidator(int8(v)); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "DatabaseDetectorResult.result": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ddru *DatabaseDetectorResultUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DatabaseDetectorResultUpdate {
	ddru.modifiers = append(ddru.modifiers, modifiers...)
	return ddru
}

func (ddru *DatabaseDetectorResultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ddru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(databasedetectorresult.Table, databasedetectorresult.Columns, sqlgraph.NewFieldSpec(databasedetectorresult.FieldID, field.TypeInt))
	if ps := ddru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ddru.mutation.UpdatedAt(); ok {
		_spec.SetField(databasedetectorresult.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ddru.mutation.Task(); ok {
		_spec.SetField(databasedetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := ddru.mutation.AddedTask(); ok {
		_spec.AddField(databasedetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := ddru.mutation.Result(); ok {
		_spec.SetField(databasedetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := ddru.mutation.AddedResult(); ok {
		_spec.AddField(databasedetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := ddru.mutation.MaxDuration(); ok {
		_spec.SetField(databasedetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := ddru.mutation.AddedMaxDuration(); ok {
		_spec.AddField(databasedetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := ddru.mutation.Messages(); ok {
		_spec.SetField(databasedetectorresult.FieldMessages, field.TypeJSON, value)
	}
	if value, ok := ddru.mutation.AppendedMessages(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, databasedetectorresult.FieldMessages, value)
		})
	}
	if value, ok := ddru.mutation.Uris(); ok {
		_spec.SetField(databasedetectorresult.FieldUris, field.TypeJSON, value)
	}
	if value, ok := ddru.mutation.AppendedUris(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, databasedetectorresult.FieldUris, value)
		})
	}
	if value, ok := ddru.mutation.Results(); ok {
		_spec.SetField(databasedetectorresult.FieldResults, field.TypeJSON, value)
	}
	if value, ok := ddru.mutation.AppendedResults(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, databasedetectorresult.FieldResults, value)
		})
	}
	_spec.AddModifiers(ddru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ddru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{databasedetectorresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ddru.mutation.done = true
	return n, nil
}

// DatabaseDetectorResultUpdateOne is the builder for updating a single DatabaseDetectorResult entity.
type DatabaseDetectorResultUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DatabaseDetectorResultMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetTask sets the "task" field.
func (ddruo *DatabaseDetectorResultUpdateOne) SetTask(i int) *DatabaseDetectorResultUpdateOne {
	ddruo.mutation.ResetTask()
	ddruo.mutation.SetTask(i)
	return ddruo
}

// AddTask adds i to the "task" field.
func (ddruo *DatabaseDetectorResultUpdateOne) AddTask(i int) *DatabaseDetectorResultUpdateOne {
	ddruo.mutation.AddTask(i)
	return ddruo
}

// SetResult sets the "result" field.
func (ddruo *DatabaseDetectorResultUpdateOne) SetResult(sr schema.DetectorResult) *DatabaseDetectorResultUpdateOne {
	ddruo.mutation.ResetResult()
	ddruo.mutation.SetResult(sr)
	return ddruo
}

// AddResult adds sr to the "result" field.
func (ddruo *DatabaseDetectorResultUpdateOne) AddResult(sr schema.DetectorResult) *DatabaseDetectorResultUpdateOne {
	ddruo.mutation.AddResult(sr)
	return ddruo
}

// SetMaxDuration sets the "maxDuration" field.
func (ddruo *DatabaseDetectorResultUpdateOne) SetMaxDuration(i int) *DatabaseDetectorResultUpdateOne {
	ddruo.mutation.ResetMaxDuration()
	ddruo.mutation.SetMaxDuration(i)
	return ddruo
}

// AddMaxDuration adds i to the "maxDuration" field.
func (ddruo *DatabaseDetectorResultUpdateOne) AddMaxDuration(i int) *DatabaseDetectorResultUpdateOne {
	ddruo.mutation.AddMaxDuration(i)
	return ddruo
}

// SetMessages sets the "messages" field.
func (ddruo *DatabaseDetectorResultUpdateOne) SetMessages(s []string) *DatabaseDetectorResultUpdateOne {
	ddruo.mutation.SetMessages(s)
	return ddruo
}

// AppendMessages appends s to the "messages" field.
func (ddruo *DatabaseDetectorResultUpdateOne) AppendMessages(s []string) *DatabaseDetectorResultUpdateOne {
	ddruo.mutation.AppendMessages(s)
	return ddruo
}

// SetUris sets the "uris" field.
func (ddruo *DatabaseDetectorResultUpdateOne) SetUris(s []string) *DatabaseDetectorResultUpdateOne {
	ddruo.mutation.SetUris(s)
	return ddruo
}

// AppendUris appends s to the "uris" field.
func (ddruo *DatabaseDetectorResultUpdateOne) AppendUris(s []string) *DatabaseDetectorResultUpdateOne {
	ddruo.mutation.AppendUris(s)
	return ddruo
}

// SetResults sets the "results" field.
func (ddruo *DatabaseDetectorResultUpdateOne) SetResults(sdsr schema.DatabaseDetectorSubResults) *DatabaseDetectorResultUpdateOne {
	ddruo.mutation.SetResults(sdsr)
	return ddruo
}

// AppendResults appends sdsr to the "results" field.
func (ddruo *DatabaseDetectorResultUpdateOne) AppendResults(sdsr schema.DatabaseDetectorSubResults) *DatabaseDetectorResultUpdateOne {
	ddruo.mutation.AppendResults(sdsr)
	return ddruo
}

// Mutation returns the DatabaseDetectorResultMutation object of the builder.
func (ddruo *DatabaseDetectorResultUpdateOne) Mutation() *DatabaseDetectorResultMutation {
	return ddruo.mutation
}

// Where appends a list predicates to the DatabaseDetectorResultUpdate builder.
func (ddruo *DatabaseDetectorResultUpdateOne) Where(ps ...predicate.DatabaseDetectorResult) *DatabaseDetectorResultUpdateOne {
	ddruo.mutation.Where(ps...)
	return ddruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ddruo *DatabaseDetectorResultUpdateOne) Select(field string, fields ...string) *DatabaseDetectorResultUpdateOne {
	ddruo.fields = append([]string{field}, fields...)
	return ddruo
}

// Save executes the query and returns the updated DatabaseDetectorResult entity.
func (ddruo *DatabaseDetectorResultUpdateOne) Save(ctx context.Context) (*DatabaseDetectorResult, error) {
	ddruo.defaults()
	return withHooks[*DatabaseDetectorResult, DatabaseDetectorResultMutation](ctx, ddruo.sqlSave, ddruo.mutation, ddruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ddruo *DatabaseDetectorResultUpdateOne) SaveX(ctx context.Context) *DatabaseDetectorResult {
	node, err := ddruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ddruo *DatabaseDetectorResultUpdateOne) Exec(ctx context.Context) error {
	_, err := ddruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddruo *DatabaseDetectorResultUpdateOne) ExecX(ctx context.Context) {
	if err := ddruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ddruo *DatabaseDetectorResultUpdateOne) defaults() {
	if _, ok := ddruo.mutation.UpdatedAt(); !ok {
		v := databasedetectorresult.UpdateDefaultUpdatedAt()
		ddruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ddruo *DatabaseDetectorResultUpdateOne) check() error {
	if v, ok := ddruo.mutation.Result(); ok {
		if err := databasedetectorresult.ResultValidator(int8(v)); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "DatabaseDetectorResult.result": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ddruo *DatabaseDetectorResultUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DatabaseDetectorResultUpdateOne {
	ddruo.modifiers = append(ddruo.modifiers, modifiers...)
	return ddruo
}

func (ddruo *DatabaseDetectorResultUpdateOne) sqlSave(ctx context.Context) (_node *DatabaseDetectorResult, err error) {
	if err := ddruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(databasedetectorresult.Table, databasedetectorresult.Columns, sqlgraph.NewFieldSpec(databasedetectorresult.FieldID, field.TypeInt))
	id, ok := ddruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DatabaseDetectorResult.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ddruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, databasedetectorresult.FieldID)
		for _, f := range fields {
			if !databasedetectorresult.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != databasedetectorresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ddruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ddruo.mutation.UpdatedAt(); ok {
		_spec.SetField(databasedetectorresult.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ddruo.mutation.Task(); ok {
		_spec.SetField(databasedetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := ddruo.mutation.AddedTask(); ok {
		_spec.AddField(databasedetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := ddruo.mutation.Result(); ok {
		_spec.SetField(databasedetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := ddruo.mutation.AddedResult(); ok {
		_spec.AddField(databasedetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := ddruo.mutation.MaxDuration(); ok {
		_spec.SetField(databasedetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := ddruo.mutation.AddedMaxDuration(); ok {
		_spec.AddField(databasedetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := ddruo.mutation.Messages(); ok {
		_spec.SetField(databasedetectorresult.FieldMessages, field.TypeJSON, value)
	}
	if value, ok := ddruo.mutation.AppendedMessages(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, databasedetectorresult.FieldMessages, value)
		})
	}
	if value, ok := ddruo.mutation.Uris(); ok {
		_spec.SetField(databasedetectorresult.FieldUris, field.TypeJSON, value)
	}
	if value, ok := ddruo.mutation.AppendedUris(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, databasedetectorresult.FieldUris, value)
		})
	}
	if value, ok := ddruo.mutation.Results(); ok {
		_spec.SetField(databasedetectorresult.FieldResults, field.TypeJSON, value)
	}
	if value, ok := ddruo.mutation.AppendedResults(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, databasedetectorresult.FieldResults, value)
		})
	}
	_spec.AddModifiers(ddruo.modifiers...)
	_node = &DatabaseDetectorResult{config: ddruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ddruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{databasedetectorresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ddruo.mutation.done = true
	return _node, nil
}