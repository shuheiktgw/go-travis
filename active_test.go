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

const testOwner = "shuheiktgw"
const testGitHubId = 83472489

func TestActiveService_FindByOwner(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/owner/%s/active", testOwner), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"include": "active.builds"})
		fmt.Fprint(w, `{"builds": [{"id":1,"number":"1","state":"created","duration":10}]}`)
	})

	opt := ActiveOption{Include: []string{"active.builds"}}
	builds, _, err := client.Active.FindByOwner(context.Background(), testOwner, &opt)

	if err != nil {
		t.Errorf("Active.FindByOwner returned error: %v", err)
	}

	want := []*Build{{Id: Uint(testBuildId), Number: String("1"), State: String(BuildStateCreated), Duration: Uint(10)}}
	if !reflect.DeepEqual(builds, want) {
		t.Errorf("Active.FindByOwner returned %+v, want %+v", builds, want)
	}
}

func TestActiveService_FindByGitHubId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/owner/github_id/%d/active", testGitHubId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"include": "active.builds"})
		fmt.Fprint(w, `{"builds": [{"id":1,"number":"1","state":"created","duration":10}]}`)
	})

	opt := ActiveOption{Include: []string{"active.builds"}}
	builds, _, err := client.Active.FindByGitHubId(context.Background(), testGitHubId, &opt)

	if err != nil {
		t.Errorf("Active.FindByGitHubId returned error: %v", err)
	}

	want := []*Build{{Id: Uint(testBuildId), Number: String("1"), State: String(BuildStateCreated), Duration: Uint(10)}}
	if !reflect.DeepEqual(builds, want) {
		t.Errorf("Active.FindByGitHubId returned %+v, want %+v", builds, want)
	}
}
