package slackbot

import (
	"io"
	"strings"

	"github.com/urfave/cli"
)

// NewEchoCommand returns a cli.Command that writes user input into w.
func NewEchoCommand(w io.Writer, options ...CommandOption) cli.Command {
	cmd := cli.Command{
		Name:            "echo",
		Usage:           "display the given message",
		ArgsUsage:       "[args...]",
		SkipFlagParsing: true,
		Action: func(c *cli.Context) error {
			text := strings.Join(c.Args(), " ")
			if text == "" {
				return NewUserInputErrorf("At least one argument is required")
			}

			return WriteString(w, text)
		},
	}

	for _, option := range options {
		cmd = option(cmd)
	}

	return cmd
}
