package slackbot

import (
	"io"
	"strings"

	"github.com/urfave/cli"
)

/*
   TODO: Should commands have:
   OnUsageError: func(context *Context, err error, isSubcommand bool) error {
           return NewUsageErrorf(err.Error())
   }

  or just use the default?
*/

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
