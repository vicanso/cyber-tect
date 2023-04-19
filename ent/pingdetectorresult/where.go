// Code generated by ent, DO NOT EDIT.

package pingdetectorresult

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/schema"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldEQ(FieldUpdatedAt, v))
}

// Task applies equality check predicate on the "task" field. It's identical to TaskEQ.
func Task(v int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldEQ(FieldTask, v))
}

// Result applies equality check predicate on the "result" field. It's identical to ResultEQ.
func Result(v schema.DetectorResult) predicate.PingDetectorResult {
	vc := int8(v)
	return predicate.PingDetectorResult(sql.FieldEQ(FieldResult, vc))
}

// MaxDuration applies equality check predicate on the "maxDuration" field. It's identical to MaxDurationEQ.
func MaxDuration(v int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldEQ(FieldMaxDuration, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldLTE(FieldUpdatedAt, v))
}

// TaskEQ applies the EQ predicate on the "task" field.
func TaskEQ(v int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldEQ(FieldTask, v))
}

// TaskNEQ applies the NEQ predicate on the "task" field.
func TaskNEQ(v int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldNEQ(FieldTask, v))
}

// TaskIn applies the In predicate on the "task" field.
func TaskIn(vs ...int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldIn(FieldTask, vs...))
}

// TaskNotIn applies the NotIn predicate on the "task" field.
func TaskNotIn(vs ...int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldNotIn(FieldTask, vs...))
}

// TaskGT applies the GT predicate on the "task" field.
func TaskGT(v int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldGT(FieldTask, v))
}

// TaskGTE applies the GTE predicate on the "task" field.
func TaskGTE(v int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldGTE(FieldTask, v))
}

// TaskLT applies the LT predicate on the "task" field.
func TaskLT(v int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldLT(FieldTask, v))
}

// TaskLTE applies the LTE predicate on the "task" field.
func TaskLTE(v int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldLTE(FieldTask, v))
}

// ResultEQ applies the EQ predicate on the "result" field.
func ResultEQ(v schema.DetectorResult) predicate.PingDetectorResult {
	vc := int8(v)
	return predicate.PingDetectorResult(sql.FieldEQ(FieldResult, vc))
}

// ResultNEQ applies the NEQ predicate on the "result" field.
func ResultNEQ(v schema.DetectorResult) predicate.PingDetectorResult {
	vc := int8(v)
	return predicate.PingDetectorResult(sql.FieldNEQ(FieldResult, vc))
}

// ResultIn applies the In predicate on the "result" field.
func ResultIn(vs ...schema.DetectorResult) predicate.PingDetectorResult {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int8(vs[i])
	}
	return predicate.PingDetectorResult(sql.FieldIn(FieldResult, v...))
}

// ResultNotIn applies the NotIn predicate on the "result" field.
func ResultNotIn(vs ...schema.DetectorResult) predicate.PingDetectorResult {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int8(vs[i])
	}
	return predicate.PingDetectorResult(sql.FieldNotIn(FieldResult, v...))
}

// ResultGT applies the GT predicate on the "result" field.
func ResultGT(v schema.DetectorResult) predicate.PingDetectorResult {
	vc := int8(v)
	return predicate.PingDetectorResult(sql.FieldGT(FieldResult, vc))
}

// ResultGTE applies the GTE predicate on the "result" field.
func ResultGTE(v schema.DetectorResult) predicate.PingDetectorResult {
	vc := int8(v)
	return predicate.PingDetectorResult(sql.FieldGTE(FieldResult, vc))
}

// ResultLT applies the LT predicate on the "result" field.
func ResultLT(v schema.DetectorResult) predicate.PingDetectorResult {
	vc := int8(v)
	return predicate.PingDetectorResult(sql.FieldLT(FieldResult, vc))
}

// ResultLTE applies the LTE predicate on the "result" field.
func ResultLTE(v schema.DetectorResult) predicate.PingDetectorResult {
	vc := int8(v)
	return predicate.PingDetectorResult(sql.FieldLTE(FieldResult, vc))
}

// MaxDurationEQ applies the EQ predicate on the "maxDuration" field.
func MaxDurationEQ(v int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldEQ(FieldMaxDuration, v))
}

// MaxDurationNEQ applies the NEQ predicate on the "maxDuration" field.
func MaxDurationNEQ(v int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldNEQ(FieldMaxDuration, v))
}

// MaxDurationIn applies the In predicate on the "maxDuration" field.
func MaxDurationIn(vs ...int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldIn(FieldMaxDuration, vs...))
}

// MaxDurationNotIn applies the NotIn predicate on the "maxDuration" field.
func MaxDurationNotIn(vs ...int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldNotIn(FieldMaxDuration, vs...))
}

// MaxDurationGT applies the GT predicate on the "maxDuration" field.
func MaxDurationGT(v int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldGT(FieldMaxDuration, v))
}

// MaxDurationGTE applies the GTE predicate on the "maxDuration" field.
func MaxDurationGTE(v int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldGTE(FieldMaxDuration, v))
}

// MaxDurationLT applies the LT predicate on the "maxDuration" field.
func MaxDurationLT(v int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldLT(FieldMaxDuration, v))
}

// MaxDurationLTE applies the LTE predicate on the "maxDuration" field.
func MaxDurationLTE(v int) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(sql.FieldLTE(FieldMaxDuration, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PingDetectorResult) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PingDetectorResult) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(func(s *sql.Selector) {
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
func Not(p predicate.PingDetectorResult) predicate.PingDetectorResult {
	return predicate.PingDetectorResult(func(s *sql.Selector) {
		p(s.Not())
	})
}