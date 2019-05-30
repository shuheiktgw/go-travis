// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"net/http"
)

// LintService handles communication with the lint endpoint
// of Travis CI API
type LintService struct {
	client *Client
}

type TravisYml struct {
	Content string `json:"content,omitempty"`
}

// Warning is a standard representation of
// a warning of .travis.yml content validation
//
// Travis CI API docs: https://developer.travis-ci.com/resource/lint
type Warning struct {
	Key     []*string `json:"key,omitempty"`
	Message *string   `json:"message,omitempty"`
	*Metadata
}

type lintResponse struct {
	Warnings []*Warning `json:"warnings,omitempty"`
}

// Lint validates the .travis.yml file and returns any warnings
//
// Travis CI API docs: https://developer.travis-ci.com/resource/lint#lint
func (es *LintService) Lint(ctx context.Context, yml *TravisYml) ([]*Warning, *http.Response, error) {
	u, err := urlWithOptions("lint", nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := es.client.NewRequest(http.MethodPost, u, yml, nil)
	if err != nil {
		return nil, nil, err
	}

	var lintResponse lintResponse
	resp, err := es.client.Do(ctx, req, &lintResponse)
	if err != nil {
		return nil, resp, err
	}

	return lintResponse.Warnings, resp, err
}
