package travis

import (
	"context"
	"fmt"
	"net/http"
)

// BranchesService handles communication with the branches
// related methods of the Travis CI API.
type BranchesService struct {
	client *Client
}

// Branch represents a Travis CI build
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

// listBranchesResponse represents the response of a call
// to the Travis CI list branches endpoint.
type listBranchesResponse struct {
	Branches []Branch `json:"branches"`
}

// getBranchResponse represents the response of a call
// to the Travis CI get branch endpoint.
type getBranchResponse struct {
	Branch *Branch `json:"branch"`
}

// List the branches of a given repository.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#builds
func (bs *BranchesService) ListFromRepository(ctx context.Context, slug string) ([]Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repos/%v/branches", slug), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var branchesResp listBranchesResponse
	resp, err := bs.client.Do(ctx, req, &branchesResp)
	if err != nil {
		return nil, resp, err
	}

	return branchesResp.Branches, resp, err
}

// Get fetches a branch based on the provided repository slug
// and it's id.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#builds
func (bs *BranchesService) Get(ctx context.Context, repoSlug string, branchId uint) (*Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repos/%v/branches/%d", repoSlug, branchId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var branchResp getBranchResponse
	resp, err := bs.client.Do(ctx, req, &branchResp)
	if err != nil {
		return nil, resp, err
	}

	return branchResp.Branch, resp, err
}

// Get fetches a branch based on the provided repository slug
// and its name.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#builds
func (bs *BranchesService) GetFromSlug(ctx context.Context, repoSlug string, branchSlug string) (*Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repos/%v/branches/%v", repoSlug, branchSlug), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var branchResp getBranchResponse
	resp, err := bs.client.Do(ctx, req, &branchResp)
	if err != nil {
		return nil, resp, err
	}

	return branchResp.Branch, resp, err
}
