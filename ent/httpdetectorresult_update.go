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
	"github.com/vicanso/cybertect/ent/httpdetectorresult"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/schema"
)

// HTTPDetectorResultUpdate is the builder for updating HTTPDetectorResult entities.
type HTTPDetectorResultUpdate struct {
	config
	hooks     []Hook
	mutation  *HTTPDetectorResultMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the HTTPDetectorResultUpdate builder.
func (hdru *HTTPDetectorResultUpdate) Where(ps ...predicate.HTTPDetectorResult) *HTTPDetectorResultUpdate {
	hdru.mutation.Where(ps...)
	return hdru
}

// SetTask sets the "task" field.
func (hdru *HTTPDetectorResultUpdate) SetTask(i int) *HTTPDetectorResultUpdate {
	hdru.mutation.ResetTask()
	hdru.mutation.SetTask(i)
	return hdru
}

// AddTask adds i to the "task" field.
func (hdru *HTTPDetectorResultUpdate) AddTask(i int) *HTTPDetectorResultUpdate {
	hdru.mutation.AddTask(i)
	return hdru
}

// SetResult sets the "result" field.
func (hdru *HTTPDetectorResultUpdate) SetResult(sr schema.DetectorResult) *HTTPDetectorResultUpdate {
	hdru.mutation.ResetResult()
	hdru.mutation.SetResult(sr)
	return hdru
}

// AddResult adds sr to the "result" field.
func (hdru *HTTPDetectorResultUpdate) AddResult(sr schema.DetectorResult) *HTTPDetectorResultUpdate {
	hdru.mutation.AddResult(sr)
	return hdru
}

// SetMaxDuration sets the "maxDuration" field.
func (hdru *HTTPDetectorResultUpdate) SetMaxDuration(i int) *HTTPDetectorResultUpdate {
	hdru.mutation.ResetMaxDuration()
	hdru.mutation.SetMaxDuration(i)
	return hdru
}

// AddMaxDuration adds i to the "maxDuration" field.
func (hdru *HTTPDetectorResultUpdate) AddMaxDuration(i int) *HTTPDetectorResultUpdate {
	hdru.mutation.AddMaxDuration(i)
	return hdru
}

// SetMessages sets the "messages" field.
func (hdru *HTTPDetectorResultUpdate) SetMessages(s []string) *HTTPDetectorResultUpdate {
	hdru.mutation.SetMessages(s)
	return hdru
}

// AppendMessages appends s to the "messages" field.
func (hdru *HTTPDetectorResultUpdate) AppendMessages(s []string) *HTTPDetectorResultUpdate {
	hdru.mutation.AppendMessages(s)
	return hdru
}

// SetURL sets the "url" field.
func (hdru *HTTPDetectorResultUpdate) SetURL(s string) *HTTPDetectorResultUpdate {
	hdru.mutation.SetURL(s)
	return hdru
}

// SetResults sets the "results" field.
func (hdru *HTTPDetectorResultUpdate) SetResults(sdsr schema.HTTPDetectorSubResults) *HTTPDetectorResultUpdate {
	hdru.mutation.SetResults(sdsr)
	return hdru
}

// AppendResults appends sdsr to the "results" field.
func (hdru *HTTPDetectorResultUpdate) AppendResults(sdsr schema.HTTPDetectorSubResults) *HTTPDetectorResultUpdate {
	hdru.mutation.AppendResults(sdsr)
	return hdru
}

// Mutation returns the HTTPDetectorResultMutation object of the builder.
func (hdru *HTTPDetectorResultUpdate) Mutation() *HTTPDetectorResultMutation {
	return hdru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hdru *HTTPDetectorResultUpdate) Save(ctx context.Context) (int, error) {
	hdru.defaults()
	return withHooks[int, HTTPDetectorResultMutation](ctx, hdru.sqlSave, hdru.mutation, hdru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hdru *HTTPDetectorResultUpdate) SaveX(ctx context.Context) int {
	affected, err := hdru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hdru *HTTPDetectorResultUpdate) Exec(ctx context.Context) error {
	_, err := hdru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hdru *HTTPDetectorResultUpdate) ExecX(ctx context.Context) {
	if err := hdru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hdru *HTTPDetectorResultUpdate) defaults() {
	if _, ok := hdru.mutation.UpdatedAt(); !ok {
		v := httpdetectorresult.UpdateDefaultUpdatedAt()
		hdru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hdru *HTTPDetectorResultUpdate) check() error {
	if v, ok := hdru.mutation.Result(); ok {
		if err := httpdetectorresult.ResultValidator(int8(v)); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "HTTPDetectorResult.result": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (hdru *HTTPDetectorResultUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *HTTPDetectorResultUpdate {
	hdru.modifiers = append(hdru.modifiers, modifiers...)
	return hdru
}

func (hdru *HTTPDetectorResultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hdru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(httpdetectorresult.Table, httpdetectorresult.Columns, sqlgraph.NewFieldSpec(httpdetectorresult.FieldID, field.TypeInt))
	if ps := hdru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hdru.mutation.UpdatedAt(); ok {
		_spec.SetField(httpdetectorresult.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := hdru.mutation.Task(); ok {
		_spec.SetField(httpdetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := hdru.mutation.AddedTask(); ok {
		_spec.AddField(httpdetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := hdru.mutation.Result(); ok {
		_spec.SetField(httpdetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := hdru.mutation.AddedResult(); ok {
		_spec.AddField(httpdetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := hdru.mutation.MaxDuration(); ok {
		_spec.SetField(httpdetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := hdru.mutation.AddedMaxDuration(); ok {
		_spec.AddField(httpdetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := hdru.mutation.Messages(); ok {
		_spec.SetField(httpdetectorresult.FieldMessages, field.TypeJSON, value)
	}
	if value, ok := hdru.mutation.AppendedMessages(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, httpdetectorresult.FieldMessages, value)
		})
	}
	if value, ok := hdru.mutation.URL(); ok {
		_spec.SetField(httpdetectorresult.FieldURL, field.TypeString, value)
	}
	if value, ok := hdru.mutation.Results(); ok {
		_spec.SetField(httpdetectorresult.FieldResults, field.TypeJSON, value)
	}
	if value, ok := hdru.mutation.AppendedResults(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, httpdetectorresult.FieldResults, value)
		})
	}
	_spec.AddModifiers(hdru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, hdru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{httpdetectorresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hdru.mutation.done = true
	return n, nil
}

// HTTPDetectorResultUpdateOne is the builder for updating a single HTTPDetectorResult entity.
type HTTPDetectorResultUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *HTTPDetectorResultMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetTask sets the "task" field.
func (hdruo *HTTPDetectorResultUpdateOne) SetTask(i int) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.ResetTask()
	hdruo.mutation.SetTask(i)
	return hdruo
}

// AddTask adds i to the "task" field.
func (hdruo *HTTPDetectorResultUpdateOne) AddTask(i int) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.AddTask(i)
	return hdruo
}

// SetResult sets the "result" field.
func (hdruo *HTTPDetectorResultUpdateOne) SetResult(sr schema.DetectorResult) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.ResetResult()
	hdruo.mutation.SetResult(sr)
	return hdruo
}

// AddResult adds sr to the "result" field.
func (hdruo *HTTPDetectorResultUpdateOne) AddResult(sr schema.DetectorResult) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.AddResult(sr)
	return hdruo
}

// SetMaxDuration sets the "maxDuration" field.
func (hdruo *HTTPDetectorResultUpdateOne) SetMaxDuration(i int) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.ResetMaxDuration()
	hdruo.mutation.SetMaxDuration(i)
	return hdruo
}

// AddMaxDuration adds i to the "maxDuration" field.
func (hdruo *HTTPDetectorResultUpdateOne) AddMaxDuration(i int) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.AddMaxDuration(i)
	return hdruo
}

// SetMessages sets the "messages" field.
func (hdruo *HTTPDetectorResultUpdateOne) SetMessages(s []string) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.SetMessages(s)
	return hdruo
}

// AppendMessages appends s to the "messages" field.
func (hdruo *HTTPDetectorResultUpdateOne) AppendMessages(s []string) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.AppendMessages(s)
	return hdruo
}

// SetURL sets the "url" field.
func (hdruo *HTTPDetectorResultUpdateOne) SetURL(s string) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.SetURL(s)
	return hdruo
}

// SetResults sets the "results" field.
func (hdruo *HTTPDetectorResultUpdateOne) SetResults(sdsr schema.HTTPDetectorSubResults) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.SetResults(sdsr)
	return hdruo
}

// AppendResults appends sdsr to the "results" field.
func (hdruo *HTTPDetectorResultUpdateOne) AppendResults(sdsr schema.HTTPDetectorSubResults) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.AppendResults(sdsr)
	return hdruo
}

// Mutation returns the HTTPDetectorResultMutation object of the builder.
func (hdruo *HTTPDetectorResultUpdateOne) Mutation() *HTTPDetectorResultMutation {
	return hdruo.mutation
}

// Where appends a list predicates to the HTTPDetectorResultUpdate builder.
func (hdruo *HTTPDetectorResultUpdateOne) Where(ps ...predicate.HTTPDetectorResult) *HTTPDetectorResultUpdateOne {
	hdruo.mutation.Where(ps...)
	return hdruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (hdruo *HTTPDetectorResultUpdateOne) Select(field string, fields ...string) *HTTPDetectorResultUpdateOne {
	hdruo.fields = append([]string{field}, fields...)
	return hdruo
}

// Save executes the query and returns the updated HTTPDetectorResult entity.
func (hdruo *HTTPDetectorResultUpdateOne) Save(ctx context.Context) (*HTTPDetectorResult, error) {
	hdruo.defaults()
	return withHooks[*HTTPDetectorResult, HTTPDetectorResultMutation](ctx, hdruo.sqlSave, hdruo.mutation, hdruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hdruo *HTTPDetectorResultUpdateOne) SaveX(ctx context.Context) *HTTPDetectorResult {
	node, err := hdruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hdruo *HTTPDetectorResultUpdateOne) Exec(ctx context.Context) error {
	_, err := hdruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hdruo *HTTPDetectorResultUpdateOne) ExecX(ctx context.Context) {
	if err := hdruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hdruo *HTTPDetectorResultUpdateOne) defaults() {
	if _, ok := hdruo.mutation.UpdatedAt(); !ok {
		v := httpdetectorresult.UpdateDefaultUpdatedAt()
		hdruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hdruo *HTTPDetectorResultUpdateOne) check() error {
	if v, ok := hdruo.mutation.Result(); ok {
		if err := httpdetectorresult.ResultValidator(int8(v)); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "HTTPDetectorResult.result": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (hdruo *HTTPDetectorResultUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *HTTPDetectorResultUpdateOne {
	hdruo.modifiers = append(hdruo.modifiers, modifiers...)
	return hdruo
}

func (hdruo *HTTPDetectorResultUpdateOne) sqlSave(ctx context.Context) (_node *HTTPDetectorResult, err error) {
	if err := hdruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(httpdetectorresult.Table, httpdetectorresult.Columns, sqlgraph.NewFieldSpec(httpdetectorresult.FieldID, field.TypeInt))
	id, ok := hdruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HTTPDetectorResult.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := hdruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, httpdetectorresult.FieldID)
		for _, f := range fields {
			if !httpdetectorresult.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != httpdetectorresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := hdruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hdruo.mutation.UpdatedAt(); ok {
		_spec.SetField(httpdetectorresult.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := hdruo.mutation.Task(); ok {
		_spec.SetField(httpdetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := hdruo.mutation.AddedTask(); ok {
		_spec.AddField(httpdetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := hdruo.mutation.Result(); ok {
		_spec.SetField(httpdetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := hdruo.mutation.AddedResult(); ok {
		_spec.AddField(httpdetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := hdruo.mutation.MaxDuration(); ok {
		_spec.SetField(httpdetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := hdruo.mutation.AddedMaxDuration(); ok {
		_spec.AddField(httpdetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := hdruo.mutation.Messages(); ok {
		_spec.SetField(httpdetectorresult.FieldMessages, field.TypeJSON, value)
	}
	if value, ok := hdruo.mutation.AppendedMessages(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, httpdetectorresult.FieldMessages, value)
		})
	}
	if value, ok := hdruo.mutation.URL(); ok {
		_spec.SetField(httpdetectorresult.FieldURL, field.TypeString, value)
	}
	if value, ok := hdruo.mutation.Results(); ok {
		_spec.SetField(httpdetectorresult.FieldResults, field.TypeJSON, value)
	}
	if value, ok := hdruo.mutation.AppendedResults(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, httpdetectorresult.FieldResults, value)
		})
	}
	_spec.AddModifiers(hdruo.modifiers...)
	_node = &HTTPDetectorResult{config: hdruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hdruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{httpdetectorresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	hdruo.mutation.done = true
	return _node, nil
}