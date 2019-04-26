// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
)

const integrationOrgId = 111

func TestOrganizationsService_Integration_Find(t *testing.T) {
	opt := OrganizationOption{Include: []string{"organization.repositories"}}
	org, res, err := integrationClient.Organizations.Find(context.TODO(), integrationOrgId, &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if got, want := org.Id, uint(integrationOrgId); got != want {
		t.Fatalf("invalid org id: want: %d, got: %d", want, got)
	}

	if r := org.Repositories[0]; !r.IsStandard() {
		t.Fatal("repository is not in a standard representation")
	}
}

func TestOrganizationsService_Integration_List(t *testing.T) {
	opt := OrganizationsOption{Limit: 50, Offset: 50, SortBy: "id", Include: []string{"organization.repositories"}}
	_, res, err := integrationClient.Organizations.List(context.TODO(), &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}
