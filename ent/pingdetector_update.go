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
	"github.com/vicanso/cybertect/ent/pingdetector"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/schema"
)

// PingDetectorUpdate is the builder for updating PingDetector entities.
type PingDetectorUpdate struct {
	config
	hooks     []Hook
	mutation  *PingDetectorMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PingDetectorUpdate builder.
func (pdu *PingDetectorUpdate) Where(ps ...predicate.PingDetector) *PingDetectorUpdate {
	pdu.mutation.Where(ps...)
	return pdu
}

// SetStatus sets the "status" field.
func (pdu *PingDetectorUpdate) SetStatus(s schema.Status) *PingDetectorUpdate {
	pdu.mutation.ResetStatus()
	pdu.mutation.SetStatus(s)
	return pdu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pdu *PingDetectorUpdate) SetNillableStatus(s *schema.Status) *PingDetectorUpdate {
	if s != nil {
		pdu.SetStatus(*s)
	}
	return pdu
}

// AddStatus adds s to the "status" field.
func (pdu *PingDetectorUpdate) AddStatus(s schema.Status) *PingDetectorUpdate {
	pdu.mutation.AddStatus(s)
	return pdu
}

// SetName sets the "name" field.
func (pdu *PingDetectorUpdate) SetName(s string) *PingDetectorUpdate {
	pdu.mutation.SetName(s)
	return pdu
}

// SetOwners sets the "owners" field.
func (pdu *PingDetectorUpdate) SetOwners(s []string) *PingDetectorUpdate {
	pdu.mutation.SetOwners(s)
	return pdu
}

// AppendOwners appends s to the "owners" field.
func (pdu *PingDetectorUpdate) AppendOwners(s []string) *PingDetectorUpdate {
	pdu.mutation.AppendOwners(s)
	return pdu
}

// SetReceivers sets the "receivers" field.
func (pdu *PingDetectorUpdate) SetReceivers(s []string) *PingDetectorUpdate {
	pdu.mutation.SetReceivers(s)
	return pdu
}

// AppendReceivers appends s to the "receivers" field.
func (pdu *PingDetectorUpdate) AppendReceivers(s []string) *PingDetectorUpdate {
	pdu.mutation.AppendReceivers(s)
	return pdu
}

// SetTimeout sets the "timeout" field.
func (pdu *PingDetectorUpdate) SetTimeout(s string) *PingDetectorUpdate {
	pdu.mutation.SetTimeout(s)
	return pdu
}

// SetInterval sets the "interval" field.
func (pdu *PingDetectorUpdate) SetInterval(s string) *PingDetectorUpdate {
	pdu.mutation.SetInterval(s)
	return pdu
}

// SetNillableInterval sets the "interval" field if the given value is not nil.
func (pdu *PingDetectorUpdate) SetNillableInterval(s *string) *PingDetectorUpdate {
	if s != nil {
		pdu.SetInterval(*s)
	}
	return pdu
}

// ClearInterval clears the value of the "interval" field.
func (pdu *PingDetectorUpdate) ClearInterval() *PingDetectorUpdate {
	pdu.mutation.ClearInterval()
	return pdu
}

// SetDescription sets the "description" field.
func (pdu *PingDetectorUpdate) SetDescription(s string) *PingDetectorUpdate {
	pdu.mutation.SetDescription(s)
	return pdu
}

// SetIps sets the "ips" field.
func (pdu *PingDetectorUpdate) SetIps(s []string) *PingDetectorUpdate {
	pdu.mutation.SetIps(s)
	return pdu
}

// AppendIps appends s to the "ips" field.
func (pdu *PingDetectorUpdate) AppendIps(s []string) *PingDetectorUpdate {
	pdu.mutation.AppendIps(s)
	return pdu
}

// Mutation returns the PingDetectorMutation object of the builder.
func (pdu *PingDetectorUpdate) Mutation() *PingDetectorMutation {
	return pdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pdu *PingDetectorUpdate) Save(ctx context.Context) (int, error) {
	pdu.defaults()
	return withHooks(ctx, pdu.sqlSave, pdu.mutation, pdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pdu *PingDetectorUpdate) SaveX(ctx context.Context) int {
	affected, err := pdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pdu *PingDetectorUpdate) Exec(ctx context.Context) error {
	_, err := pdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdu *PingDetectorUpdate) ExecX(ctx context.Context) {
	if err := pdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pdu *PingDetectorUpdate) defaults() {
	if _, ok := pdu.mutation.UpdatedAt(); !ok {
		v := pingdetector.UpdateDefaultUpdatedAt()
		pdu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pdu *PingDetectorUpdate) check() error {
	if v, ok := pdu.mutation.Status(); ok {
		if err := pingdetector.StatusValidator(int8(v)); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "PingDetector.status": %w`, err)}
		}
	}
	if v, ok := pdu.mutation.Name(); ok {
		if err := pingdetector.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "PingDetector.name": %w`, err)}
		}
	}
	if v, ok := pdu.mutation.Timeout(); ok {
		if err := pingdetector.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf(`ent: validator failed for field "PingDetector.timeout": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pdu *PingDetectorUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PingDetectorUpdate {
	pdu.modifiers = append(pdu.modifiers, modifiers...)
	return pdu
}

func (pdu *PingDetectorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pdu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(pingdetector.Table, pingdetector.Columns, sqlgraph.NewFieldSpec(pingdetector.FieldID, field.TypeInt))
	if ps := pdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pdu.mutation.UpdatedAt(); ok {
		_spec.SetField(pingdetector.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pdu.mutation.Status(); ok {
		_spec.SetField(pingdetector.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := pdu.mutation.AddedStatus(); ok {
		_spec.AddField(pingdetector.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := pdu.mutation.Name(); ok {
		_spec.SetField(pingdetector.FieldName, field.TypeString, value)
	}
	if value, ok := pdu.mutation.Owners(); ok {
		_spec.SetField(pingdetector.FieldOwners, field.TypeJSON, value)
	}
	if value, ok := pdu.mutation.AppendedOwners(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, pingdetector.FieldOwners, value)
		})
	}
	if value, ok := pdu.mutation.Receivers(); ok {
		_spec.SetField(pingdetector.FieldReceivers, field.TypeJSON, value)
	}
	if value, ok := pdu.mutation.AppendedReceivers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, pingdetector.FieldReceivers, value)
		})
	}
	if value, ok := pdu.mutation.Timeout(); ok {
		_spec.SetField(pingdetector.FieldTimeout, field.TypeString, value)
	}
	if value, ok := pdu.mutation.Interval(); ok {
		_spec.SetField(pingdetector.FieldInterval, field.TypeString, value)
	}
	if pdu.mutation.IntervalCleared() {
		_spec.ClearField(pingdetector.FieldInterval, field.TypeString)
	}
	if value, ok := pdu.mutation.Description(); ok {
		_spec.SetField(pingdetector.FieldDescription, field.TypeString, value)
	}
	if value, ok := pdu.mutation.Ips(); ok {
		_spec.SetField(pingdetector.FieldIps, field.TypeJSON, value)
	}
	if value, ok := pdu.mutation.AppendedIps(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, pingdetector.FieldIps, value)
		})
	}
	_spec.AddModifiers(pdu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pingdetector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pdu.mutation.done = true
	return n, nil
}

// PingDetectorUpdateOne is the builder for updating a single PingDetector entity.
type PingDetectorUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PingDetectorMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetStatus sets the "status" field.
func (pduo *PingDetectorUpdateOne) SetStatus(s schema.Status) *PingDetectorUpdateOne {
	pduo.mutation.ResetStatus()
	pduo.mutation.SetStatus(s)
	return pduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pduo *PingDetectorUpdateOne) SetNillableStatus(s *schema.Status) *PingDetectorUpdateOne {
	if s != nil {
		pduo.SetStatus(*s)
	}
	return pduo
}

// AddStatus adds s to the "status" field.
func (pduo *PingDetectorUpdateOne) AddStatus(s schema.Status) *PingDetectorUpdateOne {
	pduo.mutation.AddStatus(s)
	return pduo
}

// SetName sets the "name" field.
func (pduo *PingDetectorUpdateOne) SetName(s string) *PingDetectorUpdateOne {
	pduo.mutation.SetName(s)
	return pduo
}

// SetOwners sets the "owners" field.
func (pduo *PingDetectorUpdateOne) SetOwners(s []string) *PingDetectorUpdateOne {
	pduo.mutation.SetOwners(s)
	return pduo
}

// AppendOwners appends s to the "owners" field.
func (pduo *PingDetectorUpdateOne) AppendOwners(s []string) *PingDetectorUpdateOne {
	pduo.mutation.AppendOwners(s)
	return pduo
}

// SetReceivers sets the "receivers" field.
func (pduo *PingDetectorUpdateOne) SetReceivers(s []string) *PingDetectorUpdateOne {
	pduo.mutation.SetReceivers(s)
	return pduo
}

// AppendReceivers appends s to the "receivers" field.
func (pduo *PingDetectorUpdateOne) AppendReceivers(s []string) *PingDetectorUpdateOne {
	pduo.mutation.AppendReceivers(s)
	return pduo
}

// SetTimeout sets the "timeout" field.
func (pduo *PingDetectorUpdateOne) SetTimeout(s string) *PingDetectorUpdateOne {
	pduo.mutation.SetTimeout(s)
	return pduo
}

// SetInterval sets the "interval" field.
func (pduo *PingDetectorUpdateOne) SetInterval(s string) *PingDetectorUpdateOne {
	pduo.mutation.SetInterval(s)
	return pduo
}

// SetNillableInterval sets the "interval" field if the given value is not nil.
func (pduo *PingDetectorUpdateOne) SetNillableInterval(s *string) *PingDetectorUpdateOne {
	if s != nil {
		pduo.SetInterval(*s)
	}
	return pduo
}

// ClearInterval clears the value of the "interval" field.
func (pduo *PingDetectorUpdateOne) ClearInterval() *PingDetectorUpdateOne {
	pduo.mutation.ClearInterval()
	return pduo
}

// SetDescription sets the "description" field.
func (pduo *PingDetectorUpdateOne) SetDescription(s string) *PingDetectorUpdateOne {
	pduo.mutation.SetDescription(s)
	return pduo
}

// SetIps sets the "ips" field.
func (pduo *PingDetectorUpdateOne) SetIps(s []string) *PingDetectorUpdateOne {
	pduo.mutation.SetIps(s)
	return pduo
}

// AppendIps appends s to the "ips" field.
func (pduo *PingDetectorUpdateOne) AppendIps(s []string) *PingDetectorUpdateOne {
	pduo.mutation.AppendIps(s)
	return pduo
}

// Mutation returns the PingDetectorMutation object of the builder.
func (pduo *PingDetectorUpdateOne) Mutation() *PingDetectorMutation {
	return pduo.mutation
}

// Where appends a list predicates to the PingDetectorUpdate builder.
func (pduo *PingDetectorUpdateOne) Where(ps ...predicate.PingDetector) *PingDetectorUpdateOne {
	pduo.mutation.Where(ps...)
	return pduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pduo *PingDetectorUpdateOne) Select(field string, fields ...string) *PingDetectorUpdateOne {
	pduo.fields = append([]string{field}, fields...)
	return pduo
}

// Save executes the query and returns the updated PingDetector entity.
func (pduo *PingDetectorUpdateOne) Save(ctx context.Context) (*PingDetector, error) {
	pduo.defaults()
	return withHooks(ctx, pduo.sqlSave, pduo.mutation, pduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pduo *PingDetectorUpdateOne) SaveX(ctx context.Context) *PingDetector {
	node, err := pduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pduo *PingDetectorUpdateOne) Exec(ctx context.Context) error {
	_, err := pduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pduo *PingDetectorUpdateOne) ExecX(ctx context.Context) {
	if err := pduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pduo *PingDetectorUpdateOne) defaults() {
	if _, ok := pduo.mutation.UpdatedAt(); !ok {
		v := pingdetector.UpdateDefaultUpdatedAt()
		pduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pduo *PingDetectorUpdateOne) check() error {
	if v, ok := pduo.mutation.Status(); ok {
		if err := pingdetector.StatusValidator(int8(v)); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "PingDetector.status": %w`, err)}
		}
	}
	if v, ok := pduo.mutation.Name(); ok {
		if err := pingdetector.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "PingDetector.name": %w`, err)}
		}
	}
	if v, ok := pduo.mutation.Timeout(); ok {
		if err := pingdetector.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf(`ent: validator failed for field "PingDetector.timeout": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pduo *PingDetectorUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PingDetectorUpdateOne {
	pduo.modifiers = append(pduo.modifiers, modifiers...)
	return pduo
}

func (pduo *PingDetectorUpdateOne) sqlSave(ctx context.Context) (_node *PingDetector, err error) {
	if err := pduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(pingdetector.Table, pingdetector.Columns, sqlgraph.NewFieldSpec(pingdetector.FieldID, field.TypeInt))
	id, ok := pduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PingDetector.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pingdetector.FieldID)
		for _, f := range fields {
			if !pingdetector.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pingdetector.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pduo.mutation.UpdatedAt(); ok {
		_spec.SetField(pingdetector.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pduo.mutation.Status(); ok {
		_spec.SetField(pingdetector.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := pduo.mutation.AddedStatus(); ok {
		_spec.AddField(pingdetector.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := pduo.mutation.Name(); ok {
		_spec.SetField(pingdetector.FieldName, field.TypeString, value)
	}
	if value, ok := pduo.mutation.Owners(); ok {
		_spec.SetField(pingdetector.FieldOwners, field.TypeJSON, value)
	}
	if value, ok := pduo.mutation.AppendedOwners(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, pingdetector.FieldOwners, value)
		})
	}
	if value, ok := pduo.mutation.Receivers(); ok {
		_spec.SetField(pingdetector.FieldReceivers, field.TypeJSON, value)
	}
	if value, ok := pduo.mutation.AppendedReceivers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, pingdetector.FieldReceivers, value)
		})
	}
	if value, ok := pduo.mutation.Timeout(); ok {
		_spec.SetField(pingdetector.FieldTimeout, field.TypeString, value)
	}
	if value, ok := pduo.mutation.Interval(); ok {
		_spec.SetField(pingdetector.FieldInterval, field.TypeString, value)
	}
	if pduo.mutation.IntervalCleared() {
		_spec.ClearField(pingdetector.FieldInterval, field.TypeString)
	}
	if value, ok := pduo.mutation.Description(); ok {
		_spec.SetField(pingdetector.FieldDescription, field.TypeString, value)
	}
	if value, ok := pduo.mutation.Ips(); ok {
		_spec.SetField(pingdetector.FieldIps, field.TypeJSON, value)
	}
	if value, ok := pduo.mutation.AppendedIps(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, pingdetector.FieldIps, value)
		})
	}
	_spec.AddModifiers(pduo.modifiers...)
	_node = &PingDetector{config: pduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pingdetector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pduo.mutation.done = true
	return _node, nil
}
