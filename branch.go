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

// BranchesService handles communication with the branch
// related methods of the Travis CI API.
type BranchService struct {
	client *Client
}

// Branch represents a branch of a GitHub repository
//
// Travis CI API docs: https://developer.travis-ci.com/resource/branch#standard-representation
type Branch struct {
	// Name of the git branch
	Name string `json:"name,omitempty"`
	// GitHub Repository
	Repository MinimalRepository `json:"repository,omitempty"`
	// Whether or not this is the repository's default branch
	DefaultBranch bool `json:"default_branch,omitempty"`
	// Whether or not the branch still exists on GitHub
	ExistsOnGithub bool `json:"exists_on_github,omitempty"`
	// Last build on the branch
	LastBuild MinimalBuild `json:"last_build,omitempty"`
}

// MinimalBranch included when the resource is returned as part of another resource
//
// Travis CI API docs: https://developer.travis-ci.com/resource/branch#minimal-representation
type MinimalBranch struct {
	// Name of the git branch
	Name string `json:"name,omitempty"`
}

// Find fetches a branch based on the provided repository id and branch name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/branch#find
func (bs *BranchService) FindByRepoId(ctx context.Context, repoId uint, branchName string) (*Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/branch/%s", repoId, branchName), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var branch Branch
	resp, err := bs.client.Do(ctx, req, &branch)
	if err != nil {
		return nil, resp, err
	}

	return &branch, resp, err
}

// Find fetches a branch based on the provided repository slug and branch name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/branch#find
func (bs *BranchService) FindByRepoSlug(ctx context.Context, repoSlug string, branchName string) (*Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/branch/%s", url.QueryEscape(repoSlug), branchName), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var branch Branch
	resp, err := bs.client.Do(ctx, req, &branch)
	if err != nil {
		return nil, resp, err
	}

	return &branch, resp, err
}
