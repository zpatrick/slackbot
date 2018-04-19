package slackbot

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

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

// IsOrdered asserts that the specified objects are in specified order
func AssertContainsInOrder(t *testing.T, input string, expected ...string) {
	for i := 0; i < len(expected); i++ {
		previous := strings.Index(input, expected[i])
		if previous < 0 {
			t.Fatal(fmt.Sprintf("\n%s\ndoes not contain %s", input, previous))
		}

		if i < len(expected)-1 {
			next := strings.Index(input, expected[i+1])
			if next < 0 {
				t.Fatal(fmt.Sprintf("\n%s\ndoes not contain %s", input, next))
			}

			if previous > next {
				t.Fatal(fmt.Sprintf("\n%s\nindex out of order: %s comes after %s", input, previous, next))
			}
		}
	}
}
