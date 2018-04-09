package slackbot

import (
	"fmt"

	"github.com/nlopes/slack"
	"github.com/urfave/cli"
)

// NewRepeatCommand returns a cli.Command that sends the last message event sent on the specified slack channel to ch.
// The first message to return true for isCommand will be used.
// It is recommended to pass in a *slack.RTM.IncomingEvents channel as ch.
// WARNING: The client must be authenticated with app credentials!
func NewRepeatCommand(
	client SlackClient,
	channelID string,
	ch chan<- slack.RTMEvent,
	isCommand func(m slack.Message) bool,
	options ...CommandOption,
) cli.Command {
	cmd := cli.Command{
		Name:  "repeat",
		Usage: "repeat the last command sent on this channel",
		Flags: []cli.Flag{
			cli.IntFlag{
				Usage: "the number of messages to look back, between 1 and 1000",
				Value: 100,
			},
		},
		Action: func(c *cli.Context) error {
			params := &slack.GetConversationHistoryParameters{
				ChannelID: channelID,
				Limit:     c.Int("limit"),
			}

			history, err := client.GetConversationHistory(params)
			if err != nil {
				return err
			}

			for _, message := range history.Messages {
				if isCommand(message) {
					m := slack.MessageEvent(message)
					m.Channel = channelID
					e := slack.RTMEvent{
						Type: "message",
						Data: &m,
					}

					ch <- e
					return nil
				}
			}

			return fmt.Errorf("Could not find latest command on this channel")
		},
	}

	for _, option := range options {
		cmd = option(cmd)
	}

	return cmd
}
