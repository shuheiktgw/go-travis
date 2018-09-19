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

// JobsOption is query parameters to one can specify
// to find jobs
type JobsOption struct {
	// How many jobs to include in the response
	Limit int `url:"limit,omitempty"`
	// How many jobs to skip before the first entry in the response
	Offset int `url:"offset,omitempty"`
	// Attributes to sort jobs by
	SortBy []string `url:"sort_by,omitempty,brackets"`
	// // Current state of the job
	State []string `url:"state,omitempty,brackets"`
}

type getJobsResponse struct {
	Jobs []Job `json:"jobs"`
}

// FindByBuild fetches jobs based on the provided build id
//
// Travis CI API docs: https://developer.travis-ci.csom/resource/jobs#find
func (js *JobsService) FindByBuild(ctx context.Context, buildId uint) ([]Job, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/build/%d/jobs", buildId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getJobsResponse getJobsResponse
	resp, err := js.client.Do(ctx, req, &getJobsResponse)
	if err != nil {
		return nil, resp, err
	}

	return getJobsResponse.Jobs, resp, err
}

// Find fetches current user's jobs based on the provided options
// As of 2018/9/4, this endpoint returns 500 and does not seem to work correctly
// See jobs_integration_test.go, TestJobsService_Find
//
// Travis CI API docs: https://developer.travis-ci.com/resource/jobs#find
func (js *JobsService) Find(ctx context.Context, opt *JobsOption) ([]Job, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/jobs"), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getJobsResponse getJobsResponse
	resp, err := js.client.Do(ctx, req, &getJobsResponse)
	if err != nil {
		return nil, resp, err
	}

	return getJobsResponse.Jobs, resp, err
}
