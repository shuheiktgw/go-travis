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

// FindRequestsOption specifies options for
// FindRequests request.
type FindRequestsOption struct {
	// How many requests to include in the response
	Limit int `url:"limit,omitempty"`
	// How many requests to skip before the first entry in the response
	Offset int `url:"offset,omitempty"`
}

// CreateRequestOption specifies options for
// CreateRequest request.
type CreateRequestOption struct {
	// Build configuration (as parsed from .travis.yml)
	Config string `url:"config,omitempty"`
	// Travis-ci status message attached to the request
	Message string `url:"message,omitempty"`
	// Branch requested to be built
	Branch string `url:"branch,omitempty"`
	// Travis token associated with webhook on GitHub (DEPRECATED)
	Token string `url:"token,omitempty"`
}

type getRequestsResponse struct {
	Requests []Request `json:"requests"`
}

// Create endpoints actually returns following form of response.
// It is different from standard nor minimal representation of a request.
// So far, I'm not going to create a special struct to parse it, and
// just use the minimal representation of request.
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
	Request MinimalRequest `json:"request"`
}

// FindByRepoId fetches requests of given repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/requests#find
func (rs *RequestsService) FindByRepoId(ctx context.Context, repoId uint, opt *FindRequestsOption) ([]Request, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/requests", repoId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getRequestsResponse getRequestsResponse
	resp, err := rs.client.Do(ctx, req, &getRequestsResponse)
	if err != nil {
		return nil, resp, err
	}

	return getRequestsResponse.Requests, resp, err
}

// Find fetches requests of given repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/requests#find
func (rs *RequestsService) FindByRepoSlug(ctx context.Context, repoSlug string, opt *FindRequestsOption) ([]Request, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/requests", url.QueryEscape(repoSlug)), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getRequestsResponse getRequestsResponse
	resp, err := rs.client.Do(ctx, req, &getRequestsResponse)
	if err != nil {
		return nil, resp, err
	}

	return getRequestsResponse.Requests, resp, err
}

// CreateByRepoId create requests of given repository id and provided options
//
// Travis CI API docs: https://developer.travis-ci.com/resource/requests#create
func (rs *RequestsService) CreateByRepoId(ctx context.Context, repoId uint, opt *CreateRequestOption) (*MinimalRequest, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/requests", repoId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodPost, u, opt, nil)
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
func (rs *RequestsService) CreateByRepoSlug(ctx context.Context, repoSlug string, opt *CreateRequestOption) (*MinimalRequest, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/requests", url.QueryEscape(repoSlug)), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest(http.MethodPost, u, opt, nil)
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
