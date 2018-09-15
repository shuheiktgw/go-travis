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

func TestJobService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/job/%d", testJobId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1,"allow_failure":true,"number":"1","state":"created"}`)
	})

	job, _, err := client.Job.Find(context.Background(), testJobId)

	if err != nil {
		t.Errorf("Job.Find returned error: %v", err)
	}

	want := &Job{Id: testJobId, AllowFailure: true, Number: "1", State: JobStatusCreated}
	if !reflect.DeepEqual(job, want) {
		t.Errorf("Job.Find returned %+v, want %+v", job, want)
	}
}

func TestJobService_Cancel(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/job/%d/cancel", testJobId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"job":{"id":1}}`)
	})

	job, _, err := client.Job.Cancel(context.Background(), testJobId)

	if err != nil {
		t.Errorf("Job.Cancel returned error: %v", err)
	}

	want := &MinimalJob{Id: testJobId}
	if !reflect.DeepEqual(job, want) {
		t.Errorf("Job.Cancel returned %+v, want %+v", job, want)
	}
}

func TestJobService_Restart(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/job/%d/restart", testJobId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"job":{"id":1}}`)
	})

	job, _, err := client.Job.Restart(context.Background(), testJobId)

	if err != nil {
		t.Errorf("Job.Restart returned error: %v", err)
	}

	want := &MinimalJob{Id: testJobId}
	if !reflect.DeepEqual(job, want) {
		t.Errorf("Job.Restart returned %+v, want %+v", job, want)
	}
}

func TestJobService_Debug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/job/%d/debug", testJobId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"job":{"id":1}}`)
	})

	job, _, err := client.Job.Debug(context.Background(), testJobId)

	if err != nil {
		t.Errorf("Job.Debug returned error: %v", err)
	}

	want := &MinimalJob{Id: testJobId}
	if !reflect.DeepEqual(job, want) {
		t.Errorf("Job.Debug returned %+v, want %+v", job, want)
	}
}
