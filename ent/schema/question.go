package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// Question holds the schema definition for the Question entity.
type Question struct {
	ent.Schema
}

// Fields of the Question.
func (Question) Fields() []ent.Field {
	return nil
}

// Edges of the Question.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("answers", Answer.Type),
		edge.To("langs", QuestionLangs.Type),

		edge.From("quiz", Quiz.Type).Ref("questions").Unique(),
		edge.To("correct_answer", Answer.Type).Unique(),
	}
}
