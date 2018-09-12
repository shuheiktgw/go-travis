package travis

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// BranchesService handles communication with the branches
// related methods of the Travis CI API.
type BranchesService struct {
	client *Client
}

// BranchesOption specifies the optional parameters for branches endpoint
type BranchesOption struct {
	ExistsOnGithub bool   `url:"exists_on_github,omitempty"`
	Limit          int    `url:"limit,omitempty"`
	Offset         int    `url:"offset,omitempty"`
	SortBy         string `url:"sort_by,omitempty"`
}

// getBranchesResponse represents the response of a call
// to the Travis CI branches endpoint.
type getBranchesResponse struct {
	Branches []Branch `json:"branches"`
}

// FindByRepoId fetches the branches of a given repository id.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/branches#find
func (bs *BranchesService) FindByRepoId(ctx context.Context, repoId uint, opt *BranchesOption) ([]Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/branches", repoId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
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

// FindByRepoSlug fetches the branches of a given repository slug.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/branches#find
func (bs *BranchesService) FindByRepoSlug(ctx context.Context, repoSlug string, opt *BranchesOption) ([]Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/branches", url.QueryEscape(repoSlug)), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
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
