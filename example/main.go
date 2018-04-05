package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/nlopes/slack"
	"github.com/urfave/cli"
	"github.com/zpatrick/slackbot"
)

func main() {
	token := flag.String("token", "", "Bot token for the Slack API")
	flag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	client := slack.New(*token)
	rtm := client.NewRTM()
	defer rtm.Disconnect()

	ctx := context.Background()
	behaviors := []slackbot.Behavior{
		slackbot.NewLogBehavior(),
	}

	go rtm.ManageConnection()
	for e := range rtm.IncomingEvents {
		for _, behavior := range behaviors {
			if err := behavior(ctx, e); err != nil {
				log.Println(err.Error())
			}
		}

		switch data := e.Data.(type) {
		case *slack.ConnectedEvent:
			log.Printf("Slack connection successful!")
		case *slack.MessageEvent:
			text := data.Msg.Text
			if !strings.HasPrefix(strings.ToLower(text), "slackbot ") {
				continue
			}

			w := bytes.NewBuffer(nil)
			app := cli.NewApp()
			app.Name = "slackbot"
			app.Commands = []cli.Command{
				slackbot.NewEchoCommand(w),
			}
			app.Writer = slackbot.WriterFunc(func(b []byte) (n int, err error) {
				return w.Write(b)
			})
			app.CommandNotFound = func(c *cli.Context, command string) {
				text := fmt.Sprintf("Command '%s' does not exist", command)
				w.WriteString(text)
			}

			args := strings.Split(text, " ")
			if err := app.Run(args); err != nil {
				log.Println(err.Error())
				continue
			}

			m := rtm.NewOutgoingMessage(w.String(), data.Msg.Channel)
			rtm.SendMessage(m)
		}
	}
}
