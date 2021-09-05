package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("start_at"),
		field.Time("end_at"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return nil
}
