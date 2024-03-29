// Code generated by ent, DO NOT EDIT.

package configuration

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/schema"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Configuration {
	return predicate.Configuration(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Configuration {
	return predicate.Configuration(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Configuration {
	return predicate.Configuration(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Configuration {
	return predicate.Configuration(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Configuration {
	return predicate.Configuration(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Configuration {
	return predicate.Configuration(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Configuration {
	return predicate.Configuration(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v schema.Status) predicate.Configuration {
	vc := int8(v)
	return predicate.Configuration(sql.FieldEQ(FieldStatus, vc))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldName, v))
}

// Owner applies equality check predicate on the "owner" field. It's identical to OwnerEQ.
func Owner(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldOwner, v))
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldData, v))
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldStartedAt, v))
}

// EndedAt applies equality check predicate on the "ended_at" field. It's identical to EndedAtEQ.
func EndedAt(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldEndedAt, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldDescription, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v schema.Status) predicate.Configuration {
	vc := int8(v)
	return predicate.Configuration(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v schema.Status) predicate.Configuration {
	vc := int8(v)
	return predicate.Configuration(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...schema.Status) predicate.Configuration {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int8(vs[i])
	}
	return predicate.Configuration(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...schema.Status) predicate.Configuration {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int8(vs[i])
	}
	return predicate.Configuration(sql.FieldNotIn(FieldStatus, v...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v schema.Status) predicate.Configuration {
	vc := int8(v)
	return predicate.Configuration(sql.FieldGT(FieldStatus, vc))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v schema.Status) predicate.Configuration {
	vc := int8(v)
	return predicate.Configuration(sql.FieldGTE(FieldStatus, vc))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v schema.Status) predicate.Configuration {
	vc := int8(v)
	return predicate.Configuration(sql.FieldLT(FieldStatus, vc))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v schema.Status) predicate.Configuration {
	vc := int8(v)
	return predicate.Configuration(sql.FieldLTE(FieldStatus, vc))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Configuration {
	return predicate.Configuration(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Configuration {
	return predicate.Configuration(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldContainsFold(FieldName, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v Category) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v Category) predicate.Configuration {
	return predicate.Configuration(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...Category) predicate.Configuration {
	return predicate.Configuration(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...Category) predicate.Configuration {
	return predicate.Configuration(sql.FieldNotIn(FieldCategory, vs...))
}

// OwnerEQ applies the EQ predicate on the "owner" field.
func OwnerEQ(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldOwner, v))
}

// OwnerNEQ applies the NEQ predicate on the "owner" field.
func OwnerNEQ(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldNEQ(FieldOwner, v))
}

// OwnerIn applies the In predicate on the "owner" field.
func OwnerIn(vs ...string) predicate.Configuration {
	return predicate.Configuration(sql.FieldIn(FieldOwner, vs...))
}

// OwnerNotIn applies the NotIn predicate on the "owner" field.
func OwnerNotIn(vs ...string) predicate.Configuration {
	return predicate.Configuration(sql.FieldNotIn(FieldOwner, vs...))
}

// OwnerGT applies the GT predicate on the "owner" field.
func OwnerGT(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldGT(FieldOwner, v))
}

// OwnerGTE applies the GTE predicate on the "owner" field.
func OwnerGTE(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldGTE(FieldOwner, v))
}

// OwnerLT applies the LT predicate on the "owner" field.
func OwnerLT(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldLT(FieldOwner, v))
}

// OwnerLTE applies the LTE predicate on the "owner" field.
func OwnerLTE(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldLTE(FieldOwner, v))
}

// OwnerContains applies the Contains predicate on the "owner" field.
func OwnerContains(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldContains(FieldOwner, v))
}

// OwnerHasPrefix applies the HasPrefix predicate on the "owner" field.
func OwnerHasPrefix(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldHasPrefix(FieldOwner, v))
}

// OwnerHasSuffix applies the HasSuffix predicate on the "owner" field.
func OwnerHasSuffix(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldHasSuffix(FieldOwner, v))
}

// OwnerEqualFold applies the EqualFold predicate on the "owner" field.
func OwnerEqualFold(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEqualFold(FieldOwner, v))
}

// OwnerContainsFold applies the ContainsFold predicate on the "owner" field.
func OwnerContainsFold(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldContainsFold(FieldOwner, v))
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...string) predicate.Configuration {
	return predicate.Configuration(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...string) predicate.Configuration {
	return predicate.Configuration(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldLTE(FieldData, v))
}

// DataContains applies the Contains predicate on the "data" field.
func DataContains(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldContains(FieldData, v))
}

// DataHasPrefix applies the HasPrefix predicate on the "data" field.
func DataHasPrefix(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldHasPrefix(FieldData, v))
}

// DataHasSuffix applies the HasSuffix predicate on the "data" field.
func DataHasSuffix(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldHasSuffix(FieldData, v))
}

// DataEqualFold applies the EqualFold predicate on the "data" field.
func DataEqualFold(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEqualFold(FieldData, v))
}

// DataContainsFold applies the ContainsFold predicate on the "data" field.
func DataContainsFold(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldContainsFold(FieldData, v))
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldStartedAt, v))
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldNEQ(FieldStartedAt, v))
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldIn(FieldStartedAt, vs...))
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldNotIn(FieldStartedAt, vs...))
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldGT(FieldStartedAt, v))
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldGTE(FieldStartedAt, v))
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldLT(FieldStartedAt, v))
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldLTE(FieldStartedAt, v))
}

// EndedAtEQ applies the EQ predicate on the "ended_at" field.
func EndedAtEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldEndedAt, v))
}

// EndedAtNEQ applies the NEQ predicate on the "ended_at" field.
func EndedAtNEQ(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldNEQ(FieldEndedAt, v))
}

// EndedAtIn applies the In predicate on the "ended_at" field.
func EndedAtIn(vs ...time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldIn(FieldEndedAt, vs...))
}

// EndedAtNotIn applies the NotIn predicate on the "ended_at" field.
func EndedAtNotIn(vs ...time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldNotIn(FieldEndedAt, vs...))
}

// EndedAtGT applies the GT predicate on the "ended_at" field.
func EndedAtGT(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldGT(FieldEndedAt, v))
}

// EndedAtGTE applies the GTE predicate on the "ended_at" field.
func EndedAtGTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldGTE(FieldEndedAt, v))
}

// EndedAtLT applies the LT predicate on the "ended_at" field.
func EndedAtLT(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldLT(FieldEndedAt, v))
}

// EndedAtLTE applies the LTE predicate on the "ended_at" field.
func EndedAtLTE(v time.Time) predicate.Configuration {
	return predicate.Configuration(sql.FieldLTE(FieldEndedAt, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Configuration {
	return predicate.Configuration(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Configuration {
	return predicate.Configuration(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Configuration {
	return predicate.Configuration(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Configuration {
	return predicate.Configuration(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldContainsFold(FieldDescription, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Configuration) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Configuration) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
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
func Not(p predicate.Configuration) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		p(s.Not())
	})
}
