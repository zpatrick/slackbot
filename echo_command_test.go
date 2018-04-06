package slackbot

import (
	"io/ioutil"
	"strings"
	"testing"
)

/*
func TestEcho(t *testing.T) {
	cases := []struct {
		Name     string
		Input    []string
		Expected string
	}{
		{
			Name:     "one arg",
			Input:    strings.Split("slackbot echo arg0", " "),
			Expected: "arg0",
		},
		{
			Name:     "two args",
			Input:    strings.Split("slackbot echo arg0 arg1", " "),
			Expected: "arg0 arg1",
		},
			{
				Name:     "one flag",
				Input:    strings.Split("slackbot echo --flag", " "),
				Expected: "--flag",
			},
			{
				Name:     "args and flags",
				Input:    strings.Split("slackbot echo --flag arg0 arg1", " "),
				Expected: "--flag arg0 arg1",
			},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			w := bytes.NewBuffer(nil)
			cmd := NewEchoCommand(w)

			if err := NewTestApp(cmd).Run(c.Input); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, c.Expected, w.String())
		})
	}
}
*/

func TestEchoUserInputErrors(t *testing.T) {
	cases := map[string][]string{
		//"no args": strings.Split("slackbot echo", " "),
		"flag": strings.Split("slackbot echo --flag", " "),
	}

	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			cmd := NewEchoCommand(ioutil.Discard)
			if err, ok := NewTestApp(cmd).Run(input).(UserInputError); !ok {
				t.Fatal(err)
			}
		})
	}
}
