package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Service holds the schema definition for the Service entity.
type Manager struct {
	ent.Schema
}

// Fields of the Service.
func (Manager) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Time("created_at").Optional().Default(time.Now()),
		field.Time("updated_at").Optional().Default(time.Now()),
	}
}

// Edges of the Service.
func (Manager) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("properties", Property.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
