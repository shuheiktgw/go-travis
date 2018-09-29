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

func TestRequestsService_Integration_CreateAndFindById(t *testing.T) {
	createdRequest, res, err := integrationClient.Requests.CreateByRepoId(context.TODO(), integrationRepoId, &RequestBody{Message: "test", Branch: "master"})

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	request, res, err := integrationClient.Requests.FindByRepoId(context.TODO(), integrationRepoId, createdRequest.Id)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if request.Id != createdRequest.Id {
		t.Fatalf("unexpected request is retrieved: got request id: %d, want request id: %d", request.Id, createdRequest.Id)
	}
}

func TestRequestsService_Integration_CreateAndFindBySlug(t *testing.T) {
	createdRequest, res, err := integrationClient.Requests.CreateByRepoSlug(context.TODO(), integrationRepoSlug, &RequestBody{Message: "test", Branch: "master"})

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	request, res, err := integrationClient.Requests.FindByRepoSlug(context.TODO(), integrationRepoSlug, createdRequest.Id)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if request.Id != createdRequest.Id {
		t.Fatalf("unexpected request is retrieved: got request id: %d, want request id: %d", request.Id, createdRequest.Id)
	}
}

func TestRequestsService_CreateAndListById(t *testing.T) {
	createdRequest, res, err := integrationClient.Requests.CreateByRepoId(context.TODO(), integrationRepoId, &RequestBody{Message: "test", Branch: "master"})

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	requests, res, err := integrationClient.Requests.ListByRepoId(context.TODO(), integrationRepoId, &ListRequestsOption{})

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if requests[0].Id != createdRequest.Id {
		t.Fatalf("unexpected request is retrieved: got request id: %d, want request id: %d", requests[0].Id, createdRequest.Id)
	}
}

func TestRequestsService_CreateAndListBySlug(t *testing.T) {
	createdRequest, res, err := integrationClient.Requests.CreateByRepoSlug(context.TODO(), integrationRepoSlug, &RequestBody{Message: "test", Branch: "master"})

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	requests, res, err := integrationClient.Requests.ListByRepoSlug(context.TODO(), integrationRepoSlug, &ListRequestsOption{})

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if requests[0].Id != createdRequest.Id {
		t.Fatalf("unexpected request is retrieved: got request id: %d, want request id: %d", requests[0].Id, createdRequest.Id)
	}
}
