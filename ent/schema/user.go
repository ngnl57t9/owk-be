package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MinLen(4).MaxLen(25),
		field.String("username").MinLen(4).MaxLen(25).Unique(),
		field.String("password").StructTag(`json:"-"`),
		field.String("salt").StructTag(`json:"-"`),
		field.Time("updatedAt").Default(time.Now),
		field.Time("createdAt").UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("session", Session.Type).Unique().StructTag(`json:"-"`),
		edge.To("profile", Profile.Type).Unique(),
		edge.From("roles", Role.Type).Ref("users"),
	}
}
