// +build integration

package travis

import (
	"os"
)

var (
	integrationClient      *Client
	integrationTravisToken = os.Getenv("TRAVIS_API_AUTH_TOKEN")
	integrationGitHubToken = os.Getenv("TRAVIS_GITHUB_PERSONAL_ACCESS_TOKEN")
	integrationUrl         = defaultBaseURL
	integrationRepo        = "shuheiktgw/go-travis"
)

func init() {
	url := os.Getenv("TRAVIS_API_URL")
	if url != "" {
		integrationUrl = url
	}

	integrationClient = NewClient(integrationUrl, integrationTravisToken)
}
