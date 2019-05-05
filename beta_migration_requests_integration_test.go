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

func TestBetaMigrationRequestsService_Integration_List(t *testing.T) {
	opt := BetaMigrationRequestsOption{Include: []string{"beta_migration_request.organizations"}}
	requests, res, err := integrationClient.BetaMigrationRequests.List(context.TODO(), integrationUserId, &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if len(requests) == 0 {
		t.Fatal("requests cannot be empty")
	}
}

func TestBetaMigrationRequestsService_Integration_Create(t *testing.T) {
	r := BetaMigrationRequestBody{OrganizationIds: []uint{integrationOrgId}}
	request, res, err := integrationClient.BetaMigrationRequests.Create(context.TODO(), integrationUserId, &r)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if got, want := request.OwnerName, integrationGitHubOwner; got != want {
		t.Fatalf("invalid owner: got: %s, want: %s", got, want)
	}
}
