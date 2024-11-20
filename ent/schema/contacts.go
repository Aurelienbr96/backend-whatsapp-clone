package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Contact struct {
	ent.Schema
}

func (Contact) Fields() []ent.Field {
	return []ent.Field{
		// Explicit fields to reference in the index and relationships
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.UUID("owner_id", uuid.UUID{}),
		field.UUID("contact_user_id", uuid.UUID{}),
	}
}

func (Contact) Edges() []ent.Edge {
	return []ent.Edge{
		// Edges reference the explicitly defined fields
		edge.From("owner", User.Type).
			Ref("contacts").
			Unique().
			Required().
			Field("owner_id"),
		edge.To("contact_user", User.Type).
			Unique().
			Required().
			Field("contact_user_id"),
	}
}

func (Contact) Indexes() []ent.Index {
	return []ent.Index{
		// Index refers to the explicitly defined fields
		index.Fields("owner_id", "contact_user_id").
			Unique(),
	}
}
