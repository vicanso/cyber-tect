// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vicanso/cybertect/ent/databasedetectorresult"
	"github.com/vicanso/cybertect/schema"
)

// DatabaseDetectorResultCreate is the builder for creating a DatabaseDetectorResult entity.
type DatabaseDetectorResultCreate struct {
	config
	mutation *DatabaseDetectorResultMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ddrc *DatabaseDetectorResultCreate) SetCreatedAt(t time.Time) *DatabaseDetectorResultCreate {
	ddrc.mutation.SetCreatedAt(t)
	return ddrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ddrc *DatabaseDetectorResultCreate) SetNillableCreatedAt(t *time.Time) *DatabaseDetectorResultCreate {
	if t != nil {
		ddrc.SetCreatedAt(*t)
	}
	return ddrc
}

// SetUpdatedAt sets the "updated_at" field.
func (ddrc *DatabaseDetectorResultCreate) SetUpdatedAt(t time.Time) *DatabaseDetectorResultCreate {
	ddrc.mutation.SetUpdatedAt(t)
	return ddrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ddrc *DatabaseDetectorResultCreate) SetNillableUpdatedAt(t *time.Time) *DatabaseDetectorResultCreate {
	if t != nil {
		ddrc.SetUpdatedAt(*t)
	}
	return ddrc
}

// SetTask sets the "task" field.
func (ddrc *DatabaseDetectorResultCreate) SetTask(i int) *DatabaseDetectorResultCreate {
	ddrc.mutation.SetTask(i)
	return ddrc
}

// SetResult sets the "result" field.
func (ddrc *DatabaseDetectorResultCreate) SetResult(sr schema.DetectorResult) *DatabaseDetectorResultCreate {
	ddrc.mutation.SetResult(sr)
	return ddrc
}

// SetMaxDuration sets the "maxDuration" field.
func (ddrc *DatabaseDetectorResultCreate) SetMaxDuration(i int) *DatabaseDetectorResultCreate {
	ddrc.mutation.SetMaxDuration(i)
	return ddrc
}

// SetMessages sets the "messages" field.
func (ddrc *DatabaseDetectorResultCreate) SetMessages(s []string) *DatabaseDetectorResultCreate {
	ddrc.mutation.SetMessages(s)
	return ddrc
}

// SetUris sets the "uris" field.
func (ddrc *DatabaseDetectorResultCreate) SetUris(s []string) *DatabaseDetectorResultCreate {
	ddrc.mutation.SetUris(s)
	return ddrc
}

// SetResults sets the "results" field.
func (ddrc *DatabaseDetectorResultCreate) SetResults(sdsr schema.DatabaseDetectorSubResults) *DatabaseDetectorResultCreate {
	ddrc.mutation.SetResults(sdsr)
	return ddrc
}

// Mutation returns the DatabaseDetectorResultMutation object of the builder.
func (ddrc *DatabaseDetectorResultCreate) Mutation() *DatabaseDetectorResultMutation {
	return ddrc.mutation
}

// Save creates the DatabaseDetectorResult in the database.
func (ddrc *DatabaseDetectorResultCreate) Save(ctx context.Context) (*DatabaseDetectorResult, error) {
	ddrc.defaults()
	return withHooks[*DatabaseDetectorResult, DatabaseDetectorResultMutation](ctx, ddrc.sqlSave, ddrc.mutation, ddrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ddrc *DatabaseDetectorResultCreate) SaveX(ctx context.Context) *DatabaseDetectorResult {
	v, err := ddrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ddrc *DatabaseDetectorResultCreate) Exec(ctx context.Context) error {
	_, err := ddrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddrc *DatabaseDetectorResultCreate) ExecX(ctx context.Context) {
	if err := ddrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ddrc *DatabaseDetectorResultCreate) defaults() {
	if _, ok := ddrc.mutation.CreatedAt(); !ok {
		v := databasedetectorresult.DefaultCreatedAt()
		ddrc.mutation.SetCreatedAt(v)
	}
	if _, ok := ddrc.mutation.UpdatedAt(); !ok {
		v := databasedetectorresult.DefaultUpdatedAt()
		ddrc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ddrc *DatabaseDetectorResultCreate) check() error {
	if _, ok := ddrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DatabaseDetectorResult.created_at"`)}
	}
	if _, ok := ddrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DatabaseDetectorResult.updated_at"`)}
	}
	if _, ok := ddrc.mutation.Task(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required field "DatabaseDetectorResult.task"`)}
	}
	if _, ok := ddrc.mutation.Result(); !ok {
		return &ValidationError{Name: "result", err: errors.New(`ent: missing required field "DatabaseDetectorResult.result"`)}
	}
	if v, ok := ddrc.mutation.Result(); ok {
		if err := databasedetectorresult.ResultValidator(int8(v)); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf(`ent: validator failed for field "DatabaseDetectorResult.result": %w`, err)}
		}
	}
	if _, ok := ddrc.mutation.MaxDuration(); !ok {
		return &ValidationError{Name: "maxDuration", err: errors.New(`ent: missing required field "DatabaseDetectorResult.maxDuration"`)}
	}
	if _, ok := ddrc.mutation.Messages(); !ok {
		return &ValidationError{Name: "messages", err: errors.New(`ent: missing required field "DatabaseDetectorResult.messages"`)}
	}
	if _, ok := ddrc.mutation.Uris(); !ok {
		return &ValidationError{Name: "uris", err: errors.New(`ent: missing required field "DatabaseDetectorResult.uris"`)}
	}
	if _, ok := ddrc.mutation.Results(); !ok {
		return &ValidationError{Name: "results", err: errors.New(`ent: missing required field "DatabaseDetectorResult.results"`)}
	}
	return nil
}

func (ddrc *DatabaseDetectorResultCreate) sqlSave(ctx context.Context) (*DatabaseDetectorResult, error) {
	if err := ddrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ddrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ddrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ddrc.mutation.id = &_node.ID
	ddrc.mutation.done = true
	return _node, nil
}

func (ddrc *DatabaseDetectorResultCreate) createSpec() (*DatabaseDetectorResult, *sqlgraph.CreateSpec) {
	var (
		_node = &DatabaseDetectorResult{config: ddrc.config}
		_spec = sqlgraph.NewCreateSpec(databasedetectorresult.Table, sqlgraph.NewFieldSpec(databasedetectorresult.FieldID, field.TypeInt))
	)
	if value, ok := ddrc.mutation.CreatedAt(); ok {
		_spec.SetField(databasedetectorresult.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ddrc.mutation.UpdatedAt(); ok {
		_spec.SetField(databasedetectorresult.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ddrc.mutation.Task(); ok {
		_spec.SetField(databasedetectorresult.FieldTask, field.TypeInt, value)
		_node.Task = value
	}
	if value, ok := ddrc.mutation.Result(); ok {
		_spec.SetField(databasedetectorresult.FieldResult, field.TypeInt8, value)
		_node.Result = value
	}
	if value, ok := ddrc.mutation.MaxDuration(); ok {
		_spec.SetField(databasedetectorresult.FieldMaxDuration, field.TypeInt, value)
		_node.MaxDuration = value
	}
	if value, ok := ddrc.mutation.Messages(); ok {
		_spec.SetField(databasedetectorresult.FieldMessages, field.TypeJSON, value)
		_node.Messages = value
	}
	if value, ok := ddrc.mutation.Uris(); ok {
		_spec.SetField(databasedetectorresult.FieldUris, field.TypeJSON, value)
		_node.Uris = value
	}
	if value, ok := ddrc.mutation.Results(); ok {
		_spec.SetField(databasedetectorresult.FieldResults, field.TypeJSON, value)
		_node.Results = value
	}
	return _node, _spec
}

// DatabaseDetectorResultCreateBulk is the builder for creating many DatabaseDetectorResult entities in bulk.
type DatabaseDetectorResultCreateBulk struct {
	config
	builders []*DatabaseDetectorResultCreate
}

// Save creates the DatabaseDetectorResult entities in the database.
func (ddrcb *DatabaseDetectorResultCreateBulk) Save(ctx context.Context) ([]*DatabaseDetectorResult, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ddrcb.builders))
	nodes := make([]*DatabaseDetectorResult, len(ddrcb.builders))
	mutators := make([]Mutator, len(ddrcb.builders))
	for i := range ddrcb.builders {
		func(i int, root context.Context) {
			builder := ddrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DatabaseDetectorResultMutation)
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
					_, err = mutators[i+1].Mutate(root, ddrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ddrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ddrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ddrcb *DatabaseDetectorResultCreateBulk) SaveX(ctx context.Context) []*DatabaseDetectorResult {
	v, err := ddrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ddrcb *DatabaseDetectorResultCreateBulk) Exec(ctx context.Context) error {
	_, err := ddrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddrcb *DatabaseDetectorResultCreateBulk) ExecX(ctx context.Context) {
	if err := ddrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
