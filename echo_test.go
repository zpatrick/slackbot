package slackbot

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {
	cases := map[string]struct {
		Input    []string
		Expected string
	}{
		"one arg": {
			Input:    strings.Split("slackbot echo arg0", " "),
			Expected: "arg0",
		},
		"two args": {
			Input:    strings.Split("slackbot echo arg0 arg1", " "),
			Expected: "arg0 arg1",
		},
		"one flag": {
			Input:    strings.Split("slackbot echo --flag", " "),
			Expected: "--flag",
		},
		"args and flags": {
			Input:    strings.Split("slackbot echo --flag arg0 arg1", " "),
			Expected: "--flag arg0 arg1",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			w := bytes.NewBuffer(nil)
			cmd := NewEchoCommand(w)

			if err := NewTestApp(cmd).Run(c.Input); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, c.Expected, w.String())
		})
	}
}

func TestEchoUserInputErrors(t *testing.T) {
	cases := map[string][]string{
		"no args": strings.Split("slackbot echo", " "),
	}

	app := NewTestApp(NewEchoCommand(ioutil.Discard))
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			assert.IsType(t, &UserInputError{}, app.Run(args))
		})
	}
}
