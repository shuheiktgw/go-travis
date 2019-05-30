// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// SettingsService handles communication with the setting
// related methods of the Travis CI API.
type SettingsService struct {
	client *Client
}

// Setting represents a Travis CI setting.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/setting#standard-representation
type Setting struct {
	// The setting's name
	Name *string `json:"name,omitempty"`
	// The setting's value
	// Currently value can be boolean or integer
	Value interface{} `json:"value,omitempty"`
	*Metadata
}

// SettingBody is a body to update setting
type SettingBody struct {
	Name  string      `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// settingsResponse represents response from the settings endpoints
type settingsResponse struct {
	Settings []*Setting `json:"settings,omitempty"`
}

const (
	// BuildsOnlyWithTravisYmlSetting is a setting name for builds_only_with_travis_yml
	BuildsOnlyWithTravisYmlSetting = "builds_only_with_travis_yml"
	// BuildPushesSetting is a setting name for build_pushes
	BuildPushesSetting = "build_pushes"
	// BuildPullRequestsSetting is a setting name for build_pull_requests
	BuildPullRequestsSetting = "build_pull_requests"
	// MaximumNumberOfBuildsSetting is a setting name for maximum_number_of_builds
	MaximumNumberOfBuildsSetting = "maximum_number_of_builds"
	// AutoCancelPushesSetting is a setting name for auto_cancel_pushes
	AutoCancelPushesSetting = "auto_cancel_pushes"
	// AutoCancelPullRequestsSetting is a setting name for auto_cancel_pull_requests
	AutoCancelPullRequestsSetting = "auto_cancel_pull_requests"
)

// FindByRepoId fetches a setting of given repository id and setting name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/setting#find
func (ss *SettingsService) FindByRepoId(ctx context.Context, repoId uint, name string) (*Setting, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/setting/%s", repoId, name), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ss.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var setting Setting
	resp, err := ss.client.Do(ctx, req, &setting)
	if err != nil {
		return nil, resp, err
	}

	return &setting, resp, err
}

// FindByRepoSlug fetches a setting of given repository slug and setting name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/setting#find
func (ss *SettingsService) FindByRepoSlug(ctx context.Context, repoSlug string, name string) (*Setting, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/setting/%s", url.QueryEscape(repoSlug), name), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ss.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var setting Setting
	resp, err := ss.client.Do(ctx, req, &setting)
	if err != nil {
		return nil, resp, err
	}

	return &setting, resp, err
}

// ListByRepoId fetches a list of settings of given repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/settings#for_repository
func (ss *SettingsService) ListByRepoId(ctx context.Context, repoId uint) ([]*Setting, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/settings", repoId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ss.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var sr settingsResponse
	resp, err := ss.client.Do(ctx, req, &sr)
	if err != nil {
		return nil, resp, err
	}

	return sr.Settings, resp, err
}

// ListByRepoSlug fetches a list of settings of given repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/settings#for_repository
func (ss *SettingsService) ListByRepoSlug(ctx context.Context, repoSlug string) ([]*Setting, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/settings", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ss.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var sr settingsResponse
	resp, err := ss.client.Do(ctx, req, &sr)
	if err != nil {
		return nil, resp, err
	}

	return sr.Settings, resp, err
}

// UpdateByRepoId updates a setting with setting property
//
// Travis CI API docs: https://developer.travis-ci.com/resource/setting#update
func (ss *SettingsService) UpdateByRepoId(ctx context.Context, repoId uint, setting *SettingBody) (*Setting, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/setting/%s", repoId, setting.Name), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ss.client.NewRequest(
		http.MethodPatch,
		u,
		map[string]interface{}{"setting.value": setting.Value},
		nil,
	)

	if err != nil {
		return nil, nil, err
	}

	var s Setting
	resp, err := ss.client.Do(ctx, req, &s)
	if err != nil {
		return nil, resp, err
	}

	return &s, resp, err
}

// UpdateByRepoSlug updates a setting with setting property
//
// Travis CI API docs: https://developer.travis-ci.com/resource/setting#update
func (ss *SettingsService) UpdateByRepoSlug(ctx context.Context, repoSlug string, setting *SettingBody) (*Setting, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/setting/%s", url.QueryEscape(repoSlug), setting.Name), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ss.client.NewRequest(
		http.MethodPatch,
		u,
		map[string]interface{}{"setting.value": setting.Value},
		nil,
	)

	if err != nil {
		return nil, nil, err
	}

	var s Setting
	resp, err := ss.client.Do(ctx, req, &s)
	if err != nil {
		return nil, resp, err
	}

	return &s, resp, err
}
