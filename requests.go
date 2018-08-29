package travis

import (
	"context"
	"fmt"
	"net/http"
)

// RequestsService handles communication with the requests
// related methods of the Travis CI API.
type RequestsService struct {
	client *Client
}

// Request represents a Travis CI request.
// They can be used to see if and why a GitHub even has or has not triggered a new build.
type Request struct {
	Id           uint   `json:"id,omitempty"`
	RepositoryId uint   `json:"repository_id,omitempty"`
	CommitId     uint   `json:"commit_id,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	OwnerId      uint   `json:"owner_id"`
	OwnerType    string `json:"owner_type,omitempty"`
	EventType    string `json:"event_type,omitempty"`
	Result       string `json:"result,omitempty"`
	PullRequest  bool   `json:"pull_request,omitempty"`
	Branch       string `json:"branch,omitempty"`
}

type listRequestsResponse struct {
	Requests []Request       `json:"requests"`
	Commits  []MinimalCommit `json:"commits"`
}

type getRequestResponse struct {
	Request Request       `json:"request"`
	Commit  MinimalCommit `json:"commit"`
}

// RequestsListOptions specifies the optional parameters to the
// RequestsService.List method.
//
// You have to either provide RepositoryId or Slug
type RequestsListOptions struct {
	// repository id the requests belong to
	RepositoryId uint `url:"repository_id,omitempty"`

	// repository slug the requests belong to
	Slug string `url:"slug,omitempty"`

	// maximum number of requests to return (cannot be larger than 100)
	Limit uint `url:"limit,omitempty"`

	// list requests before older_than (with older_than being a request id)
	OlderThan uint `url:"older_than,omitempty"`
}

// Get fetches the request with the provided id from the Travis CI API.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#builds
func (rs *RequestsService) Get(ctx context.Context, requestId uint) (*Request, *MinimalCommit, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/requests/%d", requestId), nil)
	if err != nil {
		return nil, nil, nil, err
	}

	req, err := rs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	var reqResp getRequestResponse
	resp, err := rs.client.Do(ctx, req, &reqResp)
	if err != nil {
		return nil, nil, resp, err
	}

	return &reqResp.Request, &reqResp.Commit, resp, err
}

// List requests triggered (or not) by a repository's builds.
//
// Travis CI API docs: http://docs.travis-ci.com/api/#builds
func (rs *RequestsService) ListFromRepository(slug string, opt *RequestsListOptions) ([]Request, []MinimalCommit, *http.Response, error) {
	if opt != nil {
		opt.Slug = slug
	} else {
		opt = &RequestsListOptions{Slug: slug}
	}

	u, err := urlWithOptions("/requests", opt)
	if err != nil {
		return nil, nil, nil, err
	}

	req, err := rs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	var reqResp listRequestsResponse
	resp, err := rs.client.Do(context.TODO(), req, &reqResp)
	if err != nil {
		return nil, nil, resp, err
	}

	return reqResp.Requests, reqResp.Commits, resp, err
}
