// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBranchesService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/branches", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"exists_on_github": "true", "limit": "50"})
		fmt.Fprint(w, `{"branches": [{"name":"master","repository":{"id":1,"name":"test","slug":"shuheiktgw/test"},"default_branch":true,"exists_on_github":true}]}`)
	})

	branches, _, err := client.Branches.FindByRepoId(context.Background(), testRepoId, &BranchesOption{ExistsOnGithub: true, Limit: 50})

	if err != nil {
		t.Errorf("Branches.FindByRepoId returned error: %v", err)
	}

	want := []Branch{{Name: "master", Repository: MinimalRepository{Id: 1, Name: "test", Slug: "shuheiktgw/test"}, DefaultBranch: true, ExistsOnGithub: true}}
	if !reflect.DeepEqual(branches, want) {
		t.Errorf("Branchse.FindByRepoId returned %+v, want %+v", branches, want)
	}
}

func TestBranchesService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/branches", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"exists_on_github": "true", "limit": "50"})
		fmt.Fprint(w, `{"branches": [{"name":"master","repository":{"id":1,"name":"test","slug":"shuheiktgw/test"},"default_branch":true,"exists_on_github":true}]}`)
	})

	branches, _, err := client.Branches.FindByRepoSlug(context.Background(), testRepoSlug, &BranchesOption{ExistsOnGithub: true, Limit: 50})

	if err != nil {
		t.Errorf("Branches.indByRepoSlug returned error: %v", err)
	}

	want := []Branch{{Name: "master", Repository: MinimalRepository{Id: 1, Name: "test", Slug: "shuheiktgw/test"}, DefaultBranch: true, ExistsOnGithub: true}}
	if !reflect.DeepEqual(branches, want) {
		t.Errorf("Branchse.indByRepoSlug returned %+v, want %+v", branches, want)
	}
}
