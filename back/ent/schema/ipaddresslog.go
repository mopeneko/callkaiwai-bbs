package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// IPAddressLog holds the schema definition for the IPAddressLog entity.
type IPAddressLog struct {
	ent.Schema
}

// Fields of the IPAddressLog.
func (IPAddressLog) Fields() []ent.Field {
	return []ent.Field{
		field.Text("ip_address").Validate(func(s string) error {
			return validation.Validate(s, is.IP)
		}),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the IPAddressLog.
func (IPAddressLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).
			Ref("ip_address_log").
			Unique().
			Required(),
	}
}
