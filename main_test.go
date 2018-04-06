package slackbot

import (
	"io/ioutil"

	"github.com/urfave/cli"
)

// NewTestApp creates a new *cli.App with the specified command and reasonable defaults for testing.
func NewTestApp(cmd cli.Command) *cli.App {
	app := cli.NewApp()
	app.Commands = []cli.Command{cmd}
	app.Writer = ioutil.Discard
	app.ErrWriter = ioutil.Discard
	return app
}
