package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.Time("dob").Optional(),
		field.Int("ssn").NonNegative(),
		field.Time("lease_start_at").Default(time.Now()),
		field.Time("lease_end_at").Default(time.Now().AddDate(1, 0, 0)),
		field.Bool("is_lease_holder").Default(true),
		field.Time("created_at").Optional().Default(time.Now()),
		field.Time("updated_at").Optional().Default(time.Now()),
		field.Enum("state").Values("active", "inactive").Default("active"),
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("apartment", Apartment.Type).
			Ref("tenants").
			Unique(),
		edge.To("payments", Payment.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
