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

// CronService handles communication with the cron
// related methods of the Travis CI API.
type CronService struct {
	client *Client
}

// Cron is a standard representation of an individual cron
//
// Travis CI API docs: https://developer.travis-ci.com/resource/cron#standard-representation
type Cron struct {
	// Value uniquely identifying the cron
	Id uint `json:"id,omitempty"`
	// Github repository to which this cron belongs
	Repository MinimalRepository `json:"repository,omitempty"`
	// Git branch of repository to which this cron belongs
	Branch MinimalBranch `json:"branch,omitempty"`
	// Interval at which the cron will run (can be "daily", "weekly" or "monthly")
	Interval string `json:"interval,omitempty"`
	// Whether a cron build should run if there has been a build on this branch in the last 24 hours
	DontRunIfRecentBuildExists bool `json:"dont_run_if_recent_build_exists,omitempty"`
	// When the cron ran last
	LastRun string `json:"last_run,omitempty"`
	// When the cron is scheduled to run next
	NextRun string `json:"next_run,omitempty"`
	// When the cron was created
	CreatedAt string `json:"created_at,omitempty"`
	// Whether the cron is active or not
	Active bool `json:"active,omitempty"`
}

// CronOption specifies options for
// creating cron.
type CronOption struct {
	// Interval at which the cron will run (can be "daily", "weekly" or "monthly")
	Interval string `json:"cron.interval,omitempty"`
	// Whether a cron build should run if there has been a build on this branch in the last 24 hours
	DontRunIfRecentBuildExists bool `json:"cron.dont_run_if_recent_build_exists"`
}

const (
	CronIntervalDaily   = "daily"
	CronIntervalWeekly  = "weekly"
	CronIntervalMonthly = "monthly"
)

// CreateByRepoId creates a cron based on the provided repository id and branch name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/cron#create
func (cs *CronService) CreateByRepoId(ctx context.Context, repoId uint, branchName string, opt *CronOption) (*Cron, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/branch/%s/cron", repoId, branchName), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodPost, u, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	var cron Cron
	resp, err := cs.client.Do(ctx, req, &cron)
	if err != nil {
		return nil, resp, err
	}

	return &cron, resp, err
}

// CreateByRepoSlug creates a cron based on the provided repository slug and branch name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/cron#create
func (cs *CronService) CreateByRepoSlug(ctx context.Context, repoSlug string, branchName string, opt *CronOption) (*Cron, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/branch/%s/cron", url.QueryEscape(repoSlug), branchName), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodPost, u, opt, nil)
	if err != nil {
		return nil, nil, err
	}

	var cron Cron
	resp, err := cs.client.Do(ctx, req, &cron)
	if err != nil {
		return nil, resp, err
	}

	return &cron, resp, err
}

// Find fetches a cron based on the provided id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/cron#find
func (cs *CronService) Find(ctx context.Context, id uint) (*Cron, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/cron/%d", id), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var cron Cron
	resp, err := cs.client.Do(ctx, req, &cron)
	if err != nil {
		return nil, resp, err
	}

	return &cron, resp, err
}

// FindByRepoId fetches a cron based on the provided repository id and branch name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/cron#for_branch
func (cs *CronService) FindByRepoId(ctx context.Context, repoId uint, branch string) (*Cron, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%d/branch/%s/cron", repoId, branch), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var cron Cron
	resp, err := cs.client.Do(ctx, req, &cron)
	if err != nil {
		return nil, resp, err
	}

	return &cron, resp, err
}

// FindByRepoSlug fetches a cron based on the provided repository slug and branch name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/cron#for_branch
func (cs *CronService) FindByRepoSlug(ctx context.Context, repoSlug string, branch string) (*Cron, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/branch/%s/cron", url.QueryEscape(repoSlug), branch), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var cron Cron
	resp, err := cs.client.Do(ctx, req, &cron)
	if err != nil {
		return nil, resp, err
	}

	return &cron, resp, err
}

// Delete deletes a cron based on the provided id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/cron#delete
func (cs *CronService) Delete(ctx context.Context, id uint) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/cron/%d", id), nil)
	if err != nil {
		return nil, err
	}

	req, err := cs.client.NewRequest(http.MethodDelete, u, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := cs.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
