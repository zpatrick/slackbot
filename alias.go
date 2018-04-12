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
func NewAliasBehavior(store KeyValStore) Behavior {
	return func(ctx context.Context, e slack.RTMEvent) error {
		m, ok := e.Data.(*slack.MessageEvent)
		if !ok {
			return nil
		}

		aliases, err := store.ReadKeyValues()
		if err != nil {
			return err
		}

		for name, value := range aliases {
			m.Text = strings.Replace(m.Text, name, value, -1)
		}

		return nil
	}
}

// NewAliasCommand creates a command that allows users to add, list, and remove aliases.
func NewAliasCommand(store KeyValStore, w io.Writer, shouldProcess func(m *slack.MessageEvent) bool, options ...CommandOption) cli.Command {
	cmd := NewKVSCommand(store, w, WithName("alias"), WithUsage("manage aliases"))
	for _, option := range options {
		cmd = option(cmd)
	}

	return cmd
}
