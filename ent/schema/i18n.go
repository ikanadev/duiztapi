package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// I18n holds the schema definition for the I18n entity.
type I18n struct {
	ent.Schema
}

// Fields of the I18n.
func (I18n) Fields() []ent.Field {
	return []ent.Field{
		field.String("code"),
		field.String("language"),
	}
}

// Edges of the I18n.
func (I18n) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("answer_langs", AnswerLangs.Type),
		edge.To("question_langs", QuestionLangs.Type),
		edge.To("quiz_langs", QuizLangs.Type),
	}
}
