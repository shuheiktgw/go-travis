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
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

// CreateRequestsOption specifies options for
// CreateRequests request.
type CreateRequestsOption struct {
	// Build configuration (as parsed from .travis.yml)
	Config string `json:"config,omitempty"`
	// Travis-ci status message attached to the request
	Message string `json:"message,omitempty"`
	// Branch requested to be built
	Branch string `json:"branch,omitempty"`
	// Travis token associated with webhook on GitHub (DEPRECATED)
	Token string `json:"token,omitempty"`
}

type getRequestsResponse struct {
	Requests []Request `json:"requests"`
}

// Create endpoints actually returns following form of response.
// It is different from standard nor minimal representation of request.
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

	req, err := rs.client.NewRequest("GET", u, nil, nil)
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

	req, err := rs.client.NewRequest("GET", u, nil, nil)
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
func (rs *RequestsService) CreateByRepoId(ctx context.Context, repoId uint, opt *CreateRequestsOption) (*MinimalRequest, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/requests", repoId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("POST", u, opt, nil)
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
func (rs *RequestsService) CreateByRepoSlug(ctx context.Context, repoSlug string, opt *CreateRequestsOption) (*MinimalRequest, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/requests", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("POST", u, opt, nil)
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
