// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

package travis

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

const integrationOrgId = 111

func TestOrganizationsService_Integration_Find(t *testing.T) {
	org, res, err := integrationClient.Organizations.Find(context.TODO(), integrationOrgId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	want := &Organization{
		Id:        integrationOrgId,
		Login:     "RubyMoney",
		Name:      "RubyMoney",
		GithubId:  351550,
		AvatarUrl: "https://avatars1.githubusercontent.com/u/351550",
		Education: false,
		Metadata: Metadata{
			Type:           "organization",
			Href:           "/org/111",
			Representation: "standard",
			Permissions:    Permissions{"read": true, "sync": false},
		},
	}

	if !reflect.DeepEqual(org, want) {
		t.Errorf("Organizations.Find returned %+v, want %+v", org, want)
	}
}

func TestOrganizationsService_Integration_List(t *testing.T) {
	opt := OrganizationsOption{Limit: 50, Offset: 50, SortBy: "id"}
	_, res, err := integrationClient.Organizations.List(context.TODO(), &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}
