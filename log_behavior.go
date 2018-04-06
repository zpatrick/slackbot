package slackbot

import (
	"context"
	"log"

	"github.com/nlopes/slack"
)

// NewLogBehavior returns a Behavior that logs each event type to the specified logger
func NewLogBehavior(logger *log.Logger) Behavior {
	return func(ctx context.Context, e slack.RTMEvent) error {
		if e.Type != "" {
			logger.Print(e.Type)
		}

		return nil
	}
}
