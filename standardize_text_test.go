package slackbot

import (
	"testing"

	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

func TestStandardizeTextBehavior(t *testing.T) {
	cases := map[string]string{
		"foo":   "foo",
		"“foo”": "\"foo\"",
		"‘foo’": "'foo'",
	}

	b := NewStandardizeTextBehavior()
	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			e := NewMessageRTMEvent(input)
			if err := b(e); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, expected, e.Data.(*slack.MessageEvent).Text)
		})
	}
}
