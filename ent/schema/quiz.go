package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Quiz holds the schema definition for the Quiz entity.
type Quiz struct {
	ent.Schema
}

// Fields of the Quiz.
func (Quiz) Fields() []ent.Field {
	return []ent.Field{
		field.String("url_img"),
	}
}

// Edges of the Quiz.
func (Quiz) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("questions", Question.Type),
		edge.To("langs", QuizLangs.Type),
		edge.From("users", User.Type).Ref("quizes"),
	}
}
