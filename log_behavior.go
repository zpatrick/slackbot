package slackbot

import (
	"context"
	"log"

	"github.com/nlopes/slack"
)

func NewLogBehavior() Behavior {
	return func(ctx context.Context, e slack.RTMEvent) error {
		log.Print(e.Type)
		return nil
	}
}
