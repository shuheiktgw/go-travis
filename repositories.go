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

// RepositoriesService handles communication with the builds
// related methods of the Travis CI API.
type RepositoriesService struct {
	client *Client
}

// Repository represents a Travis CI repository
//
// Travis CI API docs: https://developer.travis-ci.com/resource/repository#standard-representation
type Repository struct {
	// Value uniquely identifying the repository
	Id *uint `json:"id"`
	// The repository's name on GitHub
	Name *string `json:"name"`
	// Same as {repository.owner.name}/{repository.name}
	Slug *string `json:"slug"`
	// The repository's description from GitHub
	Description *string `json:"description"`
	// The repository's id on GitHub
	GitHubId *uint `json:"github_id"`
	// The main programming language used according to GitHub
	GitHubLanguage *string `json:"github_language"`
	// Whether or not this repository is currently enabled on Travis CI
	Active *bool `json:"active"`
	// Whether or not this repository is private
	Private *bool `json:"private"`
	// GitHub user or organization the repository belongs to
	Owner *Owner `json:"owner"`
	// The default branch on GitHub
	DefaultBranch *Branch `json:"default_branch"`
	// Whether or not this repository is starred
	Starred *bool `json:"starred"`
	// Whether or not this repository is managed by a GitHub App installation
	ManagedByInstallation *bool `json:"managed_by_installation"`
	// Whether or not this repository runs builds on travis-ci.org (may also be null)
	ActiveOnOrg *bool `json:"active_on_org"`
	// The repository's migration_status
	MigrationStatus *string `json:"migration_status"`
	// The repository's allow_migration
	AllowMigration *bool `json:"allow_migration"`
	*Metadata
}

// RepositoriesOption is query parameters to one can specify to find repositories
type RepositoriesOption struct {
	// Filters repositories by whether or not this repository is currently enabled on Travis CI.
	Active bool `url:"active,omitempty"`
	// Filters repositories by whether or not this repository runs builds on travis-ci.org (may also be null).
	ActiveOnOrg bool `url:"active_on_org,omitempty"`
	// Filters repositories by whether or not this repository is managed by a GitHub App installation.
	ManagedByInstallation bool `url:"managed_by_installation,omitempty"`
	// Filters repositories by whether or not this repository is private.
	Private bool `url:"private,omitempty"`
	// Filters repositories by whether or not this repository is starred.
	Starred bool `url:"starred,omitempty"`
	// How many repositories to include in the response
	Limit int `url:"limit,omitempty"`
	// How many repositories to skip before the first entry in the response
	Offset int `url:"offset,omitempty"`
	// Attributes to sort repositories by
	SortBy string `url:"sort_by,omitempty"`
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// RepositoryOption is query parameters to one can specify to find repository
type RepositoryOption struct {
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

type repositoriesResponse struct {
	Repositories []*Repository `json:"repositories"`
}

// List fetches repositories of current user
//
// Travis CI API docs: https://developer.travis-ci.com/resource/repositories#for_current_user
func (rs *RepositoriesService) List(ctx context.Context, opt *RepositoriesOption) ([]*Repository, *http.Response, error) {
	u, err := urlWithOptions("repos", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var rr repositoriesResponse
	resp, err := rs.client.Do(ctx, req, &rr)
	if err != nil {
		return nil, resp, err
	}

	return rr.Repositories, resp, err
}

// ListByOwner fetches repositories base on the provided owner
//
// Travis CI API docs: https://developer.travis-ci.com/resource/repositories#for_owner
func (rs *RepositoriesService) ListByOwner(ctx context.Context, owner string, opt *RepositoriesOption) ([]*Repository, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("owner/%s/repos", owner), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var rr repositoriesResponse
	resp, err := rs.client.Do(ctx, req, &rr)
	if err != nil {
		return nil, resp, err
	}

	return rr.Repositories, resp, err
}

// ListByGitHubId fetches repositories base on the provided GitHub Id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/repositories#for_owner
func (rs *RepositoriesService) ListByGitHubId(ctx context.Context, id uint, opt *RepositoriesOption) ([]*Repository, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("owner/github_id/%d/repos", id), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var rr repositoriesResponse
	resp, err := rs.client.Do(ctx, req, &rr)
	if err != nil {
		return nil, resp, err
	}

	return rr.Repositories, resp, err
}

// Find fetches a repository based on the provided slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/repository#find
func (rs *RepositoriesService) Find(ctx context.Context, slug string, opt *RepositoryOption) (*Repository, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s", url.QueryEscape(slug)), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodGet, u, nil, nil)
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
func (rs *RepositoriesService) Activate(ctx context.Context, slug string) (*Repository, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/activate", url.QueryEscape(slug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodPost, u, nil, nil)
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
func (rs *RepositoriesService) Deactivate(ctx context.Context, slug string) (*Repository, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/deactivate", url.QueryEscape(slug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodPost, u, nil, nil)
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

// Migrate migrates a repository
//
// Travis CI API docs: https://developer.travis-ci.com/resource/repository#migrate
func (rs *RepositoriesService) Migrate(ctx context.Context, slug string) (*Repository, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/migrate", url.QueryEscape(slug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodPost, u, nil, nil)
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
func (rs *RepositoriesService) Star(ctx context.Context, slug string) (*Repository, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/star", url.QueryEscape(slug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodPost, u, nil, nil)
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
func (rs *RepositoriesService) Unstar(ctx context.Context, slug string) (*Repository, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/unstar", url.QueryEscape(slug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodPost, u, nil, nil)
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
