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

func TestCachesService_ListByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/caches", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"caches": [{"branch":"master","match":"test"}]}`)
	})

	caches, _, err := client.Caches.ListByRepoId(context.Background(), testRepoId)

	if err != nil {
		t.Errorf("Caches.FindByRepoId returned error: %v", err)
	}

	want := []*Cache{{Branch: "master", Match: "test"}}
	if !reflect.DeepEqual(caches, want) {
		t.Errorf("Caches.FindByRepoId returned %+v, want %+v", caches, want)
	}
}

func TestCachesService_ListByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/caches", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"caches": [{"branch":"master","match":"test"}]}`)
	})

	caches, _, err := client.Caches.ListByRepoSlug(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("Caches.FindByRepoSlug returned error: %v", err)
	}

	want := []*Cache{{Branch: "master", Match: "test"}}
	if !reflect.DeepEqual(caches, want) {
		t.Errorf("Caches.FindByRepoSlug returned %+v, want %+v", caches, want)
	}
}

func TestCachesService_DeleteByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/caches", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, `{"caches": [{"branch":"master","match":"test"}]}`)
	})

	caches, _, err := client.Caches.DeleteByRepoId(context.Background(), testRepoId)

	if err != nil {
		t.Errorf("Caches.DeleteByRepoId returned error: %v", err)
	}

	want := []*Cache{{Branch: "master", Match: "test"}}
	if !reflect.DeepEqual(caches, want) {
		t.Errorf("Caches.DeleteByRepoId returned %+v, want %+v", caches, want)
	}
}

func TestCachesService_DeleteByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/caches", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, `{"caches": [{"branch":"master","match":"test"}]}`)
	})

	caches, _, err := client.Caches.DeleteByRepoSlug(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("Caches.DeleteByRepoSlug returned error: %v", err)
	}

	want := []*Cache{{Branch: "master", Match: "test"}}
	if !reflect.DeepEqual(caches, want) {
		t.Errorf("Caches.DeleteByRepoSlug returned %+v, want %+v", caches, want)
	}
}
