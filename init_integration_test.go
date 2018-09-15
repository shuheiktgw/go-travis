// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
	integrationRepoSlug           = "shuheiktgwtest/go-travis-test"
	integrationRepoId        uint = 20783933
	integrationBuildId       uint = 426024083
	integrationJobId         uint = 426024084
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
