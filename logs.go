package travis

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

// LogssService handles communication with the logs
// related methods of the Travis CI API.
type LogsService struct {
	client *Client
}

// Log represents a Travis CI job log
type Log struct {
	Id    uint   `json:"id,omitempty"`
	JobId uint   `json:"job_id,omitempty"`
	Type  string `json:"type,omitempty"`
	Body  string `json:"body,omitempty"`
}

// getLogResponse represents the response of a call
// to the Travis CI get log endpoint.
type getLogResponse struct {
	Log Log `json:"log,omitempty"`
}

// Get fetches a log based on the provided id.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#logs
func (ls *LogsService) Get(ctx context.Context, logId uint) (*Log, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/logs/%d", logId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ls.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var logResp getLogResponse
	resp, err := ls.client.Do(ctx, req, &logResp)
	if err != nil {
		return nil, resp, err
	}

	return &logResp.Log, resp, err
}

// Get a job's log based on it's provided id.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#logs
func (ls *LogsService) GetByJob(ctx context.Context, jobId uint) (*Log, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/jobs/%d/log", jobId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ls.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var plainText bytes.Buffer
	resp, err := ls.client.Do(ctx, req, &plainText)
	if err != nil {
		return nil, resp, err
	}

	return &Log{JobId: jobId, Body: plainText.String()}, resp, err
}
