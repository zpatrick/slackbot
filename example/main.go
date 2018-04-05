package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
	"github.com/urfave/cli"
	"github.com/zpatrick/slackbot"
)

func main() {
	token := os.Getenv("SLACK_BOT_TOKEN")
	if token == "" {
		log.Fatal("SLACK_BOT_TOKEN not set!")
	}

	client := slack.New(token)
	rtm := client.NewRTM()
	defer rtm.Disconnect()

	behaviors := []slackbot.Behavior{
		slackbot.NewLogBehavior(),
	}

	run := func(e slack.RTMEvent) {
		event, ok := e.Data.(*slack.MessageEvent)
		if !ok {
			return
		}

		text := event.Msg.Text
		if !strings.HasPrefix(strings.ToLower(text), "slackbot ") {
			return
		}

		// let users know we are processing
		rtm.SendMessage(rtm.NewTypingMessage(event.Msg.Channel))

		w := bytes.NewBuffer(nil)
		app := cli.NewApp()
		app.Name = "slackbot"
		app.CommandNotFound = func(c *cli.Context, command string) {
			text := fmt.Sprintf("Command '%s' does not exist", command)
			w.WriteString(text)
		}
		app.Commands = []cli.Command{
			slackbot.NewEchoCommand(w),
		}

		args := strings.Split(text, " ")
		if err := app.Run(args); err != nil {
			rtm.SendMessage(rtm.NewOutgoingMessage(err.Error(), event.Msg.Channel))
			return
		}

		rtm.SendMessage(rtm.NewOutgoingMessage(w.String(), event.Msg.Channel))
	}

	go rtm.ManageConnection()
	slackbot.Run(rtm.IncomingEvents, behaviors, run)
}
