package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Applicant holds the schema definition for the Applicant entity.
type Applicant struct {
	ent.Schema
}

// Fields of the Applicant.
func (Applicant) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			Unique().
			NotEmpty(),
	}
}

// Edges of the Applicant.
func (Applicant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("launches", Launch.Type),
	}
}

// Mixin of the Order.
func (Applicant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		TimeMixin{},
	}
}

// Indexes of the Order.
func (Applicant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at", "id"),
	}
}
