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

func TestCronsService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/cron/%d", testCronId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"include": "cron.repository"})
		fmt.Fprint(w, `{"id":12345,"interval":"weekly","dont_run_if_recent_build_exists":true,"active":true}`)
	})

	opt := CronOption{Include: []string{"cron.repository"}}
	cron, _, err := client.Crons.Find(context.Background(), testCronId, &opt)

	if err != nil {
		t.Errorf("Cron.Find returned error: %v", err)
	}

	want := &Cron{Id: Uint(testCronId), Interval: String(CronIntervalWeekly), DontRunIfRecentBuildExists: Bool(true), Active: Bool(true)}
	if !reflect.DeepEqual(cron, want) {
		t.Errorf("Cron.Find returned %+v, want %+v", cron, want)
	}
}

func TestCronsService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/branch/%s/cron", testRepoId, "master"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"include": "cron.repository"})
		fmt.Fprint(w, `{"id":12345,"interval":"weekly","dont_run_if_recent_build_exists":true,"active":true}`)
	})

	opt := CronOption{Include: []string{"cron.repository"}}
	cron, _, err := client.Crons.FindByRepoId(context.Background(), testRepoId, "master", &opt)

	if err != nil {
		t.Errorf("Cron.FindByRepoId returned error: %v", err)
	}

	want := &Cron{Id: Uint(testCronId), Interval: String(CronIntervalWeekly), DontRunIfRecentBuildExists: Bool(true), Active: Bool(true)}
	if !reflect.DeepEqual(cron, want) {
		t.Errorf("Cron.FindByRepoId returned %+v, want %+v", cron, want)
	}
}

func TestCronsService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/branch/%s/cron", testRepoSlug, "master"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"include": "cron.repository"})
		fmt.Fprint(w, `{"id":12345,"interval":"weekly","dont_run_if_recent_build_exists":true,"active":true}`)
	})

	opt := CronOption{Include: []string{"cron.repository"}}
	cron, _, err := client.Crons.FindByRepoSlug(context.Background(), testRepoSlug, "master", &opt)

	if err != nil {
		t.Errorf("Cron.FindByRepoSlug returned error: %v", err)
	}

	want := &Cron{Id: Uint(testCronId), Interval: String(CronIntervalWeekly), DontRunIfRecentBuildExists: Bool(true), Active: Bool(true)}
	if !reflect.DeepEqual(cron, want) {
		t.Errorf("Cron.FindByRepoSlug returned %+v, want %+v", cron, want)
	}
}

func TestCronsService_ListByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/crons", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"limit": "5", "offset": "10"})
		fmt.Fprint(w, `{"crons":[{"id":12345,"interval":"weekly","dont_run_if_recent_build_exists":true,"active":true}]}`)
	})

	opt := CronsOption{Limit: 5, Offset: 10}
	crons, _, err := client.Crons.ListByRepoId(context.Background(), testRepoId, &opt)

	if err != nil {
		t.Errorf("Crons.FindByRepoId returned error: %v", err)
	}

	want := []*Cron{{Id: Uint(testCronId), Interval: String(CronIntervalWeekly), DontRunIfRecentBuildExists: Bool(true), Active: Bool(true)}}
	if !reflect.DeepEqual(crons, want) {
		t.Errorf("Crons.FindByRepoId returned %+v, want %+v", crons, want)
	}
}

func TestCronsService_ListByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/crons", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"limit": "5", "offset": "10"})
		fmt.Fprint(w, `{"crons":[{"id":12345,"interval":"weekly","dont_run_if_recent_build_exists":true,"active":true}]}`)
	})

	opt := CronsOption{Limit: 5, Offset: 10}
	crons, _, err := client.Crons.ListByRepoSlug(context.Background(), testRepoSlug, &opt)

	if err != nil {
		t.Errorf("Crons.FindByRepoSlug returned error: %v", err)
	}

	want := []*Cron{{Id: Uint(testCronId), Interval: String(CronIntervalWeekly), DontRunIfRecentBuildExists: Bool(true), Active: Bool(true)}}
	if !reflect.DeepEqual(crons, want) {
		t.Errorf("Crons.FindByRepoSlug returned %+v, want %+v", crons, want)
	}
}

func TestCronsService_CreateByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/branch/%s/cron", testRepoId, "master"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"cron.interval":"weekly","cron.dont_run_if_recent_build_exists":true}`+"\n")
		fmt.Fprint(w, `{"id":12345,"interval":"weekly","dont_run_if_recent_build_exists":true,"active":true}`)
	})

	opt := CronBody{Interval: CronIntervalWeekly, DontRunIfRecentBuildExists: true}
	cron, _, err := client.Crons.CreateByRepoId(context.Background(), testRepoId, "master", &opt)

	if err != nil {
		t.Errorf("Cron.CreateByRepoId returned error: %v", err)
	}

	want := &Cron{Id: Uint(testCronId), Interval: String(CronIntervalWeekly), DontRunIfRecentBuildExists: Bool(true), Active: Bool(true)}
	if !reflect.DeepEqual(cron, want) {
		t.Errorf("Cron.CreateByRepoId returned %+v, want %+v", cron, want)
	}
}

func TestCronsService_CreateByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/branch/%s/cron", testRepoSlug, "master"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"cron.interval":"weekly","cron.dont_run_if_recent_build_exists":true}`+"\n")
		fmt.Fprint(w, `{"id":12345,"interval":"weekly","dont_run_if_recent_build_exists":true,"active":true}`)
	})

	opt := CronBody{Interval: CronIntervalWeekly, DontRunIfRecentBuildExists: true}
	cron, _, err := client.Crons.CreateByRepoSlug(context.Background(), testRepoSlug, "master", &opt)

	if err != nil {
		t.Errorf("Cron.CreateByRepoId returned error: %v", err)
	}

	want := &Cron{Id: Uint(testCronId), Interval: String(CronIntervalWeekly), DontRunIfRecentBuildExists: Bool(true), Active: Bool(true)}
	if !reflect.DeepEqual(cron, want) {
		t.Errorf("Cron.CreateByRepoId returned %+v, want %+v", cron, want)
	}
}

func TestCronsService_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/cron/%d", testCronId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, `{}`)
	})

	_, err := client.Crons.Delete(context.Background(), testCronId)

	if err != nil {
		t.Errorf("Cron.Delete returned error: %v", err)
	}
}
