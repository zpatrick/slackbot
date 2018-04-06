package slackbot

import (
	"context"
	"io"
	"strings"

	"github.com/nlopes/slack"
	"github.com/urfave/cli"
)

// The AliasStore interface is used to read/write aliases to a persistent storage
type AliasStore interface {
	ReadAliases() (map[string]string, error)
	WriteAliases(map[string]string) error
}

// NewAliasBehavior creates a behavior that will replace messages' text with an alias.
func NewAliasBehavior(store AliasStore) Behavior {
	return func(ctx context.Context, e slack.RTMEvent) error {
		m, ok := e.Data.(*slack.MessageEvent)
		if !ok {
			return nil
		}

		aliases, err := store.ReadAliases()
		if err != nil {
			return err
		}

		for alias, updated := range aliases {
			m.Text = strings.Replace(m.Text, alias, updated, -1)
		}

		return nil
	}
}

// NewAliasCommand creates a command that allows users to create, list, and remove aliases.
// alias set: creates a new alias
// alias ls: lists aliases
// alias rm: removes an alias
func NewAliasCommand(store AliasStore, w io.Writer, options ...CommandOption) cli.Command {
	cmd := cli.Command{
		Name:            "alias",
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
