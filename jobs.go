// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"fmt"
	"net/http"
)

// JobsService handles communication with the jobs
// related methods of the Travis CI API.
type JobsService struct {
	client *Client
}

// Job represents a Travis CI job
//
// Travis CI API docs: https://developer.travis-ci.com/resource/job#standard-representation
type Job struct {
	// Value uniquely identifying the job
	Id *uint `json:"id,omitempty"`
	// The job's allow_failure
	AllowFailure *bool `json:"allow_failure,omitempty"`
	// Incremental number for a repository's builds
	Number *string `json:"number,omitempty"`
	// Current state of the job
	State *string `json:"state,omitempty"`
	// When the job started
	StartedAt *string `json:"started_at,omitempty"`
	// When the job finished
	FinishedAt *string `json:"finished_at,omitempty"`
	// The build the job is associated with
	Build *Build `json:"build,omitempty"`
	// Worker queue this job is/was scheduled on
	Queue *string `json:"queue,omitempty"`
	// GitHub repository the job is associated with
	Repository *Repository `json:"repository,omitempty"`
	// The commit the job is associated with
	Commit *Commit `json:"commit,omitempty"`
	// GitHub user or organization the job belongs to
	Owner *Owner `json:"owner,omitempty"`
	// The stages of the job
	Stage *Stage `json:"stage,omitempty"`
	// When the job was created
	CreatedAt *string `json:"created_at,omitempty"`
	// When the job was updated
	UpdatedAt *string `json:"updated_at,omitempty"`
	// Whether or not the job is private
	Private *bool `json:"private,omitempty"`
	// The job's config
	Config *Config `json:"config,omitempty"`
	*Metadata
}

// JobsOption is query parameters to one can specify to find jobs
type JobsOption struct {
	// How many jobs to include in the response
	Limit int `url:"limit,omitempty"`
	// How many jobs to skip before the first entry in the response
	Offset int `url:"offset,omitempty"`
	// Attributes to sort jobs by
	SortBy []string `url:"sort_by,omitempty,comma"`
	// // Current state of the job
	State []string `url:"state,omitempty,comma"`
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// JobOption is query parameters to one can specify to find job
type JobOption struct {
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

type jobsResponse struct {
	Jobs []*Job `json:"jobs"`
}

// jobResponse is only used to parse responses from Restart, Cancel and Debug
type jobResponse struct {
	Job *Job `json:"job,omitempty"`
}

const (
	// JobStatusCreated represents the job state `created`
	JobStatusCreated = "created"
	// JobStatusQueued represents the job state `queued`
	JobStatusQueued = "queued"
	// JobStatusReceived represents the job state `received`
	JobStatusReceived = "received"
	// JobStatusStarted represents the job state `started`
	JobStatusStarted = "started"
	// JobStatusCanceled represents the job state `canceled`
	JobStatusCanceled = "canceled"
	// JobStatusPassed represents the job state `passed`
	JobStatusPassed = "passed"
)

// Find fetches a job based on the provided job id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/job#find
func (js *JobsService) Find(ctx context.Context, id uint, opt *JobOption) (*Job, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("job/%d", id), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var job Job
	resp, err := js.client.Do(ctx, req, &job)
	if err != nil {
		return nil, resp, err
	}

	return &job, resp, err
}

// ListByBuild fetches jobs based on the provided build id
//
// Travis CI API docs: https://developer.travis-ci.csom/resource/jobs#find
func (js *JobsService) ListByBuild(ctx context.Context, buildId uint) ([]*Job, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("build/%d/jobs", buildId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var jr jobsResponse
	resp, err := js.client.Do(ctx, req, &jr)
	if err != nil {
		return nil, resp, err
	}

	return jr.Jobs, resp, err
}

// List fetches current user's jobs based on the provided options
// As of 2018/9/4, this endpoint returns 500 and does not seem to work correctly
// See jobs_integration_test.go, TestJobsService_Find
//
// Travis CI API docs: https://developer.travis-ci.com/resource/jobs#find
func (js *JobsService) List(ctx context.Context, opt *JobsOption) ([]*Job, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("jobs"), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var jr jobsResponse
	resp, err := js.client.Do(ctx, req, &jr)
	if err != nil {
		return nil, resp, err
	}

	return jr.Jobs, resp, err
}

// Cancel cancels a job based on the provided job id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/job#cancel
func (js *JobsService) Cancel(ctx context.Context, id uint) (*Job, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("job/%d/cancel", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest(http.MethodPost, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var jr jobResponse
	resp, err := js.client.Do(ctx, req, &jr)
	if err != nil {
		return nil, resp, err
	}

	return jr.Job, resp, err
}

// Restart restarts a job based on the provided job id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/job#restart
func (js *JobsService) Restart(ctx context.Context, id uint) (*Job, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("job/%d/restart", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest(http.MethodPost, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var jr jobResponse
	resp, err := js.client.Do(ctx, req, &jr)
	if err != nil {
		return nil, resp, err
	}

	return jr.Job, resp, err
}

// Debug restarts a job in debug mode based on the provided job id
// Debug is only available on the travis-ci.com domain, and you need
// to enable the debug feature
//
// Travis CI API docs: https://developer.travis-ci.com/resource/job#debug
func (js *JobsService) Debug(ctx context.Context, id uint) (*Job, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("job/%d/debug", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest(http.MethodPost, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var jr jobResponse
	resp, err := js.client.Do(ctx, req, &jr)
	if err != nil {
		return nil, resp, err
	}

	return jr.Job, resp, err
}
