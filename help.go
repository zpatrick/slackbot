package slackbot

import (
	"io"

	"github.com/urfave/cli"
)

// NewHelpCommand returns a cli.Command shows more detailed output than urfave/cli's default help command
func NewHelpCommand(w io.Writer, options ...CommandOption) cli.Command {
	cmd := cli.Command{
		Name:      "help",
		Usage:     "show help for the application, command, or subcommand",
		ArgsUsage: "[COMMAND [SUBCOMMAND]]",
		Action: func(c *cli.Context) error {
			// yes this is a total hack, no I'm not sorry about it
			args := append([]string{"slackbot"}, c.Args()...)
			args = append(args, "-h")
			return c.App.Run(args)
		},
	}

	for _, option := range options {
		cmd = option(cmd)
	}

	return cmd
}
