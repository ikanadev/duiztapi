package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// QuestionLangs holds the schema definition for the QuestionLangs entity.
type QuestionLangs struct {
	ent.Schema
}

// Fields of the QuestionLangs.
func (QuestionLangs) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Text("body"),
		field.Text("explanation"),
	}
}

// Edges of the QuestionLangs.
func (QuestionLangs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("i18n", I18n.Type).Ref("question_langs").Unique(),
		edge.From("question", Question.Type).Ref("langs").Unique(),
	}
}
