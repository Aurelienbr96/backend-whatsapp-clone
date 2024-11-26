// Code generated by ent, DO NOT EDIT.

package contact

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"example.com/boiletplate/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldLTE(FieldID, id))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldEQ(FieldOwnerID, v))
}

// ContactUserID applies equality check predicate on the "contact_user_id" field. It's identical to ContactUserIDEQ.
func ContactUserID(v uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldEQ(FieldContactUserID, v))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldNotIn(FieldOwnerID, vs...))
}

// ContactUserIDEQ applies the EQ predicate on the "contact_user_id" field.
func ContactUserIDEQ(v uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldEQ(FieldContactUserID, v))
}

// ContactUserIDNEQ applies the NEQ predicate on the "contact_user_id" field.
func ContactUserIDNEQ(v uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldNEQ(FieldContactUserID, v))
}

// ContactUserIDIn applies the In predicate on the "contact_user_id" field.
func ContactUserIDIn(vs ...uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldIn(FieldContactUserID, vs...))
}

// ContactUserIDNotIn applies the NotIn predicate on the "contact_user_id" field.
func ContactUserIDNotIn(vs ...uuid.UUID) predicate.Contact {
	return predicate.Contact(sql.FieldNotIn(FieldContactUserID, vs...))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasContactUser applies the HasEdge predicate on the "contact_user" edge.
func HasContactUser() predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ContactUserTable, ContactUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContactUserWith applies the HasEdge predicate on the "contact_user" edge with a given conditions (other predicates).
func HasContactUserWith(preds ...predicate.User) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		step := newContactUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Contact) predicate.Contact {
	return predicate.Contact(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Contact) predicate.Contact {
	return predicate.Contact(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Contact) predicate.Contact {
	return predicate.Contact(sql.NotPredicates(p))
}
