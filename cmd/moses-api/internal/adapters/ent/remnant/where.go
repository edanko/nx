// Code generated by ent, DO NOT EDIT.

package remnant

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/edanko/nx/cmd/moses-api/internal/adapters/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Quality applies equality check predicate on the "quality" field. It's identical to QualityEQ.
func Quality(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuality), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// Length applies equality check predicate on the "length" field. It's identical to LengthEQ.
func Length(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLength), v))
	})
}

// Width applies equality check predicate on the "width" field. It's identical to WidthEQ.
func Width(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWidth), v))
	})
}

// Thickness applies equality check predicate on the "thickness" field. It's identical to ThicknessEQ.
func Thickness(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThickness), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// QualityEQ applies the EQ predicate on the "quality" field.
func QualityEQ(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuality), v))
	})
}

// QualityNEQ applies the NEQ predicate on the "quality" field.
func QualityNEQ(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQuality), v))
	})
}

// QualityIn applies the In predicate on the "quality" field.
func QualityIn(vs ...string) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldQuality), v...))
	})
}

// QualityNotIn applies the NotIn predicate on the "quality" field.
func QualityNotIn(vs ...string) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldQuality), v...))
	})
}

// QualityGT applies the GT predicate on the "quality" field.
func QualityGT(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQuality), v))
	})
}

// QualityGTE applies the GTE predicate on the "quality" field.
func QualityGTE(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQuality), v))
	})
}

// QualityLT applies the LT predicate on the "quality" field.
func QualityLT(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQuality), v))
	})
}

// QualityLTE applies the LTE predicate on the "quality" field.
func QualityLTE(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQuality), v))
	})
}

// QualityContains applies the Contains predicate on the "quality" field.
func QualityContains(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldQuality), v))
	})
}

// QualityHasPrefix applies the HasPrefix predicate on the "quality" field.
func QualityHasPrefix(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldQuality), v))
	})
}

// QualityHasSuffix applies the HasSuffix predicate on the "quality" field.
func QualityHasSuffix(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldQuality), v))
	})
}

// QualityEqualFold applies the EqualFold predicate on the "quality" field.
func QualityEqualFold(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldQuality), v))
	})
}

// QualityContainsFold applies the ContainsFold predicate on the "quality" field.
func QualityContainsFold(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldQuality), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// LengthEQ applies the EQ predicate on the "length" field.
func LengthEQ(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLength), v))
	})
}

// LengthNEQ applies the NEQ predicate on the "length" field.
func LengthNEQ(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLength), v))
	})
}

// LengthIn applies the In predicate on the "length" field.
func LengthIn(vs ...float64) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLength), v...))
	})
}

// LengthNotIn applies the NotIn predicate on the "length" field.
func LengthNotIn(vs ...float64) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLength), v...))
	})
}

// LengthGT applies the GT predicate on the "length" field.
func LengthGT(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLength), v))
	})
}

// LengthGTE applies the GTE predicate on the "length" field.
func LengthGTE(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLength), v))
	})
}

// LengthLT applies the LT predicate on the "length" field.
func LengthLT(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLength), v))
	})
}

// LengthLTE applies the LTE predicate on the "length" field.
func LengthLTE(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLength), v))
	})
}

// WidthEQ applies the EQ predicate on the "width" field.
func WidthEQ(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWidth), v))
	})
}

// WidthNEQ applies the NEQ predicate on the "width" field.
func WidthNEQ(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWidth), v))
	})
}

// WidthIn applies the In predicate on the "width" field.
func WidthIn(vs ...float64) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWidth), v...))
	})
}

// WidthNotIn applies the NotIn predicate on the "width" field.
func WidthNotIn(vs ...float64) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWidth), v...))
	})
}

// WidthGT applies the GT predicate on the "width" field.
func WidthGT(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWidth), v))
	})
}

// WidthGTE applies the GTE predicate on the "width" field.
func WidthGTE(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWidth), v))
	})
}

// WidthLT applies the LT predicate on the "width" field.
func WidthLT(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWidth), v))
	})
}

// WidthLTE applies the LTE predicate on the "width" field.
func WidthLTE(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWidth), v))
	})
}

// ThicknessEQ applies the EQ predicate on the "thickness" field.
func ThicknessEQ(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThickness), v))
	})
}

// ThicknessNEQ applies the NEQ predicate on the "thickness" field.
func ThicknessNEQ(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThickness), v))
	})
}

// ThicknessIn applies the In predicate on the "thickness" field.
func ThicknessIn(vs ...float64) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldThickness), v...))
	})
}

// ThicknessNotIn applies the NotIn predicate on the "thickness" field.
func ThicknessNotIn(vs ...float64) predicate.Remnant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Remnant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldThickness), v...))
	})
}

// ThicknessGT applies the GT predicate on the "thickness" field.
func ThicknessGT(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThickness), v))
	})
}

// ThicknessGTE applies the GTE predicate on the "thickness" field.
func ThicknessGTE(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThickness), v))
	})
}

// ThicknessLT applies the LT predicate on the "thickness" field.
func ThicknessLT(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThickness), v))
	})
}

// ThicknessLTE applies the LTE predicate on the "thickness" field.
func ThicknessLTE(v float64) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThickness), v))
	})
}

// HasNest applies the HasEdge predicate on the "nest" edge.
func HasNest() predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NestTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, NestTable, NestColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNestWith applies the HasEdge predicate on the "nest" edge with a given conditions (other predicates).
func HasNestWith(preds ...predicate.Nest) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NestInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, NestTable, NestColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRemnantUsed applies the HasEdge predicate on the "remnant_used" edge.
func HasRemnantUsed() predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RemnantUsedTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RemnantUsedTable, RemnantUsedColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRemnantUsedWith applies the HasEdge predicate on the "remnant_used" edge with a given conditions (other predicates).
func HasRemnantUsedWith(preds ...predicate.Nest) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RemnantUsedInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RemnantUsedTable, RemnantUsedColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Remnant) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Remnant) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
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
func Not(p predicate.Remnant) predicate.Remnant {
	return predicate.Remnant(func(s *sql.Selector) {
		p(s.Not())
	})
}
