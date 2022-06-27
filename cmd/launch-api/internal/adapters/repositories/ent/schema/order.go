package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
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

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
