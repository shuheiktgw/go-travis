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

// InstallationsService handles communication with
// the installation related methods of the Travis CI API.
type InstallationsService struct {
	client *Client
}

// Installation represents a GitHub App installation
//
// Travis CI API docs: https://developer.travis-ci.com/resource/installation#standard-representation
type Installation struct {
	// The installation id
	Id *uint `json:"id,omitempty"`
	// The installation's id on GitHub
	GitHubId *uint `json:"github_id,omitempty"`
	// GitHub user or organization the installation belongs to
	Owner *Owner `json:"owner,omitempty"`
	*Metadata
}

// InstallationOption is query parameters to one can specify to find installation
type InstallationOption struct {
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// Find fetches a single GitHub installation based on the provided id.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/installation#find
func (is *InstallationsService) Find(ctx context.Context, id uint, opt *InstallationOption) (*Installation, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("installation/%d", id), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := is.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var installation Installation
	resp, err := is.client.Do(ctx, req, &installation)
	if err != nil {
		return nil, resp, err
	}

	return &installation, resp, err
}
