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

// RepositoryService handles communication with the builds
// related methods of the Travis CI API.
type RepositoryService struct {
	client *Client
}

// Repository represents a Travis CI repository
//
// Travis CI API docs: https://developer.travis-ci.com/resource/repository#standard-representation
type Repository struct {
	// Value uniquely identifying the repository
	Id uint `json:"id"`
	// The repository's name on GitHub
	Name string `json:"name"`
	// Same as {repository.owner.name}/{repository.name}
	Slug string `json:"slug"`
	// The repository's description from GitHub
	Description string `json:"description"`
	// The repository's id on GitHub
	GitHubId uint `json:"github_id"`
	// The main programming language used according to GitHub
	GitHubLanguage uint `json:"github_language"`
	// Whether or not this repository is currently enabled on Travis CI
	Active bool `json:"active"`
	// Whether or not this repository is private
	Private bool `json:"private"`
	// GitHub user or organization the repository belongs to
	Owner MinimalOwner `json:"owner"`
	// The default branch on GitHub
	DefaultBranch MinimalBranch `json:"default_branch"`
	// Whether or not this repository is starred
	Starred bool `json:"starred"`
	// Whether or not this repository is managed by a GitHub App installation
	ManagedByInstallation bool `json:"managed_by_installation"`
	// Whether or not this repository runs builds on travis-ci.org (may also be null)
	ActiveOnOrg bool `json:"active_on_org"`
}

// MinimalRepository is a minimal representation of a Travis CI repository
//
// Travis CI API docs: https://developer.travis-ci.com/resource/repository#minimal-representation
type MinimalRepository struct {
	// Value uniquely identifying the repository
	Id uint `json:"id"`
	// The repository's name on GitHub
	Name string `json:"name"`
	// Same as {repository.owner.name}/{repository.name}
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
