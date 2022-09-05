package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Apartment holds the schema definition for the Apartment entity.
type Apartment struct {
	ent.Schema
}

// Fields of the Apartment.
func (Apartment) Fields() []ent.Field {
	return []ent.Field{
		field.String("unit_number").NotEmpty(),
		field.Int("charge").NonNegative(),
		field.Time("created_at").Optional().Default(time.Now()),
		field.Time("updated_at").Optional().Default(time.Now()),
	}
}

// Edges of the Apartment.
func (Apartment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("property", Property.Type).
			Ref("apartments").
			Unique(),
		edge.To("tenants", Tenant.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("payments", Payment.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
