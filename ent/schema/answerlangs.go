package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// AnswerLangs holds the schema definition for the AnswerLangs entity.
type AnswerLangs struct {
	ent.Schema
}

// Fields of the AnswerLangs.
func (AnswerLangs) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text"),
	}
}

// Edges of the AnswerLangs.
func (AnswerLangs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("answer", Answer.Type).Ref("langs").Unique(),
		edge.From("i18n", I18n.Type).Ref("answer_langs").Unique(),
	}
}
