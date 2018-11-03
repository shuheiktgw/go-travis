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

const testJobId = 1

func TestJobsService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/job/%d", testJobId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1,"allow_failure":true,"number":"1","state":"created"}`)
	})

	job, _, err := client.Jobs.Find(context.Background(), testJobId)

	if err != nil {
		t.Errorf("Job.Find returned error: %v", err)
	}

	want := &Job{Id: testJobId, AllowFailure: true, Number: "1", State: JobStatusCreated}
	if !reflect.DeepEqual(job, want) {
		t.Errorf("Job.Find returned %+v, want %+v", job, want)
	}
}

func TestJobsService_ListByBuild(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/build/%d/jobs", testBuildId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"jobs":[{"id":1,"allow_failure":true,"number":"1","state":"created"}]}`)
	})

	job, _, err := client.Jobs.ListByBuild(context.Background(), testBuildId)

	if err != nil {
		t.Errorf("Jobs.ListByBuild returned error: %v", err)
	}

	want := []Job{{Id: testJobId, AllowFailure: true, Number: "1", State: JobStatusCreated}}
	if !reflect.DeepEqual(job, want) {
		t.Errorf("Jobs.ListByBuild returned %+v, want %+v", job, want)
	}
}

func TestJobsService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"limit": "50", "state": "created,queued"})
		fmt.Fprint(w, `{"jobs":[{"id":1,"allow_failure":true,"number":"1","state":"created"}]}`)
	})

	job, _, err := client.Jobs.List(context.Background(), &JobsOption{Limit: 50, State: []string{"created", "queued"}})

	if err != nil {
		t.Errorf("Jobs.List returned error: %v", err)
	}

	want := []Job{{Id: testJobId, AllowFailure: true, Number: "1", State: JobStatusCreated}}
	if !reflect.DeepEqual(job, want) {
		t.Errorf("Jobs.List returned %+v, want %+v", job, want)
	}
}

func TestJobsService_Cancel(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/job/%d/cancel", testJobId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"job":{"id":1}}`)
	})

	job, _, err := client.Jobs.Cancel(context.Background(), testJobId)

	if err != nil {
		t.Errorf("Job.Cancel returned error: %v", err)
	}

	want := &MinimalJob{Id: testJobId}
	if !reflect.DeepEqual(job, want) {
		t.Errorf("Job.Cancel returned %+v, want %+v", job, want)
	}
}

func TestJobsService_Restart(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/job/%d/restart", testJobId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"job":{"id":1}}`)
	})

	job, _, err := client.Jobs.Restart(context.Background(), testJobId)

	if err != nil {
		t.Errorf("Job.Restart returned error: %v", err)
	}

	want := &MinimalJob{Id: testJobId}
	if !reflect.DeepEqual(job, want) {
		t.Errorf("Job.Restart returned %+v, want %+v", job, want)
	}
}

func TestJobsService_Debug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/job/%d/debug", testJobId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"job":{"id":1}}`)
	})

	job, _, err := client.Jobs.Debug(context.Background(), testJobId)

	if err != nil {
		t.Errorf("Job.Debug returned error: %v", err)
	}

	want := &MinimalJob{Id: testJobId}
	if !reflect.DeepEqual(job, want) {
		t.Errorf("Job.Debug returned %+v, want %+v", job, want)
	}
}
