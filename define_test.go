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

func TestDefineCommand(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/words", r.URL.Path)

		query := r.URL.Query()
		assert.Equal(t, "1", query.Get("max"))
		assert.Equal(t, "d", query.Get("md"))
		assert.Equal(t, "ice cream", query.Get("sp"))

		response := []DatamuseResponse{
			{
				Definitions: []string{
					"n\tfrozen dessert containing cream and sugar and flavoring",
				},
				Word: "ice cream",
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

	w := bytes.NewBuffer(nil)
	cmd := NewDefineCommand(server.URL, w)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot define ice cream", " ")); err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, w.String(), "frozen dessert containing cream and sugar and flavoring")
}

func TestDefineCommandUserInputErrors(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal([]DatamuseResponse{})
		if err != nil {
			t.Fatal(err)
		}

		w.Write(b)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	cases := map[string][]string{
		"no args":              strings.Split("slackbot define", " "),
		"definition not found": strings.Split("slackbot define ice cream", " "),
	}

	cmd := NewDefineCommand(server.URL, ioutil.Discard)
	app := NewTestApp(cmd)
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			assert.IsType(t, &UserInputError{}, app.Run(args))
		})
	}
}
