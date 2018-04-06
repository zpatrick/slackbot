package slackbot

import (
	"context"

	"github.com/nlopes/slack"
)

// A behavior executes some functionality whenever a RTMEvent occurs.
type Behavior func(ctx context.Context, e slack.RTMEvent) error
