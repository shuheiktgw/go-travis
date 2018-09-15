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

func TestJobsService_FindByBuild(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/build/%d/jobs", testBuildId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"jobs":[{"id":1,"allow_failure":true,"number":"1","state":"created"}]}`)
	})

	job, _, err := client.Jobs.FindByBuild(context.Background(), testBuildId)

	if err != nil {
		t.Errorf("Jobs.FindByBuild returned error: %v", err)
	}

	want := []Job{{Id: testJobId, AllowFailure: true, Number: "1", State: JobStatusCreated}}
	if !reflect.DeepEqual(job, want) {
		t.Errorf("Jobs.FindByBuild returned %+v, want %+v", job, want)
	}
}

func TestJobsService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"limit": "50", "state[]": "created"})
		fmt.Fprint(w, `{"jobs":[{"id":1,"allow_failure":true,"number":"1","state":"created"}]}`)
	})

	job, _, err := client.Jobs.Find(context.Background(), &JobsOption{Limit: 50, State: []string{"created"}})

	if err != nil {
		t.Errorf("Jobs.Find returned error: %v", err)
	}

	want := []Job{{Id: testJobId, AllowFailure: true, Number: "1", State: JobStatusCreated}}
	if !reflect.DeepEqual(job, want) {
		t.Errorf("Jobs.Find returned %+v, want %+v", job, want)
	}
}
