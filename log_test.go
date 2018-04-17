package slackbot

import (
	"bytes"
	"log"
	"testing"

	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

func TestLogBehavior(t *testing.T) {
	cases := map[string]struct {
		Event    slack.RTMEvent
		Expected string
	}{
		"empty event": {},
		"connected event": {
			Event:    slack.RTMEvent{Type: "connected"},
			Expected: "connected\n",
		},
		"user typing event": {
			Event:    slack.RTMEvent{Type: "user typing"},
			Expected: "user typing\n",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			w := bytes.NewBuffer(nil)
			b := NewLogBehavior(log.New(w, "", 0))
			if err := b(c.Event); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, c.Expected, w.String())
		})
	}
}
