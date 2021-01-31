// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vicanso/cybertect/ent/dnsdetector"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/ent/schema"
)

// DNSDetectorUpdate is the builder for updating DNSDetector entities.
type DNSDetectorUpdate struct {
	config
	hooks    []Hook
	mutation *DNSDetectorMutation
}

// Where adds a new predicate for the DNSDetectorUpdate builder.
func (ddu *DNSDetectorUpdate) Where(ps ...predicate.DNSDetector) *DNSDetectorUpdate {
	ddu.mutation.predicates = append(ddu.mutation.predicates, ps...)
	return ddu
}

// SetStatus sets the "status" field.
func (ddu *DNSDetectorUpdate) SetStatus(s schema.Status) *DNSDetectorUpdate {
	ddu.mutation.ResetStatus()
	ddu.mutation.SetStatus(s)
	return ddu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ddu *DNSDetectorUpdate) SetNillableStatus(s *schema.Status) *DNSDetectorUpdate {
	if s != nil {
		ddu.SetStatus(*s)
	}
	return ddu
}

// AddStatus adds s to the "status" field.
func (ddu *DNSDetectorUpdate) AddStatus(s schema.Status) *DNSDetectorUpdate {
	ddu.mutation.AddStatus(s)
	return ddu
}

// SetName sets the "name" field.
func (ddu *DNSDetectorUpdate) SetName(s string) *DNSDetectorUpdate {
	ddu.mutation.SetName(s)
	return ddu
}

// SetOwner sets the "owner" field.
func (ddu *DNSDetectorUpdate) SetOwner(s string) *DNSDetectorUpdate {
	ddu.mutation.SetOwner(s)
	return ddu
}

// SetDescription sets the "description" field.
func (ddu *DNSDetectorUpdate) SetDescription(s string) *DNSDetectorUpdate {
	ddu.mutation.SetDescription(s)
	return ddu
}

// SetReceivers sets the "receivers" field.
func (ddu *DNSDetectorUpdate) SetReceivers(s []string) *DNSDetectorUpdate {
	ddu.mutation.SetReceivers(s)
	return ddu
}

// SetTimeout sets the "timeout" field.
func (ddu *DNSDetectorUpdate) SetTimeout(s string) *DNSDetectorUpdate {
	ddu.mutation.SetTimeout(s)
	return ddu
}

// SetHost sets the "host" field.
func (ddu *DNSDetectorUpdate) SetHost(s string) *DNSDetectorUpdate {
	ddu.mutation.SetHost(s)
	return ddu
}

// SetIps sets the "ips" field.
func (ddu *DNSDetectorUpdate) SetIps(s []string) *DNSDetectorUpdate {
	ddu.mutation.SetIps(s)
	return ddu
}

// SetServers sets the "servers" field.
func (ddu *DNSDetectorUpdate) SetServers(s []string) *DNSDetectorUpdate {
	ddu.mutation.SetServers(s)
	return ddu
}

// Mutation returns the DNSDetectorMutation object of the builder.
func (ddu *DNSDetectorUpdate) Mutation() *DNSDetectorMutation {
	return ddu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ddu *DNSDetectorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ddu.defaults()
	if len(ddu.hooks) == 0 {
		if err = ddu.check(); err != nil {
			return 0, err
		}
		affected, err = ddu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DNSDetectorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ddu.check(); err != nil {
				return 0, err
			}
			ddu.mutation = mutation
			affected, err = ddu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ddu.hooks) - 1; i >= 0; i-- {
			mut = ddu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ddu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ddu *DNSDetectorUpdate) SaveX(ctx context.Context) int {
	affected, err := ddu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ddu *DNSDetectorUpdate) Exec(ctx context.Context) error {
	_, err := ddu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddu *DNSDetectorUpdate) ExecX(ctx context.Context) {
	if err := ddu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ddu *DNSDetectorUpdate) defaults() {
	if _, ok := ddu.mutation.UpdatedAt(); !ok {
		v := dnsdetector.UpdateDefaultUpdatedAt()
		ddu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ddu *DNSDetectorUpdate) check() error {
	if v, ok := ddu.mutation.Status(); ok {
		if err := dnsdetector.StatusValidator(int8(v)); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := ddu.mutation.Name(); ok {
		if err := dnsdetector.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := ddu.mutation.Owner(); ok {
		if err := dnsdetector.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf("ent: validator failed for field \"owner\": %w", err)}
		}
	}
	if v, ok := ddu.mutation.Timeout(); ok {
		if err := dnsdetector.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf("ent: validator failed for field \"timeout\": %w", err)}
		}
	}
	if v, ok := ddu.mutation.Host(); ok {
		if err := dnsdetector.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf("ent: validator failed for field \"host\": %w", err)}
		}
	}
	return nil
}

func (ddu *DNSDetectorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dnsdetector.Table,
			Columns: dnsdetector.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dnsdetector.FieldID,
			},
		},
	}
	if ps := ddu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ddu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dnsdetector.FieldUpdatedAt,
		})
	}
	if value, ok := ddu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: dnsdetector.FieldStatus,
		})
	}
	if value, ok := ddu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: dnsdetector.FieldStatus,
		})
	}
	if value, ok := ddu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dnsdetector.FieldName,
		})
	}
	if value, ok := ddu.mutation.Owner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dnsdetector.FieldOwner,
		})
	}
	if value, ok := ddu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dnsdetector.FieldDescription,
		})
	}
	if value, ok := ddu.mutation.Receivers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dnsdetector.FieldReceivers,
		})
	}
	if value, ok := ddu.mutation.Timeout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dnsdetector.FieldTimeout,
		})
	}
	if value, ok := ddu.mutation.Host(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dnsdetector.FieldHost,
		})
	}
	if value, ok := ddu.mutation.Ips(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dnsdetector.FieldIps,
		})
	}
	if value, ok := ddu.mutation.Servers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dnsdetector.FieldServers,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ddu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dnsdetector.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DNSDetectorUpdateOne is the builder for updating a single DNSDetector entity.
type DNSDetectorUpdateOne struct {
	config
	hooks    []Hook
	mutation *DNSDetectorMutation
}

// SetStatus sets the "status" field.
func (dduo *DNSDetectorUpdateOne) SetStatus(s schema.Status) *DNSDetectorUpdateOne {
	dduo.mutation.ResetStatus()
	dduo.mutation.SetStatus(s)
	return dduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dduo *DNSDetectorUpdateOne) SetNillableStatus(s *schema.Status) *DNSDetectorUpdateOne {
	if s != nil {
		dduo.SetStatus(*s)
	}
	return dduo
}

// AddStatus adds s to the "status" field.
func (dduo *DNSDetectorUpdateOne) AddStatus(s schema.Status) *DNSDetectorUpdateOne {
	dduo.mutation.AddStatus(s)
	return dduo
}

// SetName sets the "name" field.
func (dduo *DNSDetectorUpdateOne) SetName(s string) *DNSDetectorUpdateOne {
	dduo.mutation.SetName(s)
	return dduo
}

// SetOwner sets the "owner" field.
func (dduo *DNSDetectorUpdateOne) SetOwner(s string) *DNSDetectorUpdateOne {
	dduo.mutation.SetOwner(s)
	return dduo
}

// SetDescription sets the "description" field.
func (dduo *DNSDetectorUpdateOne) SetDescription(s string) *DNSDetectorUpdateOne {
	dduo.mutation.SetDescription(s)
	return dduo
}

// SetReceivers sets the "receivers" field.
func (dduo *DNSDetectorUpdateOne) SetReceivers(s []string) *DNSDetectorUpdateOne {
	dduo.mutation.SetReceivers(s)
	return dduo
}

// SetTimeout sets the "timeout" field.
func (dduo *DNSDetectorUpdateOne) SetTimeout(s string) *DNSDetectorUpdateOne {
	dduo.mutation.SetTimeout(s)
	return dduo
}

// SetHost sets the "host" field.
func (dduo *DNSDetectorUpdateOne) SetHost(s string) *DNSDetectorUpdateOne {
	dduo.mutation.SetHost(s)
	return dduo
}

// SetIps sets the "ips" field.
func (dduo *DNSDetectorUpdateOne) SetIps(s []string) *DNSDetectorUpdateOne {
	dduo.mutation.SetIps(s)
	return dduo
}

// SetServers sets the "servers" field.
func (dduo *DNSDetectorUpdateOne) SetServers(s []string) *DNSDetectorUpdateOne {
	dduo.mutation.SetServers(s)
	return dduo
}

// Mutation returns the DNSDetectorMutation object of the builder.
func (dduo *DNSDetectorUpdateOne) Mutation() *DNSDetectorMutation {
	return dduo.mutation
}

// Save executes the query and returns the updated DNSDetector entity.
func (dduo *DNSDetectorUpdateOne) Save(ctx context.Context) (*DNSDetector, error) {
	var (
		err  error
		node *DNSDetector
	)
	dduo.defaults()
	if len(dduo.hooks) == 0 {
		if err = dduo.check(); err != nil {
			return nil, err
		}
		node, err = dduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DNSDetectorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dduo.check(); err != nil {
				return nil, err
			}
			dduo.mutation = mutation
			node, err = dduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dduo.hooks) - 1; i >= 0; i-- {
			mut = dduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dduo *DNSDetectorUpdateOne) SaveX(ctx context.Context) *DNSDetector {
	node, err := dduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dduo *DNSDetectorUpdateOne) Exec(ctx context.Context) error {
	_, err := dduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dduo *DNSDetectorUpdateOne) ExecX(ctx context.Context) {
	if err := dduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dduo *DNSDetectorUpdateOne) defaults() {
	if _, ok := dduo.mutation.UpdatedAt(); !ok {
		v := dnsdetector.UpdateDefaultUpdatedAt()
		dduo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dduo *DNSDetectorUpdateOne) check() error {
	if v, ok := dduo.mutation.Status(); ok {
		if err := dnsdetector.StatusValidator(int8(v)); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := dduo.mutation.Name(); ok {
		if err := dnsdetector.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := dduo.mutation.Owner(); ok {
		if err := dnsdetector.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf("ent: validator failed for field \"owner\": %w", err)}
		}
	}
	if v, ok := dduo.mutation.Timeout(); ok {
		if err := dnsdetector.TimeoutValidator(v); err != nil {
			return &ValidationError{Name: "timeout", err: fmt.Errorf("ent: validator failed for field \"timeout\": %w", err)}
		}
	}
	if v, ok := dduo.mutation.Host(); ok {
		if err := dnsdetector.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf("ent: validator failed for field \"host\": %w", err)}
		}
	}
	return nil
}

func (dduo *DNSDetectorUpdateOne) sqlSave(ctx context.Context) (_node *DNSDetector, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dnsdetector.Table,
			Columns: dnsdetector.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dnsdetector.FieldID,
			},
		},
	}
	id, ok := dduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing DNSDetector.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := dduo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dnsdetector.FieldUpdatedAt,
		})
	}
	if value, ok := dduo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: dnsdetector.FieldStatus,
		})
	}
	if value, ok := dduo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: dnsdetector.FieldStatus,
		})
	}
	if value, ok := dduo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dnsdetector.FieldName,
		})
	}
	if value, ok := dduo.mutation.Owner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dnsdetector.FieldOwner,
		})
	}
	if value, ok := dduo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dnsdetector.FieldDescription,
		})
	}
	if value, ok := dduo.mutation.Receivers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dnsdetector.FieldReceivers,
		})
	}
	if value, ok := dduo.mutation.Timeout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dnsdetector.FieldTimeout,
		})
	}
	if value, ok := dduo.mutation.Host(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dnsdetector.FieldHost,
		})
	}
	if value, ok := dduo.mutation.Ips(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dnsdetector.FieldIps,
		})
	}
	if value, ok := dduo.mutation.Servers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dnsdetector.FieldServers,
		})
	}
	_node = &DNSDetector{config: dduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dnsdetector.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}