package slackbot

import (
	"fmt"
	"html"
	"io"
	"net/url"
	"sort"
	"strings"

	"github.com/urfave/cli"
	"github.com/zpatrick/rclient"
)

// OpenTDBAPIEndpoint is the endpoint for the OpenTDB API
const OpenTDBAPIEndpoint = "https://opentdb.com/api.php"

// OpenTDBResponse is the response type for GET / in the OpenTDB API
type OpenTDBResponse struct {
	Questions []TriviaQuestion `json:"results"`
}

// TriviaQuestion holds information about a single trivia question
type TriviaQuestion struct {
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

// String returns a string representation of the question and all possible answers
func (t TriviaQuestion) String() string {
	// sort answers in reverse alphabetical order so we display 'True or False?'
	answers := append(t.IncorrectAnswers, t.CorrectAnswer)
	sort.Sort(sort.Reverse(sort.StringSlice(answers)))

	text := fmt.Sprintf("%s\n", t.Question)
	for i, answer := range answers {
		if i == len(answers)-1 {
			text += fmt.Sprintf("or *%s*?", answer)
			break
		}

		text += fmt.Sprintf("*%s*, ", answer)
	}

	return text
}

// The TriviaQuestionsStore interface is used to read/write trivia questions to persistent storage.
// The key in the questions map is a slack channel ID.
type TriviaStore interface {
	ReadQuestions() (questions map[string]TriviaQuestion, err error)
	WriteQuestions(questions map[string]TriviaQuestion) error
}

// The InMemoryTriviaStore type is an adapter to allow the use of a map[string]TriviaQuestion as a TriviaStore.
type InMemoryTriviaStore map[string]TriviaQuestion

// ReadQuestions is used to satisfy the TriviaStore interface.
func (s InMemoryTriviaStore) ReadQuestions() (map[string]TriviaQuestion, error) {
	return s, nil
}

// WriteTriviaQuestions is used to satisfy the TriviaStore interface.
func (s InMemoryTriviaStore) WriteQuestions(questions map[string]TriviaQuestion) error {
	s = questions
	return nil
}

// todo: for commands that require an endpoint like this, show the recommended default

// NewTriviaCommand creates a command that allows users to answer, get, and show trivia questions for the specified channel.
func NewTriviaCommand(store TriviaStore, endpoint string, channelID string, w io.Writer, options ...CommandOption) cli.Command {
	client := rclient.NewRestClient(endpoint)
	cmd := cli.Command{
		Name:  "trivia",
		Usage: "commands related to trivia",
		Subcommands: []cli.Command{
			{
				Name:      "answer",
				Usage:     "answer the current trivia question",
				ArgsUsage: "ANSWER",
				Action: func(c *cli.Context) error {
					answer := strings.Join(c.Args(), " ")
					if answer == "" {
						return NewUserInputError("Argument ANSWER is required")
					}

					questions, err := store.ReadQuestions()
					if err != nil {
						return err
					}

					question, ok := questions[channelID]
					if !ok {
						return NewUserInputError("There is no active trivia question on this channel")
					}

					var text string
					if strings.ToLower(answer) == strings.ToLower(question.CorrectAnswer) {
						text = fmt.Sprintf("*%s* is the correct answer!", answer)
					} else {
						text = fmt.Sprintf("Sorry, *%s* is not the correct answer", answer)
					}

					return WriteString(w, text)
				},
			},
			{
				Name:  "new",
				Usage: "start a new trivia question",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "difficulty",
						Value: "medium",
						Usage: "the dificulty of the question; can be 'easy', 'medium', or 'hard'",
					},
				},
				Action: func(c *cli.Context) error {
					query := url.Values{}
					query.Set("amount", "1")
					query.Set("difficulty", c.String("difficulty"))

					var response OpenTDBResponse
					if err := client.Get("", &response, rclient.Query(query)); err != nil {
						return err
					}

					if len(response.Questions) == 0 {
						return fmt.Errorf("No trivia questions returned by the api!")
					}

					// opentdb returns html-encoded json
					q := response.Questions[0]
					incorrectAnswers := make([]string, len(q.IncorrectAnswers))
					for i, incorrectAnswer := range q.IncorrectAnswers {
						incorrectAnswers[i] = html.UnescapeString(incorrectAnswer)
					}

					question := TriviaQuestion{
						Question:         html.UnescapeString(q.Question),
						CorrectAnswer:    html.UnescapeString(q.CorrectAnswer),
						IncorrectAnswers: incorrectAnswers,
					}

					questions, err := store.ReadQuestions()
					if err != nil {
						return err
					}

					questions[channelID] = question
					if err := store.WriteQuestions(questions); err != nil {
						return err
					}

					return WriteString(w, question.String())
				},
			},
			{
				Name:  "show",
				Usage: "show the current trivia question",
				Action: func(c *cli.Context) error {
					questions, err := store.ReadQuestions()
					if err != nil {
						return err
					}

					question, ok := questions[channelID]
					if !ok {
						return NewUserInputError("There is no active trivia question on this channel")
					}

					return WriteString(w, question.String())
				},
			},
		},
	}

	for _, option := range options {
		cmd = option(cmd)
	}

	return cmd
}
