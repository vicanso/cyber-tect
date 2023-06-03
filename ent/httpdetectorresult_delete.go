// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vicanso/cybertect/ent/httpdetectorresult"
	"github.com/vicanso/cybertect/ent/predicate"
)

// HTTPDetectorResultDelete is the builder for deleting a HTTPDetectorResult entity.
type HTTPDetectorResultDelete struct {
	config
	hooks    []Hook
	mutation *HTTPDetectorResultMutation
}

// Where appends a list predicates to the HTTPDetectorResultDelete builder.
func (hdrd *HTTPDetectorResultDelete) Where(ps ...predicate.HTTPDetectorResult) *HTTPDetectorResultDelete {
	hdrd.mutation.Where(ps...)
	return hdrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hdrd *HTTPDetectorResultDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, hdrd.sqlExec, hdrd.mutation, hdrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (hdrd *HTTPDetectorResultDelete) ExecX(ctx context.Context) int {
	n, err := hdrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hdrd *HTTPDetectorResultDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(httpdetectorresult.Table, sqlgraph.NewFieldSpec(httpdetectorresult.FieldID, field.TypeInt))
	if ps := hdrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hdrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	hdrd.mutation.done = true
	return affected, err
}

// HTTPDetectorResultDeleteOne is the builder for deleting a single HTTPDetectorResult entity.
type HTTPDetectorResultDeleteOne struct {
	hdrd *HTTPDetectorResultDelete
}

// Where appends a list predicates to the HTTPDetectorResultDelete builder.
func (hdrdo *HTTPDetectorResultDeleteOne) Where(ps ...predicate.HTTPDetectorResult) *HTTPDetectorResultDeleteOne {
	hdrdo.hdrd.mutation.Where(ps...)
	return hdrdo
}

// Exec executes the deletion query.
func (hdrdo *HTTPDetectorResultDeleteOne) Exec(ctx context.Context) error {
	n, err := hdrdo.hdrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{httpdetectorresult.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hdrdo *HTTPDetectorResultDeleteOne) ExecX(ctx context.Context) {
	if err := hdrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
