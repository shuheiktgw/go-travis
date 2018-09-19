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

func TestCachesService_Integration_FindByRepoId(t *testing.T) {
	_, res, err := integrationClient.Caches.FindByRepoId(context.TODO(), integrationRepoId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}

func TestCachesService_Integration_FindByRepoSlug(t *testing.T) {
	_, res, err := integrationClient.Caches.FindByRepoSlug(context.TODO(), integrationRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}

func TestCachesService_Integration_DeleteByRepoId(t *testing.T) {
	_, res, err := integrationClient.Caches.DeleteByRepoId(context.TODO(), integrationRepoId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}

func TestCachesService_Integration_DeleteByRepoSlug(t *testing.T) {
	_, res, err := integrationClient.Caches.DeleteByRepoSlug(context.TODO(), integrationRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}
