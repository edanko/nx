package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Kind holds the schema definition for the Kind entity.
type Kind struct {
	ent.Schema
}

// Fields of the Kind.
func (Kind) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).
			Default(uuid.New).
			Immutable().
			Unique(),
		field.Text("name").
			Unique().
			NotEmpty(),
		field.Text("description").
			Optional().
			Nillable(),
		field.Enum("status").
			Values("published", "draft").
			Default("draft"),
	}
}

// Edges of the Kind.
func (Kind) Edges() []ent.Edge {
	return nil
}
