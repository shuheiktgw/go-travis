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

// CronsService handles communication with the crons
// related methods of the Travis CI API.
type CronsService struct {
	client *Client
}

// Cron is a standard representation of an individual cron
//
// Travis CI API docs: https://developer.travis-ci.com/resource/cron#standard-representation
type Cron struct {
	// Value uniquely identifying the cron
	Id *uint `json:"id,omitempty"`
	// Github repository to which this cron belongs
	Repository *Repository `json:"repository,omitempty"`
	// Git branch of repository to which this cron belongs
	Branch *Branch `json:"branch,omitempty"`
	// Interval at which the cron will run (can be "daily", "weekly" or "monthly")
	Interval *string `json:"interval,omitempty"`
	// Whether a cron build should run if there has been a build on this branch in the last 24 hours
	DontRunIfRecentBuildExists *bool `json:"dont_run_if_recent_build_exists,omitempty"`
	// When the cron ran last
	LastRun *string `json:"last_run,omitempty"`
	// When the cron is scheduled to run next
	NextRun *string `json:"next_run,omitempty"`
	// When the cron was created
	CreatedAt *string `json:"created_at,omitempty"`
	// Whether the cron is active or not
	Active *bool `json:"active,omitempty"`
	*Metadata
}

// CronBody specifies body for
// creating cron.
type CronBody struct {
	// Interval at which the cron will run (can be "daily", "weekly" or "monthly")
	Interval string `json:"cron.interval,omitempty"`
	// Whether a cron build should run if there has been a build on this branch in the last 24 hours
	DontRunIfRecentBuildExists bool `json:"cron.dont_run_if_recent_build_exists"`
}

// CronOption specifies options for
// fetching cron.
type CronOption struct {
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// CronsOption specifies options for
// fetching crons.
type CronsOption struct {
	// How many crons to include in the response
	Limit int `url:"limit,omitempty"`
	// How many crons to skip before the first entry in the response
	Offset int `url:"offset,omitempty"`
	// List of attributes to eager load
	Include []string `url:"include,omitempty,comma"`
}

// cronsResponse represents a response
// from crons endpoints
type cronsResponse struct {
	Crons []*Cron `json:"crons,omitempty"`
}

const (
	CronIntervalDaily   = "daily"
	CronIntervalWeekly  = "weekly"
	CronIntervalMonthly = "monthly"
)

// Find fetches a cron based on the provided id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/cron#find
func (cs *CronsService) Find(ctx context.Context, id uint, opt *CronOption) (*Cron, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("cron/%d", id), opt)
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
func (cs *CronsService) FindByRepoId(ctx context.Context, repoId uint, branch string, opt *CronOption) (*Cron, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/branch/%s/cron", repoId, branch), opt)
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
func (cs *CronsService) FindByRepoSlug(ctx context.Context, repoSlug string, branch string, opt *CronOption) (*Cron, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/branch/%s/cron", url.QueryEscape(repoSlug), branch), opt)
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

// ListByRepoId fetches crons based on the provided repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/crons#for_repository
func (cs *CronsService) ListByRepoId(ctx context.Context, repoId uint, opt *CronsOption) ([]*Cron, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/crons", repoId), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var cr cronsResponse
	resp, err := cs.client.Do(ctx, req, &cr)
	if err != nil {
		return nil, resp, err
	}

	return cr.Crons, resp, err
}

// ListByRepoSlug fetches crons based on the provided repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/crons#for_repository
func (cs *CronsService) ListByRepoSlug(ctx context.Context, repoSlug string, opt *CronsOption) ([]*Cron, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/crons", url.QueryEscape(repoSlug)), opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var cr cronsResponse
	resp, err := cs.client.Do(ctx, req, &cr)
	if err != nil {
		return nil, resp, err
	}

	return cr.Crons, resp, err
}

// CreateByRepoId creates a cron based on the provided repository id and branch name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/cron#create
func (cs *CronsService) CreateByRepoId(ctx context.Context, repoId uint, branchName string, cron *CronBody) (*Cron, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/branch/%s/cron", repoId, branchName), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodPost, u, cron, nil)
	if err != nil {
		return nil, nil, err
	}

	var c Cron
	resp, err := cs.client.Do(ctx, req, &c)
	if err != nil {
		return nil, resp, err
	}

	return &c, resp, err
}

// CreateByRepoSlug creates a cron based on the provided repository slug and branch name
//
// Travis CI API docs: https://developer.travis-ci.com/resource/cron#create
func (cs *CronsService) CreateByRepoSlug(ctx context.Context, repoSlug string, branchName string, cron *CronBody) (*Cron, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/branch/%s/cron", url.QueryEscape(repoSlug), branchName), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := cs.client.NewRequest(http.MethodPost, u, cron, nil)
	if err != nil {
		return nil, nil, err
	}

	var c Cron
	resp, err := cs.client.Do(ctx, req, &c)
	if err != nil {
		return nil, resp, err
	}

	return &c, resp, err
}

// Delete deletes a cron based on the provided id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/cron#delete
func (cs *CronsService) Delete(ctx context.Context, id uint) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("cron/%d", id), nil)
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
