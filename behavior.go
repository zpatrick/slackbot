package slackbot

import (
	"github.com/nlopes/slack"
)

// A behavior executes some functionality whenever a RTMEvent occurs.
type Behavior func(e slack.RTMEvent) error
