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

func TestBetaMigrationRequestsService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/user/%d/beta_migration_requests", testUserId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"include": "beta_migration_requests.organizations"})
		fmt.Fprint(w, `{"beta_migration_requests":[{"id":1,"owner_id":2,"owner_name":"test","owner_type":"User"}]}`)
	})

	opt := BetaMigrationRequestsOption{Include: []string{"beta_migration_requests.organizations"}}
	requests, _, err := client.BetaMigrationRequests.List(context.Background(), testUserId, &opt)

	if err != nil {
		t.Errorf("BetaMigrationRequest.List returned error: %v", err)
	}

	want := []*BetaMigrationRequest{{Id: 1, OwnerId: 2, OwnerName: "test", OwnerType: "User"}}
	if !reflect.DeepEqual(requests, want) {
		t.Errorf("BetaMigrationRequest.List returned %+v, want %+v", requests, want)
	}
}

func TestBetaMigrationRequestsService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/user/%d/beta_migration_request", testUserId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"beta_migration_request.organizations":[1]}`+"\n")
		fmt.Fprint(w, `{"id":1,"owner_id":2,"owner_name":"test","owner_type":"User"}`)
	})

	rb := BetaMigrationRequestBody{OrganizationIds: []uint{1}}
	request, _, err := client.BetaMigrationRequests.Create(context.Background(), testUserId, &rb)

	if err != nil {
		t.Errorf("BetaMigrationRequests.Create returned error: %v", err)
	}

	want := &BetaMigrationRequest{Id: 1, OwnerId: 2, OwnerName: "test", OwnerType: "User"}
	if !reflect.DeepEqual(request, want) {
		t.Errorf("BetaMigrationRequests.Create returned %+v, want %+v", request, want)
	}
}
