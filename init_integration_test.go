// +build integration

package travis

import (
	"os"
)

var (
	integrationClient        *Client
	integrationGitHubToken        = os.Getenv("TRAVIS_GITHUB_PERSONAL_ACCESS_TOKEN")
	integrationGitHubOwner        = "shuheiktgwtest"
	integrationGitHubOwnerId uint = 41975784
	integrationRepo               = "shuheiktgwtest/go-travis-test"
	integrationRepoId        uint = 20783933
	integrationTravisToken        = os.Getenv("TRAVIS_API_AUTH_TOKEN")
	integrationUrl                = defaultBaseURL
)

func init() {
	url := os.Getenv("TRAVIS_API_URL")
	if url != "" {
		integrationUrl = url
	}

	integrationClient = NewClient(integrationUrl, integrationTravisToken)
}
