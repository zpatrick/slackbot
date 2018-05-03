package slackbot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriviaAnswer(t *testing.T) {
	store := InMemoryTriviaStore{
		"channel_id": {
			Question:         "What is the best movie of all time?",
			CorrectAnswer:    "White Chicks",
			IncorrectAnswers: []string{"Casablanca", "Citizen Kane"},
		},
	}

	cases := map[string]struct {
		Args      []string
		IsCorrect bool
	}{
		"lower case correct answer": {
			Args:      strings.Split("slackbot trivia answer white chicks", " "),
			IsCorrect: true,
		},
		"titled correct answer": {
			Args:      strings.Split("slackbot trivia answer White Chicks", " "),
			IsCorrect: true,
		},
		"incorrect answer from options ": {
			Args:      strings.Split("slackbot trivia answer Casablanca", " "),
			IsCorrect: false,
		},
		"incorrect random answer": {
			Args:      strings.Split("slackbot trivia answer Gigli", " "),
			IsCorrect: false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			w := bytes.NewBuffer(nil)
			cmd := NewTriviaCommand(store, "", "channel_id", w)
			if err := NewTestApp(cmd).Run(c.Args); err != nil {
				t.Fatal(err)
			}

			if c.IsCorrect {
				assert.Contains(t, w.String(), "is the correct answer")
			} else {
				assert.Contains(t, w.String(), "is not the correct answer")
			}
		})
	}
}

func TestTriviaAnswerUserInputErrors(t *testing.T) {
	cases := map[string][]string{
		"missing ANSWER arg":     strings.Split("slackbot trivia answer", " "),
		"no question on channel": strings.Split("slackbot trivia answer foo", " "),
	}

	app := NewTestApp(NewTriviaCommand(InMemoryTriviaStore{}, "", "", ioutil.Discard))
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			assert.IsType(t, &UserInputError{}, app.Run(args))
		})
	}
}

func TestTriviaNewWithDefaults(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/", r.URL.Path)

		query := r.URL.Query()
		assert.Equal(t, "1", query.Get("amount"))
		assert.Equal(t, "medium", query.Get("difficulty"))

		// opentdb send escaped html
		response := OpenTDBResponse{
			Questions: []TriviaQuestion{
				{
					Question:         "Which is the world&#39;s &quot;greatest&quot; band of all time?",
					CorrectAnswer:    "Hoobastank",
					IncorrectAnswers: []string{"Smashmouth", "Matchbox Twenty", "My Chemical Romance"},
				},
			},
		}

		b, err := json.Marshal(response)
		if err != nil {
			t.Fatal(err)
		}

		w.Write(b)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	store := InMemoryTriviaStore{}
	w := bytes.NewBuffer(nil)
	cmd := NewTriviaCommand(store, server.URL, "channel_id", w)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot trivia new", " ")); err != nil {
		t.Fatal(err)
	}

	expected := InMemoryTriviaStore{
		"channel_id": {
			Question:         "Which is the world's \"greatest\" band of all time?",
			CorrectAnswer:    "Hoobastank",
			IncorrectAnswers: []string{"Smashmouth", "Matchbox Twenty", "My Chemical Romance"},
		},
	}

	assert.Contains(t, w.String(), expected["channel_id"].String())
	assert.Equal(t, expected, store)
}

func TestTriviaNewWithDifficultyFlag(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		assert.Equal(t, "hard", query.Get("difficulty"))

		response := OpenTDBResponse{
			Questions: make([]TriviaQuestion, 1),
		}

		b, err := json.Marshal(response)
		if err != nil {
			t.Fatal(err)
		}

		w.Write(b)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	cmd := NewTriviaCommand(InMemoryTriviaStore{}, server.URL, "", ioutil.Discard)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot trivia new --difficulty hard", " ")); err != nil {
		t.Fatal(err)
	}
}

func TestTriviaShow(t *testing.T) {
	t.Skip("TODO")
}

func TestTriviaShowUserInputErrors(t *testing.T) {
	t.Skip("TODO")
}
