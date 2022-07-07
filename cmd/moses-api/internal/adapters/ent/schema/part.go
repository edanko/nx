package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Part holds the schema definition for the Part entity.
type Part struct {
	ent.Schema
}

// Fields of the Part.
func (Part) Fields() []ent.Field {
	return []ent.Field{
		field.String("project").
			NotEmpty(),
		field.String("section").
			NotEmpty(),
		field.String("pos").
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
		field.Int64("count").
			Positive(),
		field.Int64("nested"),
	}
}

// Edges of the Part.
func (Part) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("nests", Nest.Type).
			Ref("parts"),
	}
}

// Mixin of the Part.
func (Part) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		TimeMixin{},
	}
}

// Indexes of the Part.
func (Part) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at", "id"),
	}
}
