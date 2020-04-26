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

// BuildsService handles communication with the builds
// related methods of the Travis CI API.
type BuildsService struct {
	client *Client
}

// Build represents a Travis CI build
//
// Travis CI API docs: https://developer.travis-ci.com/resource/build#standard-representation
type Build struct {
	// Value uniquely identifying the build
	Id *uint `json:"id,omitempty"`
	// Incremental number for a repository's builds
	Number *string `json:"number,omitempty"`
	// Current state of the build
	State *string `json:"state,omitempty"`
	// Wall clock time in seconds
	Duration *int64 `json:"duration,omitempty"`
	// Event that triggered the build
	EventType *string `json:"event_type,omitempty"`
	// State of the previous build (useful to see if state changed)
	PreviousState *string `json:"previous_state,omitempty"`
	// Title of the build's pull request
	PullRequestTitle *string `json:"pull_request_title,omitempty"`
	// Number of the build's pull request
	PullRequestNumber *uint `json:"pull_request_number,omitempty"`
	// When the build started
	StartedAt *string `json:"started_at,omitempty"`
	// When the build finished
	FinishedAt *string `json:"finished_at,omitempty"`
	// The last time the build was updated
	UpdatedAt *string `json:"updated_at,omitempty"`
	// Whether or not the build is private
	Private *bool `json:"private,omitempty"`
	// GitHub repository the build is associated with
	Repository *Repository `json:"repository,omitempty"`
	// The branch the build is associated with
	Branch *Branch `json:"branch,omitempty"`
	// The build's tag
	Tag *Tag `json:"tag,omitempty"`
	// The commit the build is associated with
	Commit *Commit `json:"commit,omitempty"`
	// List of jobs that are part of the build's matrix
	Jobs []*Job `json:"jobs,omitempty"`
	// The stages of the build
	Stages []*Stage `json:"stages,omitempty"`
	// The User or Organization that created the build
	CreatedBy *Owner `json:"created_by,omitempty"`
	*Metadata
}

// BuildsOption specifies the optional parameters for builds endpoint
type BuildsOption struct {
	// How many builds to include in the response
	Limit int `url:"limit,omitempty"`
	// How many builds to skip before the first entry in the response
	Offset int `url:"offset,omitempty"`
	// Attributes to sort builds by
	SortBy string `url:"sort_by,omitempty"`
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// BuildsByRepoOption specifies the optional parameters for builds endpoint
type BuildsByRepoOption struct {
	// Filters builds by name of the git branch
	BranchName []string `url:"branch.name,omitempty,comma"`
	// The User or Organization that created the build
	CreatedBy []string `url:"created_by,omitempty,comma"`
	// Event that triggered the build
	EventType []string `url:"event_type,omitempty,comma"`
	// State of the previous build (useful to see if state changed)
	PreviousState []string `url:"previous_state,omitempty,comma"`
	// Current state of the build
	State []string `url:"state,omitempty,comma"`
	// How many builds to include in the response
	Limit int `url:"limit,omitempty"`
	// How many builds to skip before the first entry in the response
	Offset int `url:"offset,omitempty"`
	// Attributes to sort builds by
	SortBy string `url:"sort_by,omitempty"`
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// BuildOption specifies the optional parameters for build endpoint
type BuildOption struct {
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

type buildsResponse struct {
	Builds []*Build `json:"builds"`
}

// buildResponse is only used to parse responses from Restart, Cancel
type buildResponse struct {
	Build *Build `json:"build,omitempty"`
}

const (
	// BuildStateCreated represents the build state `created`
	BuildStateCreated = "created"
	// BuildStateReceived represents the build state `received`
	BuildStateReceived = "received"
	// BuildStateStarted represents the build state `started`
	BuildStateStarted = "started"
	// BuildStatePassed represents the build state `passed`
	BuildStatePassed = "passed"
	// BuildStateFailed represents the build state `failed`
	BuildStateFailed = "failed"
	// BuildStateErrored represents the build state `errored`
	BuildStateErrored = "errored"
	// BuildStateCanceled represents the build state `canceled`
	BuildStateCanceled = "canceled"
)

const (
	// BuildEventTypePush represents the build event type `push`
	BuildEventTypePush = "push"
	// BuildEventTypePullRequest represents the build event type `pull_request`
	BuildEventTypePullRequest = "pull_request"
)

// Find fetches a build based on the provided build id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/build#find
func (bs *BuildsService) Find(ctx context.Context, id uint, opt *BuildOption) (*Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("build/%d", id), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var build Build
	resp, err := bs.client.Do(ctx, req, &build)
	if err != nil {
		return nil, resp, err
	}

	return &build, resp, err
}

// List fetches current user's builds based on the provided options
//
// Travis CI API docs: https://developer.travis-ci.com/resource/builds#for_current_user
func (bs *BuildsService) List(ctx context.Context, opt *BuildsOption) ([]*Build, *http.Response, error) {
	u, err := urlWithOptions("builds", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var br buildsResponse
	resp, err := bs.client.Do(ctx, req, &br)
	if err != nil {
		return nil, resp, err
	}

	return br.Builds, resp, err
}

// ListByRepoId fetches current user's builds based on the repository id and options
//
// Travis CI API docs: https://developer.travis-ci.com/resource/builds#find
func (bs *BuildsService) ListByRepoId(ctx context.Context, repoId uint, opt *BuildsByRepoOption) ([]*Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/builds", repoId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var br buildsResponse
	resp, err := bs.client.Do(ctx, req, &br)
	if err != nil {
		return nil, resp, err
	}

	return br.Builds, resp, err
}

// ListByRepoSlug fetches current user's builds based on the repository slug and options
//
// Travis CI API docs: https://developer.travis-ci.com/resource/builds#find
func (bs *BuildsService) ListByRepoSlug(ctx context.Context, repoSlug string, opt *BuildsByRepoOption) ([]*Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/builds", url.QueryEscape(repoSlug)), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var br buildsResponse
	resp, err := bs.client.Do(ctx, req, &br)
	if err != nil {
		return nil, resp, err
	}

	return br.Builds, resp, err
}

// Cancel cancels a build based on the provided build id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/build#cancel
func (bs *BuildsService) Cancel(ctx context.Context, id uint) (*Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("build/%d/cancel", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodPost, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var response buildResponse
	resp, err := bs.client.Do(ctx, req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Build, resp, err
}

// Restart restarts a build based on the provided build id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/build#restart
func (bs *BuildsService) Restart(ctx context.Context, id uint) (*Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("build/%d/restart", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodPost, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var br buildResponse
	resp, err := bs.client.Do(ctx, req, &br)
	if err != nil {
		return nil, resp, err
	}

	return br.Build, resp, err
}
