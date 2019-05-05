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

func TestRequestsService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/request/%d", testRepoId, testRequestId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"include": "request.repository"})
		fmt.Fprint(w, `{"id":1,"state":"processed","result":"rejected"}`)
	})

	opt := RequestOption{Include: []string{"request.repository"}}
	repo, _, err := client.Requests.FindByRepoId(context.Background(), testRepoId, testRequestId, &opt)

	if err != nil {
		t.Errorf("RequestService.FindByRepoId returned error: %v", err)
	}

	want := &Request{Id: 1, State: "processed", Result: "rejected"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("RequestService.FindByRepoId returned %+v, want %+v", repo, want)
	}
}

func TestRequestsService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/request/%d", testRepoSlug, testRequestId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"include": "request.repository"})
		fmt.Fprint(w, `{"id":1,"state":"processed","result":"rejected"}`)
	})

	opt := RequestOption{Include: []string{"request.repository"}}
	repo, _, err := client.Requests.FindByRepoSlug(context.Background(), testRepoSlug, testRequestId, &opt)

	if err != nil {
		t.Errorf("RequestService.FindByRepoId returned error: %v", err)
	}

	want := &Request{Id: 1, State: "processed", Result: "rejected"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("RequestService.FindByRepoId returned %+v, want %+v", repo, want)
	}
}

func TestRequestsService_ListByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/requests", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"limit": "5", "offset": "5", "include": "request.repository"})
		fmt.Fprint(w, `{"requests": [{"id":1,"state":"processed","result":"rejected"}]}`)
	})

	opt := RequestsOption{Limit: 5, Offset: 5, Include: []string{"request.repository"}}
	repos, _, err := client.Requests.ListByRepoId(context.Background(), testRepoId, &opt)

	if err != nil {
		t.Errorf("RequestsService.FindByRepoId returned error: %v", err)
	}

	want := []*Request{{Id: 1, State: "processed", Result: "rejected"}}
	if !reflect.DeepEqual(repos, want) {
		t.Errorf("RequestsService.FindByRepoId returned %+v, want %+v", repos, want)
	}
}

func TestRequestsService_ListByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/requests", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"limit": "5", "offset": "5", "include": "request.repository"})
		fmt.Fprint(w, `{"requests": [{"id":1,"state":"processed","result":"rejected"}]}`)
	})

	opt := RequestsOption{Limit: 5, Offset: 5, Include: []string{"request.repository"}}
	repos, _, err := client.Requests.ListByRepoSlug(context.Background(), testRepoSlug, &opt)

	if err != nil {
		t.Errorf("RequestsService.FindByRepoSlug returned error: %v", err)
	}

	want := []*Request{{Id: 1, State: "processed", Result: "rejected"}}
	if !reflect.DeepEqual(repos, want) {
		t.Errorf("RequestsService.FindByRepoSlug returned %+v, want %+v", repos, want)
	}
}

func TestRequestsService_CreateByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/requests", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"config":"testConfig","message":"testMessage","branch":"master","token":"testToken"}`+"\n")
		fmt.Fprint(w, `{"request": {"id":1,"message":"message!"}}`)
	})

	repo, _, err := client.Requests.CreateByRepoId(context.Background(), testRepoId, &RequestBody{Config: "testConfig", Message: "testMessage", Branch: "master", Token: "testToken"})

	if err != nil {
		t.Errorf("RequestsService.CreateByRepoId returned error: %v", err)
	}

	want := &Request{Id: 1, Message: "message!"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("RequestsService.CreateByRepoId returned %+v, want %+v", repo, want)
	}
}

func TestRequestsService_CreateByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/requests", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"config":"testConfig","message":"testMessage","branch":"master","token":"testToken"}`+"\n")
		fmt.Fprint(w, `{"request": {"id":1,"message":"message!"}}`)
	})

	repo, _, err := client.Requests.CreateByRepoSlug(context.Background(), testRepoSlug, &RequestBody{Config: "testConfig", Message: "testMessage", Branch: "master", Token: "testToken"})

	if err != nil {
		t.Errorf("RequestsService.CreateByRepoSlug returned error: %v", err)
	}

	want := &Request{Id: 1, Message: "message!"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("RequestsService.CreateByRepoSlug returned %+v, want %+v", repo, want)
	}
}
