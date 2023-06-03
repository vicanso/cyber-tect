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
	"github.com/vicanso/cybertect/ent/httpdetector"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/schema"
)

// HTTPDetectorUpdate is the builder for updating HTTPDetector entities.
type HTTPDetectorUpdate struct {
	config
	hooks     []Hook
	mutation  *HTTPDetectorMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the HTTPDetectorUpdate builder.
func (hdu *HTTPDetectorUpdate) Where(ps ...predicate.HTTPDetector) *HTTPDetectorUpdate {
	hdu.mutation.Where(ps...)
	return hdu
}

// SetStatus sets the "status" field.
func (hdu *HTTPDetectorUpdate) SetStatus(s schema.Status) *HTTPDetectorUpdate {
	hdu.mutation.ResetStatus()
	hdu.mutation.SetStatus(s)
	return hdu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hdu *HTTPDetectorUpdate) SetNillableStatus(s *schema.Status) *HTTPDetectorUpdate {
	if s != nil {
		hdu.SetStatus(*s)
	}
	return hdu
}

// AddStatus adds s to the "status" field.
func (hdu *HTTPDetectorUpdate) AddStatus(s schema.Status) *HTTPDetectorUpdate {
	hdu.mutation.AddStatus(s)
	return hdu
}

// SetName sets the "name" field.
func (hdu *HTTPDetectorUpdate) SetName(s string) *HTTPDetectorUpdate {
	hdu.mutation.SetName(s)
	return hdu
}

// SetOwners sets the "owners" field.
func (hdu *HTTPDetectorUpdate) SetOwners(s []string) *HTTPDetectorUpdate {
	hdu.mutation.SetOwners(s)
	return hdu
}

// AppendOwners appends s to the "owners" field.
func (hdu *HTTPDetectorUpdate) AppendOwners(s []string) *HTTPDetectorUpdate {
	hdu.mutation.AppendOwners(s)
	return hdu
}

// SetReceivers sets the "receivers" field.
func (hdu *HTTPDetectorUpdate) SetReceivers(s []string) *HTTPDetectorUpdate {
	hdu.mutation.SetReceivers(s)
	return hdu
}

// AppendReceivers appends s to the "receivers" field.
func (hdu *HTTPDetectorUpdate) AppendReceivers(s []string) *HTTPDetectorUpdate {
	hdu.mutation.AppendReceivers(s)
	return hdu
}

// SetTimeout sets the "timeout" field.
func (hdu *HTTPDetectorUpdate) SetTimeout(s string) *HTTPDetectorUpdate {
	hdu.mutation.SetTimeout(s)
	return hdu
}

// SetInterval sets the "interval" field.
func (hdu *HTTPDetectorUpdate) SetInterval(s string) *HTTPDetectorUpdate {
	hdu.mutation.SetInterval(s)
	return hdu
}

// SetNillableInterval sets the "interval" field if the given value is not nil.
func (hdu *HTTPDetectorUpdate) SetNillableInterval(s *string) *HTTPDetectorUpdate {
	if s != nil {
		hdu.SetInterval(*s)
	}
	return hdu
}

// ClearInterval clears the value of the "interval" field.
func (hdu *HTTPDetectorUpdate) ClearInterval() *HTTPDetectorUpdate {
	hdu.mutation.ClearInterval()
	return hdu
}

// SetDescription sets the "description" field.
func (hdu *HTTPDetectorUpdate) SetDescription(s string) *HTTPDetectorUpdate {
	hdu.mutation.SetDescription(s)
	return hdu
}

// SetIps sets the "ips" field.
func (hdu *HTTPDetectorUpdate) SetIps(s []string) *HTTPDetectorUpdate {
	hdu.mutation.SetIps(s)
	return hdu
}

// AppendIps appends s to the "ips" field.
func (hdu *HTTPDetectorUpdate) AppendIps(s []string) *HTTPDetectorUpdate {
	hdu.mutation.AppendIps(s)
	return hdu
}

// SetURL sets the "url" field.
func (hdu *HTTPDetectorUpdate) SetURL(s string) *HTTPDetectorUpdate {
	hdu.mutation.SetURL(s)
	return hdu
}

// SetScript sets the "script" field.
func (hdu *HTTPDetectorUpdate) SetScript(s string) *HTTPDetectorUpdate {
	hdu.mutation.SetScript(s)
	return hdu
}

// SetNillableScript sets the "script" field if the given value is not nil.
func (hdu *HTTPDetectorUpdate) SetNillableScript(s *string) *HTTPDetectorUpdate {
	if s != nil {
		hdu.SetScript(*s)
	}
	return hdu
}

// ClearScript clears the value of the "script" field.
func (hdu *HTTPDetectorUpdate) ClearScript() *HTTPDetectorUpdate {
	hdu.mutation.ClearScript()
	return hdu
}

// SetProxies sets the "proxies" field.
func (hdu *HTTPDetectorUpdate) SetProxies(s []string) *HTTPDetectorUpdate {
	hdu.mutation.SetProxies(s)
	return hdu
}

// AppendProxies appends s to the "proxies" field.
func (hdu *HTTPDetectorUpdate) AppendProxies(s []string) *HTTPDetectorUpdate {
	hdu.mutation.AppendProxies(s)
	return hdu
}

// ClearProxies clears the value of the "proxies" field.
func (hdu *HTTPDetectorUpdate) ClearProxies() *HTTPDetectorUpdate {
	hdu.mutation.ClearProxies()
	return hdu
}

// SetRandomQueryString sets the "randomQueryString" field.
func (hdu *HTTPDetectorUpdate) SetRandomQueryString(i int8) *HTTPDetectorUpdate {
	hdu.mutation.ResetRandomQueryString()
	hdu.mutation.SetRandomQueryString(i)
	return hdu
}

// SetNillableRandomQueryString sets the "randomQueryString" field if the given value is not nil.
func (hdu *HTTPDetectorUpdate) SetNillableRandomQueryString(i *int8) *HTTPDetectorUpdate {
	if i != nil {
		hdu.SetRandomQueryString(*i)
	}
	return hdu
}

// AddRandomQueryString adds i to the "randomQueryString" field.
func (hdu *HTTPDetectorUpdate) AddRandomQueryString(i int8) *HTTPDetectorUpdate {
	hdu.mutation.AddRandomQueryString(i)
	return hdu
}

// ClearRandomQueryString clears the value of the "randomQueryString" field.
func (hdu *HTTPDetectorUpdate) ClearRandomQueryString() *HTTPDetectorUpdate {
	hdu.mutation.ClearRandomQueryString()
	return hdu
}

// Mutation returns the HTTPDetectorMutation object of the builder.
func (hdu *HTTPDetectorUpdate) Mutation() *HTTPDetectorMutation {
	return hdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hdu *HTTPDetectorUpdate) Save(ctx context.Context) (int, error) {
	hdu.defaults()
	return withHooks(ctx, hdu.sqlSave, hdu.mutation, hdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hdu *HTTPDetectorUpdate) SaveX(ctx context.Context) int {
	affected, err := hdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hdu *HTTPDetectorUpdate) Exec(ctx context.Context) error {
	_, err := hdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hdu *HTTPDetectorUpdate) ExecX(ctx context.Context) {
	if err := hdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hdu *HTTPDetectorUpdate) defaults() {
	if _, ok := hdu.mutation.UpdatedAt(); !ok {
		v := httpdetector.UpdateDefaultUpdatedAt()
		hdu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hdu *HTTPDetectorUpdate) check() error {
	if v, ok := hdu.mutation.Status(); ok {
		if err := httpdetector.StatusValidator(int8(v)); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "HTTPDetector.status": %w`, err)}
		}
	}
	if v, ok := hdu.mutation.Name(); ok {
		if err := httpdetector.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "HTTPDetector.name": %w`, err)}
		}
	}
	if v, ok := hdu.mutation.Timeout(); ok {
		if err := httpdetector.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf(`ent: validator failed for field "HTTPDetector.timeout": %w`, err)}
		}
	}
	if v, ok := hdu.mutation.URL(); ok {
		if err := httpdetector.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "HTTPDetector.url": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (hdu *HTTPDetectorUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *HTTPDetectorUpdate {
	hdu.modifiers = append(hdu.modifiers, modifiers...)
	return hdu
}

func (hdu *HTTPDetectorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hdu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(httpdetector.Table, httpdetector.Columns, sqlgraph.NewFieldSpec(httpdetector.FieldID, field.TypeInt))
	if ps := hdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hdu.mutation.UpdatedAt(); ok {
		_spec.SetField(httpdetector.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := hdu.mutation.Status(); ok {
		_spec.SetField(httpdetector.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := hdu.mutation.AddedStatus(); ok {
		_spec.AddField(httpdetector.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := hdu.mutation.Name(); ok {
		_spec.SetField(httpdetector.FieldName, field.TypeString, value)
	}
	if value, ok := hdu.mutation.Owners(); ok {
		_spec.SetField(httpdetector.FieldOwners, field.TypeJSON, value)
	}
	if value, ok := hdu.mutation.AppendedOwners(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, httpdetector.FieldOwners, value)
		})
	}
	if value, ok := hdu.mutation.Receivers(); ok {
		_spec.SetField(httpdetector.FieldReceivers, field.TypeJSON, value)
	}
	if value, ok := hdu.mutation.AppendedReceivers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, httpdetector.FieldReceivers, value)
		})
	}
	if value, ok := hdu.mutation.Timeout(); ok {
		_spec.SetField(httpdetector.FieldTimeout, field.TypeString, value)
	}
	if value, ok := hdu.mutation.Interval(); ok {
		_spec.SetField(httpdetector.FieldInterval, field.TypeString, value)
	}
	if hdu.mutation.IntervalCleared() {
		_spec.ClearField(httpdetector.FieldInterval, field.TypeString)
	}
	if value, ok := hdu.mutation.Description(); ok {
		_spec.SetField(httpdetector.FieldDescription, field.TypeString, value)
	}
	if value, ok := hdu.mutation.Ips(); ok {
		_spec.SetField(httpdetector.FieldIps, field.TypeJSON, value)
	}
	if value, ok := hdu.mutation.AppendedIps(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, httpdetector.FieldIps, value)
		})
	}
	if value, ok := hdu.mutation.URL(); ok {
		_spec.SetField(httpdetector.FieldURL, field.TypeString, value)
	}
	if value, ok := hdu.mutation.Script(); ok {
		_spec.SetField(httpdetector.FieldScript, field.TypeString, value)
	}
	if hdu.mutation.ScriptCleared() {
		_spec.ClearField(httpdetector.FieldScript, field.TypeString)
	}
	if value, ok := hdu.mutation.Proxies(); ok {
		_spec.SetField(httpdetector.FieldProxies, field.TypeJSON, value)
	}
	if value, ok := hdu.mutation.AppendedProxies(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, httpdetector.FieldProxies, value)
		})
	}
	if hdu.mutation.ProxiesCleared() {
		_spec.ClearField(httpdetector.FieldProxies, field.TypeJSON)
	}
	if value, ok := hdu.mutation.RandomQueryString(); ok {
		_spec.SetField(httpdetector.FieldRandomQueryString, field.TypeInt8, value)
	}
	if value, ok := hdu.mutation.AddedRandomQueryString(); ok {
		_spec.AddField(httpdetector.FieldRandomQueryString, field.TypeInt8, value)
	}
	if hdu.mutation.RandomQueryStringCleared() {
		_spec.ClearField(httpdetector.FieldRandomQueryString, field.TypeInt8)
	}
	_spec.AddModifiers(hdu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, hdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{httpdetector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hdu.mutation.done = true
	return n, nil
}

// HTTPDetectorUpdateOne is the builder for updating a single HTTPDetector entity.
type HTTPDetectorUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *HTTPDetectorMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetStatus sets the "status" field.
func (hduo *HTTPDetectorUpdateOne) SetStatus(s schema.Status) *HTTPDetectorUpdateOne {
	hduo.mutation.ResetStatus()
	hduo.mutation.SetStatus(s)
	return hduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hduo *HTTPDetectorUpdateOne) SetNillableStatus(s *schema.Status) *HTTPDetectorUpdateOne {
	if s != nil {
		hduo.SetStatus(*s)
	}
	return hduo
}

// AddStatus adds s to the "status" field.
func (hduo *HTTPDetectorUpdateOne) AddStatus(s schema.Status) *HTTPDetectorUpdateOne {
	hduo.mutation.AddStatus(s)
	return hduo
}

// SetName sets the "name" field.
func (hduo *HTTPDetectorUpdateOne) SetName(s string) *HTTPDetectorUpdateOne {
	hduo.mutation.SetName(s)
	return hduo
}

// SetOwners sets the "owners" field.
func (hduo *HTTPDetectorUpdateOne) SetOwners(s []string) *HTTPDetectorUpdateOne {
	hduo.mutation.SetOwners(s)
	return hduo
}

// AppendOwners appends s to the "owners" field.
func (hduo *HTTPDetectorUpdateOne) AppendOwners(s []string) *HTTPDetectorUpdateOne {
	hduo.mutation.AppendOwners(s)
	return hduo
}

// SetReceivers sets the "receivers" field.
func (hduo *HTTPDetectorUpdateOne) SetReceivers(s []string) *HTTPDetectorUpdateOne {
	hduo.mutation.SetReceivers(s)
	return hduo
}

// AppendReceivers appends s to the "receivers" field.
func (hduo *HTTPDetectorUpdateOne) AppendReceivers(s []string) *HTTPDetectorUpdateOne {
	hduo.mutation.AppendReceivers(s)
	return hduo
}

// SetTimeout sets the "timeout" field.
func (hduo *HTTPDetectorUpdateOne) SetTimeout(s string) *HTTPDetectorUpdateOne {
	hduo.mutation.SetTimeout(s)
	return hduo
}

// SetInterval sets the "interval" field.
func (hduo *HTTPDetectorUpdateOne) SetInterval(s string) *HTTPDetectorUpdateOne {
	hduo.mutation.SetInterval(s)
	return hduo
}

// SetNillableInterval sets the "interval" field if the given value is not nil.
func (hduo *HTTPDetectorUpdateOne) SetNillableInterval(s *string) *HTTPDetectorUpdateOne {
	if s != nil {
		hduo.SetInterval(*s)
	}
	return hduo
}

// ClearInterval clears the value of the "interval" field.
func (hduo *HTTPDetectorUpdateOne) ClearInterval() *HTTPDetectorUpdateOne {
	hduo.mutation.ClearInterval()
	return hduo
}

// SetDescription sets the "description" field.
func (hduo *HTTPDetectorUpdateOne) SetDescription(s string) *HTTPDetectorUpdateOne {
	hduo.mutation.SetDescription(s)
	return hduo
}

// SetIps sets the "ips" field.
func (hduo *HTTPDetectorUpdateOne) SetIps(s []string) *HTTPDetectorUpdateOne {
	hduo.mutation.SetIps(s)
	return hduo
}

// AppendIps appends s to the "ips" field.
func (hduo *HTTPDetectorUpdateOne) AppendIps(s []string) *HTTPDetectorUpdateOne {
	hduo.mutation.AppendIps(s)
	return hduo
}

// SetURL sets the "url" field.
func (hduo *HTTPDetectorUpdateOne) SetURL(s string) *HTTPDetectorUpdateOne {
	hduo.mutation.SetURL(s)
	return hduo
}

// SetScript sets the "script" field.
func (hduo *HTTPDetectorUpdateOne) SetScript(s string) *HTTPDetectorUpdateOne {
	hduo.mutation.SetScript(s)
	return hduo
}

// SetNillableScript sets the "script" field if the given value is not nil.
func (hduo *HTTPDetectorUpdateOne) SetNillableScript(s *string) *HTTPDetectorUpdateOne {
	if s != nil {
		hduo.SetScript(*s)
	}
	return hduo
}

// ClearScript clears the value of the "script" field.
func (hduo *HTTPDetectorUpdateOne) ClearScript() *HTTPDetectorUpdateOne {
	hduo.mutation.ClearScript()
	return hduo
}

// SetProxies sets the "proxies" field.
func (hduo *HTTPDetectorUpdateOne) SetProxies(s []string) *HTTPDetectorUpdateOne {
	hduo.mutation.SetProxies(s)
	return hduo
}

// AppendProxies appends s to the "proxies" field.
func (hduo *HTTPDetectorUpdateOne) AppendProxies(s []string) *HTTPDetectorUpdateOne {
	hduo.mutation.AppendProxies(s)
	return hduo
}

// ClearProxies clears the value of the "proxies" field.
func (hduo *HTTPDetectorUpdateOne) ClearProxies() *HTTPDetectorUpdateOne {
	hduo.mutation.ClearProxies()
	return hduo
}

// SetRandomQueryString sets the "randomQueryString" field.
func (hduo *HTTPDetectorUpdateOne) SetRandomQueryString(i int8) *HTTPDetectorUpdateOne {
	hduo.mutation.ResetRandomQueryString()
	hduo.mutation.SetRandomQueryString(i)
	return hduo
}

// SetNillableRandomQueryString sets the "randomQueryString" field if the given value is not nil.
func (hduo *HTTPDetectorUpdateOne) SetNillableRandomQueryString(i *int8) *HTTPDetectorUpdateOne {
	if i != nil {
		hduo.SetRandomQueryString(*i)
	}
	return hduo
}

// AddRandomQueryString adds i to the "randomQueryString" field.
func (hduo *HTTPDetectorUpdateOne) AddRandomQueryString(i int8) *HTTPDetectorUpdateOne {
	hduo.mutation.AddRandomQueryString(i)
	return hduo
}

// ClearRandomQueryString clears the value of the "randomQueryString" field.
func (hduo *HTTPDetectorUpdateOne) ClearRandomQueryString() *HTTPDetectorUpdateOne {
	hduo.mutation.ClearRandomQueryString()
	return hduo
}

// Mutation returns the HTTPDetectorMutation object of the builder.
func (hduo *HTTPDetectorUpdateOne) Mutation() *HTTPDetectorMutation {
	return hduo.mutation
}

// Where appends a list predicates to the HTTPDetectorUpdate builder.
func (hduo *HTTPDetectorUpdateOne) Where(ps ...predicate.HTTPDetector) *HTTPDetectorUpdateOne {
	hduo.mutation.Where(ps...)
	return hduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (hduo *HTTPDetectorUpdateOne) Select(field string, fields ...string) *HTTPDetectorUpdateOne {
	hduo.fields = append([]string{field}, fields...)
	return hduo
}

// Save executes the query and returns the updated HTTPDetector entity.
func (hduo *HTTPDetectorUpdateOne) Save(ctx context.Context) (*HTTPDetector, error) {
	hduo.defaults()
	return withHooks(ctx, hduo.sqlSave, hduo.mutation, hduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hduo *HTTPDetectorUpdateOne) SaveX(ctx context.Context) *HTTPDetector {
	node, err := hduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hduo *HTTPDetectorUpdateOne) Exec(ctx context.Context) error {
	_, err := hduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hduo *HTTPDetectorUpdateOne) ExecX(ctx context.Context) {
	if err := hduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hduo *HTTPDetectorUpdateOne) defaults() {
	if _, ok := hduo.mutation.UpdatedAt(); !ok {
		v := httpdetector.UpdateDefaultUpdatedAt()
		hduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hduo *HTTPDetectorUpdateOne) check() error {
	if v, ok := hduo.mutation.Status(); ok {
		if err := httpdetector.StatusValidator(int8(v)); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "HTTPDetector.status": %w`, err)}
		}
	}
	if v, ok := hduo.mutation.Name(); ok {
		if err := httpdetector.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "HTTPDetector.name": %w`, err)}
		}
	}
	if v, ok := hduo.mutation.Timeout(); ok {
		if err := httpdetector.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf(`ent: validator failed for field "HTTPDetector.timeout": %w`, err)}
		}
	}
	if v, ok := hduo.mutation.URL(); ok {
		if err := httpdetector.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "HTTPDetector.url": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (hduo *HTTPDetectorUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *HTTPDetectorUpdateOne {
	hduo.modifiers = append(hduo.modifiers, modifiers...)
	return hduo
}

func (hduo *HTTPDetectorUpdateOne) sqlSave(ctx context.Context) (_node *HTTPDetector, err error) {
	if err := hduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(httpdetector.Table, httpdetector.Columns, sqlgraph.NewFieldSpec(httpdetector.FieldID, field.TypeInt))
	id, ok := hduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HTTPDetector.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := hduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, httpdetector.FieldID)
		for _, f := range fields {
			if !httpdetector.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != httpdetector.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := hduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hduo.mutation.UpdatedAt(); ok {
		_spec.SetField(httpdetector.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := hduo.mutation.Status(); ok {
		_spec.SetField(httpdetector.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := hduo.mutation.AddedStatus(); ok {
		_spec.AddField(httpdetector.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := hduo.mutation.Name(); ok {
		_spec.SetField(httpdetector.FieldName, field.TypeString, value)
	}
	if value, ok := hduo.mutation.Owners(); ok {
		_spec.SetField(httpdetector.FieldOwners, field.TypeJSON, value)
	}
	if value, ok := hduo.mutation.AppendedOwners(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, httpdetector.FieldOwners, value)
		})
	}
	if value, ok := hduo.mutation.Receivers(); ok {
		_spec.SetField(httpdetector.FieldReceivers, field.TypeJSON, value)
	}
	if value, ok := hduo.mutation.AppendedReceivers(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, httpdetector.FieldReceivers, value)
		})
	}
	if value, ok := hduo.mutation.Timeout(); ok {
		_spec.SetField(httpdetector.FieldTimeout, field.TypeString, value)
	}
	if value, ok := hduo.mutation.Interval(); ok {
		_spec.SetField(httpdetector.FieldInterval, field.TypeString, value)
	}
	if hduo.mutation.IntervalCleared() {
		_spec.ClearField(httpdetector.FieldInterval, field.TypeString)
	}
	if value, ok := hduo.mutation.Description(); ok {
		_spec.SetField(httpdetector.FieldDescription, field.TypeString, value)
	}
	if value, ok := hduo.mutation.Ips(); ok {
		_spec.SetField(httpdetector.FieldIps, field.TypeJSON, value)
	}
	if value, ok := hduo.mutation.AppendedIps(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, httpdetector.FieldIps, value)
		})
	}
	if value, ok := hduo.mutation.URL(); ok {
		_spec.SetField(httpdetector.FieldURL, field.TypeString, value)
	}
	if value, ok := hduo.mutation.Script(); ok {
		_spec.SetField(httpdetector.FieldScript, field.TypeString, value)
	}
	if hduo.mutation.ScriptCleared() {
		_spec.ClearField(httpdetector.FieldScript, field.TypeString)
	}
	if value, ok := hduo.mutation.Proxies(); ok {
		_spec.SetField(httpdetector.FieldProxies, field.TypeJSON, value)
	}
	if value, ok := hduo.mutation.AppendedProxies(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, httpdetector.FieldProxies, value)
		})
	}
	if hduo.mutation.ProxiesCleared() {
		_spec.ClearField(httpdetector.FieldProxies, field.TypeJSON)
	}
	if value, ok := hduo.mutation.RandomQueryString(); ok {
		_spec.SetField(httpdetector.FieldRandomQueryString, field.TypeInt8, value)
	}
	if value, ok := hduo.mutation.AddedRandomQueryString(); ok {
		_spec.AddField(httpdetector.FieldRandomQueryString, field.TypeInt8, value)
	}
	if hduo.mutation.RandomQueryStringCleared() {
		_spec.ClearField(httpdetector.FieldRandomQueryString, field.TypeInt8)
	}
	_spec.AddModifiers(hduo.modifiers...)
	_node = &HTTPDetector{config: hduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{httpdetector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	hduo.mutation.done = true
	return _node, nil
}
