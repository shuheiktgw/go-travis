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

type JobOption struct {
	Limit  int      `url:"limit,omitempty"`
	Offset int      `url:"offset,omitempty"`
	SortBy []string `url:"sort_by,omitempty,brackets"`
	State  []string `url:"state,omitempty,brackets"`
}

type getJobsResponse struct {
	Jobs []Job `json:"jobs"`
}

// FindByBuild fetches jobs based on the provided build id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/jobs#find
func (js *JobsService) FindByBuild(ctx context.Context, buildId uint) ([]Job, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/build/%d/jobs", buildId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest("GET", u, nil, nil)
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
func (js *JobsService) Find(ctx context.Context, opt *JobOption) ([]Job, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/jobs"), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := js.client.NewRequest("GET", u, nil, nil)
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
