package travis

import (
	"context"
	"fmt"
	"net/http"
)

// BuildService handles communication with the builds
// related methods of the Travis CI API.
type BuildService struct {
	client *Client
}

// Build represents a Travis CI build
//
// https://developer.travis-ci.com/resource/build#standard-representation
type Build struct {
	Id                uint              `json:"id,omitempty"`
	Number            string            `json:"number,omitempty"`
	State             string            `json:"state,omitempty"`
	Duration          uint              `json:"duration,omitempty"`
	EventType         string            `json:"event_type,omitempty"`
	PreviousState     string            `json:"previous_state,omitempty"`
	PullRequestTitle  string            `json:"pull_request_title,omitempty"`
	PullRequestNumber uint              `json:"pull_request_number,omitempty"`
	StartedAt         string            `json:"started_at,omitempty"`
	FinishedAt        string            `json:"finished_at,omitempty"`
	UpdatedAt         string            `json:"updated_at,omitempty"`
	Private           bool              `json:"private,omitempty"`
	Repository        MinimalRepository `json:"repository,omitempty"`
	Branch            MinimalBranch     `json:"branch,omitempty"`
	Tag               string            `json:"tag,omitempty"`
	Commit            MinimalCommit     `json:"commit,omitempty"`
	Jobs              []MinimalJob      `json:"jobs,omitempty"`
	Stages            []MinimalStage    `json:"stages,omitempty"`
	CreatedBy         MinimalOwner      `json:"owner,omitempty"`
}

// MinimalBuild is a minimal representation of a Travis CI build
//
// https://developer.travis-ci.com/resource/build#minimal-representation
type MinimalBuild struct {
	Id                uint   `json:"id,omitempty"`
	Number            string `json:"number,omitempty"`
	State             string `json:"state,omitempty"`
	Duration          uint   `json:"duration,omitempty"`
	EventType         string `json:"event_type,omitempty"`
	PreviousState     string `json:"previous_state,omitempty"`
	PullRequestTitle  string `json:"pull_request_title,omitempty"`
	PullRequestNumber uint   `json:"pull_request_number,omitempty"`
	StartedAt         string `json:"started_at,omitempty"`
	FinishedAt        string `json:"finished_at,omitempty"`
	Private           bool   `json:"private,omitempty"`
}

// buildResponse is only used to parse responses from Restart, Cancel
type buildResponse struct {
	Build MinimalBuild `json:"build,omitempty"`
}

// Find fetches a build based on the provided build id
//
// https://developer.travis-ci.com/resource/build#find
func (bs *BuildService) Find(ctx context.Context, id uint) (*Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/build/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
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

// Cancel cancels a build based on the provided build id
//
// https://developer.travis-ci.com/resource/build#cancel
func (bs *BuildService) Cancel(ctx context.Context, id uint) (*MinimalBuild, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/build/%d/cancel", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("POST", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var response buildResponse
	resp, err := bs.client.Do(ctx, req, &response)
	if err != nil {
		return nil, resp, err
	}

	return &response.Build, resp, err
}

// Restart restarts a build based on the provided build id
//
// https://developer.travis-ci.com/resource/build#restart
func (bs *BuildService) Restart(ctx context.Context, id uint) (*MinimalBuild, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/build/%d/restart", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("POST", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var response buildResponse
	resp, err := bs.client.Do(ctx, req, &response)
	if err != nil {
		return nil, resp, err
	}

	return &response.Build, resp, err
}
