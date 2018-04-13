package slackbot

import (
	"context"
	"io"
	"strings"

	"github.com/nlopes/slack"
	"github.com/urfave/cli"
)

// todo: long description about how aliases work and an example using them

// NewAliasBehavior creates a behavior that will replace messages' text with an alias.
func NewAliasBehavior(store KeyValStore, shouldProcess func(m *slack.MessageEvent) bool) Behavior {
	return func(ctx context.Context, e slack.RTMEvent) error {
		m, ok := e.Data.(*slack.MessageEvent)
		if !ok {
			return nil
		}

		if !shouldProcess(m) {
			return nil
		}

		// only run aliases on the command portion of the text
		args := strings.Split(m.Text, " ")
		if len(args) < 2 {
			return nil
		}

		aliases, err := store.ReadKeyValues()
		if err != nil {
			return err
		}

		key := args[1]
		val, ok := aliases[key]
		if !ok {
			return nil
		}

		args[1] = val
		m.Text = strings.Join(args, " ")
		return nil
	}
}

// NewAliasCommand creates a command that allows users to add, list, and remove aliases.
func NewAliasCommand(store KeyValStore, w io.Writer, options ...CommandOption) cli.Command {
	cmd := NewKVSCommand(store, w, WithName("alias"), WithUsage("manage aliases"))
	for _, option := range options {
		cmd = option(cmd)
	}

	return cmd
}
