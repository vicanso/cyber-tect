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
	"github.com/vicanso/cybertect/ent/tcpdetector"
	"github.com/vicanso/cybertect/schema"
)

// TCPDetectorUpdate is the builder for updating TCPDetector entities.
type TCPDetectorUpdate struct {
	config
	hooks     []Hook
	mutation  *TCPDetectorMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TCPDetectorUpdate builder.
func (tdu *TCPDetectorUpdate) Where(ps ...predicate.TCPDetector) *TCPDetectorUpdate {
	tdu.mutation.Where(ps...)
	return tdu
}

// SetStatus sets the "status" field.
func (tdu *TCPDetectorUpdate) SetStatus(s schema.Status) *TCPDetectorUpdate {
	tdu.mutation.ResetStatus()
	tdu.mutation.SetStatus(s)
	return tdu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tdu *TCPDetectorUpdate) SetNillableStatus(s *schema.Status) *TCPDetectorUpdate {
	if s != nil {
		tdu.SetStatus(*s)
	}
	return tdu
}

// AddStatus adds s to the "status" field.
func (tdu *TCPDetectorUpdate) AddStatus(s schema.Status) *TCPDetectorUpdate {
	tdu.mutation.AddStatus(s)
	return tdu
}

// SetName sets the "name" field.
func (tdu *TCPDetectorUpdate) SetName(s string) *TCPDetectorUpdate {
	tdu.mutation.SetName(s)
	return tdu
}

// SetOwners sets the "owners" field.
func (tdu *TCPDetectorUpdate) SetOwners(s []string) *TCPDetectorUpdate {
	tdu.mutation.SetOwners(s)
	return tdu
}

// AppendOwners appends s to the "owners" field.
func (tdu *TCPDetectorUpdate) AppendOwners(s []string) *TCPDetectorUpdate {
	tdu.mutation.AppendOwners(s)
	return tdu
}

// SetReceivers sets the "receivers" field.
func (tdu *TCPDetectorUpdate) SetReceivers(s []string) *TCPDetectorUpdate {
	tdu.mutation.SetReceivers(s)
	return tdu
}

// AppendReceivers appends s to the "receivers" field.
func (tdu *TCPDetectorUpdate) AppendReceivers(s []string) *TCPDetectorUpdate {
	tdu.mutation.AppendReceivers(s)
	return tdu
}

// SetTimeout sets the "timeout" field.
func (tdu *TCPDetectorUpdate) SetTimeout(s string) *TCPDetectorUpdate {
	tdu.mutation.SetTimeout(s)
	return tdu
}

// SetInterval sets the "interval" field.
func (tdu *TCPDetectorUpdate) SetInterval(s string) *TCPDetectorUpdate {
	tdu.mutation.SetInterval(s)
	return tdu
}

// SetNillableInterval sets the "interval" field if the given value is not nil.
func (tdu *TCPDetectorUpdate) SetNillableInterval(s *string) *TCPDetectorUpdate {
	if s != nil {
		tdu.SetInterval(*s)
	}
	return tdu
}

// ClearInterval clears the value of the "interval" field.
func (tdu *TCPDetectorUpdate) ClearInterval() *TCPDetectorUpdate {
	tdu.mutation.ClearInterval()
	return tdu
}

// SetDescription sets the "description" field.
func (tdu *TCPDetectorUpdate) SetDescription(s string) *TCPDetectorUpdate {
	tdu.mutation.SetDescription(s)
	return tdu
}

// SetAddrs sets the "addrs" field.
func (tdu *TCPDetectorUpdate) SetAddrs(s []string) *TCPDetectorUpdate {
	tdu.mutation.SetAddrs(s)
	return tdu
}

// AppendAddrs appends s to the "addrs" field.
func (tdu *TCPDetectorUpdate) AppendAddrs(s []string) *TCPDetectorUpdate {
	tdu.mutation.AppendAddrs(s)
	return tdu
}

// Mutation returns the TCPDetectorMutation object of the builder.
func (tdu *TCPDetectorUpdate) Mutation() *TCPDetectorMutation {
	return tdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tdu *TCPDetectorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tdu.defaults()
	if len(tdu.hooks) == 0 {
		if err = tdu.check(); err != nil {
			return 0, err
		}
		affected, err = tdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TCPDetectorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tdu.check(); err != nil {
				return 0, err
			}
			tdu.mutation = mutation
			affected, err = tdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tdu.hooks) - 1; i >= 0; i-- {
			if tdu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tdu *TCPDetectorUpdate) SaveX(ctx context.Context) int {
	affected, err := tdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tdu *TCPDetectorUpdate) Exec(ctx context.Context) error {
	_, err := tdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdu *TCPDetectorUpdate) ExecX(ctx context.Context) {
	if err := tdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tdu *TCPDetectorUpdate) defaults() {
	if _, ok := tdu.mutation.UpdatedAt(); !ok {
		v := tcpdetector.UpdateDefaultUpdatedAt()
		tdu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tdu *TCPDetectorUpdate) check() error {
	if v, ok := tdu.mutation.Status(); ok {
		if err := tcpdetector.StatusValidator(int8(v)); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "TCPDetector.status": %w`, err)}
		}
	}
	if v, ok := tdu.mutation.Name(); ok {
		if err := tcpdetector.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TCPDetector.name": %w`, err)}
		}
	}
	if v, ok := tdu.mutation.Timeout(); ok {
		if err := tcpdetector.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf(`ent: validator failed for field "TCPDetector.timeout": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tdu *TCPDetectorUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TCPDetectorUpdate {
	tdu.modifiers = append(tdu.modifiers, modifiers...)
	return tdu
}

func (tdu *TCPDetectorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tcpdetector.Table,
			Columns: tcpdetector.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tcpdetector.FieldID,
			},
		},
	}
	if ps := tdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tdu.mutation.UpdatedAt(); ok {
		_spec.SetField(tcpdetector.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tdu.mutation.Status(); ok {
		_spec.SetField(tcpdetector.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := tdu.mutation.AddedStatus(); ok {
		_spec.AddField(tcpdetector.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := tdu.mutation.Name(); ok {
		_spec.SetField(tcpdetector.FieldName, field.TypeString, value)
	}
	if value, ok := tdu.mutation.Owners(); ok {
		_spec.SetField(tcpdetector.FieldOwners, field.TypeJSON, value)
	}
	if value, ok := tdu.mutation.AppendedOwners(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, tcpdetector.FieldOwners, value)
		})
	}
	if value, ok := tdu.mutation.Receivers(); ok {
		_spec.SetField(tcpdetector.FieldReceivers, field.TypeJSON, value)
	}
	if value, ok := tdu.mutation.AppendedReceivers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, tcpdetector.FieldReceivers, value)
		})
	}
	if value, ok := tdu.mutation.Timeout(); ok {
		_spec.SetField(tcpdetector.FieldTimeout, field.TypeString, value)
	}
	if value, ok := tdu.mutation.Interval(); ok {
		_spec.SetField(tcpdetector.FieldInterval, field.TypeString, value)
	}
	if tdu.mutation.IntervalCleared() {
		_spec.ClearField(tcpdetector.FieldInterval, field.TypeString)
	}
	if value, ok := tdu.mutation.Description(); ok {
		_spec.SetField(tcpdetector.FieldDescription, field.TypeString, value)
	}
	if value, ok := tdu.mutation.Addrs(); ok {
		_spec.SetField(tcpdetector.FieldAddrs, field.TypeJSON, value)
	}
	if value, ok := tdu.mutation.AppendedAddrs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, tcpdetector.FieldAddrs, value)
		})
	}
	_spec.AddModifiers(tdu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tcpdetector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TCPDetectorUpdateOne is the builder for updating a single TCPDetector entity.
type TCPDetectorUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TCPDetectorMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetStatus sets the "status" field.
func (tduo *TCPDetectorUpdateOne) SetStatus(s schema.Status) *TCPDetectorUpdateOne {
	tduo.mutation.ResetStatus()
	tduo.mutation.SetStatus(s)
	return tduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tduo *TCPDetectorUpdateOne) SetNillableStatus(s *schema.Status) *TCPDetectorUpdateOne {
	if s != nil {
		tduo.SetStatus(*s)
	}
	return tduo
}

// AddStatus adds s to the "status" field.
func (tduo *TCPDetectorUpdateOne) AddStatus(s schema.Status) *TCPDetectorUpdateOne {
	tduo.mutation.AddStatus(s)
	return tduo
}

// SetName sets the "name" field.
func (tduo *TCPDetectorUpdateOne) SetName(s string) *TCPDetectorUpdateOne {
	tduo.mutation.SetName(s)
	return tduo
}

// SetOwners sets the "owners" field.
func (tduo *TCPDetectorUpdateOne) SetOwners(s []string) *TCPDetectorUpdateOne {
	tduo.mutation.SetOwners(s)
	return tduo
}

// AppendOwners appends s to the "owners" field.
func (tduo *TCPDetectorUpdateOne) AppendOwners(s []string) *TCPDetectorUpdateOne {
	tduo.mutation.AppendOwners(s)
	return tduo
}

// SetReceivers sets the "receivers" field.
func (tduo *TCPDetectorUpdateOne) SetReceivers(s []string) *TCPDetectorUpdateOne {
	tduo.mutation.SetReceivers(s)
	return tduo
}

// AppendReceivers appends s to the "receivers" field.
func (tduo *TCPDetectorUpdateOne) AppendReceivers(s []string) *TCPDetectorUpdateOne {
	tduo.mutation.AppendReceivers(s)
	return tduo
}

// SetTimeout sets the "timeout" field.
func (tduo *TCPDetectorUpdateOne) SetTimeout(s string) *TCPDetectorUpdateOne {
	tduo.mutation.SetTimeout(s)
	return tduo
}

// SetInterval sets the "interval" field.
func (tduo *TCPDetectorUpdateOne) SetInterval(s string) *TCPDetectorUpdateOne {
	tduo.mutation.SetInterval(s)
	return tduo
}

// SetNillableInterval sets the "interval" field if the given value is not nil.
func (tduo *TCPDetectorUpdateOne) SetNillableInterval(s *string) *TCPDetectorUpdateOne {
	if s != nil {
		tduo.SetInterval(*s)
	}
	return tduo
}

// ClearInterval clears the value of the "interval" field.
func (tduo *TCPDetectorUpdateOne) ClearInterval() *TCPDetectorUpdateOne {
	tduo.mutation.ClearInterval()
	return tduo
}

// SetDescription sets the "description" field.
func (tduo *TCPDetectorUpdateOne) SetDescription(s string) *TCPDetectorUpdateOne {
	tduo.mutation.SetDescription(s)
	return tduo
}

// SetAddrs sets the "addrs" field.
func (tduo *TCPDetectorUpdateOne) SetAddrs(s []string) *TCPDetectorUpdateOne {
	tduo.mutation.SetAddrs(s)
	return tduo
}

// AppendAddrs appends s to the "addrs" field.
func (tduo *TCPDetectorUpdateOne) AppendAddrs(s []string) *TCPDetectorUpdateOne {
	tduo.mutation.AppendAddrs(s)
	return tduo
}

// Mutation returns the TCPDetectorMutation object of the builder.
func (tduo *TCPDetectorUpdateOne) Mutation() *TCPDetectorMutation {
	return tduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tduo *TCPDetectorUpdateOne) Select(field string, fields ...string) *TCPDetectorUpdateOne {
	tduo.fields = append([]string{field}, fields...)
	return tduo
}

// Save executes the query and returns the updated TCPDetector entity.
func (tduo *TCPDetectorUpdateOne) Save(ctx context.Context) (*TCPDetector, error) {
	var (
		err  error
		node *TCPDetector
	)
	tduo.defaults()
	if len(tduo.hooks) == 0 {
		if err = tduo.check(); err != nil {
			return nil, err
		}
		node, err = tduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TCPDetectorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tduo.check(); err != nil {
				return nil, err
			}
			tduo.mutation = mutation
			node, err = tduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tduo.hooks) - 1; i >= 0; i-- {
			if tduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TCPDetector)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TCPDetectorMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tduo *TCPDetectorUpdateOne) SaveX(ctx context.Context) *TCPDetector {
	node, err := tduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tduo *TCPDetectorUpdateOne) Exec(ctx context.Context) error {
	_, err := tduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tduo *TCPDetectorUpdateOne) ExecX(ctx context.Context) {
	if err := tduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tduo *TCPDetectorUpdateOne) defaults() {
	if _, ok := tduo.mutation.UpdatedAt(); !ok {
		v := tcpdetector.UpdateDefaultUpdatedAt()
		tduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tduo *TCPDetectorUpdateOne) check() error {
	if v, ok := tduo.mutation.Status(); ok {
		if err := tcpdetector.StatusValidator(int8(v)); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "TCPDetector.status": %w`, err)}
		}
	}
	if v, ok := tduo.mutation.Name(); ok {
		if err := tcpdetector.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TCPDetector.name": %w`, err)}
		}
	}
	if v, ok := tduo.mutation.Timeout(); ok {
		if err := tcpdetector.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf(`ent: validator failed for field "TCPDetector.timeout": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tduo *TCPDetectorUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TCPDetectorUpdateOne {
	tduo.modifiers = append(tduo.modifiers, modifiers...)
	return tduo
}

func (tduo *TCPDetectorUpdateOne) sqlSave(ctx context.Context) (_node *TCPDetector, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tcpdetector.Table,
			Columns: tcpdetector.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tcpdetector.FieldID,
			},
		},
	}
	id, ok := tduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TCPDetector.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tcpdetector.FieldID)
		for _, f := range fields {
			if !tcpdetector.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tcpdetector.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tduo.mutation.UpdatedAt(); ok {
		_spec.SetField(tcpdetector.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tduo.mutation.Status(); ok {
		_spec.SetField(tcpdetector.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := tduo.mutation.AddedStatus(); ok {
		_spec.AddField(tcpdetector.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := tduo.mutation.Name(); ok {
		_spec.SetField(tcpdetector.FieldName, field.TypeString, value)
	}
	if value, ok := tduo.mutation.Owners(); ok {
		_spec.SetField(tcpdetector.FieldOwners, field.TypeJSON, value)
	}
	if value, ok := tduo.mutation.AppendedOwners(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, tcpdetector.FieldOwners, value)
		})
	}
	if value, ok := tduo.mutation.Receivers(); ok {
		_spec.SetField(tcpdetector.FieldReceivers, field.TypeJSON, value)
	}
	if value, ok := tduo.mutation.AppendedReceivers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, tcpdetector.FieldReceivers, value)
		})
	}
	if value, ok := tduo.mutation.Timeout(); ok {
		_spec.SetField(tcpdetector.FieldTimeout, field.TypeString, value)
	}
	if value, ok := tduo.mutation.Interval(); ok {
		_spec.SetField(tcpdetector.FieldInterval, field.TypeString, value)
	}
	if tduo.mutation.IntervalCleared() {
		_spec.ClearField(tcpdetector.FieldInterval, field.TypeString)
	}
	if value, ok := tduo.mutation.Description(); ok {
		_spec.SetField(tcpdetector.FieldDescription, field.TypeString, value)
	}
	if value, ok := tduo.mutation.Addrs(); ok {
		_spec.SetField(tcpdetector.FieldAddrs, field.TypeJSON, value)
	}
	if value, ok := tduo.mutation.AppendedAddrs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, tcpdetector.FieldAddrs, value)
		})
	}
	_spec.AddModifiers(tduo.modifiers...)
	_node = &TCPDetector{config: tduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tcpdetector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
