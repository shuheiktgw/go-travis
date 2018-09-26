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

// PreferenceService handles communication with the
// preference related methods of the Travis CI API.
type PreferenceService struct {
	client *Client
}

// Preference is a standard representation of an individual preference
//
// Travis CI API docs: https://developer.travis-ci.com/resource/preference#standard-representation
type Preference struct {
	// The preference's name
	Name string `json:"name,omitempty"`
	// The preference's value
	Value bool `json:"value"`
}

// Find fetches the current user's preference based on
// the provided preference name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/preference#find
func (ps *PreferenceService) Find(ctx context.Context, name string) (*Preference, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/preference/%s", name), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ps.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var preference Preference
	resp, err := ps.client.Do(ctx, req, &preference)
	if err != nil {
		return nil, resp, err
	}

	return &preference, resp, err
}

// Update updates the current user's preference based on
// the provided preference property
//
// Travis CI API docs: https://developer.travis-ci.com/resource/preference#update
func (ps *PreferenceService) Update(ctx context.Context, preference *Preference) (*Preference, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/preference/%s", preference.Name), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ps.client.NewRequest(http.MethodPatch, u, preference, nil)
	if err != nil {
		return nil, nil, err
	}

	var p Preference
	resp, err := ps.client.Do(ctx, req, &p)
	if err != nil {
		return nil, resp, err
	}

	return &p, resp, err
}
