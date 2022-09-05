package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount").NonNegative(),
		field.Time("date"),
		field.Int("owner_id"),
		field.Time("created_at").Optional().Default(time.Now()),
		field.Time("updated_at").Optional().Default(time.Now()),
		field.Enum("state").Values("processed", "unprocessed").Default("unprocessed"),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("payments").
			Required().
			Field("owner_id").
			Unique(),
		edge.From("apartment", Apartment.Type).
			Ref("payments").
			Unique(),
	}
}
