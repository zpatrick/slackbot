package slackbot

import (
	"fmt"

	"github.com/nlopes/slack"
	"github.com/urfave/cli"
)

// NewDeleteCommand retuns a command that deletes the last message sent by the specified bot on the specified channel.
// WARNING: The client must be authenticated with both app and bot credentials!
// It is highly recommended to use a *DualSlackClient for this paramter.
func NewDeleteCommand(client SlackClient, botID, channelID string, options ...CommandOption) cli.Command {
	cmd := cli.Command{
		Name:  "delete",
		Usage: "delete the last message sent by this bot",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "limit",
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

			var lastMessageTimestamp string
			for _, message := range history.Messages {
				if message.User == botID {
					lastMessageTimestamp = message.Timestamp
					break
				}
			}

			if lastMessageTimestamp == "" {
				return fmt.Errorf("Failed to find last message sent by this bot")
			}

			if _, _, err := client.DeleteMessage(channelID, lastMessageTimestamp); err != nil {
				return err
			}

			return nil
		},
	}

	for _, option := range options {
		cmd = option(cmd)
	}

	return cmd
}
