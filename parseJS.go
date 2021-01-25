package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/vmkevv/duiztapi/ent"
	"github.com/vmkevv/duiztapi/ent/i18n"
	"github.com/vmkevv/duiztapi/ent/quiz"
)

type resource struct {
	URL      string
	language string
}

type question struct {
	title         string
	text          string
	options       []string
	correctOption int
	explanation   string
}

var resources []resource = []resource{
	{
		URL:      "https://raw.githubusercontent.com/lydiahallie/javascript-questions/master/README.md",
		language: "en_US",
	},
}

func parseJs() error {
	for _, resource := range resources {
		body, err := getURLBody(resource.URL)
		if err != nil {
			return err
		}
		questions := scanBody(body)
		err = saveQuestions(resource.language, questions)
		if err != nil {
			return err
		}
	}
	return nil
}

func scanBody(body io.Reader) []question {
	var questions []question
	scanner := bufio.NewReader(body)
	var item string
	finded := false
	for {
		line, err := scanner.ReadString('\n')
		if err == io.EOF {
			question := parseItem(item)
			questions = append(questions, question)
			break
		}
		item += line
		if strings.Contains(line, "######") {
			if finded {
				question := parseItem(item)
				questions = append(questions, question)
			}
			finded = true
			item = line
		}
	}
	return questions
}

func getURLBody(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func saveQuestions(lang string, questions []question) error {
	ctx := context.Background()
	client, err := ent.Open("postgres", "postgres://postgres:12345@localhost:5432/duiztdb?sslmode=disable")
	if err != nil {
		return err
	}
	defer client.Close()
	language, err := client.I18n.Query().Where(i18n.CodeEQ(lang)).Only(ctx)
	if err != nil {
		return err
	}
	jsQuiz, err := client.Quiz.Query().Where(quiz.IDEQ(1)).Only(ctx)
	if err != nil {
		return err
	}
	fmt.Println(language)
	fmt.Println(jsQuiz)
	for _, question := range questions {
		dbQuestion, err := client.Question.Create().SetQuiz(jsQuiz).Save(ctx)
		if err != nil {
			return err
		}
		dbQuestionLang, err := client.QuestionLangs.
			Create().
			SetI18n(language).
			SetQuestion(dbQuestion).
			SetTitle(question.title).
			SetBody(question.text).
			SetExplanation(question.explanation).Save(ctx)
		if err != nil {
			return err
		}
		for i, option := range question.options {
			dbAnswer, err := client.Answer.Create().SetQuestion(dbQuestion).Save(ctx)
			if err != nil {
				return err
			}
			_, err = client.AnswerLangs.Create().SetI18n(language).SetAnswer(dbAnswer).SetText(option).Save(ctx)
			if err != nil {
				return err
			}
			if i == question.correctOption {
				dbQuestion, err = dbQuestion.Update().SetCorrectAnswer(dbAnswer).Save(ctx)
				if err != nil {
					return err
				}
			}
		}
		fmt.Println(dbQuestionLang)
	}
	return nil
}

func parseItem(line string) question {
	var q question
	var title string
	dotIndex := strings.IndexRune(line, '.')
	title = line[dotIndex+2 : strings.IndexRune(line, '\n')]
	q.title = title

	var questionText string
	questionInitIndex := strings.Index(line, "```javascript")
	if questionInitIndex != -1 {
		questionEndIndex := strings.Index(line[questionInitIndex+3:], "```")
		questionText = line[questionInitIndex : questionEndIndex+questionInitIndex+6]
	}
	q.text = questionText

	var options []string
	expr := regexp.MustCompile("- [A-Z]:")
	answersIndex := expr.FindAllIndex([]byte(line), len(line))
	for _, iniEnd := range answersIndex {
		option := line[iniEnd[1]+1 : iniEnd[1]+strings.IndexRune(line[iniEnd[1]:], '\n')]
		options = append(options, option)
	}
	q.options = options

	answerIndex := strings.Index(line, "\n#### ")
	answerIndex = answerIndex + strings.IndexRune(line[answerIndex:], ':') + 2
	answer := line[answerIndex : answerIndex+1]
	aIndex := answer[0] - 65
	q.correctOption = int(aIndex)

	explanation := line[answerIndex+2 : answerIndex+2+strings.Index(line[answerIndex+2:], "</p>")]
	q.explanation = explanation
	return q
}
