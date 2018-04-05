package slackbot

import "github.com/nlopes/slack"

type Behavior func(e slack.RTMEvent)
