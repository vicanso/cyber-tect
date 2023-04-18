// Code generated by ent, DO NOT EDIT.

package dnsdetectorresult

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/schema"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Task applies equality check predicate on the "task" field. It's identical to TaskEQ.
func Task(v int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTask), v))
	})
}

// Result applies equality check predicate on the "result" field. It's identical to ResultEQ.
func Result(v schema.DetectorResult) predicate.DNSDetectorResult {
	vc := int8(v)
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResult), vc))
	})
}

// MaxDuration applies equality check predicate on the "maxDuration" field. It's identical to MaxDurationEQ.
func MaxDuration(v int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxDuration), v))
	})
}

// Host applies equality check predicate on the "host" field. It's identical to HostEQ.
func Host(v string) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHost), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DNSDetectorResult {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DNSDetectorResult {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DNSDetectorResult {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DNSDetectorResult {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// TaskEQ applies the EQ predicate on the "task" field.
func TaskEQ(v int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTask), v))
	})
}

// TaskNEQ applies the NEQ predicate on the "task" field.
func TaskNEQ(v int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTask), v))
	})
}

// TaskIn applies the In predicate on the "task" field.
func TaskIn(vs ...int) predicate.DNSDetectorResult {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTask), v...))
	})
}

// TaskNotIn applies the NotIn predicate on the "task" field.
func TaskNotIn(vs ...int) predicate.DNSDetectorResult {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTask), v...))
	})
}

// TaskGT applies the GT predicate on the "task" field.
func TaskGT(v int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTask), v))
	})
}

// TaskGTE applies the GTE predicate on the "task" field.
func TaskGTE(v int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTask), v))
	})
}

// TaskLT applies the LT predicate on the "task" field.
func TaskLT(v int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTask), v))
	})
}

// TaskLTE applies the LTE predicate on the "task" field.
func TaskLTE(v int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTask), v))
	})
}

// ResultEQ applies the EQ predicate on the "result" field.
func ResultEQ(v schema.DetectorResult) predicate.DNSDetectorResult {
	vc := int8(v)
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResult), vc))
	})
}

// ResultNEQ applies the NEQ predicate on the "result" field.
func ResultNEQ(v schema.DetectorResult) predicate.DNSDetectorResult {
	vc := int8(v)
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResult), vc))
	})
}

// ResultIn applies the In predicate on the "result" field.
func ResultIn(vs ...schema.DetectorResult) predicate.DNSDetectorResult {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int8(vs[i])
	}
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldResult), v...))
	})
}

// ResultNotIn applies the NotIn predicate on the "result" field.
func ResultNotIn(vs ...schema.DetectorResult) predicate.DNSDetectorResult {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int8(vs[i])
	}
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldResult), v...))
	})
}

// ResultGT applies the GT predicate on the "result" field.
func ResultGT(v schema.DetectorResult) predicate.DNSDetectorResult {
	vc := int8(v)
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResult), vc))
	})
}

// ResultGTE applies the GTE predicate on the "result" field.
func ResultGTE(v schema.DetectorResult) predicate.DNSDetectorResult {
	vc := int8(v)
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResult), vc))
	})
}

// ResultLT applies the LT predicate on the "result" field.
func ResultLT(v schema.DetectorResult) predicate.DNSDetectorResult {
	vc := int8(v)
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResult), vc))
	})
}

// ResultLTE applies the LTE predicate on the "result" field.
func ResultLTE(v schema.DetectorResult) predicate.DNSDetectorResult {
	vc := int8(v)
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResult), vc))
	})
}

// MaxDurationEQ applies the EQ predicate on the "maxDuration" field.
func MaxDurationEQ(v int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxDuration), v))
	})
}

// MaxDurationNEQ applies the NEQ predicate on the "maxDuration" field.
func MaxDurationNEQ(v int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaxDuration), v))
	})
}

// MaxDurationIn applies the In predicate on the "maxDuration" field.
func MaxDurationIn(vs ...int) predicate.DNSDetectorResult {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMaxDuration), v...))
	})
}

// MaxDurationNotIn applies the NotIn predicate on the "maxDuration" field.
func MaxDurationNotIn(vs ...int) predicate.DNSDetectorResult {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMaxDuration), v...))
	})
}

// MaxDurationGT applies the GT predicate on the "maxDuration" field.
func MaxDurationGT(v int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMaxDuration), v))
	})
}

// MaxDurationGTE applies the GTE predicate on the "maxDuration" field.
func MaxDurationGTE(v int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMaxDuration), v))
	})
}

// MaxDurationLT applies the LT predicate on the "maxDuration" field.
func MaxDurationLT(v int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMaxDuration), v))
	})
}

// MaxDurationLTE applies the LTE predicate on the "maxDuration" field.
func MaxDurationLTE(v int) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMaxDuration), v))
	})
}

// HostEQ applies the EQ predicate on the "host" field.
func HostEQ(v string) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHost), v))
	})
}

// HostNEQ applies the NEQ predicate on the "host" field.
func HostNEQ(v string) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHost), v))
	})
}

// HostIn applies the In predicate on the "host" field.
func HostIn(vs ...string) predicate.DNSDetectorResult {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHost), v...))
	})
}

// HostNotIn applies the NotIn predicate on the "host" field.
func HostNotIn(vs ...string) predicate.DNSDetectorResult {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHost), v...))
	})
}

// HostGT applies the GT predicate on the "host" field.
func HostGT(v string) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHost), v))
	})
}

// HostGTE applies the GTE predicate on the "host" field.
func HostGTE(v string) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHost), v))
	})
}

// HostLT applies the LT predicate on the "host" field.
func HostLT(v string) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHost), v))
	})
}

// HostLTE applies the LTE predicate on the "host" field.
func HostLTE(v string) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHost), v))
	})
}

// HostContains applies the Contains predicate on the "host" field.
func HostContains(v string) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHost), v))
	})
}

// HostHasPrefix applies the HasPrefix predicate on the "host" field.
func HostHasPrefix(v string) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHost), v))
	})
}

// HostHasSuffix applies the HasSuffix predicate on the "host" field.
func HostHasSuffix(v string) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHost), v))
	})
}

// HostEqualFold applies the EqualFold predicate on the "host" field.
func HostEqualFold(v string) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHost), v))
	})
}

// HostContainsFold applies the ContainsFold predicate on the "host" field.
func HostContainsFold(v string) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHost), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DNSDetectorResult) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DNSDetectorResult) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DNSDetectorResult) predicate.DNSDetectorResult {
	return predicate.DNSDetectorResult(func(s *sql.Selector) {
		p(s.Not())
	})
}
