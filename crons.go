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

// CronsService handles communication with the crons
// related methods of the Travis CI API.
type CronsService struct {
	client *Client
}

// CronsOption specifies options for
// fetching crons.
type CronsOption struct {
	// How many crons to include in the response
	Limit int `url:"limit,omitempty"`
	// How many crons to skip before the first entry in the response
	Offset int `url:"offset,omitempty"`
}

// getCronsResponse represents a response
// from crons endpoints
type getCronsResponse struct {
	Crons []Cron `json:"crons,omitempty"`
}

// FindByRepoId fetches crons based on the provided repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/crons#for_repository
func (cs *CronsService) FindByRepoId(ctx context.Context, repoId uint, opt *CronsOption) ([]Cron, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/crons", repoId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getCronsResponse getCronsResponse
	resp, err := cs.client.Do(ctx, req, &getCronsResponse)
	if err != nil {
		return nil, resp, err
	}

	return getCronsResponse.Crons, resp, err
}

// FindByRepoId fetches crons based on the provided repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/crons#for_repository
func (cs *CronsService) FindByRepoSlug(ctx context.Context, repoSlug string, opt *CronsOption) ([]Cron, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/crons", url.QueryEscape(repoSlug)), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getCronsResponse getCronsResponse
	resp, err := cs.client.Do(ctx, req, &getCronsResponse)
	if err != nil {
		return nil, resp, err
	}

	return getCronsResponse.Crons, resp, err
}
