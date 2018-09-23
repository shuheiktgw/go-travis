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

// EnvVarsService handles communication with the env vars endpoints
// of Travis CI API
type EnvVarsService struct {
	client *Client
}

// getEnvVarsResponse represents the response of a call
// to the Travis CI env vars endpoint.
type getEnvVarsResponse struct {
	EnvVars []EnvVar `json:"env_vars"`
}

// FindByRepoId fetches environment variables based on the given repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_vars#for_repository
func (es *EnvVarsService) FindByRepoId(ctx context.Context, repoId uint) ([]EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/env_vars", repoId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := es.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getEnvVarsResponse getEnvVarsResponse
	resp, err := es.client.Do(ctx, req, &getEnvVarsResponse)
	if err != nil {
		return nil, resp, err
	}

	return getEnvVarsResponse.EnvVars, resp, err
}

// FindByRepoSlug fetches environment variables based on the given repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_vars#for_repository
func (es *EnvVarsService) FindByRepoSlug(ctx context.Context, repoSlug string) ([]EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/env_vars", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := es.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getEnvVarsResponse getEnvVarsResponse
	resp, err := es.client.Do(ctx, req, &getEnvVarsResponse)
	if err != nil {
		return nil, resp, err
	}

	return getEnvVarsResponse.EnvVars, resp, err
}

// CreateByRepoId creates environment variable based on the given repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_vars#create
func (es *EnvVarsService) CreateByRepoId(ctx context.Context, repoId uint, envVar *EnvVarBody) (*EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/env_vars", repoId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := es.client.NewRequest(http.MethodPost, u, envVar, nil)
	if err != nil {
		return nil, nil, err
	}

	var e EnvVar
	resp, err := es.client.Do(ctx, req, &e)
	if err != nil {
		return nil, resp, err
	}

	return &e, resp, err
}

// CreateByRepoSlug creates environment variable based on the given repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_vars#create
func (es *EnvVarsService) CreateByRepoSlug(ctx context.Context, repoSlug string, envVar *EnvVarBody) (*EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/env_vars", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := es.client.NewRequest(http.MethodPost, u, envVar, nil)
	if err != nil {
		return nil, nil, err
	}

	var e EnvVar
	resp, err := es.client.Do(ctx, req, &e)
	if err != nil {
		return nil, resp, err
	}

	return &e, resp, err
}
