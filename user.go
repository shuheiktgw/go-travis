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

// UserService handles communication with the users
// related methods of the Travis CI API.
type UserService struct {
	client *Client
}

// User represents a Travis CI user.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/user#standard-representation
type User struct {
	// Value uniquely identifying the user
	Id uint `json:"id,omitempty"`
	// Login set on Github
	Login string `json:"login,omitempty"`
	// Name set on GitHub
	Name string `json:"name,omitempty"`
	// Id set on GitHub
	GithubId uint `json:"github_id,omitempty"`
	// Avatar URL set on GitHub
	AvatarUrl string `json:"avatar_url,omitempty"`
	// Whether or not the user has an education account
	Education bool `json:"education,omitempty"`
	// Whether or not the user is currently being synced with Github
	IsSyncing bool `json:"is_syncing,omitempty"`
	// The last time the user was synced with GitHub
	SyncedAt string `json:"synced_at,omitempty"`
	Metadata
}

// Current fetches the currently authenticated user from Travis CI API.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/user#current
func (us *UserService) Current(ctx context.Context) (*User, *http.Response, error) {
	u, err := urlWithOptions("/user", nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := us.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var user User
	resp, err := us.client.Do(ctx, req, &user)
	if err != nil {
		return nil, resp, err
	}

	return &user, resp, err
}

// Get fetches the user with the provided id from the Travis CI API.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/user#find
func (us *UserService) Find(ctx context.Context, id uint) (*User, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/user/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := us.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var user User
	resp, err := us.client.Do(ctx, req, &user)
	if err != nil {
		return nil, resp, err
	}

	return &user, resp, err
}

// Sync triggers a new sync with GitHub.
// Might return status 409 if the user is currently syncing.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/user#sync
func (us *UserService) Sync(ctx context.Context, id uint) (*User, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/user/%d/sync", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := us.client.NewRequest(http.MethodPost, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var user User
	resp, err := us.client.Do(ctx, req, &user)
	if err != nil {
		return nil, nil, err
	}

	return &user, resp, err
}
