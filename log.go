package travis

import (
	"context"
	"fmt"
	"net/http"
)

// LogService handles communication with the logs
// related methods of the Travis CI API.
type LogService struct {
	client *Client
}

// Log represents a Travis CI job log
type Log struct {
	// The log's id
	Id uint `json:"id,omitempty"`
	// The content of the log
	Content string `json:"content,omitempty"`
	// The log parts that form the log
	LogParts []LogPart `json:"log_parts,omitempty"`
}

// 	LogPart is parts that form the log
type LogPart struct {
	Content string `json:"content,omitempty"`
	Final   bool   `json:"final,omitempty"`
	Number  uint   `json:"number,omitempty"`
}

// FindByJob fetches a job's log based on it's provided id.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/log#find
func (ls *LogService) FindByJob(ctx context.Context, jobId uint) (*Log, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/job/%d/log", jobId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ls.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var log Log
	resp, err := ls.client.Do(ctx, req, &log)
	if err != nil {
		return nil, resp, err
	}

	return &log, resp, err
}
