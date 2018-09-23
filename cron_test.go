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

const testCronId = 12345

func TestCronService_CreateByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/branch/%s/cron", testRepoId, "master"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"cron.interval":"weekly","cron.dont_run_if_recent_build_exists":true}`+"\n")
		fmt.Fprint(w, `{"id":12345,"interval":"weekly","dont_run_if_recent_build_exists":true,"active":true}`)
	})

	opt := CronOption{Interval: CronIntervalWeekly, DontRunIfRecentBuildExists: true}
	cron, _, err := client.Cron.CreateByRepoId(context.Background(), testRepoId, "master", &opt)

	if err != nil {
		t.Errorf("Cron.CreateByRepoId returned error: %v", err)
	}

	want := &Cron{Id: testCronId, Interval: CronIntervalWeekly, DontRunIfRecentBuildExists: true, Active: true}
	if !reflect.DeepEqual(cron, want) {
		t.Errorf("Cron.CreateByRepoId returned %+v, want %+v", cron, want)
	}
}

func TestCronService_CreateByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/branch/%s/cron", testRepoSlug, "master"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"cron.interval":"weekly","cron.dont_run_if_recent_build_exists":true}`+"\n")
		fmt.Fprint(w, `{"id":12345,"interval":"weekly","dont_run_if_recent_build_exists":true,"active":true}`)
	})

	opt := CronOption{Interval: CronIntervalWeekly, DontRunIfRecentBuildExists: true}
	cron, _, err := client.Cron.CreateByRepoSlug(context.Background(), testRepoSlug, "master", &opt)

	if err != nil {
		t.Errorf("Cron.CreateByRepoId returned error: %v", err)
	}

	want := &Cron{Id: testCronId, Interval: CronIntervalWeekly, DontRunIfRecentBuildExists: true, Active: true}
	if !reflect.DeepEqual(cron, want) {
		t.Errorf("Cron.CreateByRepoId returned %+v, want %+v", cron, want)
	}
}

func TestCronService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/cron/%d", testCronId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":12345,"interval":"weekly","dont_run_if_recent_build_exists":true,"active":true}`)
	})

	cron, _, err := client.Cron.Find(context.Background(), testCronId)

	if err != nil {
		t.Errorf("Cron.Find returned error: %v", err)
	}

	want := &Cron{Id: testCronId, Interval: CronIntervalWeekly, DontRunIfRecentBuildExists: true, Active: true}
	if !reflect.DeepEqual(cron, want) {
		t.Errorf("Cron.Find returned %+v, want %+v", cron, want)
	}
}

func TestCronService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/branch/%s/cron", testRepoId, "master"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":12345,"interval":"weekly","dont_run_if_recent_build_exists":true,"active":true}`)
	})

	cron, _, err := client.Cron.FindByRepoId(context.Background(), testRepoId, "master")

	if err != nil {
		t.Errorf("Cron.FindByRepoId returned error: %v", err)
	}

	want := &Cron{Id: testCronId, Interval: CronIntervalWeekly, DontRunIfRecentBuildExists: true, Active: true}
	if !reflect.DeepEqual(cron, want) {
		t.Errorf("Cron.FindByRepoId returned %+v, want %+v", cron, want)
	}
}

func TestCronService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/branch/%s/cron", testRepoSlug, "master"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":12345,"interval":"weekly","dont_run_if_recent_build_exists":true,"active":true}`)
	})

	cron, _, err := client.Cron.FindByRepoSlug(context.Background(), testRepoSlug, "master")

	if err != nil {
		t.Errorf("Cron.FindByRepoSlug returned error: %v", err)
	}

	want := &Cron{Id: testCronId, Interval: CronIntervalWeekly, DontRunIfRecentBuildExists: true, Active: true}
	if !reflect.DeepEqual(cron, want) {
		t.Errorf("Cron.FindByRepoSlug returned %+v, want %+v", cron, want)
	}
}

func TestCronService_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/cron/%d", testCronId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, `{}`)
	})

	_, err := client.Cron.Delete(context.Background(), testCronId)

	if err != nil {
		t.Errorf("Cron.Delete returned error: %v", err)
	}
}
