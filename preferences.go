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

// PreferencesService handles communication with the
// preferences related methods of the Travis CI API.
type PreferencesService struct {
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

type getPreferencesResponse struct {
	Preferences []Preference `json:"preferences,omitempty"`
}

// Find fetches the current user's preference based on
// the provided preference name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/preference#find
func (ps *PreferencesService) Find(ctx context.Context, name string) (*Preference, *http.Response, error) {
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

// List fetches the current user's preferences
//
// Travis CI API docs: https://developer.travis-ci.com/resource/preferences#for_user
func (ps *PreferencesService) List(ctx context.Context) ([]Preference, *http.Response, error) {
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

// Update updates the current user's preference based on
// the provided preference property
//
// Travis CI API docs: https://developer.travis-ci.com/resource/preference#update
func (ps *PreferencesService) Update(ctx context.Context, preference *Preference) (*Preference, *http.Response, error) {
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
