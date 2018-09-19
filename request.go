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

// RequestService handles communication with the request
// related methods of the Travis CI API.
type RequestService struct {
	client *Client
}

// Request represents a Travis CI request.
// They can be used to see if and why a GitHub even has or has not triggered a new build.
//
// // Travis CI API docs: https://developer.travis-ci.com/resource/request#standard-representation
type Request struct {
	// Value uniquely identifying the request
	Id uint `json:"id,omitempty"`
	// The state of a request (eg. whether it has been processed or not)
	State string `json:"state,omitempty"`
	// The result of the request (eg. rejected or approved)
	Result string `json:"result,omitempty"`
	// Travis-ci status message attached to the request.
	Message string `json:"message,omitempty"`
	// GitHub user or organization the request belongs to
	Repository MinimalRepository `json:"repository,omitempty"`
	// Name of the branch requested to be built
	BranchName string `json:"branch_name,omitempty"`
	// The commit the request is associated with
	Commit MinimalCommit `json:"commit,omitempty"`
	// The request's builds
	Builds []MinimalBuild `json:"builds,omitempty"`
	// GitHub user or organization the request belongs to
	Owner MinimalOwner `json:"owner,omitempty"`
	// When Travis CI created the request
	CreatedAt string `json:"created_at,omitempty"`
	// Origin of request (push, pull request, api)
	EventType string `json:"event_type,omitempty"`
}

// MinimalRequest is a minimal representation a Travis CI request.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/request#minimal-representation
type MinimalRequest struct {
	// Value uniquely identifying the request
	Id uint `json:"id,omitempty"`
	// The state of a request (eg. whether it has been processed or not)
	State string `json:"state,omitempty"`
	// The result of the request (eg. rejected or approved)
	Result string `json:"result,omitempty"`
	// Travis-ci status message attached to the request.
	Message string `json:"message,omitempty"`
}

// Find fetches request of given repository id and request id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/request#find
func (rs *RequestService) FindByRepoId(ctx context.Context, repoId uint, id uint) (*Request, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/request/%d", repoId, id), nil)
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

// Find fetches request of given repository slug and request id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/request#find
func (rs *RequestService) FindByRepoSlug(ctx context.Context, repoSlug string, id uint) (*Request, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/request/%d", url.QueryEscape(repoSlug), id), nil)
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
