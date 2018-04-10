package slackbot

import (
	"fmt"
	"io/ioutil"

	"github.com/nlopes/slack"
	"github.com/urfave/cli"
)

// NewTestApp creates a new *cli.App with the specified command and reasonable defaults for testing.
func NewTestApp(cmd cli.Command) *cli.App {
	app := cli.NewApp()
	app.Commands = []cli.Command{cmd}
	app.Writer = ioutil.Discard
	app.ErrWriter = ioutil.Discard
	return app
}

// NewMesageRTMEvent is a helper function that creates a slack.RTMEvent with the formatted message
func NewMessageRTMEvent(format string, tokens ...interface{}) slack.RTMEvent {
	return slack.RTMEvent{
		Type: "message",
		Data: &slack.MessageEvent{
			Msg: slack.Msg{
				Text: fmt.Sprintf(format, tokens...),
			},
		},
	}
}

// NewMesageChannelRTMEvent is a helper function that creates a slack.RTMEvent with the formatted message
// and the specified channelID.
func NewMessageChannelRTMEvent(channelID string, format string, tokens ...interface{}) slack.RTMEvent {
	return slack.RTMEvent{
		Type: "message",
		Data: &slack.MessageEvent{
			Msg: slack.Msg{
				Channel: channelID,
				Text:    fmt.Sprintf(format, tokens...),
			},
		},
	}
}
