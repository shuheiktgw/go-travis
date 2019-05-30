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

// OwnerService handles communication with the GitHub owner
// related methods of the Travis CI API.
type OwnerService struct {
	client *Client
}

// Owner represents a GitHub Repository
//
// Travis CI API docs: https://developer.travis-ci.com/resource/owner#standard-representation
type Owner struct {
	// Value uniquely identifying the owner
	Id *uint `json:"id"`
	// User or organization login set on GitHub
	Login *string `json:"login"`
	// User or organization name set on GitHub
	Name *string `json:"name"`
	// User or organization id set on GitHub
	GitHubId *uint `json:"github_id"`
	// Link to user or organization avatar (image) set on GitHub
	AvatarUrl *string `json:"avatar_url"`
	// Whether or not the owner has an education account
	Education *bool `json:"education"`
	// Repositories belonging to this account
	Repositories []*Repository `json:"repositories"`
	// Installation belonging to the owner
	Installation *Installation `json:"installation"`
	*Metadata
}

// OwnerOption specifies the optional parameters for owner endpoint
type OwnerOption struct {
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// Find fetches a owner based on the provided login
// Login is user or organization login set on GitHub
//
// Travis CI API docs: https://developer.travis-ci.com/resource/owner#find
func (os *OwnerService) FindByLogin(ctx context.Context, login string, opt *OwnerOption) (*Owner, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("owner/%s", login), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := os.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var owner Owner
	resp, err := os.client.Do(ctx, req, &owner)
	if err != nil {
		return nil, resp, err
	}

	return &owner, resp, err
}

// Find fetches a owner based on the provided GitHub id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/owner#find
func (os *OwnerService) FindByGitHubId(ctx context.Context, githubId uint, opt *OwnerOption) (*Owner, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("owner/github_id/%d", githubId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := os.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var owner Owner
	resp, err := os.client.Do(ctx, req, &owner)
	if err != nil {
		return nil, resp, err
	}

	return &owner, resp, err
}
