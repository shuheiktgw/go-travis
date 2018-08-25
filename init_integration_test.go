// +build integration

package travis

import (
	"os"
)

var (
	integrationClient      *Client
	integrationToken       string = os.Getenv("TRAVIS_API_AUTH_TOKEN")
	integrationGitHubToken string = os.Getenv("TRAVIS_GITHUB_PERSONAL_ACCESS_TOKEN")
	integrationUrl         string = defaultBaseURL
	integrationRepo        string = "shuheiktgw/go-travis"
)

func init() {
	url := os.Getenv("TRAVIS_API_URL")
	if url != "" {
		integrationUrl = url
	}

	integrationClient = NewClient(integrationUrl, integrationToken)
}
