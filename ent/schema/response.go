package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Response holds the schema definition for the Response entity.
type Response struct {
	ent.Schema
}

// Fields of the UserResp.
func (Response) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the UserResp.
func (Response) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("responses").Unique(),
		edge.From("answer", Answer.Type).Ref("responses").Unique(),
	}
}
