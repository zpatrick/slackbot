package slackbot

import (
	"testing"

	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

func TestExpandSlackbotBehavior(t *testing.T) {
	cases := map[string]struct {
		Event  slack.RTMEvent
		Assert func(t *testing.T, e slack.RTMEvent)
	}{
		"non-message event": {
			Event:  slack.RTMEvent{},
			Assert: func(t *testing.T, e slack.RTMEvent) {},
		},
		"empty message event": {
			Event: NewMessageRTMEvent(""),
			Assert: func(t *testing.T, e slack.RTMEvent) {
				assert.Equal(t, "", e.Data.(*slack.MessageEvent).Text)
			},
		},
		"'!' expanded to 'slackbot '": {
			Event: NewMessageRTMEvent("!echo foo"),
			Assert: func(t *testing.T, e slack.RTMEvent) {
				assert.Equal(t, "slackbot echo foo", e.Data.(*slack.MessageEvent).Text)
			},
		},
		"only first '!' is expanded": {
			Event: NewMessageRTMEvent("!echo foo!"),
			Assert: func(t *testing.T, e slack.RTMEvent) {
				assert.Equal(t, "slackbot echo foo!", e.Data.(*slack.MessageEvent).Text)
			},
		},
	}

	b := NewExpandPromptBehavior("!", "slackbot ")
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if err := b(c.Event); err != nil {
				t.Fatal(err)
			}

			c.Assert(t, c.Event)
		})
	}
}
