// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

package travis

import (
	"context"
	"testing"
)

func TestAuthenticationService_UsingGithubToken(t *testing.T) {
	token, res, err := integrationClient.Authentication.UsingGithubToken(context.TODO(), integrationGitHubToken)

	if err != nil {
		t.Fatalf("UsingGithubToken fails: %s", err)
	}

	if got, want := res.StatusCode, 200; got != want {
		t.Fatalf("UsingGithubToken fails: invalid http response %s", res.Status)
	}

	if token == "" {
		t.Fatalf("UsingGithubToken fails: token is empty")
	}
}
