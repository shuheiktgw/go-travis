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
	JobStatusCreated  = "created"
	JobStatusQueued   = "queued"
	JobStatusReceived = "received"
	JobStatusStarted  = "started"
	JobStatusCanceled = "canceled"
	JobStatusPassed   = "passed"
)

// Job represents a Travis CI job
//
// https://developer.travis-ci.com/resource/job#standard-representation
type Job struct {
	Id           uint              `json:"id,omitempty"`
	AllowFailure bool              `json:"allow_failure,omitempty"`
	Number       string            `json:"number,omitempty"`
	State        string            `json:"state,omitempty"`
	StartedAt    string            `json:"started_at,omitempty"`
	FinishedAt   string            `json:"finished_at,omitempty"`
	Build        MinimalBuild      `json:"build,omitempty"`
	Queue        string            `json:"queue,omitempty"`
	Repository   MinimalRepository `json:"repository,omitempty"`
	Commit       MinimalCommit     `json:"commit,omitempty"`
	Owner        MinimalOwner      `json:"owner,omitempty"`
	Stage        MinimalStage      `json:"stage,omitempty"`
	CreatedAt    string            `json:"created_at,omitempty"`
	UpdatedAt    string            `json:"updated_at,omitempty"`
	Private      bool              `json:"private,omitempty"`
}

// MinimalJob is a minimal representation of a Travis CI job
//
// https://developer.travis-ci.com/resource/job#standard-representation
type MinimalJob struct {
	Id uint `json:"id,omitempty"`
}

// jobResponse is only used to parse responses from Restart, Cancel and Debug
type jobResponse struct {
	Job MinimalJob `json:"job,omitempty"`
}

// Find fetches a job based on the provided job id
//
// https://developer.travis-ci.com/resource/job#find
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
// https://developer.travis-ci.com/resource/job#cancel
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
// https://developer.travis-ci.com/resource/job#restart
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
// https://developer.travis-ci.com/resource/job#debug
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
