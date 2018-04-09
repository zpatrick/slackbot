package slackbot

import (
	"io"
	"net/url"
	"strings"

	"github.com/urfave/cli"
	"github.com/zpatrick/rclient"
)

// DatamuseAPIEndpoint is the endpoint for the Datamuse API
const DatamuseAPIEndpoint = "https://api.datamuse.com"

// DatamuseResponse is the response type for GET /words in the Datamuse API
type DatamuseResponse struct {
	Definitions []string `json:"defs"`
	Word        string   `json:"word"`
}

// NewDefineCommand returns a cli.Command that defines the given word or phrase using the Datamuse API.
func NewDefineCommand(endpoint string, w io.Writer) cli.Command {
	client := rclient.NewRestClient(endpoint)
	return cli.Command{
		Name:      "define",
		Usage:     "define a word or phrase",
		ArgsUsage: "args...",
		Action: func(c *cli.Context) error {
			args := strings.Join(c.Args(), " ")
			if args == "" {
				return NewUserInputErrorf("At least one argument is required")
			}

			query := url.Values{}
			query.Set("sp", args)
			query.Set("max", "1")
			query.Set("md", "d")

			var responses []DatamuseResponse
			if err := client.Get("/words", &responses, rclient.Query(query)); err != nil {
				return err
			}

			if len(responses) == 0 || len(responses[0].Definitions) == 0 {
				return NewUserInputErrorf("Could not find any definitions for *%s*", args)
			}

			// strip the first 2 characters as it always contains the part of spech followed by a tab, e.g.
			// "n\tsome noun definition"
			definition := responses[0].Definitions[0][2:]
			return WriteStringf(w, "*%s*: %s\n", responses[0].Word, definition)
		},
	}
}
