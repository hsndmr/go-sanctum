package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PersonalAccessToken holds the schema definition for the PersonalAccessToken entity.
type PersonalAccessToken struct {
	ent.Schema
}

// Fields of the PersonalAccessToken.
func (PersonalAccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("user_id"),
		field.String("token").MaxLen(64),
		field.JSON("abilities", []string{}).
			Optional(),
		field.Time("expiration_at"),
		field.Time("last_used_at").
			Optional().
			Nillable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the PersonalAccessToken.
func (PersonalAccessToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
				Ref("personal_access_tokens").
				Unique().
				Required().
				Field("user_id"),
	}
}

func(PersonalAccessToken) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("token"),
	}
}