package slackbot

import (
	"context"

	"github.com/nlopes/slack"
)

// todo: is this even necessary?
func Run(
	incomingEvents chan slack.RTMEvent,
	ctx context.Context,
	behaviors []Behavior,
	run func(ctx context.Context, e slack.RTMEvent),
) {
	for e := range incomingEvents {
		for _, behavior := range behaviors {
			behavior(ctx, e)
		}

		run(ctx, e)
	}
}
