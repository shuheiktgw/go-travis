// +build integration

package travis

import (
	"fmt"
	"os"
)

var (
	integrationClient *Client
	integrationToken  string
	integrationUrl    string = defaultBaseURL
	integrationRepo   string = "Ableton/go-travis"
)

func init() {
	url := os.Getenv("TRAVIS_API_URL")
	if url != "" {
		integrationUrl = url
	}

	integrationToken := os.Getenv("TRAVIS_API_AUTH_TOKEN")
	if integrationToken == "" {
		fmt.Println(
			"TRAVIS_API_AUTH_TOKEN environment variable not set. ",
			"Unable to authenticate the integration tests client. ",
			"Some tests won't run!",
		)
	}

	integrationClient = NewClient(integrationUrl, integrationToken)
}
