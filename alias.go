package slackbot

import (
	"context"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/nlopes/slack"
	"github.com/urfave/cli"
)

// The AliasStore interface is used to read/write aliases to persistent storage
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

		for name, value := range aliases {
			m.Text = strings.Replace(m.Text, name, value, -1)
		}

		return nil
	}
}

// NewAliasCommand creates a command that allows users to create, list, and remove aliases.
// alias set: set an alias
// alias ls: lists aliases
// alias rm: removes an alias
func NewAliasCommand(store AliasStore, w io.Writer, options ...CommandOption) cli.Command {
	cmd := cli.Command{
		Name:  "alias",
		Usage: "manage aliases",
		Subcommands: []cli.Command{
			{
				Name:  "ls",
				Usage: "list aliases",
				Action: func(c *cli.Context) error {
					aliases, err := store.ReadAliases()
					if err != nil {
						return err
					}

					if len(aliases) == 0 {
						return WriteString(w, "There are currently no aliases")
					}

					keys := make([]string, 0, len(aliases))
					for k := range aliases {
						keys = append(keys, k)
					}

					var text string
					sort.Strings(keys)
					for _, key := range keys {
						text += fmt.Sprintf("*%s*: %s\n", key, aliases[key])
					}

					return WriteString(w, text)
				},
			},
			{
				Name:      "rm",
				Usage:     "remove an alias",
				ArgsUsage: "NAME",
				Action: func(c *cli.Context) error {
					name := c.Args().Get(0)
					if name == "" {
						return NewUserInputError("Argument NAME is required")
					}

					aliases, err := store.ReadAliases()
					if err != nil {
						return err
					}

					if _, ok := aliases[name]; !ok {
						return NewUserInputErrorf("No aliases with the name *%s* exist", name)
					}

					delete(aliases, name)
					if err := store.WriteAliases(aliases); err != nil {
						return err
					}

					return WriteStringf(w, "Alias *%s* has been deleted", name)
				},
			},
			{
				Name:      "set",
				Usage:     "set an alias",
				ArgsUsage: "NAME VALUE",
				Action: func(c *cli.Context) error {
					name := c.Args().Get(0)
					if name == "" {
						return NewUserInputError("Argument NAME is required")
					}

					value := c.Args().Get(1)
					if value == "" {
						return NewUserInputError("Argument VALUE is required")
					}

					aliases, err := store.ReadAliases()
					if err != nil {
						return err
					}

					aliases[name] = value
					if err := store.WriteAliases(aliases); err != nil {
						return err
					}

					return WriteStringf(w, "Alias *%s* has been set", name)
				},
			},
		},
	}

	for _, option := range options {
		cmd = option(cmd)
	}

	return cmd
}
