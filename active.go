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

// activeResponse represents the response of a call
// to the Travis CI active endpoint.
type activeResponse struct {
	Builds []*Build `json:"builds"`
}

// ActiveOption specifies the optional parameters for active endpoint
type ActiveOption struct {
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// FindByOwner fetches active builds based on the owner's name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/active#for_owner
func (as *ActiveService) FindByOwner(ctx context.Context, owner string, opt *ActiveOption) ([]*Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("owner/%s/active", owner), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := as.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var ar activeResponse
	resp, err := as.client.Do(ctx, req, &ar)
	if err != nil {
		return nil, resp, err
	}

	return ar.Builds, resp, err
}

// FindByGitHubId fetches active builds based on the owner's GitHub id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/active#for_owner
func (as *ActiveService) FindByGitHubId(ctx context.Context, githubId uint, opt *ActiveOption) ([]*Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("owner/github_id/%d/active", githubId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := as.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var ar activeResponse
	resp, err := as.client.Do(ctx, req, &ar)
	if err != nil {
		return nil, resp, err
	}

	return ar.Builds, resp, err
}
