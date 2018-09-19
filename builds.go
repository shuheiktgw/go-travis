// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// BuildsService handles communication with the builds
// related methods of the Travis CI API.
type BuildsService struct {
	client *Client
}

// BuildsOption specifies the optional parameters for builds endpoint
type BuildsOption struct {
	// How many builds to include in the response
	Limit int `url:"limit,omitempty"`
	// How many builds to skip before the first entry in the response
	Offset int `url:"offset,omitempty"`
	// Attributes to sort builds by
	SortBy string `url:"sort_by,omitempty"`
}

// BuildsByRepositoryOption specifies the optional parameters for builds endpoint
type BuildsByRepositoryOption struct {
	// The User or Organization that created the build
	CreatedBy []string `url:"created_by,omitempty,brackets"`
	// Event that triggered the build
	EventType []string `url:"event_type,omitempty,brackets"`
	// State of the previous build (useful to see if state changed)
	PreviousState []string `url:"previous_state,omitempty,brackets"`
	// Current state of the build
	State []string `url:"state,omitempty,brackets"`
	// How many builds to include in the response
	Limit int `url:"limit,omitempty"`
	// How many builds to skip before the first entry in the response
	Offset int `url:"offset,omitempty"`
	// Attributes to sort builds by
	SortBy string `url:"sort_by,omitempty"`
}

type getBuildsResponse struct {
	Builds []Build `json:"builds"`
}

// Find fetches current user's builds based on the provided options
//
// Travis CI API docs: https://developer.travis-ci.com/resource/builds#for_current_user
func (bs *BuildsService) Find(ctx context.Context, opt *BuildsOption) ([]Build, *http.Response, error) {
	u, err := urlWithOptions("/builds", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getBuildsResponse getBuildsResponse
	resp, err := bs.client.Do(ctx, req, &getBuildsResponse)
	if err != nil {
		return nil, resp, err
	}

	return getBuildsResponse.Builds, resp, err
}

// FindByRepoId fetches current user's builds based on the repository id and options
//
// Travis CI API docs: https://developer.travis-ci.com/resource/builds#find
func (bs *BuildsService) FindByRepoId(ctx context.Context, repoId uint, opt *BuildsByRepositoryOption) ([]Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/builds", repoId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getBuildsResponse getBuildsResponse
	resp, err := bs.client.Do(ctx, req, &getBuildsResponse)
	if err != nil {
		return nil, resp, err
	}

	return getBuildsResponse.Builds, resp, err
}

// FindByRepoSlug fetches current user's builds based on the repository slug and options
//
// Travis CI API docs: https://developer.travis-ci.com/resource/builds#find
func (bs *BuildsService) FindByRepoSlug(ctx context.Context, repoSlug string, opt *BuildsByRepositoryOption) ([]Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/builds", url.QueryEscape(repoSlug)), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getBuildsResponse getBuildsResponse
	resp, err := bs.client.Do(ctx, req, &getBuildsResponse)
	if err != nil {
		return nil, resp, err
	}

	return getBuildsResponse.Builds, resp, err
}
