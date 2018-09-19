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

// EnvVarService handles communication with the env var endpoints
// of Travis CI API
type EnvVarService struct {
	client *Client
}

// EnvVar is a standard representation of a environment variable on Travis CI
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_var#standard-representation
type EnvVar struct {
	// The environment variable id
	Id string `json:"id,omitempty"`
	// The environment variable name, e.g. FOO
	Name string `json:"name,omitempty"`
	// The environment variable's value, e.g. bar
	Value string `json:"value,omitempty"`
	// Whether this environment variable should be publicly visible or not
	Public bool `json:"public,omitempty"`
}

// EnvVarOption specifies options for
// creating and updating env var.
type EnvVarOption struct {
	// The environment variable name, e.g. FOO
	Name string `url:"env_var.name,omitempty"`
	// The environment variable's value, e.g. bar
	Value string `url:"env_var.value,omitempty"`
	// Whether this environment variable should be publicly visible or not
	Public bool `url:"env_var.public"`
}

// FindByRepoId fetches environment variable based on the given repository id and env var id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_var#find
func (es *EnvVarService) FindByRepoId(ctx context.Context, repoId uint, id string) (*EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/env_var/%s", repoId, id), nil)
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
func (es *EnvVarService) FindByRepoSlug(ctx context.Context, repoSlug string, id string) (*EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/env_var/%s", url.QueryEscape(repoSlug), id), nil)
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

// UpdateByRepoId updates environment variable based on the given option
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_var#update
func (es *EnvVarService) UpdateByRepoId(ctx context.Context, repoId uint, id string, opt *EnvVarOption) (*EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/env_var/%s", repoId, id), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := es.client.NewRequest(http.MethodPatch, u, nil, nil)
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

// UpdateByRepoSlug updates environment variable based on the given option
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_var#update
func (es *EnvVarService) UpdateByRepoSlug(ctx context.Context, repoSlug string, id string, opt *EnvVarOption) (*EnvVar, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/env_var/%s", url.QueryEscape(repoSlug), id), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := es.client.NewRequest(http.MethodPatch, u, nil, nil)
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

// DeleteByRepoId deletes environment variable based on the given repository id and the env var id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/env_var#delete
func (es *EnvVarService) DeleteByRepoId(ctx context.Context, repoId uint, id string) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/env_var/%s", repoId, id), nil)
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
func (es *EnvVarService) DeleteByRepoSlug(ctx context.Context, repoSlug string, id string) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/env_var/%s", url.QueryEscape(repoSlug), id), nil)
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
