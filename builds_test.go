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

const testBuildId = 1

func TestBuildsService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/build/%d", testBuildId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1,"number":"1","state":"created","duration":10}`)
	})

	build, _, err := client.Builds.Find(context.Background(), testBuildId)

	if err != nil {
		t.Errorf("Build.Find returned error: %v", err)
	}

	want := &Build{Id: testBuildId, Number: "1", State: BuildStateCreated, Duration: 10}
	if !reflect.DeepEqual(build, want) {
		t.Errorf("Build.Find returned %+v, want %+v", build, want)
	}
}

func TestBuildsService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/builds", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"limit": "50", "sort_by": "id"})
		fmt.Fprint(w, `{"builds": [{"id":1,"number":"1","state":"created","duration":10}]}`)
	})

	builds, _, err := client.Builds.List(context.Background(), &BuildsOption{Limit: 50, SortBy: "id"})

	if err != nil {
		t.Errorf("Builds.Find returned error: %v", err)
	}

	want := []Build{{Id: testBuildId, Number: "1", State: BuildStateCreated, Duration: 10}}
	if !reflect.DeepEqual(builds, want) {
		t.Errorf("Builds.Find returned %+v, want %+v", builds, want)
	}
}

func TestBuildsService_ListByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/builds", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"limit": "50", "sort_by": "id"})
		fmt.Fprint(w, `{"builds": [{"id":1,"number":"1","state":"created","duration":10}]}`)
	})

	builds, _, err := client.Builds.ListByRepoId(context.Background(), testRepoId, &BuildsByRepositoryOption{Limit: 50, SortBy: "id"})

	if err != nil {
		t.Errorf("Builds.FindByRepoId returned error: %v", err)
	}

	want := []Build{{Id: testBuildId, Number: "1", State: BuildStateCreated, Duration: 10}}
	if !reflect.DeepEqual(builds, want) {
		t.Errorf("Builds.FindByRepoId returned %+v, want %+v", builds, want)
	}
}

func TestBuildsService_ListByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/builds", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"limit": "50", "sort_by": "id"})
		fmt.Fprint(w, `{"builds": [{"id":1,"number":"1","state":"created","duration":10}]}`)
	})

	builds, _, err := client.Builds.ListByRepoSlug(context.Background(), testRepoSlug, &BuildsByRepositoryOption{Limit: 50, SortBy: "id"})

	if err != nil {
		t.Errorf("Builds.FindByRepoSlug returned error: %v", err)
	}

	want := []Build{{Id: testBuildId, Number: "1", State: BuildStateCreated, Duration: 10}}
	if !reflect.DeepEqual(builds, want) {
		t.Errorf("Builds.FindByRepoSlug returned %+v, want %+v", builds, want)
	}
}

func TestBuildsService_Cancel(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/build/%d/cancel", testBuildId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"build":{"id":1,"number":"1","state":"created","duration":10}}`)
	})

	build, _, err := client.Builds.Cancel(context.Background(), testBuildId)

	if err != nil {
		t.Errorf("Build.Cancel returned error: %v", err)
	}

	want := &MinimalBuild{Id: testBuildId, Number: "1", State: BuildStateCreated, Duration: 10}
	if !reflect.DeepEqual(build, want) {
		t.Errorf("Build.Cancel returned %+v, want %+v", build, want)
	}
}

func TestBuildsService_Restart(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/build/%d/restart", testBuildId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"build":{"id":1,"number":"1","state":"created","duration":10}}`)
	})

	build, _, err := client.Builds.Restart(context.Background(), testBuildId)

	if err != nil {
		t.Errorf("Build.Restart returned error: %v", err)
	}

	want := &MinimalBuild{Id: testBuildId, Number: "1", State: BuildStateCreated, Duration: 10}
	if !reflect.DeepEqual(build, want) {
		t.Errorf("Build.Restart returned %+v, want %+v", build, want)
	}
}
