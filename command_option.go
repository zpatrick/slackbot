package slackbot

import "github.com/urfave/cli"

type CommandOption func(cmd cli.Command) cli.Command

func WithName(name string) CommandOption {
	return func(cmd cli.Command) cli.Command {
		cmd.Name = name
		return cmd
	}
}
