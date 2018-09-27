// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
)

func TestEmailSubscriptionService_Integration_UnsubscribeAndSubscribeByRepoSlug(t *testing.T) {
	res, err := integrationClient.EmailSubscription.UnsubscribeByRepoSlug(context.TODO(), integrationRepoSlug)

	if err != nil {
		t.Fatalf("EmailSubscription.UnsubscribeByRepoSlug unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("EmailSubscription.UnsubscribeByRepoSlug invalid http status: %s", res.Status)
	}

	res, err = integrationClient.EmailSubscription.SubscribeByRepoSlug(context.TODO(), integrationRepoSlug)

	if err != nil {
		t.Fatalf("EmailSubscription.SubscribeByRepoSlug unexpected error occured: %s", err)
	}

	// This seems very wired so I created a PR and hope it will be merged soon
	// https://github.com/travis-ci/travis-api/pull/829
	// TODO: Fix the status code once the PR is merged
	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("EmailSubscription.SubscribeByRepoSlug invalid http status: %s", res.Status)
	}
}

func TestEmailSubscriptionService_Integration_UnsubscribeAndSubscribeByRepoId(t *testing.T) {
	res, err := integrationClient.EmailSubscription.UnsubscribeByRepoId(context.TODO(), integrationRepoId)

	if err != nil {
		t.Fatalf("EmailSubscription.UnsubscribeByRepoId unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("EmailSubscription.UnsubscribeByRepoId invalid http status: %s", res.Status)
	}

	res, err = integrationClient.EmailSubscription.SubscribeByRepoId(context.TODO(), integrationRepoId)

	if err != nil {
		t.Fatalf("EmailSubscription.SubscribeByRepoId unexpected error occured: %s", err)
	}

	// TODO: Fix the status code once the PR is merged
	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("EmailSubscription.SubscribeByRepoId invalid http status: %s", res.Status)
	}
}
