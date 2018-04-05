package slackbot

import (
	"github.com/nlopes/slack"
)

func Run(incomingEvents chan slack.RTMEvent, behaviors []Behavior, run func(e slack.RTMEvent)) {
	for e := range incomingEvents {
		for _, behavior := range behaviors {
			behavior(e)
		}

		run(e)
	}
}
