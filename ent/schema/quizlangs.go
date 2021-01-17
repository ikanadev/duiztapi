package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// QuizLangs holds the schema definition for the QuizLangs entity.
type QuizLangs struct {
	ent.Schema
}

// Fields of the QuizLangs.
func (QuizLangs) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Text("description"),
	}
}

// Edges of the QuizLangs.
func (QuizLangs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("quiz", Quiz.Type).Ref("langs").Unique(),
		edge.From("i18n", I18n.Type).Ref("quiz_langs").Unique(),
	}
}
