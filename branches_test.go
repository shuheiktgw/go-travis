// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

var (
	testRepoId          uint = 12345
	testRepoSlug             = "shuheiktgw/go-travis-test"
	testEscapedRepoSlug      = url.QueryEscape(testRepoSlug)
)

func TestBranchesService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/branch/%s", testRepoId, "master"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"include": "branch.recent_builds"})
		fmt.Fprint(w, `{"name":"master","repository":{"id":1,"name":"test","slug":"shuheiktgw/test"},"default_branch":true,"exists_on_github":true}`)
	})

	opt := BranchOption{Include: []string{"branch.recent_builds"}}
	branch, _, err := client.Branches.FindByRepoId(context.Background(), testRepoId, "master", &opt)

	if err != nil {
		t.Errorf("Branch.FindByRepoId returned error: %v", err)
	}

	want := &Branch{Name: "master", Repository: &Repository{Id: 1, Name: "test", Slug: "shuheiktgw/test"}, DefaultBranch: true, ExistsOnGithub: true}
	if !reflect.DeepEqual(branch, want) {
		t.Errorf("Branch.FindByRepoId returned %+v, want %+v", branch, want)
	}
}

func TestBranchesService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/branch/%s", testEscapedRepoSlug, "master"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"include": "branch.recent_builds"})
		fmt.Fprint(w, `{"name":"master","repository":{"id":1,"name":"test","slug":"shuheiktgw/test"},"default_branch":true,"exists_on_github":true}`)
	})

	opt := BranchOption{Include: []string{"branch.recent_builds"}}
	branch, _, err := client.Branches.FindByRepoSlug(context.Background(), testEscapedRepoSlug, "master", &opt)

	if err != nil {
		t.Errorf("Branch.FindByRepoId returned error: %v", err)
	}

	want := &Branch{Name: "master", Repository: &Repository{Id: 1, Name: "test", Slug: "shuheiktgw/test"}, DefaultBranch: true, ExistsOnGithub: true}
	if !reflect.DeepEqual(branch, want) {
		t.Errorf("Branch.FindByRepoId returned %+v, want %+v", branch, want)
	}
}

func TestBranchesService_ListByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/branches", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"exists_on_github": "true", "limit": "50", "include": "branch.recent_builds"})
		fmt.Fprint(w, `{"branches": [{"name":"master","repository":{"id":1,"name":"test","slug":"shuheiktgw/test"},"default_branch":true,"exists_on_github":true}]}`)
	})

	opt := BranchesOption{ExistsOnGithub: true, Limit: 50, Include: []string{"branch.recent_builds"}}
	branches, _, err := client.Branches.ListByRepoId(context.Background(), testRepoId, &opt)

	if err != nil {
		t.Errorf("Branches.FindByRepoId returned error: %v", err)
	}

	want := []*Branch{{Name: "master", Repository: &Repository{Id: 1, Name: "test", Slug: "shuheiktgw/test"}, DefaultBranch: true, ExistsOnGithub: true}}
	if !reflect.DeepEqual(branches, want) {
		t.Errorf("Branchse.FindByRepoId returned %+v, want %+v", branches, want)
	}
}

func TestBranchesService_ListByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/branches", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"exists_on_github": "true", "limit": "50", "include": "branch.recent_builds"})
		fmt.Fprint(w, `{"branches": [{"name":"master","repository":{"id":1,"name":"test","slug":"shuheiktgw/test"},"default_branch":true,"exists_on_github":true}]}`)
	})

	opt := BranchesOption{ExistsOnGithub: true, Limit: 50, Include: []string{"branch.recent_builds"}}
	branches, _, err := client.Branches.ListByRepoSlug(context.Background(), testRepoSlug, &opt)

	if err != nil {
		t.Errorf("Branches.indByRepoSlug returned error: %v", err)
	}

	want := []*Branch{{Name: "master", Repository: &Repository{Id: 1, Name: "test", Slug: "shuheiktgw/test"}, DefaultBranch: true, ExistsOnGithub: true}}
	if !reflect.DeepEqual(branches, want) {
		t.Errorf("Branchse.indByRepoSlug returned %+v, want %+v", branches, want)
	}
}
