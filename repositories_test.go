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

func TestRepositoriesService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"active_on_org": "true", "starred": "true", "private": "true"})
		fmt.Fprint(w, `{"repositories": [{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}]}`)
	})

	opt := RepositoriesOption{ActiveOnOrg: true, Starred: true, Private: true}
	repos, _, err := client.Repositories.List(context.Background(), &opt)

	if err != nil {
		t.Errorf("Repository.List returned error: %v", err)
	}

	want := &Repository{Id: 1, Name: "go-travis-test", Slug: "shuheiktgw/go-travis-test"}
	if !reflect.DeepEqual(repos[0], want) {
		t.Errorf("Repository.List returned %+v, want %+v", repos[0], want)
	}
}

func TestRepositoriesService_ListByOwner(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	owner := "shuheiktgw"
	mux.HandleFunc(fmt.Sprintf("/owner/%s/repos", owner), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"active_on_org": "true", "starred": "true", "private": "true"})
		fmt.Fprint(w, `{"repositories": [{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}]}`)
	})

	opt := RepositoriesOption{ActiveOnOrg: true, Starred: true, Private: true}
	repos, _, err := client.Repositories.ListByOwner(context.Background(), owner, &opt)

	if err != nil {
		t.Errorf("Repository.ListByOwner returned error: %v", err)
	}

	want := &Repository{Id: 1, Name: "go-travis-test", Slug: "shuheiktgw/go-travis-test"}
	if !reflect.DeepEqual(repos[0], want) {
		t.Errorf("Repository.ListByOwner returned %+v, want %+v", repos[0], want)
	}
}

func TestRepositoriesService_ListByGitHubId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	var id uint = 9999
	mux.HandleFunc(fmt.Sprintf("/owner/github_id/%d/repos", id), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"active_on_org": "true", "starred": "true", "private": "true"})
		fmt.Fprint(w, `{"repositories": [{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}]}`)
	})

	opt := RepositoriesOption{ActiveOnOrg: true, Starred: true, Private: true}
	repos, _, err := client.Repositories.ListByGitHubId(context.Background(), id, &opt)

	if err != nil {
		t.Errorf("Repository.ListByOwner returned error: %v", err)
	}

	want := &Repository{Id: 1, Name: "go-travis-test", Slug: "shuheiktgw/go-travis-test"}
	if !reflect.DeepEqual(repos[0], want) {
		t.Errorf("Repository.ListByOwner returned %+v, want %+v", repos[0], want)
	}
}

func TestRepositoriesService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"include": "repository.default_branch"})
		fmt.Fprint(w, `{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}`)
	})

	opt := RepositoryOption{Include: []string{"repository.default_branch"}}
	repo, _, err := client.Repositories.Find(context.Background(), testRepoSlug, &opt)

	if err != nil {
		t.Errorf("Repository.Find returned error: %v", err)
	}

	want := &Repository{Id: 1, Name: "go-travis-test", Slug: "shuheiktgw/go-travis-test"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("Repository.Find returned %+v, want %+v", repo, want)
	}
}

func TestRepositoriesService_Activate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/activate", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}`)
	})

	repo, _, err := client.Repositories.Activate(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("Repository.Activate returned error: %v", err)
	}

	want := &Repository{Id: 1, Name: "go-travis-test", Slug: "shuheiktgw/go-travis-test"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("Repository.Activate returned %+v, want %+v", repo, want)
	}
}

func TestRepositoriesService_Deactivate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/deactivate", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}`)
	})

	repo, _, err := client.Repositories.Deactivate(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("Repository.Deactivate returned error: %v", err)
	}

	want := &Repository{Id: 1, Name: "go-travis-test", Slug: "shuheiktgw/go-travis-test"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("Repository.Deactivate returned %+v, want %+v", repo, want)
	}
}

func TestRepositoriesService_Migrate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/migrate", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}`)
	})

	repo, _, err := client.Repositories.Migrate(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("Repository.Migrate returned error: %v", err)
	}

	want := &Repository{Id: 1, Name: "go-travis-test", Slug: "shuheiktgw/go-travis-test"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("Repository.Migrate returned %+v, want %+v", repo, want)
	}
}

func TestRepositoriesService_Star(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/star", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}`)
	})

	repo, _, err := client.Repositories.Star(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("Repository.Star returned error: %v", err)
	}

	want := &Repository{Id: 1, Name: "go-travis-test", Slug: "shuheiktgw/go-travis-test"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("Repository.Star returned %+v, want %+v", repo, want)
	}
}

func TestRepositoriesService_Unstar(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/unstar", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}`)
	})

	repo, _, err := client.Repositories.Unstar(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("Repository.Unstar returned error: %v", err)
	}

	want := &Repository{Id: 1, Name: "go-travis-test", Slug: "shuheiktgw/go-travis-test"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("Repository.Unstar returned %+v, want %+v", repo, want)
	}
}
