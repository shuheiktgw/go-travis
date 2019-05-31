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

func TestOwnerService_Integration_FindByLogin(t *testing.T) {
	opt := OwnerOption{Include: []string{"owner.repositories"}}
	owner, res, err := integrationClient.Owner.FindByLogin(context.TODO(), integrationGitHubOwner, &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if *owner.Login != integrationGitHubOwner {
		t.Fatalf("unexpected owner returned: want %s: got %s", integrationGitHubOwner, *owner.Login)
	}

	if len(owner.Repositories) == 0 {
		t.Fatal("repositories are empty")
	}

	if !owner.Repositories[0].IsStandard() {
		t.Fatal("repository is not in a standard representation")
	}
}

func TestOwnerService_Integration_FindByGitHubId(t *testing.T) {
	opt := OwnerOption{Include: []string{"owner.installation"}}
	owner, res, err := integrationClient.Owner.FindByGitHubId(context.TODO(), integrationGitHubOwnerId, &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if *owner.GitHubId != integrationGitHubOwnerId {
		t.Fatalf("unexpected owner returned: want %s: got %s", integrationGitHubOwner, *owner.Login)
	}
}
