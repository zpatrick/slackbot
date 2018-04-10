package slackbot

import (
	"fmt"
	"io"
	"sort"

	"github.com/urfave/cli"
)

// The KeyValStore interface is used to read/write key:val pairs to persistent storage
type KeyValStore interface {
	ReadKeyValues() (kvs map[string]string, err error)
	WriteKeyValues(kvs map[string]string) error
}

// The InMemoryKeyValStore type is an adapter to allow the use of a map[string]string as a KeyValStore.
type InMemoryKeyValStore map[string]string

// ReadKeyValues is used to satisfy the KeyValStore interface.
func (s InMemoryKeyValStore) ReadKeyValues() (map[string]string, error) {
	return s, nil
}

// WriteKeyValues is used to satisfy the KeyValStore interface.
func (s InMemoryKeyValStore) WriteKeyValues(kvs map[string]string) error {
	s = kvs
	return nil
}

// NewKVSCommand creates a command that allows users to add, list, and remove entries in a KeyValStore.
func NewKVSCommand(store KeyValStore, w io.Writer, options ...CommandOption) cli.Command {
	cmd := cli.Command{
		Name:  "kvs",
		Usage: "manage a key-val store",
		Subcommands: []cli.Command{
			{
				Name:  "ls",
				Usage: "list entries in the store",
				Flags: []cli.Flag{
					cli.IntFlag{
						Name:  "limit",
						Value: 50,
						Usage: "The maximum number of entries to display",
					},
					cli.BoolFlag{
						Name:  "ascending",
						Usage: "Show entries in reverse-alphabetical order",
					},
				},
				Action: func(c *cli.Context) error {
					kvs, err := store.ReadKeyValues()
					if err != nil {
						return err
					}

					if len(kvs) == 0 {
						return WriteString(w, "There are currently no entries in the store")
					}

					keys := make(sort.StringSlice, 0, len(kvs))
					for k := range kvs {
						keys = append(keys, k)
					}

					if c.Bool("ascending") {
						sort.Sort(sort.Reverse(keys))
					} else {
						sort.Strings(keys)
					}

					var text string
					for i := 0; i < len(keys) && i < c.Int("limit"); i++ {
						key := keys[i]
						text += fmt.Sprintf("*%s*: %s\n", key, kvs[key])
					}

					return WriteString(w, text)
				},
			},
			{
				Name:      "rm",
				Usage:     "remove an entry from the store",
				ArgsUsage: "KEY",
				Action: func(c *cli.Context) error {
					key := c.Args().Get(0)
					if key == "" {
						return NewUserInputError("Argument KEY is required")
					}

					kvs, err := store.ReadKeyValues()
					if err != nil {
						return err
					}

					if _, ok := kvs[key]; !ok {
						return NewUserInputErrorf("No entry for *%s* exists", key)
					}

					delete(kvs, key)
					if err := store.WriteKeyValues(kvs); err != nil {
						return err
					}

					return WriteStringf(w, "Entry *%s* has been deleted", key)
				},
			},
			{
				Name:      "add",
				Usage:     "add an entry to the store",
				ArgsUsage: "KEY VAL",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "force",
						Usage: "overwrite entry if it already exists",
					},
				},
				Action: func(c *cli.Context) error {
					key := c.Args().Get(0)
					if key == "" {
						return NewUserInputError("Argument KEY is required")
					}

					val := c.Args().Get(1)
					if val == "" {
						return NewUserInputError("Argument VAL is required")
					}

					kvs, err := store.ReadKeyValues()
					if err != nil {
						return err
					}

					if _, ok := kvs[key]; ok && !c.Bool("force") {
						return NewUserInputErrorf("An entry for *%s* already exists", key)
					}

					kvs[key] = val
					if err := store.WriteKeyValues(kvs); err != nil {
						return err
					}

					return WriteStringf(w, "Entry *%s* has been added", key)
				},
			},
		},
	}

	for _, option := range options {
		cmd = option(cmd)
	}

	return cmd
}
