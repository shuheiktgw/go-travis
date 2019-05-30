// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"net/http"
)

// BroadcastsService handles communication with the
// broadcasts related methods of the Travis CI API.
type BroadcastsService struct {
	client *Client
}

// Broadcast is a standard representation of an individual broadcast
//
// Travis CI API docs: https://developer.travis-ci.com/resource/broadcast#standard-representation
type Broadcast struct {
	// Value uniquely identifying the broadcast
	Id *uint `json:"id,omitempty"`
	// Message to display to the user
	Message *string `json:"message,omitempty"`
	// Broadcast category (used for icon and color)
	Category *string `json:"category,omitempty"`
	// Whether or not the broadcast should still be displayed
	Active *bool `json:"active,omitempty"`
	// When the broadcast was created
	CreatedAt *string `json:"created_at,omitempty"`
	// Either a user, organization or repository, or null for global
	Recipient interface{} `json:"recipient,omitempty"`
	*Metadata
}

// BroadcastsOption specifies the optional parameters for broadcasts endpoint
type BroadcastsOption struct {
	// Filters broadcasts by whether or not the broadcast should still be displayed
	Active bool `url:"active"`
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// broadcastsResponse represents a response
// from broadcast endpoints
type broadcastsResponse struct {
	Broadcasts []*Broadcast `json:"broadcasts,omitempty"`
}

// List fetches a list of broadcasts for the current user
//
// Travis CI API docs: https://developer.travis-ci.com/resource/broadcasts#for_current_user
func (bs *BroadcastsService) List(ctx context.Context, opt *BroadcastsOption) ([]*Broadcast, *http.Response, error) {
	u, err := urlWithOptions("broadcasts", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var br broadcastsResponse
	resp, err := bs.client.Do(ctx, req, &br)
	if err != nil {
		return nil, resp, err
	}

	return br.Broadcasts, resp, err
}
