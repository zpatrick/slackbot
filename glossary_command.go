package slackbot

import (
	"io"

	"github.com/urfave/cli"
)

// todo: Use StoreAdapter in quintilesims/slackbot

type GlossaryStore interface {
	ReadGlossary() (map[string]string, error)
	WriteGlossary(map[string]string) error
}

func NewGlossaryCommand(store GlossaryStore, w io.Writer, options ...CommandOption) cli.Command {
	// todo: glossary commands get, set, ls, rm
	cmd := cli.Command{}

	for _, option := range options {
		cmd = option(cmd)
	}

	return cmd
}
