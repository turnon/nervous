// Code generated by entc, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/turnon/nervous/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
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
func NameNotIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
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
func NameGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartAt), v))
	})
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartAt), v...))
	})
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartAt), v...))
	})
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartAt), v))
	})
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartAt), v))
	})
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartAt), v))
	})
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartAt), v))
	})
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndAt), v))
	})
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndAt), v...))
	})
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndAt), v...))
	})
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndAt), v))
	})
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndAt), v))
	})
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndAt), v))
	})
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndAt), v))
	})
}

// HasTag applies the HasEdge predicate on the "tag" edge.
func HasTag() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TagTable, TagColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagWith applies the HasEdge predicate on the "tag" edge with a given conditions (other predicates).
func HasTagWith(preds ...predicate.Tag) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TagTable, TagColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
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
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		p(s.Not())
	})
}
