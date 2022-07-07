package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Remnant holds the schema definition for the Remnant entity.
type Remnant struct {
	ent.Schema
}

// Fields of the Remnant.
func (Remnant) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			Unique().
			NotEmpty(),
		field.String("quality").
			NotEmpty(),
		field.String("type").
			NotEmpty(),
		field.Float("length").
			Positive(),
		field.Float("width").
			Positive(),
		field.Float("thickness").
			Positive(),
	}
}

// Edges of the Remnant.
func (Remnant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("nest", Nest.Type).
			Ref("remnant").
			Unique().
			Required(),
		edge.To("remnant_used", Nest.Type),
	}
}

// Mixin of the Remnant.
func (Remnant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		TimeMixin{},
	}
}

// Indexes of the Remnant.
func (Remnant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at", "id"),
	}
}
