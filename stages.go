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

// StagesService handles communication with the stage
// related methods of the Travis CI API.
type StagesService struct {
	client *Client
}

// Stage is a standard representation of an individual stage
//
// Travis CI API docs: https://developer.travis-ci.com/resource/stage#standard-representation
type Stage struct {
	// Value uniquely identifying the stage
	Id *uint `json:"id,omitempty"`
	// Incremental number for a stage
	Number *uint `json:"number,omitempty"`
	// The name of the stage
	Name *string `json:"name,omitempty"`
	// Current state of the stage
	State *string `json:"state,omitempty"`
	// When the stage started
	StartedAt *string `json:"started_at,omitempty"`
	// When the stage finished
	FinishedAt *string `json:"finished_at,omitempty"`
	// The jobs of a stage.
	Jobs []*Job `json:"jobs,omitempty"`
	*Metadata
}

// StagesOption is query parameters to one can specify to list stages
type StagesOption struct {
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

type stagesResponse struct {
	Stages []*Stage `json:"stages"`
}

// ListByBuild fetches stages of the build
//
// Travis CI API docs: https://developer.travis-ci.com/resource/stages#find
func (ss *StagesService) ListByBuild(ctx context.Context, buildId uint, opt *StagesOption) ([]*Stage, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("build/%d/stages", buildId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := ss.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var sr stagesResponse
	resp, err := ss.client.Do(ctx, req, &sr)
	if err != nil {
		return nil, resp, err
	}

	return sr.Stages, resp, err
}
