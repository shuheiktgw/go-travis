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

func TestEmailSubscriptionsService_SubscribeByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/email_subscription", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	_, err := client.EmailSubscriptions.SubscribeByRepoId(context.Background(), testRepoId)

	if err != nil {
		t.Errorf("EmailSubscriptions.SubscribeByRepoId returned error: %v", err)
	}
}

func TestEmailSubscriptionsService_SubscribeByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/email_subscription", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	_, err := client.EmailSubscriptions.SubscribeByRepoSlug(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("EmailSubscriptions.SubscribeByRepoSlug returned error: %v", err)
	}
}

func TestEmailSubscriptionsService_UnsubscribeByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/email_subscription", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.EmailSubscriptions.UnsubscribeByRepoId(context.Background(), testRepoId)

	if err != nil {
		t.Errorf("EmailSubscriptions.UnsubscribeByRepoId returned error: %v", err)
	}
}

func TestEmailSubscriptionsService_UnsubscribeByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/email_subscription", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.EmailSubscriptions.UnsubscribeByRepoSlug(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("EmailSubscriptions.UnsubscribeByRepoId returned error: %v", err)
	}
}
