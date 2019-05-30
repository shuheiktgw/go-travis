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

// EmailSubscriptionService handles communication with the
// email subscription related methods of the Travis CI API.
type EmailSubscriptionsService struct {
	client *Client
}

// SubscribeByRepoId enables an email subscription of the repository based on the provided repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/email_subscription#resubscribe
func (es *EmailSubscriptionsService) SubscribeByRepoId(ctx context.Context, repoId uint) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/email_subscription", repoId), nil)
	if err != nil {
		return nil, err
	}

	req, err := es.client.NewRequest(http.MethodPost, u, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := es.client.Do(ctx, req, nil)
	return resp, err
}

// SubscribeByRepoSlug enables an email subscription of the repository based on the provided repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/email_subscription#resubscribe
func (es *EmailSubscriptionsService) SubscribeByRepoSlug(ctx context.Context, repoSlug string) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/email_subscription", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, err
	}

	req, err := es.client.NewRequest(http.MethodPost, u, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := es.client.Do(ctx, req, nil)
	return resp, err
}

// UnsubscribeByRepoId disables an email subscription of the repository based on the provided repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/email_subscription#unsubscribe
func (es *EmailSubscriptionsService) UnsubscribeByRepoId(ctx context.Context, repoId uint) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/email_subscription", repoId), nil)
	if err != nil {
		return nil, err
	}

	req, err := es.client.NewRequest(http.MethodDelete, u, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := es.client.Do(ctx, req, nil)
	return resp, err
}

// UnsubscribeByRepoSlug disables an email subscription of the repository based on the provided repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/email_subscription#unsubscribe
func (es *EmailSubscriptionsService) UnsubscribeByRepoSlug(ctx context.Context, repoSlug string) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/email_subscription", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, err
	}

	req, err := es.client.NewRequest(http.MethodDelete, u, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := es.client.Do(ctx, req, nil)
	return resp, err
}
