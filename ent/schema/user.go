package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StorageKey("user_id"),
		field.String("ID"),
		field.String("user_name"),
		field.String("user_type"),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").UpdateDefault(time.Now()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
