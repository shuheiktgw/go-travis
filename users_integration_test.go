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

func TestUserService_Integration_Current(t *testing.T) {
	opt := UserOption{Include: []string{"user.repositories", "user.installation", "user.emails"}}
	user, res, err := integrationClient.User.Current(context.TODO(), &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if *user.Id != integrationUserId {
		t.Fatalf("unexpected user returned: want user_id: %d, got user_id %d", integrationUserId, user.Id)
	}

	if !user.Repositories[0].IsStandard() {
		t.Fatal("repository is not in a standard representation")
	}

	if len(user.Emails) == 0 {
		t.Fatal("emails are empty")
	}
}

func TestUserService_Integration_Find(t *testing.T) {
	opt := UserOption{Include: []string{"user.repositories", "user.installation", "user.emails"}}
	user, res, err := integrationClient.User.Find(context.TODO(), integrationUserId, &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if *user.Id != integrationUserId {
		t.Fatalf("unexpected user returned: want user_id: %d, got user_id %d", integrationUserId, user.Id)
	}

	if !user.Repositories[0].IsStandard() {
		t.Fatal("repository is not in a standard representation")
	}

	if len(user.Emails) == 0 {
		t.Fatal("emails are empty")
	}
}

func TestUserService_Integration_Sync(t *testing.T) {
	user, res, err := integrationClient.User.Sync(context.TODO(), integrationUserId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if *user.Id != integrationUserId {
		t.Fatalf("UserService.Find returned id %+v, want %+v", user.Id, integrationUserId)
	}
}
