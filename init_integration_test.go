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
	integrationBuildId       uint = 426024083
	integrationClient        *Client
	integrationGitHubOwner        = "shuheiktgwtest"
	integrationGitHubOwnerId uint = 41975784
	integrationGitHubToken        = os.Getenv("TRAVIS_GITHUB_PERSONAL_ACCESS_TOKEN")
	integrationJobId         uint = 426024084
	integrationRepoSlug           = "shuheiktgwtest/go-travis-test"
	integrationRepoId        uint = 20783933
	integrationTravisToken        = os.Getenv("TRAVIS_API_AUTH_TOKEN")
	integrationUrl                = ApiOrgUrl
	integrationUserId        uint = 1362503
)

func init() {
	url := os.Getenv("TRAVIS_API_URL")
	if url != "" {
		integrationUrl = url
	}

	integrationClient = NewClient(integrationUrl, integrationTravisToken)
}
