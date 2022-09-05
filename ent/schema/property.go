package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Property holds the schema definition for the Property entity.
type Property struct {
	ent.Schema
}

// Fields of the Property.
func (Property) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("address").NotEmpty(),
		field.Int("units_count").NonNegative(),
		field.Time("created_at").Optional().Default(time.Now()),
		field.Time("updated_at").Optional().Default(time.Now()),
	}
}

// Edges of the Property.
func (Property) Edges() []ent.Edge {
	return []ent.Edge{
		// // Create an inverse-edge called "manager" of type `Manager`
		// 		// and reference it to the "properties" edge (in Manager schema)
		// 		// explicitly using the `Ref` method.
		edge.From("manager", Manager.Type).
			Ref("properties").
			// 			// setting the edge to unique, ensure
			// 			// that a property can have only one owner.
			Unique(),
		// Edges of the Apartment.
		edge.To("apartments", Apartment.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
