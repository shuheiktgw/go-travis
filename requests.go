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

// RequestsService handles communication with the requests
// related methods of the Travis CI API.
type RequestsService struct {
	client *Client
}

// Request represents a Travis CI request.
// They can be used to see if and why a GitHub even has or has not triggered a new build.
//
// // Travis CI API docs: https://developer.travis-ci.com/resource/request#standard-representation
type Request struct {
	// Value uniquely identifying the request
	Id *uint `json:"id,omitempty"`
	// The state of a request (eg. whether it has been processed or not)
	State *string `json:"state,omitempty"`
	// The result of the request (eg. rejected or approved)
	Result *string `json:"result,omitempty"`
	// Travis-ci status message attached to the request.
	Message *string `json:"message,omitempty"`
	// GitHub user or organization the request belongs to
	Repository *Repository `json:"repository,omitempty"`
	// Name of the branch requested to be built
	BranchName *string `json:"branch_name,omitempty"`
	// The commit the request is associated with
	Commit *Commit `json:"commit,omitempty"`
	// The request's builds
	Builds []*Build `json:"builds,omitempty"`
	// GitHub user or organization the request belongs to
	Owner *Owner `json:"owner,omitempty"`
	// When Travis CI created the request
	CreatedAt *string `json:"created_at,omitempty"`
	// Origin of request (push, pull request, api)
	EventType *string `json:"event_type,omitempty"`
	// The base commit the request is associated with
	BaseCommit *string `json:"base_commit,omitempty"`
	// The head commit the request is associated with
	HeadCommit *string `json:"head_commit,omitempty"`
	*Metadata
}

// RequestsOption specifies options for
// listing requests.
type RequestsOption struct {
	// How many requests to include in the response
	Limit int `url:"limit,omitempty"`
	// How many requests to skip before the first entry in the response
	Offset int `url:"offset,omitempty"`
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// RequestOption specifies options for
// finding a request.
type RequestOption struct {
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// RequestBody specifies body for
// creating request.
type RequestBody struct {
	// Build configuration (as parsed from .travis.yml)
	Config interface{} `json:"config,omitempty"`
	// Travis-ci status message attached to the request
	Message string `json:"message,omitempty"`
	// Branch requested to be built
	Branch string `json:"branch,omitempty"`
	// Travis token associated with webhook on GitHub (DEPRECATED)
	Token string `json:"token,omitempty"`
}

type requestsResponse struct {
	Requests []*Request `json:"requests"`
}

// FindByRepoId fetches request of given repository id and request id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/request#find
func (rs *RequestsService) FindByRepoId(ctx context.Context, repoId uint, id uint, opt *RequestOption) (*Request, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/request/%d", repoId, id), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var request Request
	resp, err := rs.client.Do(ctx, req, &request)
	if err != nil {
		return nil, resp, err
	}

	return &request, resp, err
}

// FindByRepoSlug fetches request of given repository slug and request id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/request#find
func (rs *RequestsService) FindByRepoSlug(ctx context.Context, repoSlug string, id uint, opt *RequestOption) (*Request, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/request/%d", url.QueryEscape(repoSlug), id), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var request Request
	resp, err := rs.client.Do(ctx, req, &request)
	if err != nil {
		return nil, resp, err
	}

	return &request, resp, err
}

// Create endpoints actually returns following form of response.
// It is different from standard nor minimal representation of a request.
// So far, I'm not going to create a special struct to parse it, and
// just use the standard representation of a request.
//
//{
//  "@type":              "pending",
//  "remaining_requests": 1,
//  "repository":         {
//    "@type":            "repository",
//    "@href":            "/repo/1",
//    "@representation":  "minimal",
//    "id":               1,
//    "name":             "test",
//    "slug":             "owner/repo"
//  },
//  "request":            {
//    "repository":       {
//      "id":             1,
//      "owner_name":     "owner",
//      "name":           "repo"
//    },
//    "user":             {
//      "id":             1
//    },
//    "id":               1,
//    "message":          "Override the commit message: this is an api request",
//    "branch":           "master",
//    "config":           { }
//  },
//  "resource_type":      "request"
//}
type createRequestResponse struct {
	Request Request `json:"request"`
}

// ListByRepoId fetches requests of given repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/requests#find
func (rs *RequestsService) ListByRepoId(ctx context.Context, repoId uint, opt *RequestsOption) ([]*Request, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/requests", repoId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var rr requestsResponse
	resp, err := rs.client.Do(ctx, req, &rr)
	if err != nil {
		return nil, resp, err
	}

	return rr.Requests, resp, err
}

// ListByRepoSlug fetches requests of given repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/requests#find
func (rs *RequestsService) ListByRepoSlug(ctx context.Context, repoSlug string, opt *RequestsOption) ([]*Request, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/requests", url.QueryEscape(repoSlug)), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var rr requestsResponse
	resp, err := rs.client.Do(ctx, req, &rr)
	if err != nil {
		return nil, resp, err
	}

	return rr.Requests, resp, err
}

// CreateByRepoId create requests of given repository id and provided options
//
// Travis CI API docs: https://developer.travis-ci.com/resource/requests#create
func (rs *RequestsService) CreateByRepoId(ctx context.Context, repoId uint, request *RequestBody) (*Request, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/requests", repoId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodPost, u, request, nil)
	if err != nil {
		return nil, nil, err
	}

	var createRequestResponse createRequestResponse
	resp, err := rs.client.Do(ctx, req, &createRequestResponse)
	if err != nil {
		return nil, resp, err
	}

	return &createRequestResponse.Request, resp, err
}

// CreateByRepoSlug create requests of given repository slug and provided options
//
// Travis CI API docs: https://developer.travis-ci.com/resource/requests#create
func (rs *RequestsService) CreateByRepoSlug(ctx context.Context, repoSlug string, request *RequestBody) (*Request, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/requests", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodPost, u, request, nil)
	if err != nil {
		return nil, nil, err
	}

	var createRequestResponse createRequestResponse
	resp, err := rs.client.Do(ctx, req, &createRequestResponse)
	if err != nil {
		return nil, resp, err
	}

	return &createRequestResponse.Request, resp, err
}
