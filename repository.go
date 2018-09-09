package travis

import (
	"context"
	"fmt"
	"net/http"

	"net/url"
)

// RepositoryService handles communication with the builds
// related methods of the Travis CI API.
type RepositoryService struct {
	client *Client
}

// Repository represents a Travis CI repository
//
// https://developer.travis-ci.com/resource/repository#standard-representation
type Repository struct {
	Id                    uint          `json:"id"`
	Name                  string        `json:"name"`
	Slug                  string        `json:"slug"`
	Description           string        `json:"description"`
	GitHubId              uint          `json:"github_id"`
	GitHubLanguage        uint          `json:"github_language"`
	Active                bool          `json:"active"`
	Private               bool          `json:"private"`
	Owner                 MinimalOwner  `json:"owner"`
	DefaultBranch         MinimalBranch `json:"default_branch"`
	Starred               bool          `json:"starred"`
	ManagedByInstallation bool          `json:"managed_by_installation"`
	ActiveOnOrg           bool          `json:"active_on_org"`
}

// MinimalRepository is a minimal representation of a Travis CI repository
//
// https://developer.travis-ci.com/resource/repository#minimal-representation
type MinimalRepository struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Find fetches a repository based on the provided slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/repository#find
func (rs *RepositoryService) Find(ctx context.Context, slug string) (*Repository, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s", url.QueryEscape(slug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var repository Repository
	resp, err := rs.client.Do(ctx, req, &repository)
	if err != nil {
		return nil, resp, err
	}

	return &repository, resp, err
}

// Activate activates Travis CI on the specified repository
//
// Travis CI API docs: https://developer.travis-ci.com/resource/repository#activate
func (rs *RepositoryService) Activate(ctx context.Context, slug string) (*Repository, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/activate", url.QueryEscape(slug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("POST", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var repository Repository
	resp, err := rs.client.Do(ctx, req, &repository)
	if err != nil {
		return nil, resp, err
	}

	return &repository, resp, err
}

// Deactivate deactivates Travis CI on the specified repository
//
// Travis CI API docs: https://developer.travis-ci.com/resource/repository#deactivate
func (rs *RepositoryService) Deactivate(ctx context.Context, slug string) (*Repository, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/deactivate", url.QueryEscape(slug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("POST", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var repository Repository
	resp, err := rs.client.Do(ctx, req, &repository)
	if err != nil {
		return nil, resp, err
	}

	return &repository, resp, err
}

// Star stars a repository based on the currently logged in user
//
// Travis CI API docs: https://developer.travis-ci.com/resource/repository#star
func (rs *RepositoryService) Star(ctx context.Context, slug string) (*Repository, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/star", url.QueryEscape(slug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("POST", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var repository Repository
	resp, err := rs.client.Do(ctx, req, &repository)
	if err != nil {
		return nil, resp, err
	}

	return &repository, resp, err
}

// Unstar unstars a repository based on the currently logged in user
//
// Travis CI API docs: https://developer.travis-ci.com/resource/repository#unstar
func (rs *RepositoryService) Unstar(ctx context.Context, slug string) (*Repository, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/unstar", url.QueryEscape(slug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("POST", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var repository Repository
	resp, err := rs.client.Do(ctx, req, &repository)
	if err != nil {
		return nil, resp, err
	}

	return &repository, resp, err
}
