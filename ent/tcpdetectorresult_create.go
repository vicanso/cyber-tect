// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vicanso/cybertect/ent/tcpdetectorresult"
	"github.com/vicanso/cybertect/schema"
)

// TCPDetectorResultCreate is the builder for creating a TCPDetectorResult entity.
type TCPDetectorResultCreate struct {
	config
	mutation *TCPDetectorResultMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tdrc *TCPDetectorResultCreate) SetCreatedAt(t time.Time) *TCPDetectorResultCreate {
	tdrc.mutation.SetCreatedAt(t)
	return tdrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tdrc *TCPDetectorResultCreate) SetNillableCreatedAt(t *time.Time) *TCPDetectorResultCreate {
	if t != nil {
		tdrc.SetCreatedAt(*t)
	}
	return tdrc
}

// SetUpdatedAt sets the "updated_at" field.
func (tdrc *TCPDetectorResultCreate) SetUpdatedAt(t time.Time) *TCPDetectorResultCreate {
	tdrc.mutation.SetUpdatedAt(t)
	return tdrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tdrc *TCPDetectorResultCreate) SetNillableUpdatedAt(t *time.Time) *TCPDetectorResultCreate {
	if t != nil {
		tdrc.SetUpdatedAt(*t)
	}
	return tdrc
}

// SetTask sets the "task" field.
func (tdrc *TCPDetectorResultCreate) SetTask(i int) *TCPDetectorResultCreate {
	tdrc.mutation.SetTask(i)
	return tdrc
}

// SetResult sets the "result" field.
func (tdrc *TCPDetectorResultCreate) SetResult(sr schema.DetectorResult) *TCPDetectorResultCreate {
	tdrc.mutation.SetResult(sr)
	return tdrc
}

// SetMaxDuration sets the "maxDuration" field.
func (tdrc *TCPDetectorResultCreate) SetMaxDuration(i int) *TCPDetectorResultCreate {
	tdrc.mutation.SetMaxDuration(i)
	return tdrc
}

// SetMessages sets the "messages" field.
func (tdrc *TCPDetectorResultCreate) SetMessages(s []string) *TCPDetectorResultCreate {
	tdrc.mutation.SetMessages(s)
	return tdrc
}

// SetAddrs sets the "addrs" field.
func (tdrc *TCPDetectorResultCreate) SetAddrs(s []string) *TCPDetectorResultCreate {
	tdrc.mutation.SetAddrs(s)
	return tdrc
}

// SetResults sets the "results" field.
func (tdrc *TCPDetectorResultCreate) SetResults(sdsr schema.TCPDetectorSubResults) *TCPDetectorResultCreate {
	tdrc.mutation.SetResults(sdsr)
	return tdrc
}

// Mutation returns the TCPDetectorResultMutation object of the builder.
func (tdrc *TCPDetectorResultCreate) Mutation() *TCPDetectorResultMutation {
	return tdrc.mutation
}

// Save creates the TCPDetectorResult in the database.
func (tdrc *TCPDetectorResultCreate) Save(ctx context.Context) (*TCPDetectorResult, error) {
	tdrc.defaults()
	return withHooks[*TCPDetectorResult, TCPDetectorResultMutation](ctx, tdrc.sqlSave, tdrc.mutation, tdrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tdrc *TCPDetectorResultCreate) SaveX(ctx context.Context) *TCPDetectorResult {
	v, err := tdrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tdrc *TCPDetectorResultCreate) Exec(ctx context.Context) error {
	_, err := tdrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdrc *TCPDetectorResultCreate) ExecX(ctx context.Context) {
	if err := tdrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tdrc *TCPDetectorResultCreate) defaults() {
	if _, ok := tdrc.mutation.CreatedAt(); !ok {
		v := tcpdetectorresult.DefaultCreatedAt()
		tdrc.mutation.SetCreatedAt(v)
	}
	if _, ok := tdrc.mutation.UpdatedAt(); !ok {
		v := tcpdetectorresult.DefaultUpdatedAt()
		tdrc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tdrc *TCPDetectorResultCreate) check() error {
	if _, ok := tdrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TCPDetectorResult.created_at"`)}
	}
	if _, ok := tdrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TCPDetectorResult.updated_at"`)}
	}
	if _, ok := tdrc.mutation.Task(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required field "TCPDetectorResult.task"`)}
	}
	if _, ok := tdrc.mutation.Result(); !ok {
		return &ValidationError{Name: "result", err: errors.New(`ent: missing required field "TCPDetectorResult.result"`)}
	}
	if v, ok := tdrc.mutation.Result(); ok {
		if err := tcpdetectorresult.ResultValidator(int8(v)); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "TCPDetectorResult.result": %w`, err)}
		}
	}
	if _, ok := tdrc.mutation.MaxDuration(); !ok {
		return &ValidationError{Name: "maxDuration", err: errors.New(`ent: missing required field "TCPDetectorResult.maxDuration"`)}
	}
	if _, ok := tdrc.mutation.Messages(); !ok {
		return &ValidationError{Name: "messages", err: errors.New(`ent: missing required field "TCPDetectorResult.messages"`)}
	}
	if _, ok := tdrc.mutation.Addrs(); !ok {
		return &ValidationError{Name: "addrs", err: errors.New(`ent: missing required field "TCPDetectorResult.addrs"`)}
	}
	if _, ok := tdrc.mutation.Results(); !ok {
		return &ValidationError{Name: "results", err: errors.New(`ent: missing required field "TCPDetectorResult.results"`)}
	}
	return nil
}

func (tdrc *TCPDetectorResultCreate) sqlSave(ctx context.Context) (*TCPDetectorResult, error) {
	if err := tdrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tdrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tdrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tdrc.mutation.id = &_node.ID
	tdrc.mutation.done = true
	return _node, nil
}

func (tdrc *TCPDetectorResultCreate) createSpec() (*TCPDetectorResult, *sqlgraph.CreateSpec) {
	var (
		_node = &TCPDetectorResult{config: tdrc.config}
		_spec = sqlgraph.NewCreateSpec(tcpdetectorresult.Table, sqlgraph.NewFieldSpec(tcpdetectorresult.FieldID, field.TypeInt))
	)
	if value, ok := tdrc.mutation.CreatedAt(); ok {
		_spec.SetField(tcpdetectorresult.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tdrc.mutation.UpdatedAt(); ok {
		_spec.SetField(tcpdetectorresult.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tdrc.mutation.Task(); ok {
		_spec.SetField(tcpdetectorresult.FieldTask, field.TypeInt, value)
		_node.Task = value
	}
	if value, ok := tdrc.mutation.Result(); ok {
		_spec.SetField(tcpdetectorresult.FieldResult, field.TypeInt8, value)
		_node.Result = value
	}
	if value, ok := tdrc.mutation.MaxDuration(); ok {
		_spec.SetField(tcpdetectorresult.FieldMaxDuration, field.TypeInt, value)
		_node.MaxDuration = value
	}
	if value, ok := tdrc.mutation.Messages(); ok {
		_spec.SetField(tcpdetectorresult.FieldMessages, field.TypeJSON, value)
		_node.Messages = value
	}
	if value, ok := tdrc.mutation.Addrs(); ok {
		_spec.SetField(tcpdetectorresult.FieldAddrs, field.TypeJSON, value)
		_node.Addrs = value
	}
	if value, ok := tdrc.mutation.Results(); ok {
		_spec.SetField(tcpdetectorresult.FieldResults, field.TypeJSON, value)
		_node.Results = value
	}
	return _node, _spec
}

// TCPDetectorResultCreateBulk is the builder for creating many TCPDetectorResult entities in bulk.
type TCPDetectorResultCreateBulk struct {
	config
	builders []*TCPDetectorResultCreate
}

// Save creates the TCPDetectorResult entities in the database.
func (tdrcb *TCPDetectorResultCreateBulk) Save(ctx context.Context) ([]*TCPDetectorResult, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tdrcb.builders))
	nodes := make([]*TCPDetectorResult, len(tdrcb.builders))
	mutators := make([]Mutator, len(tdrcb.builders))
	for i := range tdrcb.builders {
		func(i int, root context.Context) {
			builder := tdrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TCPDetectorResultMutation)
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
					_, err = mutators[i+1].Mutate(root, tdrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tdrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tdrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tdrcb *TCPDetectorResultCreateBulk) SaveX(ctx context.Context) []*TCPDetectorResult {
	v, err := tdrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tdrcb *TCPDetectorResultCreateBulk) Exec(ctx context.Context) error {
	_, err := tdrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdrcb *TCPDetectorResultCreateBulk) ExecX(ctx context.Context) {
	if err := tdrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
