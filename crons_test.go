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

func TestCronsService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/crons", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"limit": "5", "offset": "10"})
		fmt.Fprint(w, `{"crons":[{"id":12345,"interval":"weekly","dont_run_if_recent_build_exists":true,"active":true}]}`)
	})

	opt := CronsOption{Limit: 5, Offset: 10}
	crons, _, err := client.Crons.FindByRepoId(context.Background(), testRepoId, &opt)

	if err != nil {
		t.Errorf("Crons.FindByRepoId returned error: %v", err)
	}

	want := []Cron{{Id: testCronId, Interval: CronIntervalWeekly, DontRunIfRecentBuildExists: true, Active: true}}
	if !reflect.DeepEqual(crons, want) {
		t.Errorf("Crons.FindByRepoId returned %+v, want %+v", crons, want)
	}
}
func TestCronsService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/crons", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"limit": "5", "offset": "10"})
		fmt.Fprint(w, `{"crons":[{"id":12345,"interval":"weekly","dont_run_if_recent_build_exists":true,"active":true}]}`)
	})

	opt := CronsOption{Limit: 5, Offset: 10}
	crons, _, err := client.Crons.FindByRepoSlug(context.Background(), testRepoSlug, &opt)

	if err != nil {
		t.Errorf("Crons.FindByRepoSlug returned error: %v", err)
	}

	want := []Cron{{Id: testCronId, Interval: CronIntervalWeekly, DontRunIfRecentBuildExists: true, Active: true}}
	if !reflect.DeepEqual(crons, want) {
		t.Errorf("Crons.FindByRepoSlug returned %+v, want %+v", crons, want)
	}
}
