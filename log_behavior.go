package slackbot

import (
	"log"

	"github.com/nlopes/slack"
)

func NewLogBehavior() func(e slack.RTMEvent) {
	return func(e slack.RTMEvent) {
		log.Print(e.Type)
	}
}
