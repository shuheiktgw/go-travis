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

func TestRepositoryService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}`)
	})

	repo, _, err := client.Repository.Find(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("Repository.Find returned error: %v", err)
	}

	want := &Repository{Id: 1, Name: "go-travis-test", Slug: "shuheiktgw/go-travis-test"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("Repository.Find returned %+v, want %+v", repo, want)
	}
}

func TestRepositoryService_Activate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/activate", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}`)
	})

	repo, _, err := client.Repository.Activate(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("Repository.Activate returned error: %v", err)
	}

	want := &Repository{Id: 1, Name: "go-travis-test", Slug: "shuheiktgw/go-travis-test"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("Repository.Activate returned %+v, want %+v", repo, want)
	}
}

func TestRepositoryService_Deactivate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/deactivate", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}`)
	})

	repo, _, err := client.Repository.Deactivate(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("Repository.Deactivate returned error: %v", err)
	}

	want := &Repository{Id: 1, Name: "go-travis-test", Slug: "shuheiktgw/go-travis-test"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("Repository.Deactivate returned %+v, want %+v", repo, want)
	}
}

func TestRepositoryService_Star(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/star", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}`)
	})

	repo, _, err := client.Repository.Star(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("Repository.Star returned error: %v", err)
	}

	want := &Repository{Id: 1, Name: "go-travis-test", Slug: "shuheiktgw/go-travis-test"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("Repository.Star returned %+v, want %+v", repo, want)
	}
}

func TestRepositoryService_Unstar(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/unstar", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"id":1,"name":"go-travis-test","slug":"shuheiktgw/go-travis-test"}`)
	})

	repo, _, err := client.Repository.Unstar(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("Repository.Unstar returned error: %v", err)
	}

	want := &Repository{Id: 1, Name: "go-travis-test", Slug: "shuheiktgw/go-travis-test"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("Repository.Unstar returned %+v, want %+v", repo, want)
	}
}
