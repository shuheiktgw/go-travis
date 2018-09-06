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
	ExistOnGithub bool   `url:"exists_on_github,omitempty"`
	Limit         int    `url:"limit,omitempty"`
	Offset        int    `url:"offset,omitempty"`
	SortBy        string `url:"sort_by,omitempty"`
}

// getBranchesResponse represents the response of a call
// to the Travis CI branches endpoint.
type getBranchesResponse struct {
	Branches []Branch `json:"branches"`
}

// FindByRepositoryId fetches the branches of a given repository id.
// As of 2018/9/7 this endpoints returns 404 and does not seem to work correctly
//
// Travis CI API docs: https://developer.travis-ci.com/resource/branches#find
func (bs *BranchesService) FindByRepositoryId(ctx context.Context, repositoryId uint, opt *BranchesOption) ([]Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repos/%d/branches", repositoryId), opt)
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

// FindByRepositorySlug fetches the branches of a given repository slug.
// As of 2018/9/7 this endpoints returns 404 and does not seem to work correctly
//
// Travis CI API docs: https://developer.travis-ci.com/resource/branches#find
func (bs *BranchesService) FindByRepositorySlug(ctx context.Context, repositorySlug string, opt *BranchesOption) ([]Branch, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repos/%s/branches", url.QueryEscape(repositorySlug)), opt)
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
