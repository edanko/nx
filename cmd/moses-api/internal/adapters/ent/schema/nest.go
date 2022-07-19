package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Nest holds the schema definition for the Nest entity.
type Nest struct {
	ent.Schema
}

// Fields of the Nest.
func (Nest) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			Unique().
			NotEmpty(),
		field.Float("length").
			Positive(),
	}
}

// Edges of the Nest.
func (Nest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parts", Part.Type).
			Required(),
		edge.To("remnant", Remnant.Type).
			Unique(),
		edge.From("remnant_used", Remnant.Type).
			Ref("remnant_used").
			Unique(),
	}
}

// Mixin of the Nest.
func (Nest) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		TimeMixin{},
	}
}

// Indexes of the Order.
func (Nest) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at", "id"),
	}
}
