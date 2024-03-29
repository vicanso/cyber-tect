// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vicanso/cybertect/ent/httpdetectorresult"
	"github.com/vicanso/cybertect/schema"
)

// HTTPDetectorResultCreate is the builder for creating a HTTPDetectorResult entity.
type HTTPDetectorResultCreate struct {
	config
	mutation *HTTPDetectorResultMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (hdrc *HTTPDetectorResultCreate) SetCreatedAt(t time.Time) *HTTPDetectorResultCreate {
	hdrc.mutation.SetCreatedAt(t)
	return hdrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hdrc *HTTPDetectorResultCreate) SetNillableCreatedAt(t *time.Time) *HTTPDetectorResultCreate {
	if t != nil {
		hdrc.SetCreatedAt(*t)
	}
	return hdrc
}

// SetUpdatedAt sets the "updated_at" field.
func (hdrc *HTTPDetectorResultCreate) SetUpdatedAt(t time.Time) *HTTPDetectorResultCreate {
	hdrc.mutation.SetUpdatedAt(t)
	return hdrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hdrc *HTTPDetectorResultCreate) SetNillableUpdatedAt(t *time.Time) *HTTPDetectorResultCreate {
	if t != nil {
		hdrc.SetUpdatedAt(*t)
	}
	return hdrc
}

// SetTask sets the "task" field.
func (hdrc *HTTPDetectorResultCreate) SetTask(i int) *HTTPDetectorResultCreate {
	hdrc.mutation.SetTask(i)
	return hdrc
}

// SetResult sets the "result" field.
func (hdrc *HTTPDetectorResultCreate) SetResult(sr schema.DetectorResult) *HTTPDetectorResultCreate {
	hdrc.mutation.SetResult(sr)
	return hdrc
}

// SetMaxDuration sets the "maxDuration" field.
func (hdrc *HTTPDetectorResultCreate) SetMaxDuration(i int) *HTTPDetectorResultCreate {
	hdrc.mutation.SetMaxDuration(i)
	return hdrc
}

// SetMessages sets the "messages" field.
func (hdrc *HTTPDetectorResultCreate) SetMessages(s []string) *HTTPDetectorResultCreate {
	hdrc.mutation.SetMessages(s)
	return hdrc
}

// SetURL sets the "url" field.
func (hdrc *HTTPDetectorResultCreate) SetURL(s string) *HTTPDetectorResultCreate {
	hdrc.mutation.SetURL(s)
	return hdrc
}

// SetResults sets the "results" field.
func (hdrc *HTTPDetectorResultCreate) SetResults(sdsr schema.HTTPDetectorSubResults) *HTTPDetectorResultCreate {
	hdrc.mutation.SetResults(sdsr)
	return hdrc
}

// Mutation returns the HTTPDetectorResultMutation object of the builder.
func (hdrc *HTTPDetectorResultCreate) Mutation() *HTTPDetectorResultMutation {
	return hdrc.mutation
}

// Save creates the HTTPDetectorResult in the database.
func (hdrc *HTTPDetectorResultCreate) Save(ctx context.Context) (*HTTPDetectorResult, error) {
	hdrc.defaults()
	return withHooks(ctx, hdrc.sqlSave, hdrc.mutation, hdrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hdrc *HTTPDetectorResultCreate) SaveX(ctx context.Context) *HTTPDetectorResult {
	v, err := hdrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hdrc *HTTPDetectorResultCreate) Exec(ctx context.Context) error {
	_, err := hdrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hdrc *HTTPDetectorResultCreate) ExecX(ctx context.Context) {
	if err := hdrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hdrc *HTTPDetectorResultCreate) defaults() {
	if _, ok := hdrc.mutation.CreatedAt(); !ok {
		v := httpdetectorresult.DefaultCreatedAt()
		hdrc.mutation.SetCreatedAt(v)
	}
	if _, ok := hdrc.mutation.UpdatedAt(); !ok {
		v := httpdetectorresult.DefaultUpdatedAt()
		hdrc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hdrc *HTTPDetectorResultCreate) check() error {
	if _, ok := hdrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "HTTPDetectorResult.created_at"`)}
	}
	if _, ok := hdrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "HTTPDetectorResult.updated_at"`)}
	}
	if _, ok := hdrc.mutation.Task(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required field "HTTPDetectorResult.task"`)}
	}
	if _, ok := hdrc.mutation.Result(); !ok {
		return &ValidationError{Name: "result", err: errors.New(`ent: missing required field "HTTPDetectorResult.result"`)}
	}
	if v, ok := hdrc.mutation.Result(); ok {
		if err := httpdetectorresult.ResultValidator(int8(v)); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "HTTPDetectorResult.result": %w`, err)}
		}
	}
	if _, ok := hdrc.mutation.MaxDuration(); !ok {
		return &ValidationError{Name: "maxDuration", err: errors.New(`ent: missing required field "HTTPDetectorResult.maxDuration"`)}
	}
	if _, ok := hdrc.mutation.Messages(); !ok {
		return &ValidationError{Name: "messages", err: errors.New(`ent: missing required field "HTTPDetectorResult.messages"`)}
	}
	if _, ok := hdrc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "HTTPDetectorResult.url"`)}
	}
	if _, ok := hdrc.mutation.Results(); !ok {
		return &ValidationError{Name: "results", err: errors.New(`ent: missing required field "HTTPDetectorResult.results"`)}
	}
	return nil
}

func (hdrc *HTTPDetectorResultCreate) sqlSave(ctx context.Context) (*HTTPDetectorResult, error) {
	if err := hdrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hdrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hdrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	hdrc.mutation.id = &_node.ID
	hdrc.mutation.done = true
	return _node, nil
}

func (hdrc *HTTPDetectorResultCreate) createSpec() (*HTTPDetectorResult, *sqlgraph.CreateSpec) {
	var (
		_node = &HTTPDetectorResult{config: hdrc.config}
		_spec = sqlgraph.NewCreateSpec(httpdetectorresult.Table, sqlgraph.NewFieldSpec(httpdetectorresult.FieldID, field.TypeInt))
	)
	if value, ok := hdrc.mutation.CreatedAt(); ok {
		_spec.SetField(httpdetectorresult.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := hdrc.mutation.UpdatedAt(); ok {
		_spec.SetField(httpdetectorresult.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := hdrc.mutation.Task(); ok {
		_spec.SetField(httpdetectorresult.FieldTask, field.TypeInt, value)
		_node.Task = value
	}
	if value, ok := hdrc.mutation.Result(); ok {
		_spec.SetField(httpdetectorresult.FieldResult, field.TypeInt8, value)
		_node.Result = value
	}
	if value, ok := hdrc.mutation.MaxDuration(); ok {
		_spec.SetField(httpdetectorresult.FieldMaxDuration, field.TypeInt, value)
		_node.MaxDuration = value
	}
	if value, ok := hdrc.mutation.Messages(); ok {
		_spec.SetField(httpdetectorresult.FieldMessages, field.TypeJSON, value)
		_node.Messages = value
	}
	if value, ok := hdrc.mutation.URL(); ok {
		_spec.SetField(httpdetectorresult.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := hdrc.mutation.Results(); ok {
		_spec.SetField(httpdetectorresult.FieldResults, field.TypeJSON, value)
		_node.Results = value
	}
	return _node, _spec
}

// HTTPDetectorResultCreateBulk is the builder for creating many HTTPDetectorResult entities in bulk.
type HTTPDetectorResultCreateBulk struct {
	config
	builders []*HTTPDetectorResultCreate
}

// Save creates the HTTPDetectorResult entities in the database.
func (hdrcb *HTTPDetectorResultCreateBulk) Save(ctx context.Context) ([]*HTTPDetectorResult, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hdrcb.builders))
	nodes := make([]*HTTPDetectorResult, len(hdrcb.builders))
	mutators := make([]Mutator, len(hdrcb.builders))
	for i := range hdrcb.builders {
		func(i int, root context.Context) {
			builder := hdrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HTTPDetectorResultMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hdrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hdrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hdrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hdrcb *HTTPDetectorResultCreateBulk) SaveX(ctx context.Context) []*HTTPDetectorResult {
	v, err := hdrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hdrcb *HTTPDetectorResultCreateBulk) Exec(ctx context.Context) error {
	_, err := hdrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hdrcb *HTTPDetectorResultCreateBulk) ExecX(ctx context.Context) {
	if err := hdrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
