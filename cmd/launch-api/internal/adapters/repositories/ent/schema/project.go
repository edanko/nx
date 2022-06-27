package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
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

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return nil
}
