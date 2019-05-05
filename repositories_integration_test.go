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

func TestRepositoriesService_Integration_List(t *testing.T) {
	opt := RepositoriesOption{Active: true, Include: []string{"repository.default_branch"}}
	repos, res, err := integrationClient.Repositories.List(context.TODO(), &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if len(repos) == 0 {
		t.Fatal("repositories are empty")
	}
}

func TestRepositoriesService_Integration_ListByOwner(t *testing.T) {
	opt := RepositoriesOption{Active: true, Include: []string{"repository.default_branch"}}
	repos, res, err := integrationClient.Repositories.ListByOwner(context.TODO(), integrationGitHubOwner, &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if len(repos) == 0 {
		t.Fatal("repositories are empty")
	}
}

func TestRepositoriesService_Integration_ListGitHubId(t *testing.T) {
	opt := RepositoriesOption{Active: true, Include: []string{"repository.default_branch"}}
	repos, res, err := integrationClient.Repositories.ListByGitHubId(context.TODO(), integrationGitHubOwnerId, &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if len(repos) == 0 {
		t.Fatal("repositories are empty")
	}
}

func TestRepositoriesService_Integration_Find(t *testing.T) {
	opt := RepositoryOption{Include: []string{"repository.default_branch"}}
	repo, res, err := integrationClient.Repositories.Find(context.TODO(), integrationRepoSlug, &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if got, want := repo.Slug, integrationRepoSlug; got != want {
		t.Fatalf("unexpected repository returned: want %s: got %s", want, got)
	}

	if !repo.DefaultBranch.IsStandard() {
		t.Fatalf("default_branch is in a standard representation")
	}
}

func TestRepositoriesService_Integration_Activation(t *testing.T) {
	repo, res, err := integrationClient.Repositories.Deactivate(context.TODO(), integrationRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if repo.Slug != integrationRepoSlug {
		t.Fatalf("unexpected repository returned: want %s: got %s", integrationRepoSlug, repo.Slug)
	}

	repo, res, err = integrationClient.Repositories.Activate(context.TODO(), integrationRepoSlug)

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

func TestRepositoriesService_Integration_Migrate(t *testing.T) {
	_, res, err := integrationClient.Repositories.Migrate(context.TODO(), integrationRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	// The repository is not allowed to migrate as of May 5th, 2019
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}

func TestRepositoriesService_Integration_Star(t *testing.T) {
	repo, res, err := integrationClient.Repositories.Star(context.TODO(), integrationRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if repo.Slug != integrationRepoSlug {
		t.Fatalf("unexpected repository returned: want %s: got %s", integrationRepoSlug, repo.Slug)
	}

	repo, res, err = integrationClient.Repositories.Unstar(context.TODO(), integrationRepoSlug)

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
