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

func TestOwnerService_FindByLogin(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/owner/shuheiktgw", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1,"login":"shuheiktgw","github_id":1}`)
	})

	owner, _, err := client.Owner.FindByLogin(context.Background(), "shuheiktgw")

	if err != nil {
		t.Errorf("Owner.FindByLogin returned error: %v", err)
	}

	want := &Owner{Id: 1, Login: "shuheiktgw", GitHubId: 1}
	if !reflect.DeepEqual(owner, want) {
		t.Errorf("Owner.FindByLogin returned %+v, want %+v", owner, want)
	}
}

func TestOwnerService_FindByGitHubId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/owner/github_id/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1,"login":"shuheiktgw","github_id":1}`)
	})

	owner, _, err := client.Owner.FindByGitHubId(context.Background(), 1)

	if err != nil {
		t.Errorf("Owner.FindByGitHubId returned error: %v", err)
	}

	want := &Owner{Id: 1, Login: "shuheiktgw", GitHubId: 1}
	if !reflect.DeepEqual(owner, want) {
		t.Errorf("Owner.FindByGitHubId returned %+v, want %+v", owner, want)
	}
}
