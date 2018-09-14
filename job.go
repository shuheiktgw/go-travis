package travis

import (
	"context"
	"fmt"
	"net/http"
)

// JobService handles communication with the job
// related methods of the Travis CI API.
type JobService struct {
	client *Client
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

// Job represents a Travis CI job
//
// Travis CI API docs: https://developer.travis-ci.com/resource/job#standard-representation
type Job struct {
	// Value uniquely identifying the job
	Id uint `json:"id,omitempty"`
	// The job's allow_failure
	AllowFailure bool `json:"allow_failure,omitempty"`
	// Incremental number for a repository's builds
	Number string `json:"number,omitempty"`
	// Current state of the job
	State string `json:"state,omitempty"`
	// When the job started
	StartedAt string `json:"started_at,omitempty"`
	// When the job finished
	FinishedAt string `json:"finished_at,omitempty"`
	// The build the job is associated with
	Build MinimalBuild `json:"build,omitempty"`
	// Worker queue this job is/was scheduled on
	Queue string `json:"queue,omitempty"`
	// GitHub repository the job is associated with
	Repository MinimalRepository `json:"repository,omitempty"`
	// The commit the job is associated with
	Commit MinimalCommit `json:"commit,omitempty"`
	// GitHub user or organization the job belongs to
	Owner MinimalOwner `json:"owner,omitempty"`
	// The stages of the job
	Stage MinimalStage `json:"stage,omitempty"`
	// When the job was created
	CreatedAt string `json:"created_at,omitempty"`
	// When the job was updated
	UpdatedAt string `json:"updated_at,omitempty"`
	// Whether or not the job is private
	Private bool `json:"private,omitempty"`
}

// MinimalJob is a minimal representation of a Travis CI job
//
// Travis CI API docs: https://developer.travis-ci.com/resource/job#standard-representation
type MinimalJob struct {
	// Value uniquely identifying the job
	Id uint `json:"id,omitempty"`
}

// jobResponse is only used to parse responses from Restart, Cancel and Debug
type jobResponse struct {
	Job MinimalJob `json:"job,omitempty"`
}

// Find fetches a job based on the provided job id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/job#find
func (js *JobService) Find(ctx context.Context, id uint) (*Job, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/job/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest("GET", u, nil, nil)
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

// Cancel cancels a job based on the provided job id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/job#cancel
func (js *JobService) Cancel(ctx context.Context, id uint) (*MinimalJob, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/job/%d/cancel", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest("POST", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var jobResponse jobResponse
	resp, err := js.client.Do(ctx, req, &jobResponse)
	if err != nil {
		return nil, resp, err
	}

	return &jobResponse.Job, resp, err
}

// Restart restarts a job based on the provided job id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/job#restart
func (js *JobService) Restart(ctx context.Context, id uint) (*MinimalJob, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/job/%d/restart", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest("POST", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var jobResponse jobResponse
	resp, err := js.client.Do(ctx, req, &jobResponse)
	if err != nil {
		return nil, resp, err
	}

	return &jobResponse.Job, resp, err
}

// Debug restarts a job in debug mode based on the provided job id
// Debug is only available on the travis-ci.com domain, and you need
// to enable the debug feature
//
// Travis CI API docs: https://developer.travis-ci.com/resource/job#debug
func (js *JobService) Debug(ctx context.Context, id uint) (*MinimalJob, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/job/%d/debug", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest("POST", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var jobResponse jobResponse
	resp, err := js.client.Do(ctx, req, &jobResponse)
	if err != nil {
		return nil, resp, err
	}

	return &jobResponse.Job, resp, err
}
