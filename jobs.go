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

type getJobsResponse struct {
	Jobs []Job `json:"jobs"`
}

// FindByBuild fetches a jobs based on the provided build id
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
