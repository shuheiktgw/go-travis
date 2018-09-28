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

func TestBranchesService_Integration_FindByRepoId(t *testing.T) {
	branch, res, err := integrationClient.Branches.FindByRepoId(context.TODO(), integrationRepoId, "master")

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if branch.Name != "master" {
		t.Fatalf("unexpected branch returned: want %s: got %s", "master", branch.Name)
	}

	if branch.Repository.Id != integrationRepoId {
		t.Fatalf("unexpected branch returned: want %d: got %d", integrationRepoId, branch.Repository.Id)
	}
}

func TestBranchesService_Integration_FindByRepoSlug(t *testing.T) {
	branch, res, err := integrationClient.Branches.FindByRepoSlug(context.TODO(), integrationRepoSlug, "master")

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if branch.Name != "master" {
		t.Fatalf("unexpected branch returned: want %s: got %s", "master", branch.Name)
	}

	if branch.Repository.Slug != integrationRepoSlug {
		t.Fatalf("unexpected branch returned: want %s: got %s", integrationRepoSlug, branch.Repository.Slug)
	}
}

func TestBranchesService_Integration_ListByRepoId(t *testing.T) {
	cases := []*ListBranchesOption{
		{},
		{Limit: 1},
		{SortBy: "id"},
		{Offset: 0},
	}

	for i, opt := range cases {
		branches, res, err := integrationClient.Branches.ListByRepoId(context.TODO(), integrationRepoId, opt)

		if err != nil {
			t.Fatalf("#%d unexpected error occured: %s", i, err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("#%d invalid http status: %s", i, res.Status)
		}

		if len(branches) == 0 {
			t.Fatalf("#%d returned empty branches", i)
		}
	}
}

func TestBranchesService_Integration_ListByRepoSlug(t *testing.T) {
	cases := []*ListBranchesOption{
		{},
		{Limit: 1},
		{SortBy: "id"},
		{Offset: 0},
	}

	for i, opt := range cases {
		branches, res, err := integrationClient.Branches.ListByRepoSlug(context.TODO(), integrationRepoSlug, opt)

		if err != nil {
			t.Fatalf("#%d unexpected error occured: %s", i, err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("#%d invalid http status: %s", i, res.Status)
		}

		if len(branches) == 0 {
			t.Fatalf("#%d returned empty branches", i)
		}
	}
}
