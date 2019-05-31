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

// CachesService handles communication with the caches endpoints
// of Travis CI API
type CachesService struct {
	client *Client
}

// Cache is a standard representation of a cache on Travis CI
//
// Travis CI API docs: https://developer.travis-ci.com/resource/caches#attributes
type Cache struct {
	// The branch the cache belongs to
	Branch *string `json:"branch,omitempty"`
	// The string to match against the cache name
	Match *string `json:"match,omitempty"`
	*Metadata
}

// cachesResponse represents the response of a call
// to the Travis CI caches endpoint.
type cachesResponse struct {
	Caches []*Cache `json:"caches"`
}

// ListByRepoId fetches caches based on the given repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/caches#find
func (cs *CachesService) ListByRepoId(ctx context.Context, repoId uint) ([]*Cache, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/caches", repoId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var cachesResponse cachesResponse
	resp, err := cs.client.Do(ctx, req, &cachesResponse)
	if err != nil {
		return nil, resp, err
	}

	return cachesResponse.Caches, resp, err
}

// ListByRepoSlug fetches caches based on the given repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/caches#find
func (cs *CachesService) ListByRepoSlug(ctx context.Context, repoSlug string) ([]*Cache, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/caches", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var cachesResponse cachesResponse
	resp, err := cs.client.Do(ctx, req, &cachesResponse)
	if err != nil {
		return nil, resp, err
	}

	return cachesResponse.Caches, resp, err
}

// DeleteByRepoId deletes caches based on the given repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/caches#delete
func (cs *CachesService) DeleteByRepoId(ctx context.Context, repoId uint) ([]*Cache, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/caches", repoId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodDelete, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var cachesResponse cachesResponse
	resp, err := cs.client.Do(ctx, req, &cachesResponse)
	if err != nil {
		return nil, resp, err
	}

	return cachesResponse.Caches, resp, err
}

// DeleteByRepoSlug deletes caches based on the given repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/caches#delete
func (cs *CachesService) DeleteByRepoSlug(ctx context.Context, repoSlug string) ([]*Cache, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/caches", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodDelete, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var cachesResponse cachesResponse
	resp, err := cs.client.Do(ctx, req, &cachesResponse)
	if err != nil {
		return nil, resp, err
	}

	return cachesResponse.Caches, resp, err
}
