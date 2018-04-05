package slackbot

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
	"github.com/urfave/cli"
)

// NewDeleteCommand retuns a command that deletes the last message sent by the specified bot on the specified channel.
// WARNING: The client must be authenticated with both app and bot credentials!
// It is highly recommended you use a *DualSlackClient for this paramter.
func NewDeleteCommand(client SlackClient, botID, channelID string, options ...CommandOption) cli.Command {
	cmd := cli.Command{
		Name:  "delete",
		Usage: "delete the last message sent by the slackbot",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "count",
				Usage: "number of messages to look back, between 1 and 1000",
				Value: 100,
			},
		},
		Action: func(c *cli.Context) error {
			var getHistory func(string, slack.HistoryParameters) (*slack.History, error)
			switch {
			case strings.HasPrefix(channelID, "C"):
				getHistory = client.GetChannelHistory
			case strings.HasPrefix(channelID, "D"):
				getHistory = client.GetIMHistory
			case strings.HasPrefix(channelID, "G"):
				getHistory = client.GetGroupHistory
			default:
				return fmt.Errorf("Cannot find channel type for '%s'", channelID)
			}

			params := slack.NewHistoryParameters()
			params.Count = c.Int("count")

			history, err := getHistory(channelID, params)
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
				return fmt.Errorf("Failed to find last message sent by bot %s", botID)
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
