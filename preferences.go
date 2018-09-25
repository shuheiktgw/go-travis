// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"net/http"
)

// PreferencesService handles communication with the
// preferences related methods of the Travis CI API.
type PreferencesService struct {
	client *Client
}

type getPreferencesResponse struct {
	Preferences []Preference `json:"preferences,omitempty"`
}

// Find fetches the current user's preferences
//
// Travis CI API docs: https://developer.travis-ci.com/resource/preferences#for_user
func (ps *PreferencesService) Find(ctx context.Context) ([]Preference, *http.Response, error) {
	u, err := urlWithOptions("/preferences", nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ps.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var getPreferencesResponse getPreferencesResponse
	resp, err := ps.client.Do(ctx, req, &getPreferencesResponse)
	if err != nil {
		return nil, resp, err
	}

	return getPreferencesResponse.Preferences, resp, err
}
