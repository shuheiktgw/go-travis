// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestEmailSubscriptionService_SubscribeByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/email_subscription", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	_, err := client.EmailSubscription.SubscribeByRepoId(context.Background(), testRepoId)

	if err != nil {
		t.Errorf("EmailSubscription.SubscribeByRepoId returned error: %v", err)
	}
}

func TestEmailSubscriptionService_SubscribeByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/email_subscription", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	_, err := client.EmailSubscription.SubscribeByRepoSlug(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("EmailSubscription.SubscribeByRepoSlug returned error: %v", err)
	}
}

func TestEmailSubscriptionService_UnsubscribeByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/email_subscription", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.EmailSubscription.UnsubscribeByRepoId(context.Background(), testRepoId)

	if err != nil {
		t.Errorf("EmailSubscription.UnsubscribeByRepoId returned error: %v", err)
	}
}

func TestEmailSubscriptionService_UnsubscribeByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/email_subscription", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.EmailSubscription.UnsubscribeByRepoSlug(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("EmailSubscription.UnsubscribeByRepoId returned error: %v", err)
	}
}
