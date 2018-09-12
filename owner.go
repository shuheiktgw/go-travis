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
// https://developer.travis-ci.com/resource/owner#standard-representation
type Owner struct {
	Id        uint   `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	GitHubId  uint   `json:"github_id"`
	AvatarUrl string `json:"avatar_url"`
	Education bool   `json:"education"`
}

// MinimalOwner represents a minimal GitHub Owner
//
// https://developer.travis-ci.com/resource/owner#minimal-representation
type MinimalOwner struct {
	Id    uint   `json:"id"`
	Login string `json:"login"`
}

// Find fetches a owner based on the provided login
// Login is user or organization login set on GitHub
//
// Travis CI API docs: https://developer.travis-ci.com/resource/owner#find
func (os *OwnerService) FindByLogin(ctx context.Context, login string) (*Owner, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/owner/%s", login), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := os.client.NewRequest("GET", u, nil, nil)
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
func (os *OwnerService) FindByGitHubId(ctx context.Context, githubId uint) (*Owner, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/owner/github_id/%d", githubId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := os.client.NewRequest("GET", u, nil, nil)
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
