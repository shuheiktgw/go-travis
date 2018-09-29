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

func TestRepositoriesService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}`)
	})

	repo, _, err := client.Repositories.Find(context.Background(), testRepoSlug)

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
