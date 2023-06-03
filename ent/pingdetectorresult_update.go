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
	"github.com/vicanso/cybertect/ent/pingdetectorresult"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/schema"
)

// PingDetectorResultUpdate is the builder for updating PingDetectorResult entities.
type PingDetectorResultUpdate struct {
	config
	hooks     []Hook
	mutation  *PingDetectorResultMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PingDetectorResultUpdate builder.
func (pdru *PingDetectorResultUpdate) Where(ps ...predicate.PingDetectorResult) *PingDetectorResultUpdate {
	pdru.mutation.Where(ps...)
	return pdru
}

// SetTask sets the "task" field.
func (pdru *PingDetectorResultUpdate) SetTask(i int) *PingDetectorResultUpdate {
	pdru.mutation.ResetTask()
	pdru.mutation.SetTask(i)
	return pdru
}

// AddTask adds i to the "task" field.
func (pdru *PingDetectorResultUpdate) AddTask(i int) *PingDetectorResultUpdate {
	pdru.mutation.AddTask(i)
	return pdru
}

// SetResult sets the "result" field.
func (pdru *PingDetectorResultUpdate) SetResult(sr schema.DetectorResult) *PingDetectorResultUpdate {
	pdru.mutation.ResetResult()
	pdru.mutation.SetResult(sr)
	return pdru
}

// AddResult adds sr to the "result" field.
func (pdru *PingDetectorResultUpdate) AddResult(sr schema.DetectorResult) *PingDetectorResultUpdate {
	pdru.mutation.AddResult(sr)
	return pdru
}

// SetMaxDuration sets the "maxDuration" field.
func (pdru *PingDetectorResultUpdate) SetMaxDuration(i int) *PingDetectorResultUpdate {
	pdru.mutation.ResetMaxDuration()
	pdru.mutation.SetMaxDuration(i)
	return pdru
}

// AddMaxDuration adds i to the "maxDuration" field.
func (pdru *PingDetectorResultUpdate) AddMaxDuration(i int) *PingDetectorResultUpdate {
	pdru.mutation.AddMaxDuration(i)
	return pdru
}

// SetMessages sets the "messages" field.
func (pdru *PingDetectorResultUpdate) SetMessages(s []string) *PingDetectorResultUpdate {
	pdru.mutation.SetMessages(s)
	return pdru
}

// AppendMessages appends s to the "messages" field.
func (pdru *PingDetectorResultUpdate) AppendMessages(s []string) *PingDetectorResultUpdate {
	pdru.mutation.AppendMessages(s)
	return pdru
}

// SetIps sets the "ips" field.
func (pdru *PingDetectorResultUpdate) SetIps(s []string) *PingDetectorResultUpdate {
	pdru.mutation.SetIps(s)
	return pdru
}

// AppendIps appends s to the "ips" field.
func (pdru *PingDetectorResultUpdate) AppendIps(s []string) *PingDetectorResultUpdate {
	pdru.mutation.AppendIps(s)
	return pdru
}

// SetResults sets the "results" field.
func (pdru *PingDetectorResultUpdate) SetResults(sdsr schema.PingDetectorSubResults) *PingDetectorResultUpdate {
	pdru.mutation.SetResults(sdsr)
	return pdru
}

// AppendResults appends sdsr to the "results" field.
func (pdru *PingDetectorResultUpdate) AppendResults(sdsr schema.PingDetectorSubResults) *PingDetectorResultUpdate {
	pdru.mutation.AppendResults(sdsr)
	return pdru
}

// Mutation returns the PingDetectorResultMutation object of the builder.
func (pdru *PingDetectorResultUpdate) Mutation() *PingDetectorResultMutation {
	return pdru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pdru *PingDetectorResultUpdate) Save(ctx context.Context) (int, error) {
	pdru.defaults()
	return withHooks(ctx, pdru.sqlSave, pdru.mutation, pdru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pdru *PingDetectorResultUpdate) SaveX(ctx context.Context) int {
	affected, err := pdru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pdru *PingDetectorResultUpdate) Exec(ctx context.Context) error {
	_, err := pdru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdru *PingDetectorResultUpdate) ExecX(ctx context.Context) {
	if err := pdru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pdru *PingDetectorResultUpdate) defaults() {
	if _, ok := pdru.mutation.UpdatedAt(); !ok {
		v := pingdetectorresult.UpdateDefaultUpdatedAt()
		pdru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pdru *PingDetectorResultUpdate) check() error {
	if v, ok := pdru.mutation.Result(); ok {
		if err := pingdetectorresult.ResultValidator(int8(v)); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "PingDetectorResult.result": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pdru *PingDetectorResultUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PingDetectorResultUpdate {
	pdru.modifiers = append(pdru.modifiers, modifiers...)
	return pdru
}

func (pdru *PingDetectorResultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pdru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(pingdetectorresult.Table, pingdetectorresult.Columns, sqlgraph.NewFieldSpec(pingdetectorresult.FieldID, field.TypeInt))
	if ps := pdru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pdru.mutation.UpdatedAt(); ok {
		_spec.SetField(pingdetectorresult.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pdru.mutation.Task(); ok {
		_spec.SetField(pingdetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := pdru.mutation.AddedTask(); ok {
		_spec.AddField(pingdetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := pdru.mutation.Result(); ok {
		_spec.SetField(pingdetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := pdru.mutation.AddedResult(); ok {
		_spec.AddField(pingdetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := pdru.mutation.MaxDuration(); ok {
		_spec.SetField(pingdetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := pdru.mutation.AddedMaxDuration(); ok {
		_spec.AddField(pingdetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := pdru.mutation.Messages(); ok {
		_spec.SetField(pingdetectorresult.FieldMessages, field.TypeJSON, value)
	}
	if value, ok := pdru.mutation.AppendedMessages(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, pingdetectorresult.FieldMessages, value)
		})
	}
	if value, ok := pdru.mutation.Ips(); ok {
		_spec.SetField(pingdetectorresult.FieldIps, field.TypeJSON, value)
	}
	if value, ok := pdru.mutation.AppendedIps(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, pingdetectorresult.FieldIps, value)
		})
	}
	if value, ok := pdru.mutation.Results(); ok {
		_spec.SetField(pingdetectorresult.FieldResults, field.TypeJSON, value)
	}
	if value, ok := pdru.mutation.AppendedResults(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, pingdetectorresult.FieldResults, value)
		})
	}
	_spec.AddModifiers(pdru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pdru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pingdetectorresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pdru.mutation.done = true
	return n, nil
}

// PingDetectorResultUpdateOne is the builder for updating a single PingDetectorResult entity.
type PingDetectorResultUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PingDetectorResultMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetTask sets the "task" field.
func (pdruo *PingDetectorResultUpdateOne) SetTask(i int) *PingDetectorResultUpdateOne {
	pdruo.mutation.ResetTask()
	pdruo.mutation.SetTask(i)
	return pdruo
}

// AddTask adds i to the "task" field.
func (pdruo *PingDetectorResultUpdateOne) AddTask(i int) *PingDetectorResultUpdateOne {
	pdruo.mutation.AddTask(i)
	return pdruo
}

// SetResult sets the "result" field.
func (pdruo *PingDetectorResultUpdateOne) SetResult(sr schema.DetectorResult) *PingDetectorResultUpdateOne {
	pdruo.mutation.ResetResult()
	pdruo.mutation.SetResult(sr)
	return pdruo
}

// AddResult adds sr to the "result" field.
func (pdruo *PingDetectorResultUpdateOne) AddResult(sr schema.DetectorResult) *PingDetectorResultUpdateOne {
	pdruo.mutation.AddResult(sr)
	return pdruo
}

// SetMaxDuration sets the "maxDuration" field.
func (pdruo *PingDetectorResultUpdateOne) SetMaxDuration(i int) *PingDetectorResultUpdateOne {
	pdruo.mutation.ResetMaxDuration()
	pdruo.mutation.SetMaxDuration(i)
	return pdruo
}

// AddMaxDuration adds i to the "maxDuration" field.
func (pdruo *PingDetectorResultUpdateOne) AddMaxDuration(i int) *PingDetectorResultUpdateOne {
	pdruo.mutation.AddMaxDuration(i)
	return pdruo
}

// SetMessages sets the "messages" field.
func (pdruo *PingDetectorResultUpdateOne) SetMessages(s []string) *PingDetectorResultUpdateOne {
	pdruo.mutation.SetMessages(s)
	return pdruo
}

// AppendMessages appends s to the "messages" field.
func (pdruo *PingDetectorResultUpdateOne) AppendMessages(s []string) *PingDetectorResultUpdateOne {
	pdruo.mutation.AppendMessages(s)
	return pdruo
}

// SetIps sets the "ips" field.
func (pdruo *PingDetectorResultUpdateOne) SetIps(s []string) *PingDetectorResultUpdateOne {
	pdruo.mutation.SetIps(s)
	return pdruo
}

// AppendIps appends s to the "ips" field.
func (pdruo *PingDetectorResultUpdateOne) AppendIps(s []string) *PingDetectorResultUpdateOne {
	pdruo.mutation.AppendIps(s)
	return pdruo
}

// SetResults sets the "results" field.
func (pdruo *PingDetectorResultUpdateOne) SetResults(sdsr schema.PingDetectorSubResults) *PingDetectorResultUpdateOne {
	pdruo.mutation.SetResults(sdsr)
	return pdruo
}

// AppendResults appends sdsr to the "results" field.
func (pdruo *PingDetectorResultUpdateOne) AppendResults(sdsr schema.PingDetectorSubResults) *PingDetectorResultUpdateOne {
	pdruo.mutation.AppendResults(sdsr)
	return pdruo
}

// Mutation returns the PingDetectorResultMutation object of the builder.
func (pdruo *PingDetectorResultUpdateOne) Mutation() *PingDetectorResultMutation {
	return pdruo.mutation
}

// Where appends a list predicates to the PingDetectorResultUpdate builder.
func (pdruo *PingDetectorResultUpdateOne) Where(ps ...predicate.PingDetectorResult) *PingDetectorResultUpdateOne {
	pdruo.mutation.Where(ps...)
	return pdruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pdruo *PingDetectorResultUpdateOne) Select(field string, fields ...string) *PingDetectorResultUpdateOne {
	pdruo.fields = append([]string{field}, fields...)
	return pdruo
}

// Save executes the query and returns the updated PingDetectorResult entity.
func (pdruo *PingDetectorResultUpdateOne) Save(ctx context.Context) (*PingDetectorResult, error) {
	pdruo.defaults()
	return withHooks(ctx, pdruo.sqlSave, pdruo.mutation, pdruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pdruo *PingDetectorResultUpdateOne) SaveX(ctx context.Context) *PingDetectorResult {
	node, err := pdruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pdruo *PingDetectorResultUpdateOne) Exec(ctx context.Context) error {
	_, err := pdruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdruo *PingDetectorResultUpdateOne) ExecX(ctx context.Context) {
	if err := pdruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pdruo *PingDetectorResultUpdateOne) defaults() {
	if _, ok := pdruo.mutation.UpdatedAt(); !ok {
		v := pingdetectorresult.UpdateDefaultUpdatedAt()
		pdruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pdruo *PingDetectorResultUpdateOne) check() error {
	if v, ok := pdruo.mutation.Result(); ok {
		if err := pingdetectorresult.ResultValidator(int8(v)); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "PingDetectorResult.result": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pdruo *PingDetectorResultUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PingDetectorResultUpdateOne {
	pdruo.modifiers = append(pdruo.modifiers, modifiers...)
	return pdruo
}

func (pdruo *PingDetectorResultUpdateOne) sqlSave(ctx context.Context) (_node *PingDetectorResult, err error) {
	if err := pdruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(pingdetectorresult.Table, pingdetectorresult.Columns, sqlgraph.NewFieldSpec(pingdetectorresult.FieldID, field.TypeInt))
	id, ok := pdruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PingDetectorResult.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pdruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pingdetectorresult.FieldID)
		for _, f := range fields {
			if !pingdetectorresult.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pingdetectorresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pdruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pdruo.mutation.UpdatedAt(); ok {
		_spec.SetField(pingdetectorresult.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pdruo.mutation.Task(); ok {
		_spec.SetField(pingdetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := pdruo.mutation.AddedTask(); ok {
		_spec.AddField(pingdetectorresult.FieldTask, field.TypeInt, value)
	}
	if value, ok := pdruo.mutation.Result(); ok {
		_spec.SetField(pingdetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := pdruo.mutation.AddedResult(); ok {
		_spec.AddField(pingdetectorresult.FieldResult, field.TypeInt8, value)
	}
	if value, ok := pdruo.mutation.MaxDuration(); ok {
		_spec.SetField(pingdetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := pdruo.mutation.AddedMaxDuration(); ok {
		_spec.AddField(pingdetectorresult.FieldMaxDuration, field.TypeInt, value)
	}
	if value, ok := pdruo.mutation.Messages(); ok {
		_spec.SetField(pingdetectorresult.FieldMessages, field.TypeJSON, value)
	}
	if value, ok := pdruo.mutation.AppendedMessages(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, pingdetectorresult.FieldMessages, value)
		})
	}
	if value, ok := pdruo.mutation.Ips(); ok {
		_spec.SetField(pingdetectorresult.FieldIps, field.TypeJSON, value)
	}
	if value, ok := pdruo.mutation.AppendedIps(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, pingdetectorresult.FieldIps, value)
		})
	}
	if value, ok := pdruo.mutation.Results(); ok {
		_spec.SetField(pingdetectorresult.FieldResults, field.TypeJSON, value)
	}
	if value, ok := pdruo.mutation.AppendedResults(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, pingdetectorresult.FieldResults, value)
		})
	}
	_spec.AddModifiers(pdruo.modifiers...)
	_node = &PingDetectorResult{config: pdruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pdruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pingdetectorresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pdruo.mutation.done = true
	return _node, nil
}
