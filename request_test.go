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

const testRequestId = 12345

func TestRequestService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/request/%d", testRepoId, testRequestId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1,"state":"processed","result":"rejected"}`)
	})

	repo, _, err := client.Request.FindByRepoId(context.Background(), testRepoId, testRequestId)

	if err != nil {
		t.Errorf("RequestService.FindByRepoId returned error: %v", err)
	}

	want := &Request{Id: 1, State: "processed", Result: "rejected"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("RequestService.FindByRepoId returned %+v, want %+v", repo, want)
	}
}

func TestRequestService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/request/%d", testRepoSlug, testRequestId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1,"state":"processed","result":"rejected"}`)
	})

	repo, _, err := client.Request.FindByRepoSlug(context.Background(), testRepoSlug, testRequestId)

	if err != nil {
		t.Errorf("RequestService.FindByRepoId returned error: %v", err)
	}

	want := &Request{Id: 1, State: "processed", Result: "rejected"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("RequestService.FindByRepoId returned %+v, want %+v", repo, want)
	}
}
