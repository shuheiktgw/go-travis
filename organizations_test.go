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

func TestOrganizationsService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/org/111", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":111,"login":"TestOrg","name":"TestOrg","github_id":12345,"avatar_url":"https:///test.com","education":false}`)
	})

	org, _, err := client.Organizations.Find(context.Background(), 111)

	if err != nil {
		t.Errorf("Organizations.Find returned error: %v", err)
	}

	want := &Organization{Id: 111, Login: "TestOrg", Name: "TestOrg", GithubId: 12345, AvatarUrl: "https:///test.com", Education: false}
	if !reflect.DeepEqual(org, want) {
		t.Errorf("Organizations.Find returned %+v, want %+v", org, want)
	}
}

func TestOrganizationsService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/orgs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"limit": "50", "offset": "50", "sort_by": "id"})
		fmt.Fprint(w, `{"organizations":[{"id":111,"login":"TestOrg","name":"TestOrg","github_id":12345,"avatar_url":"https:///test.com","education":false}]}`)
	})

	opt := OrganizationsOption{Limit: 50, Offset: 50, SortBy: "id"}
	orgs, _, err := client.Organizations.List(context.Background(), &opt)

	if err != nil {
		t.Errorf("Organizations.List returned error: %v", err)
	}

	want := []Organization{{Id: 111, Login: "TestOrg", Name: "TestOrg", GithubId: 12345, AvatarUrl: "https:///test.com", Education: false}}
	if !reflect.DeepEqual(orgs, want) {
		t.Errorf("Organizations.List returned %+v, want %+v", orgs, want)
	}
}
