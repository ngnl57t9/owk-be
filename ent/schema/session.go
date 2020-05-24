package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"time"
)

// Session holds the schema definition for the Profile entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("token"),
		field.Time("expiredAt"),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").UpdateDefault(time.Now),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("session").Unique().Required(),
	}
}
