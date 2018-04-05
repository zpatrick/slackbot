package slackbot

import (
	"testing"

	"github.com/nlopes/slack"
)

func TestSlackClientSatiesfiesInterface(t *testing.T) {
	var _ SlackClient = &slack.Client{}
}
