package slackbot

import (
	"context"
	"strings"

	"github.com/nlopes/slack"
)

// NewExpandPromptBehavior creates a behavior that will replace any message's text that starts with
// prompt and replace it with expanded.
// This can be used to make it easier for users to execute commands.
// For example, `NewExpandPromptBehavior("!", "slackbot ")` would convert `!echo foo` to `slackbot echo foo`.
func NewExpandPromptBehavior(prompt, expanded string) Behavior {
	return func(ctx context.Context, e slack.RTMEvent) error {
		m, ok := e.Data.(*slack.MessageEvent)
		if !ok {
			return nil
		}

		if strings.HasPrefix(m.Text, prompt) {
			m.Text = strings.Replace(m.Text, prompt, expanded, 1)
		}

		return nil
	}
}
