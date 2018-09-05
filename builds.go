package travis

import (
	"context"
	"fmt"
	"net/http"
)

// BuildsService handles communication with the builds
// related methods of the Travis CI API.
type BuildsService struct {
	client *Client
}

// BuildListOptions specifies the optional parameters for builds endpoint
type BuildsOption struct {
	Limit  int      `url:"limit,omitempty"`
	Offset int      `url:"offset,omitempty"`
	SortBy []string `url:"sort_by,omitempty,brackets"`
}

type getBuildsResponse struct {
	Builds []Build `json:"builds"`
}

// Find fetches current user's builds based on the provided options
//
// Travis CI API docs: https://developer.travis-ci.com/resource/builds#for_current_user
func (bs *BuildsService) Find(ctx context.Context, opt *BuildsOption) ([]Build, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/builds"), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getBuildsResponse getBuildsResponse
	resp, err := bs.client.Do(ctx, req, &getBuildsResponse)
	if err != nil {
		return nil, resp, err
	}

	return getBuildsResponse.Builds, resp, err
}
