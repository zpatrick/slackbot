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

func TestGIFCommandWithDefaults(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/v1/search", r.URL.Path)

		query := r.URL.Query()
		assert.Equal(t, "api_key", query.Get("key"))
		assert.Equal(t, "dogs playing poker", query.Get("q"))
		assert.Equal(t, "minimal", query.Get("mediafilter"))
		assert.Equal(t, "20", query.Get("limit"))
		assert.Equal(t, "strict", query.Get("safesearch"))

		response := TenorSearchResponse{
			GIFs: []TenorGIF{
				{URL: "gif_url"},
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
	cmd := NewGIFCommand(server.URL, "api_key", w)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot gif dogs playing poker", " ")); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "gif_url", w.String())
}

func TestGIFCommandWithExplicitFlag(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		assert.Equal(t, "off", query.Get("safesearch"))

		response := TenorSearchResponse{
			GIFs: make([]TenorGIF, 1),
		}

		b, err := json.Marshal(response)
		if err != nil {
			t.Fatal(err)
		}

		w.Write(b)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	cmd := NewGIFCommand(server.URL, "", ioutil.Discard)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot gif --explicit dogs", " ")); err != nil {
		t.Fatal(err)
	}
}

func TestGIFCommandWithLimitFlag(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		assert.Equal(t, "5", query.Get("limit"))

		response := TenorSearchResponse{
			GIFs: make([]TenorGIF, 1),
		}

		b, err := json.Marshal(response)
		if err != nil {
			t.Fatal(err)
		}

		w.Write(b)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	cmd := NewGIFCommand(server.URL, "", ioutil.Discard)
	if err := NewTestApp(cmd).Run(strings.Split("slackbot gif --limit 5 dogs", " ")); err != nil {
		t.Fatal(err)
	}
}

func TestGIFCommandWithRandomFlag(t *testing.T) {
	results := map[string]int{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := TenorSearchResponse{
			GIFs: []TenorGIF{
				{URL: "one"},
				{URL: "two"},
				{URL: "three"},
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

	for i := 0; i < 1000; i++ {
		w := bytes.NewBuffer(nil)
		cmd := NewGIFCommand(server.URL, "", w)
		if err := NewTestApp(cmd).Run(strings.Split("slackbot gif --random dogs", " ")); err != nil {
			t.Fatal(err)
		}

		results[w.String()]++
	}

	// it seems reasonable to assume each URL will be returned at least once after 1000 iterations
	assert.NotZero(t, results["one"])
	assert.NotZero(t, results["two"])
	assert.NotZero(t, results["three"])
}

func TestGIFCommandUserInputErrors(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(TenorSearchResponse{})
		if err != nil {
			t.Fatal(err)
		}

		w.Write(b)
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	cases := map[string][]string{
		"no args":       strings.Split("slackbot gif", " "),
		"gif not found": strings.Split("slackbot gif dogs", " "),
	}

	app := NewTestApp(NewGIFCommand(server.URL, "", ioutil.Discard))
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			assert.IsType(t, &UserInputError{}, app.Run(args))
		})
	}
}
