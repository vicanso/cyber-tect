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
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/ent/tcpdetectorresult"
	"github.com/vicanso/cybertect/schema"
)

// TCPDetectorResultUpdate is the builder for updating TCPDetectorResult entities.
type TCPDetectorResultUpdate struct {
	config
	hooks     []Hook
	mutation  *TCPDetectorResultMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TCPDetectorResultUpdate builder.
func (tdru *TCPDetectorResultUpdate) Where(ps ...predicate.TCPDetectorResult) *TCPDetectorResultUpdate {
	tdru.mutation.Where(ps...)
	return tdru
}

// SetTask sets the "task" field.
func (tdru *TCPDetectorResultUpdate) SetTask(i int) *TCPDetectorResultUpdate {
	tdru.mutation.ResetTask()
	tdru.mutation.SetTask(i)
	return tdru
}

// AddTask adds i to the "task" field.
func (tdru *TCPDetectorResultUpdate) AddTask(i int) *TCPDetectorResultUpdate {
	tdru.mutation.AddTask(i)
	return tdru
}

// SetResult sets the "result" field.
func (tdru *TCPDetectorResultUpdate) SetResult(sr schema.DetectorResult) *TCPDetectorResultUpdate {
	tdru.mutation.ResetResult()
	tdru.mutation.SetResult(sr)
	return tdru
}

// AddResult adds sr to the "result" field.
func (tdru *TCPDetectorResultUpdate) AddResult(sr schema.DetectorResult) *TCPDetectorResultUpdate {
	tdru.mutation.AddResult(sr)
	return tdru
}

// SetMaxDuration sets the "maxDuration" field.
func (tdru *TCPDetectorResultUpdate) SetMaxDuration(i int) *TCPDetectorResultUpdate {
	tdru.mutation.ResetMaxDuration()
	tdru.mutation.SetMaxDuration(i)
	return tdru
}

// AddMaxDuration adds i to the "maxDuration" field.
func (tdru *TCPDetectorResultUpdate) AddMaxDuration(i int) *TCPDetectorResultUpdate {
	tdru.mutation.AddMaxDuration(i)
	return tdru
}

// SetMessages sets the "messages" field.
func (tdru *TCPDetectorResultUpdate) SetMessages(s []string) *TCPDetectorResultUpdate {
	tdru.mutation.SetMessages(s)
	return tdru
}

// AppendMessages appends s to the "messages" field.
func (tdru *TCPDetectorResultUpdate) AppendMessages(s []string) *TCPDetectorResultUpdate {
	tdru.mutation.AppendMessages(s)
	return tdru
}

// SetAddrs sets the "addrs" field.
func (tdru *TCPDetectorResultUpdate) SetAddrs(s []string) *TCPDetectorResultUpdate {
	tdru.mutation.SetAddrs(s)
	return tdru
}

// AppendAddrs appends s to the "addrs" field.
func (tdru *TCPDetectorResultUpdate) AppendAddrs(s []string) *TCPDetectorResultUpdate {
	tdru.mutation.AppendAddrs(s)
	return tdru
}

// SetResults sets the "results" field.
func (tdru *TCPDetectorResultUpdate) SetResults(sdsr schema.TCPDetectorSubResults) *TCPDetectorResultUpdate {
	tdru.mutation.SetResults(sdsr)
	return tdru
}

// AppendResults appends sdsr to the "results" field.
func (tdru *TCPDetectorResultUpdate) AppendResults(sdsr schema.TCPDetectorSubResults) *TCPDetectorResultUpdate {
	tdru.mutation.AppendResults(sdsr)
	return tdru
}

// Mutation returns the TCPDetectorResultMutation object of the builder.
func (tdru *TCPDetectorResultUpdate) Mutation() *TCPDetectorResultMutation {
	return tdru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tdru *TCPDetectorResultUpdate) Save(ctx context.Context) (int, error) {
	tdru.defaults()
	return withHooks(ctx, tdru.sqlSave, tdru.mutation, tdru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tdru *TCPDetectorResultUpdate) SaveX(ctx context.Context) int {
	affected, err := tdru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tdru *TCPDetectorResultUpdate) Exec(ctx context.Context) error {
	_, err := tdru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdru *TCPDetectorResultUpdate) ExecX(ctx context.Context) {
	if err := tdru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tdru *TCPDetectorResultUpdate) defaults() {
	if _, ok := tdru.mutation.UpdatedAt(); !ok {
		v := tcpdetectorresult.UpdateDefaultUpdatedAt()
		tdru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tdru *TCPDetectorResultUpdate) check() error {
	if v, ok := tdru.mutation.Result(); ok {
		if err := tcpdetectorresult.ResultValidator(int8(v)); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "TCPDetectorResult.result": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tdru *TCPDetectorResultUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TCPDetectorResultUpdate {
	tdru.modifiers = append(tdru.modifiers, modifiers...)
	return tdru
}

func (tdru *TCPDetectorResultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tdru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tcpdetectorresult.Table, tcpdetectorresult.Columns, sqlgraph.NewFieldSpec(tcpdetectorresult.FieldID, field.TypeInt))
	if ps := tdru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tdru.mutation.UpdatedAt(); ok {
		_spec.SetField(tcpdetectorresult.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tdru.mutation.Task(); ok {
		_spec.SetField(tcpdetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := tdru.mutation.AddedTask(); ok {
		_spec.AddField(tcpdetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := tdru.mutation.Result(); ok {
		_spec.SetField(tcpdetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := tdru.mutation.AddedResult(); ok {
		_spec.AddField(tcpdetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := tdru.mutation.MaxDuration(); ok {
		_spec.SetField(tcpdetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := tdru.mutation.AddedMaxDuration(); ok {
		_spec.AddField(tcpdetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := tdru.mutation.Messages(); ok {
		_spec.SetField(tcpdetectorresult.FieldMessages, field.TypeJSON, value)
	}
	if value, ok := tdru.mutation.AppendedMessages(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, tcpdetectorresult.FieldMessages, value)
		})
	}
	if value, ok := tdru.mutation.Addrs(); ok {
		_spec.SetField(tcpdetectorresult.FieldAddrs, field.TypeJSON, value)
	}
	if value, ok := tdru.mutation.AppendedAddrs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, tcpdetectorresult.FieldAddrs, value)
		})
	}
	if value, ok := tdru.mutation.Results(); ok {
		_spec.SetField(tcpdetectorresult.FieldResults, field.TypeJSON, value)
	}
	if value, ok := tdru.mutation.AppendedResults(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, tcpdetectorresult.FieldResults, value)
		})
	}
	_spec.AddModifiers(tdru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tdru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tcpdetectorresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tdru.mutation.done = true
	return n, nil
}

// TCPDetectorResultUpdateOne is the builder for updating a single TCPDetectorResult entity.
type TCPDetectorResultUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TCPDetectorResultMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetTask sets the "task" field.
func (tdruo *TCPDetectorResultUpdateOne) SetTask(i int) *TCPDetectorResultUpdateOne {
	tdruo.mutation.ResetTask()
	tdruo.mutation.SetTask(i)
	return tdruo
}

// AddTask adds i to the "task" field.
func (tdruo *TCPDetectorResultUpdateOne) AddTask(i int) *TCPDetectorResultUpdateOne {
	tdruo.mutation.AddTask(i)
	return tdruo
}

// SetResult sets the "result" field.
func (tdruo *TCPDetectorResultUpdateOne) SetResult(sr schema.DetectorResult) *TCPDetectorResultUpdateOne {
	tdruo.mutation.ResetResult()
	tdruo.mutation.SetResult(sr)
	return tdruo
}

// AddResult adds sr to the "result" field.
func (tdruo *TCPDetectorResultUpdateOne) AddResult(sr schema.DetectorResult) *TCPDetectorResultUpdateOne {
	tdruo.mutation.AddResult(sr)
	return tdruo
}

// SetMaxDuration sets the "maxDuration" field.
func (tdruo *TCPDetectorResultUpdateOne) SetMaxDuration(i int) *TCPDetectorResultUpdateOne {
	tdruo.mutation.ResetMaxDuration()
	tdruo.mutation.SetMaxDuration(i)
	return tdruo
}

// AddMaxDuration adds i to the "maxDuration" field.
func (tdruo *TCPDetectorResultUpdateOne) AddMaxDuration(i int) *TCPDetectorResultUpdateOne {
	tdruo.mutation.AddMaxDuration(i)
	return tdruo
}

// SetMessages sets the "messages" field.
func (tdruo *TCPDetectorResultUpdateOne) SetMessages(s []string) *TCPDetectorResultUpdateOne {
	tdruo.mutation.SetMessages(s)
	return tdruo
}

// AppendMessages appends s to the "messages" field.
func (tdruo *TCPDetectorResultUpdateOne) AppendMessages(s []string) *TCPDetectorResultUpdateOne {
	tdruo.mutation.AppendMessages(s)
	return tdruo
}

// SetAddrs sets the "addrs" field.
func (tdruo *TCPDetectorResultUpdateOne) SetAddrs(s []string) *TCPDetectorResultUpdateOne {
	tdruo.mutation.SetAddrs(s)
	return tdruo
}

// AppendAddrs appends s to the "addrs" field.
func (tdruo *TCPDetectorResultUpdateOne) AppendAddrs(s []string) *TCPDetectorResultUpdateOne {
	tdruo.mutation.AppendAddrs(s)
	return tdruo
}

// SetResults sets the "results" field.
func (tdruo *TCPDetectorResultUpdateOne) SetResults(sdsr schema.TCPDetectorSubResults) *TCPDetectorResultUpdateOne {
	tdruo.mutation.SetResults(sdsr)
	return tdruo
}

// AppendResults appends sdsr to the "results" field.
func (tdruo *TCPDetectorResultUpdateOne) AppendResults(sdsr schema.TCPDetectorSubResults) *TCPDetectorResultUpdateOne {
	tdruo.mutation.AppendResults(sdsr)
	return tdruo
}

// Mutation returns the TCPDetectorResultMutation object of the builder.
func (tdruo *TCPDetectorResultUpdateOne) Mutation() *TCPDetectorResultMutation {
	return tdruo.mutation
}

// Where appends a list predicates to the TCPDetectorResultUpdate builder.
func (tdruo *TCPDetectorResultUpdateOne) Where(ps ...predicate.TCPDetectorResult) *TCPDetectorResultUpdateOne {
	tdruo.mutation.Where(ps...)
	return tdruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tdruo *TCPDetectorResultUpdateOne) Select(field string, fields ...string) *TCPDetectorResultUpdateOne {
	tdruo.fields = append([]string{field}, fields...)
	return tdruo
}

// Save executes the query and returns the updated TCPDetectorResult entity.
func (tdruo *TCPDetectorResultUpdateOne) Save(ctx context.Context) (*TCPDetectorResult, error) {
	tdruo.defaults()
	return withHooks(ctx, tdruo.sqlSave, tdruo.mutation, tdruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tdruo *TCPDetectorResultUpdateOne) SaveX(ctx context.Context) *TCPDetectorResult {
	node, err := tdruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tdruo *TCPDetectorResultUpdateOne) Exec(ctx context.Context) error {
	_, err := tdruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdruo *TCPDetectorResultUpdateOne) ExecX(ctx context.Context) {
	if err := tdruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tdruo *TCPDetectorResultUpdateOne) defaults() {
	if _, ok := tdruo.mutation.UpdatedAt(); !ok {
		v := tcpdetectorresult.UpdateDefaultUpdatedAt()
		tdruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tdruo *TCPDetectorResultUpdateOne) check() error {
	if v, ok := tdruo.mutation.Result(); ok {
		if err := tcpdetectorresult.ResultValidator(int8(v)); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "TCPDetectorResult.result": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tdruo *TCPDetectorResultUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TCPDetectorResultUpdateOne {
	tdruo.modifiers = append(tdruo.modifiers, modifiers...)
	return tdruo
}

func (tdruo *TCPDetectorResultUpdateOne) sqlSave(ctx context.Context) (_node *TCPDetectorResult, err error) {
	if err := tdruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tcpdetectorresult.Table, tcpdetectorresult.Columns, sqlgraph.NewFieldSpec(tcpdetectorresult.FieldID, field.TypeInt))
	id, ok := tdruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TCPDetectorResult.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tdruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tcpdetectorresult.FieldID)
		for _, f := range fields {
			if !tcpdetectorresult.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tcpdetectorresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tdruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tdruo.mutation.UpdatedAt(); ok {
		_spec.SetField(tcpdetectorresult.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tdruo.mutation.Task(); ok {
		_spec.SetField(tcpdetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := tdruo.mutation.AddedTask(); ok {
		_spec.AddField(tcpdetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := tdruo.mutation.Result(); ok {
		_spec.SetField(tcpdetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := tdruo.mutation.AddedResult(); ok {
		_spec.AddField(tcpdetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := tdruo.mutation.MaxDuration(); ok {
		_spec.SetField(tcpdetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := tdruo.mutation.AddedMaxDuration(); ok {
		_spec.AddField(tcpdetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := tdruo.mutation.Messages(); ok {
		_spec.SetField(tcpdetectorresult.FieldMessages, field.TypeJSON, value)
	}
	if value, ok := tdruo.mutation.AppendedMessages(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, tcpdetectorresult.FieldMessages, value)
		})
	}
	if value, ok := tdruo.mutation.Addrs(); ok {
		_spec.SetField(tcpdetectorresult.FieldAddrs, field.TypeJSON, value)
	}
	if value, ok := tdruo.mutation.AppendedAddrs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, tcpdetectorresult.FieldAddrs, value)
		})
	}
	if value, ok := tdruo.mutation.Results(); ok {
		_spec.SetField(tcpdetectorresult.FieldResults, field.TypeJSON, value)
	}
	if value, ok := tdruo.mutation.AppendedResults(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, tcpdetectorresult.FieldResults, value)
		})
	}
	_spec.AddModifiers(tdruo.modifiers...)
	_node = &TCPDetectorResult{config: tdruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tdruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tcpdetectorresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tdruo.mutation.done = true
	return _node, nil
}
