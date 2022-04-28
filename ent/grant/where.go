// Code generated by entc, DO NOT EDIT.

package grant

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/shifty11/cosmos-gov/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Grantee applies equality check predicate on the "grantee" field. It's identical to GranteeEQ.
func Grantee(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGrantee), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiresAt), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// GranteeEQ applies the EQ predicate on the "grantee" field.
func GranteeEQ(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGrantee), v))
	})
}

// GranteeNEQ applies the NEQ predicate on the "grantee" field.
func GranteeNEQ(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGrantee), v))
	})
}

// GranteeIn applies the In predicate on the "grantee" field.
func GranteeIn(vs ...string) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGrantee), v...))
	})
}

// GranteeNotIn applies the NotIn predicate on the "grantee" field.
func GranteeNotIn(vs ...string) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGrantee), v...))
	})
}

// GranteeGT applies the GT predicate on the "grantee" field.
func GranteeGT(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGrantee), v))
	})
}

// GranteeGTE applies the GTE predicate on the "grantee" field.
func GranteeGTE(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGrantee), v))
	})
}

// GranteeLT applies the LT predicate on the "grantee" field.
func GranteeLT(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGrantee), v))
	})
}

// GranteeLTE applies the LTE predicate on the "grantee" field.
func GranteeLTE(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGrantee), v))
	})
}

// GranteeContains applies the Contains predicate on the "grantee" field.
func GranteeContains(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGrantee), v))
	})
}

// GranteeHasPrefix applies the HasPrefix predicate on the "grantee" field.
func GranteeHasPrefix(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGrantee), v))
	})
}

// GranteeHasSuffix applies the HasSuffix predicate on the "grantee" field.
func GranteeHasSuffix(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGrantee), v))
	})
}

// GranteeEqualFold applies the EqualFold predicate on the "grantee" field.
func GranteeEqualFold(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGrantee), v))
	})
}

// GranteeContainsFold applies the ContainsFold predicate on the "grantee" field.
func GranteeContainsFold(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGrantee), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
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
func TypeNotIn(vs ...string) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
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
func TypeGT(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpiresAt), v...))
	})
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.Grant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpiresAt), v...))
	})
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpiresAt), v))
	})
}

// HasGranter applies the HasEdge predicate on the "granter" edge.
func HasGranter() predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GranterTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GranterTable, GranterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGranterWith applies the HasEdge predicate on the "granter" edge with a given conditions (other predicates).
func HasGranterWith(preds ...predicate.Wallet) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GranterInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GranterTable, GranterColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Grant) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Grant) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
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
func Not(p predicate.Grant) predicate.Grant {
	return predicate.Grant(func(s *sql.Selector) {
		p(s.Not())
	})
}