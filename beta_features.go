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

// BetaFeaturesService handles communication with the
// beta feature related methods of the Travis CI API.
type BetaFeaturesService struct {
	client *Client
}

// BetaFeature is a standard representation of an individual beta feature
//
// Travis CI API docs: https://developer.travis-ci.com/resource/beta_feature#attributes
type BetaFeature struct {
	// Value uniquely identifying the beta feature
	Id *uint `json:"id,omitempty"`
	// The name of the feature
	Name *string `json:"name,omitempty"`
	// Longer description of the feature
	Description *string `json:"description,omitempty"`
	// Indicates if the user has this feature turned on
	Enabled *bool `json:"enabled,omitempty"`
	// Url for users to leave Travis CI feedback on this feature
	FeedbackUrl *string `json:"feedback_url,omitempty"`
	*Metadata
}

// betaFeaturesResponse represents a response
// from organizations endpoints
type betaFeaturesResponse struct {
	BetaFeatures []*BetaFeature `json:"beta_features,omitempty"`
}

// List fetches a list of beta features available to a user
//
// Travis CI API docs: https://developer.travis-ci.com/resource/beta_features#find
func (bs *BetaFeaturesService) List(ctx context.Context, userId uint) ([]*BetaFeature, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("user/%d/beta_features", userId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var br betaFeaturesResponse
	resp, err := bs.client.Do(ctx, req, &br)
	if err != nil {
		return nil, resp, err
	}

	return br.BetaFeatures, resp, err
}

// Update updates a user's beta_feature
//
// Travis CI API docs: https://developer.travis-ci.com/resource/beta_feature#update
func (bs *BetaFeaturesService) Update(ctx context.Context, userId uint, id uint, enabled bool) (*BetaFeature, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("user/%d/beta_feature/%d", userId, id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodPatch, u, map[string]bool{"enabled": enabled}, nil)
	if err != nil {
		return nil, nil, err
	}

	var feature BetaFeature
	resp, err := bs.client.Do(ctx, req, &feature)
	if err != nil {
		return nil, resp, err
	}

	return &feature, resp, err
}

// Delete delete a user's beta feature
//
// Travis CI API docs: https://developer.travis-ci.com/resource/beta_feature#delete
func (bs *BetaFeaturesService) Delete(ctx context.Context, userId uint, id uint) (*BetaFeature, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("user/%d/beta_feature/%d", userId, id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodDelete, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var feature BetaFeature
	resp, err := bs.client.Do(ctx, req, &feature)
	if err != nil {
		return nil, resp, err
	}

	return &feature, resp, err
}
