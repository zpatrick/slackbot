package slackbot

import (
	"io"
	"math/rand"
	"net/url"
	"strconv"
	"strings"

	"github.com/urfave/cli"
	"github.com/zpatrick/rclient"
)

// TenorAPIEndpoint is the endpoint for the Tenor API
const TenorAPIEndpoint = "https://api.tenor.com/"

// TenorSearchResponse is the response type for GET /v1/search in the Tenor API
type TenorSearchResponse struct {
	GIFs []TenorGIF `json:"results"`
}

// TenorGIF holds information about a Gif from Tenor
type TenorGIF struct {
	URL string `json:"itemurl"`
}

// NewGIFCommand returns a cli.Command that displays a gif using the Tenor API.
// Key is your Tenor API key.
func NewGIFCommand(endpoint, key string, w io.Writer) cli.Command {
	client := rclient.NewRestClient(endpoint)
	return cli.Command{
		Name:      "gif",
		Usage:     "display a gif",
		ArgsUsage: "args...",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "explicit",
				Usage: "enable explicit content",
			},
			cli.IntFlag{
				Name:  "limit",
				Value: 1,
				Usage: "limit number of results returned (max size 50)",
			},
			cli.BoolFlag{
				Name:  "random",
				Usage: "return a random gif from the set of results",
			},
		},
		Action: func(c *cli.Context) error {
			args := strings.Join(c.Args(), " ")
			if args == "" {
				return NewUserInputError("At least one argument is required")
			}

			query := url.Values{}
			query.Set("key", key)
			query.Set("q", args)
			query.Set("mediafilter", "minimal")
			query.Set("limit", strconv.Itoa(c.Int("limit")))
			query.Set("safesearch", "strict")
			if c.Bool("explicit") {
				query.Set("safesearch", "off")
			}

			var response TenorSearchResponse
			if err := client.Get("/v1/search", &response, rclient.Query(query)); err != nil {
				return err
			}

			if len(response.GIFs) == 0 {
				return NewUserInputErrorf("No gifs found for *%s*", args)
			}

			url := response.GIFs[0].URL
			if c.Bool("random") {
				i := rand.Intn(len(response.GIFs))
				url = response.GIFs[i].URL
			}

			return WriteString(w, url)
		},
	}
}
