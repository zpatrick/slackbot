package slackbot

import (
	"strings"

	"github.com/nlopes/slack"
)

// NewStandardizeTextBehavior creates a behavior that standardizes the text in slack message events.
// Currently, this means converting single and double quotes to ' and ", respectively.
func NewStandardizeTextBehavior() Behavior {
	replacer := strings.NewReplacer(
		"‘", "'",
		"’", "'",
		"“", "\"",
		"”", "\"")

	return func(e slack.RTMEvent) error {
		m, ok := e.Data.(*slack.MessageEvent)
		if !ok {
			return nil
		}

		m.Text = replacer.Replace(m.Text)
		return nil
	}
}
