// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"fmt"
	"net/http"
)

// BuildsService handles communication with the builds
// related methods of the Travis CI API.
type AuthenticationService struct {
	client *Client
}

// AccessToken is a token to access Travis CI API
type AccessToken string

type accessTokenResponse struct {
	Token AccessToken `json:"access_token"`
}

// UsingGithubToken will generate a Travis CI API authentication
// token and call the UsingTravisToken method with it, leaving your
// client authenticated and ready to use.
func (as *AuthenticationService) UsingGithubToken(ctx context.Context, githubToken string) (AccessToken, *http.Response, error) {
	if githubToken == "" {
		return "", nil, fmt.Errorf("unable to authenticate client; empty github token provided")
	}

	b := map[string]string{"github_token": githubToken}
	h := map[string]string{"Accept": mediaTypeV2}

	req, err := as.client.NewRequest(http.MethodPost, "/auth/github", b, h)
	if err != nil {
		return "", nil, err
	}

	// This is the only place you need to access Travis API v2.1
	// See https://github.com/travis-ci/travis-ci/issues/9273
	// FIXME Use API V3 once compatible API is implemented
	req.Header.Del("Travis-API-Version")

	atr := &accessTokenResponse{}
	resp, err := as.client.Do(ctx, req, atr)
	if err != nil {
		return "", nil, err
	}

	as.UsingTravisToken(string(atr.Token))

	return atr.Token, resp, err
}

// UsingTravisToken will format and write provided
// travisToken in the AuthenticationService client's headers.
func (as *AuthenticationService) UsingTravisToken(travisToken string) error {
	if travisToken == "" {
		return fmt.Errorf("unable to authenticate client; empty travis token provided")
	}

	as.client.Headers["Authorization"] = "token " + travisToken

	return nil
}
