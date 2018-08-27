package travis

import (
	"context"
	"fmt"
	"net/http"

	"net/url"

	"github.com/pkg/errors"
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

// MinimalOwner represents a GitHub Owner
//
// https://developer.travis-ci.com/resource/owner#minimal-representation
type MinimalOwner struct {
	Id    uint   `json:"id"`
	Login string `json:"login"`
}

// RepositoryOption specifies the optional parameters for the
// RepositoryService.
type RepositoryOption struct {
	// list of repository ids to fetch, cannot be combined with other parameters
	Id uint `url:"id,omitempty"`

	// filter by slug
	Slug string `url:"slug,omitempty"`
}

// // Identifier returns repository's identifier, either id or slug
func (ro *RepositoryOption) Identifier() (string, error) {
	if ro.Id != 0 {
		return fmt.Sprint(ro.Id), nil
	}

	if ro.Slug != "" {
		return url.QueryEscape(ro.Slug), nil
	}

	return "", errors.New("empty repository option: you need to specify either id or slug")
}

// Find fetches a repository based on the provided option
//
// Travis CI API docs: https://developer.travis-ci.com/resource/repository#find
func (rs *RepositoryService) Find(ctx context.Context, ro *RepositoryOption) (*Repository, *http.Response, error) {
	identifier, err := ro.Identifier()

	if err != nil {
		return nil, nil, err
	}

	u, err := urlWithOptions(fmt.Sprintf("/repo/%s", identifier), nil)
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
func (rs *RepositoryService) Activate(ctx context.Context, ro *RepositoryOption) (*Repository, *http.Response, error) {
	identifier, err := ro.Identifier()

	if err != nil {
		return nil, nil, err
	}

	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/activate", identifier), nil)
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
func (rs *RepositoryService) Deactivate(ctx context.Context, ro *RepositoryOption) (*Repository, *http.Response, error) {
	identifier, err := ro.Identifier()

	if err != nil {
		return nil, nil, err
	}

	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/deactivate", identifier), nil)
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
