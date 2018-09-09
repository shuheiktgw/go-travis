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
// https://developer.travis-ci.com/resource/branch#standard-representation
type Branch struct {
	Name           string     `json:"name,omitempty"`
	Repository     Repository `json:"repository,omitempty"`
	DefaultBranch  bool       `json:"default_branch,omitempty"`
	ExistsOnGithub bool       `json:"exists_on_github,omitempty"`
	LastBuild      Build      `json:"last_build,omitempty"`
}

// MinimalBranch included when the resource is returned as part of another resource
//
// https://developer.travis-ci.com/resource/branch#minimal-representation
type MinimalBranch struct {
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
