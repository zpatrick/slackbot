package slackbot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUserID(t *testing.T) {
	cases := map[string]string{
		"<@uid>":       "uid",
		"<@IF903kvLS>": "IF903kvLS",
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			result, err := ParseUserID(input)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, expected, result)
		})
	}
}

func TestParseUserIDErrors(t *testing.T) {
	cases := map[string]string{
		"empty":               "",
		"not escaped":         "uid",
		"only @ sign":         "@uid",
		"missing @ sign":      "<uid>",
		"missing <":           "@uid>",
		"missing >":           "<@uid",
		"leading whitespace":  " <@uid>",
		"trailing whitespace": "<@uid> ",
	}

	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := ParseUserID(input); err == nil {
				t.Fatal("Error was nil!")
			}
		})
	}
}

func TestEscapeUserID(t *testing.T) {
	cases := map[string]string{
		"":          "<@>",
		"uid":       "<@uid>",
		"IF903kvLS": "<@IF903kvLS>",
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			assert.Equal(t, expected, EscapeUserID(input))
		})
	}
}
