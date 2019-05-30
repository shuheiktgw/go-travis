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

// OrganizationsService handles communication with the
// organization related methods of the Travis CI API.
type OrganizationsService struct {
	client *Client
}

// Organization is a standard representation of an individual organization
//
// Travis CI API docs: https://developer.travis-ci.com/resource/organization#standard-representation
type Organization struct {
	// Value uniquely identifying the organization
	Id *uint `json:"id,omitempty"`
	// Login set on GitHub
	Login *string `json:"login,omitempty"`
	// Name set on GitHub
	Name *string `json:"name,omitempty"`
	// Id set on GitHub
	GithubId *uint `json:"github_id,omitempty"`
	// Avatar_url set on GitHub
	AvatarUrl *string `json:"avatar_url,omitempty"`
	// Whether or not the organization has an education account
	Education *bool `json:"education,omitempty"`
	// Repositories belonging to this organization.
	Repositories []*Repository `json:"repositories,omitempty"`
	// Installation belonging to the organization
	Installation *Installation `json:"installation,omitempty"`
	*Metadata
}

// OrganizationOption specifies the optional parameters for organization endpoint
type OrganizationOption struct {
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// OrganizationsOption specifies the optional parameters for organizations endpoint
type OrganizationsOption struct {
	// How many organizations to include in the response
	Limit int `url:"limit,omitempty"`
	// How many organizations to skip before the first entry in the response
	Offset int `url:"offset,omitempty"`
	// Attributes to sort organizations by
	SortBy string `url:"sort_by,omitempty"`
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// organizationsResponse represents a response
// from organizations endpoints
type organizationsResponse struct {
	Organizations []*Organization `json:"organizations,omitempty"`
}

// Find fetches an organization with the given id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/organization#find
func (os *OrganizationsService) Find(ctx context.Context, id uint, opt *OrganizationOption) (*Organization, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("org/%d", id), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := os.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var org Organization
	resp, err := os.client.Do(ctx, req, &org)
	if err != nil {
		return nil, resp, err
	}

	return &org, resp, err
}

// List fetches a list of organizations the current user is a member of
//
// Travis CI API docs: https://developer.travis-ci.com/resource/organizations#for_current_user
func (os *OrganizationsService) List(ctx context.Context, opt *OrganizationsOption) ([]*Organization, *http.Response, error) {
	u, err := urlWithOptions("orgs", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := os.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var or organizationsResponse
	resp, err := os.client.Do(ctx, req, &or)
	if err != nil {
		return nil, resp, err
	}

	return or.Organizations, resp, err
}
