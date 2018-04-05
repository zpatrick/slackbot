package slackbot

import (
	"io"
	"strings"

	"github.com/urfave/cli"
)

func NewEchoCommand(w io.Writer, options ...CommandOption) cli.Command {
	cmd := cli.Command{
		Name:            "echo",
		Usage:           "display the given message",
		ArgsUsage:       "[args...]",
		SkipFlagParsing: true,
		Action: func(c *cli.Context) error {
			text := strings.Join(c.Args(), " ")
			return WriteString(w, text)
		},
	}

	for _, option := range options {
		cmd = option(cmd)
	}

	return cmd
}
