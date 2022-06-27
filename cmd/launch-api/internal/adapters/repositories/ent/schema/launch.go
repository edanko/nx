package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Launch holds the schema definition for the Launch entity.
type Launch struct {
	ent.Schema
}

// Fields of the Launch.
func (Launch) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.Enum("status").
			Values("todo", "started", "completed").
			Default("todo"),

		field.Text("applicant").
			NotEmpty(),
		field.Text("reason").
			NotEmpty(),

		field.Text("description").
			Optional().
			Nillable(),
	}
}

// Edges of the Launch.
func (Launch) Edges() []ent.Edge {
	return nil
}
