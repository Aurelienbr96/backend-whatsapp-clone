package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("phone_number").Unique(),
		field.String("name"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
