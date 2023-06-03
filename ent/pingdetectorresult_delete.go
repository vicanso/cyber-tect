// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vicanso/cybertect/ent/pingdetectorresult"
	"github.com/vicanso/cybertect/ent/predicate"
)

// PingDetectorResultDelete is the builder for deleting a PingDetectorResult entity.
type PingDetectorResultDelete struct {
	config
	hooks    []Hook
	mutation *PingDetectorResultMutation
}

// Where appends a list predicates to the PingDetectorResultDelete builder.
func (pdrd *PingDetectorResultDelete) Where(ps ...predicate.PingDetectorResult) *PingDetectorResultDelete {
	pdrd.mutation.Where(ps...)
	return pdrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pdrd *PingDetectorResultDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pdrd.sqlExec, pdrd.mutation, pdrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pdrd *PingDetectorResultDelete) ExecX(ctx context.Context) int {
	n, err := pdrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pdrd *PingDetectorResultDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(pingdetectorresult.Table, sqlgraph.NewFieldSpec(pingdetectorresult.FieldID, field.TypeInt))
	if ps := pdrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pdrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pdrd.mutation.done = true
	return affected, err
}

// PingDetectorResultDeleteOne is the builder for deleting a single PingDetectorResult entity.
type PingDetectorResultDeleteOne struct {
	pdrd *PingDetectorResultDelete
}

// Where appends a list predicates to the PingDetectorResultDelete builder.
func (pdrdo *PingDetectorResultDeleteOne) Where(ps ...predicate.PingDetectorResult) *PingDetectorResultDeleteOne {
	pdrdo.pdrd.mutation.Where(ps...)
	return pdrdo
}

// Exec executes the deletion query.
func (pdrdo *PingDetectorResultDeleteOne) Exec(ctx context.Context) error {
	n, err := pdrdo.pdrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{pingdetectorresult.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdrdo *PingDetectorResultDeleteOne) ExecX(ctx context.Context) {
	if err := pdrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
