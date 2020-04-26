// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

package travis

import (
	"os"
	"strconv"
)

var (
	integrationBuildId       uint
	integrationClient        *Client
	integrationGitHubOwner        = "shuheiktgwtest"
	integrationGitHubOwnerId uint = 41975784
	integrationGitHubToken        = os.Getenv("TRAVIS_GITHUB_PERSONAL_ACCESS_TOKEN")
	integrationJobId         uint = 218491904
	integrationRepoSlug           = "shuheiktgwtest/go-travis-test"
	integrationRepoId        uint = 20783933
	integrationTravisToken        = os.Getenv("TRAVIS_API_AUTH_TOKEN")
	integrationUrl                = ApiComUrl
	integrationUserId        uint = 1362503
)

func init() {
	url := os.Getenv("TRAVIS_API_URL")
	if url != "" {
		integrationUrl = url
	}

	integrationClient = NewClient(integrationUrl, integrationTravisToken)
	integrationBuildId = toUint(os.Getenv("TRAVIS_INTEGRATION_BUILD_ID"))
}

func toUint(s string) uint {
	i ,err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return uint(i)
}
