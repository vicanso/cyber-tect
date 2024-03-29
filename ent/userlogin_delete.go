// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/ent/userlogin"
)

// UserLoginDelete is the builder for deleting a UserLogin entity.
type UserLoginDelete struct {
	config
	hooks    []Hook
	mutation *UserLoginMutation
}

// Where appends a list predicates to the UserLoginDelete builder.
func (uld *UserLoginDelete) Where(ps ...predicate.UserLogin) *UserLoginDelete {
	uld.mutation.Where(ps...)
	return uld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uld *UserLoginDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, uld.sqlExec, uld.mutation, uld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (uld *UserLoginDelete) ExecX(ctx context.Context) int {
	n, err := uld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uld *UserLoginDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userlogin.Table, sqlgraph.NewFieldSpec(userlogin.FieldID, field.TypeInt))
	if ps := uld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, uld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	uld.mutation.done = true
	return affected, err
}

// UserLoginDeleteOne is the builder for deleting a single UserLogin entity.
type UserLoginDeleteOne struct {
	uld *UserLoginDelete
}

// Where appends a list predicates to the UserLoginDelete builder.
func (uldo *UserLoginDeleteOne) Where(ps ...predicate.UserLogin) *UserLoginDeleteOne {
	uldo.uld.mutation.Where(ps...)
	return uldo
}

// Exec executes the deletion query.
func (uldo *UserLoginDeleteOne) Exec(ctx context.Context) error {
	n, err := uldo.uld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userlogin.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uldo *UserLoginDeleteOne) ExecX(ctx context.Context) {
	if err := uldo.Exec(ctx); err != nil {
		panic(err)
	}
}
