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

func TestRepositoryService_Integration_Find(t *testing.T) {
	repo, res, err := integrationClient.Repository.Find(context.TODO(), integrationRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if got, want := repo.Slug, integrationRepoSlug; got != want {
		t.Fatalf("unexpected repository returned: want %s: got %s", want, got)
	}
}

func TestRepositoryService_Integration_Activation(t *testing.T) {
	t.Parallel()

	repo, res, err := integrationClient.Repository.Deactivate(context.TODO(), integrationRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if repo.Slug != integrationRepoSlug {
		t.Fatalf("unexpected repository returned: want %s: got %s", integrationRepoSlug, repo.Slug)
	}

	repo, res, err = integrationClient.Repository.Activate(context.TODO(), integrationRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if repo.Slug != integrationRepoSlug {
		t.Fatalf("unexpected repository returned: want %s: got %s", integrationRepoSlug, repo.Slug)
	}
}

func TestRepositoryService_Integration_Star(t *testing.T) {
	t.Parallel()

	repo, res, err := integrationClient.Repository.Star(context.TODO(), integrationRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if repo.Slug != integrationRepoSlug {
		t.Fatalf("unexpected repository returned: want %s: got %s", integrationRepoSlug, repo.Slug)
	}

	repo, res, err = integrationClient.Repository.Unstar(context.TODO(), integrationRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if repo.Slug != integrationRepoSlug {
		t.Fatalf("unexpected repository returned: want %s: got %s", integrationRepoSlug, repo.Slug)
	}
}
