package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Gender int

const (
	GenderMale   Gender = 1
	GenderFemale Gender = 2
	GenderOthers Gender = 3
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			Validate(func(s string) error {
				return validation.Validate(s, validation.Required)
			}),
		field.Int("gender").
			Validate(func(i int) error {
				ii := Gender(i)
				return validation.Validate(ii, validation.Required, validation.In(GenderMale, GenderFemale, GenderOthers))
			}),
		field.Text("introduction").Optional(),
		field.Text("tweet_url").
			Optional().
			Validate(func(s string) error {
				return validation.Validate(s, is.URL)
			}),
		field.Text("tiktok_url").
			Optional().
			Validate(func(s string) error {
				return validation.Validate(s, is.URL)
			}),
		field.Text("contact_url").
			Optional().
			Validate(func(s string) error {
				return validation.Validate(s, is.URL)
			}),
		field.Text("contact_id").Optional(),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ip_address_log", IPAddressLog.Type),
	}
}
