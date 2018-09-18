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

// ActiveService handles communication with the active endpoints
// of Travis CI API
type ActiveService struct {
	client *Client
}

// getActiveResponse represents the response of a call
// to the Travis CI active endpoint.
type getActiveResponse struct {
	Builds []Build `json:"builds"`
}

// FindByOwner fetches active builds based on the owner's name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/active#for_owner
func (as *ActiveService) FindByOwner(ctx context.Context, owner string) ([]Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/owner/%s/active", owner), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := as.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getActiveResponse getActiveResponse
	resp, err := as.client.Do(ctx, req, &getActiveResponse)
	if err != nil {
		return nil, resp, err
	}

	return getActiveResponse.Builds, resp, err
}

// FindByGitHubId fetches active builds based on the owner's GitHub id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/active#for_owner
func (as *ActiveService) FindByGitHubId(ctx context.Context, githubId uint) ([]Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/owner/github_id/%d/active", githubId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := as.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getActiveResponse getActiveResponse
	resp, err := as.client.Do(ctx, req, &getActiveResponse)
	if err != nil {
		return nil, resp, err
	}

	return getActiveResponse.Builds, resp, err
}
