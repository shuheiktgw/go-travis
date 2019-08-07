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

// EnvVar is a standard representation of a environment variable on Travis CI
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_var#standard-representation
type EnvVar struct {
	// The environment variable id
	Id *string `json:"id,omitempty"`
	// The environment variable name, e.g. FOO
	Name *string `json:"name,omitempty"`
	// The environment variable's value, e.g. bar
	Value *string `json:"value,omitempty"`
	// Whether this environment variable should be publicly visible or not
	Public *bool `json:"public,omitempty"`
	// The env_var's branch.
	Branch *string `json:"branch,omitempty"`
	*Metadata
}

// EnvVarBody specifies options for
// creating and updating env var.
type EnvVarBody struct {
	// The environment variable name, e.g. FOO
	Name string `json:"env_var.name,omitempty"`
	// The environment variable's value, e.g. bar
	Value string `json:"env_var.value,omitempty"`
	// Whether this environment variable should be publicly visible or not
	Public bool `json:"env_var.public"`
	// The env_var's branch.
	Branch string `json:"env_var.branch,omitempty"`
}

// envVarsResponse represents the response of a call
// to the Travis CI env vars endpoint.
type envVarsResponse struct {
	EnvVars []*EnvVar `json:"env_vars"`
}

// FindByRepoId fetches environment variable based on the given repository id and env var id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_var#find
func (es *EnvVarsService) FindByRepoId(ctx context.Context, repoId uint, id string) (*EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/env_var/%s", repoId, id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := es.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var envVar EnvVar
	resp, err := es.client.Do(ctx, req, &envVar)
	if err != nil {
		return nil, resp, err
	}

	return &envVar, resp, err
}

// FindByRepoSlug fetches environment variable based on the given repository slug and env var id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_var#find
func (es *EnvVarsService) FindByRepoSlug(ctx context.Context, repoSlug string, id string) (*EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/env_var/%s", url.QueryEscape(repoSlug), id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := es.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var envVar EnvVar
	resp, err := es.client.Do(ctx, req, &envVar)
	if err != nil {
		return nil, resp, err
	}

	return &envVar, resp, err
}

// ListByRepoId fetches environment variables based on the given repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_vars#for_repository
func (es *EnvVarsService) ListByRepoId(ctx context.Context, repoId uint) ([]*EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/env_vars", repoId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := es.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var er envVarsResponse
	resp, err := es.client.Do(ctx, req, &er)
	if err != nil {
		return nil, resp, err
	}

	return er.EnvVars, resp, err
}

// ListByRepoSlug fetches environment variables based on the given repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_vars#for_repository
func (es *EnvVarsService) ListByRepoSlug(ctx context.Context, repoSlug string) ([]*EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/env_vars", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := es.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var er envVarsResponse
	resp, err := es.client.Do(ctx, req, &er)
	if err != nil {
		return nil, resp, err
	}

	return er.EnvVars, resp, err
}

// CreateByRepoId creates environment variable based on the given repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_vars#create
func (es *EnvVarsService) CreateByRepoId(ctx context.Context, repoId uint, envVar *EnvVarBody) (*EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/env_vars", repoId), nil)
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
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/env_vars", url.QueryEscape(repoSlug)), nil)
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

// UpdateByRepoId updates environment variable based on the given option
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_var#update
func (es *EnvVarsService) UpdateByRepoId(ctx context.Context, repoId uint, id string, envVar *EnvVarBody) (*EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/env_var/%s", repoId, id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := es.client.NewRequest(http.MethodPatch, u, envVar, nil)
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

// UpdateByRepoSlug updates environment variable based on the given option
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_var#update
func (es *EnvVarsService) UpdateByRepoSlug(ctx context.Context, repoSlug string, id string, envVar *EnvVarBody) (*EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/env_var/%s", url.QueryEscape(repoSlug), id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := es.client.NewRequest(http.MethodPatch, u, envVar, nil)
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

// DeleteByRepoId deletes environment variable based on the given repository id and the env var id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_var#delete
func (es *EnvVarsService) DeleteByRepoId(ctx context.Context, repoId uint, id string) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/env_var/%s", repoId, id), nil)
	if err != nil {
		return nil, err
	}

	req, err := es.client.NewRequest(http.MethodDelete, u, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := es.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// DeleteByRepoSlug deletes environment variable based on the given repository slug and the env var id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_var#delete
func (es *EnvVarsService) DeleteByRepoSlug(ctx context.Context, repoSlug string, id string) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/env_var/%s", url.QueryEscape(repoSlug), id), nil)
	if err != nil {
		return nil, err
	}

	req, err := es.client.NewRequest(http.MethodDelete, u, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := es.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
