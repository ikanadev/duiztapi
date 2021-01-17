package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// Answer holds the schema definition for the Answer entity.
type Answer struct {
	ent.Schema
}

// Fields of the Answer.
func (Answer) Fields() []ent.Field {
	return nil
}

// Edges of the Answer.
func (Answer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("langs", AnswerLangs.Type),
		edge.To("responses", Response.Type),

		edge.From("question", Question.Type).
			Ref("answers").Unique(),
	}
}
