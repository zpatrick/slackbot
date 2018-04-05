package slackbot

import (
	"context"

	"github.com/nlopes/slack"
)

type Behavior func(ctx context.Context, e slack.RTMEvent) error
