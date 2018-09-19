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

const integrationEnvVarId = "88ee9d56-62bb-4093-a278-0c5cfd1e5cd5"

func TestEnvVarService_Integration_FindByRepoId(t *testing.T) {
	envVar, res, err := integrationClient.EnvVar.FindByRepoId(context.TODO(), integrationRepoId, integrationEnvVarId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if envVar.Id != integrationEnvVarId {
		t.Fatalf("unexpected env var id returned: want %s got %s", integrationEnvVarId, envVar.Id)
	}
}

func TestEnvVarService_Integration_FindByRepoSlug(t *testing.T) {
	envVar, res, err := integrationClient.EnvVar.FindByRepoSlug(context.TODO(), integrationRepoSlug, integrationEnvVarId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if envVar.Id != integrationEnvVarId {
		t.Fatalf("unexpected env var id returned: want %s got %s", integrationEnvVarId, envVar.Id)
	}
}
