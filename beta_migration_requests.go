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

// BetaMigrationRequestsService handles communication with the
// beta migration requests related methods of the Travis CI API.
type BetaMigrationRequestsService struct {
	client *Client
}

// BetaMigrationRequest is a standard representation of an individual beta migration request
//
// Travis CI API docs: https://developer.travis-ci.com/resource/beta_migration_request#attributes
type BetaMigrationRequest struct {
	// The beta_migration_request's id
	Id *uint `json:"id,omitempty"`
	// The beta_migration_request's owner_id
	OwnerId *uint `json:"owner_id,omitempty"`
	// The beta_migration_request's owner_name
	OwnerName *string `json:"owner_name,omitempty"`
	// Longer description of the feature
	OwnerType *string `json:"owner_type,omitempty"`
	// The beta_migration_request's accepted_at
	AcceptedAt *string `json:"accepted_at,omitempty"`
	// The beta_migration_request's organizations
	Organizations []*Organization `json:"organizations,omitempty"`
	*Metadata
}

// BetaMigrationRequestOption specifies options for
// finding a beta migration requests.
type BetaMigrationRequestsOption struct {
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// BetaMigrationRequestBody specifies body for
// creating a beta migration request.
type BetaMigrationRequestBody struct {
	// The beta_migration_request's organizations
	OrganizationIds []uint `json:"beta_migration_request.organizations"`
}

type betaMigrationRequestsResponse struct {
	BetaMigrationRequests []*BetaMigrationRequest `json:"beta_migration_requests,omitempty"`
}

// List fetches a list of beta migration requests created by the user
//
// Travis CI API docs: https://developer.travis-ci.com/resource/beta_migration_requests#find
func (bs *BetaMigrationRequestsService) List(ctx context.Context, userId uint, opt *BetaMigrationRequestsOption) ([]*BetaMigrationRequest, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("user/%d/beta_migration_requests", userId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var bmrr betaMigrationRequestsResponse
	resp, err := bs.client.Do(ctx, req, &bmrr)
	if err != nil {
		return nil, resp, err
	}

	return bmrr.BetaMigrationRequests, resp, err
}

// Create creates a beta migration request
//
// Travis CI API docs: https://developer.travis-ci.com/resource/beta_migration_request#create
func (bs *BetaMigrationRequestsService) Create(ctx context.Context, userId uint, request *BetaMigrationRequestBody) (*BetaMigrationRequest, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("user/%d/beta_migration_request", userId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodPost, u, request, nil)
	if err != nil {
		return nil, nil, err
	}

	var bmr BetaMigrationRequest
	resp, err := bs.client.Do(ctx, req, &bmr)
	if err != nil {
		return nil, resp, err
	}

	return &bmr, resp, err
}
