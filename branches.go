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
type BranchesService struct {
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

// ListBranchesOption specifies the optional parameters for branches endpoint
type ListBranchesOption struct {
	// Whether or not the branch still exists on GitHub
	ExistsOnGithub bool `url:"exists_on_github,omitempty"`
	// How many branches to include in the response
	Limit int `url:"limit,omitempty"`
	// How many branches to skip before the first entry in the response
	Offset int `url:"offset,omitempty"`
	// Attributes to sort branches by
	SortBy string `url:"sort_by,omitempty"`
}

// getBranchesResponse represents the response of a call
// to the Travis CI branches endpoint.
type getBranchesResponse struct {
	Branches []Branch `json:"branches"`
}

// FindByRepoId fetches a branch based on the provided repository id and branch name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/branch#find
func (bs *BranchesService) FindByRepoId(ctx context.Context, repoId uint, branchName string) (*Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/branch/%s", repoId, branchName), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodGet, u, nil, nil)
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

// FindByRepoSlug fetches a branch based on the provided repository slug and branch name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/branch#find
func (bs *BranchesService) FindByRepoSlug(ctx context.Context, repoSlug string, branchName string) (*Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/branch/%s", url.QueryEscape(repoSlug), branchName), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodGet, u, nil, nil)
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

// ListByRepoId fetches the branches of a given repository id.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/branches#find
func (bs *BranchesService) ListByRepoId(ctx context.Context, repoId uint, opt *ListBranchesOption) ([]Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/branches", repoId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getBranchesResponse getBranchesResponse
	resp, err := bs.client.Do(ctx, req, &getBranchesResponse)
	if err != nil {
		return nil, resp, err
	}

	return getBranchesResponse.Branches, resp, err
}

// ListByRepoSlug fetches the branches of a given repository slug.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/branches#find
func (bs *BranchesService) ListByRepoSlug(ctx context.Context, repoSlug string, opt *ListBranchesOption) ([]Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/branches", url.QueryEscape(repoSlug)), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getBranchesResponse getBranchesResponse
	resp, err := bs.client.Do(ctx, req, &getBranchesResponse)
	if err != nil {
		return nil, resp, err
	}

	return getBranchesResponse.Branches, resp, err
}
