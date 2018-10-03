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

const integrationUserId uint = 1362503

func TestUserService_Integration_Current(t *testing.T) {
	user, res, err := integrationClient.User.Current(context.TODO())

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if user.Id != integrationUserId {
		t.Fatalf("unexpected user returned: want user_id: %d, got user_id %d", integrationUserId, user.Id)
	}
}

func TestUserService_Integration_Find(t *testing.T) {
	user, res, err := integrationClient.User.Find(context.TODO(), integrationUserId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if user.Id != integrationUserId {
		t.Fatalf("unexpected user returned: want user_id: %d, got user_id %d", integrationUserId, user.Id)
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

	if user.Id != integrationUserId {
		t.Fatalf("UserService.Find returned id %+v, want %+v", user.Id, integrationUserId)
	}
}
