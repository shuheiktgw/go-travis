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
	"time"
)

const integrationEnvVarId = "88ee9d56-62bb-4093-a278-0c5cfd1e5cd5"

func TestEnvVarsService_Integration_FindByRepoId(t *testing.T) {
	envVar, res, err := integrationClient.EnvVars.FindByRepoId(context.TODO(), integrationRepoId, integrationEnvVarId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if *envVar.Id != integrationEnvVarId {
		t.Fatalf("unexpected env var id returned: want %s got %s", integrationEnvVarId, envVar.Id)
	}
}

func TestEnvVarsService_Integration_FindByRepoSlug(t *testing.T) {
	envVar, res, err := integrationClient.EnvVars.FindByRepoSlug(context.TODO(), integrationRepoSlug, integrationEnvVarId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if *envVar.Id != integrationEnvVarId {
		t.Fatalf("unexpected env var id returned: want %s got %s", integrationEnvVarId, envVar.Id)
	}
}

func TestEnvVarsService_Integration_ListByRepoId(t *testing.T) {
	vars, res, err := integrationClient.EnvVars.ListByRepoId(context.TODO(), integrationRepoId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if len(vars) == 0 {
		t.Fatal("env vars are not supposed to be empty")
	}
}

func TestEnvVarsService_Integration_ListByRepoSlug(t *testing.T) {
	vars, res, err := integrationClient.EnvVars.ListByRepoSlug(context.TODO(), integrationRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if len(vars) == 0 {
		t.Fatal("env vars are not supposed to be empty")
	}
}

func TestEnvVarsService_Integration_CreateAndUpdateAndDeleteEnvVarByRepoId(t *testing.T) {
	// Create
	body := EnvVarBody{Name: "TEST", Value: "test", Public: true}
	envVar, res, err := integrationClient.EnvVars.CreateByRepoId(context.TODO(), integrationRepoId, &body)

	if err != nil {
		t.Fatalf("EnvVars.CreateByRepoId returned unexpected error: %s", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("EnvVars.CreateByRepoId returned invalid http status: %s", res.Status)
	}

	if *envVar.Name != body.Name || *envVar.Value != body.Value || *envVar.Public != body.Public {
		t.Fatalf("EnvVars.CreateByRepoId returned invalid EnvVar: %v", envVar)
	}

	// Be nice to the API
	time.Sleep(2 * time.Second)

	// Update
	body = EnvVarBody{Name: "NEW_TEST", Value: "new_test", Public: false}
	envVar, res, err = integrationClient.EnvVars.UpdateByRepoId(context.TODO(), integrationRepoId, *envVar.Id, &body)

	if err != nil {
		t.Fatalf("EnvVars.UpdateByRepoId returned unexpected error: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("EnvVars.UpdateByRepoId returned invalid http status: %s", res.Status)
	}

	if *envVar.Name != body.Name || envVar.Value != nil || *envVar.Public != body.Public {
		t.Fatalf("EnvVars.UpdateByRepoId returned invalid EnvVar: %v", envVar)
	}

	// Be nice to the API
	time.Sleep(2 * time.Second)

	// Delete
	res, err = integrationClient.EnvVars.DeleteByRepoId(context.TODO(), integrationRepoId, *envVar.Id)

	if err != nil {
		t.Fatalf("EnvVars.DeleteByRepoId returned unexpected error: %s", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("EnvVars.DeleteByRepoId returned invalid http status: %s", res.Status)
	}
}

func TestEnvVarsService_Integration_CreateAndUpdateAndDeleteEnvVarByRepoSlug(t *testing.T) {
	// Create
	body := EnvVarBody{Name: "TEST", Value: "test", Public: true}
	envVar, res, err := integrationClient.EnvVars.CreateByRepoSlug(context.TODO(), integrationRepoSlug, &body)

	if err != nil {
		t.Fatalf("EnvVars.CreateByRepoSlug returned unexpected error: %s", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("EnvVars.CreateByRepoSlug returned invalid http status: %s", res.Status)
	}

	if *envVar.Name != body.Name || *envVar.Value != body.Value || *envVar.Public != body.Public {
		t.Fatalf("EnvVars.CreateByRepoSlug returned invalid EnvVar: %v", envVar)
	}

	// Be nice to the API
	time.Sleep(2 * time.Second)

	// Update
	body = EnvVarBody{Name: "NEW_TEST", Value: "new_test", Public: false}
	envVar, res, err = integrationClient.EnvVars.UpdateByRepoSlug(context.TODO(), integrationRepoSlug, *envVar.Id, &body)

	if err != nil {
		t.Fatalf("EnvVar.UpdateByRepoSlug returned unexpected error: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("EnvVar.UpdateByRepoSlug returned invalid http status: %s", res.Status)
	}

	if *envVar.Name != body.Name || envVar.Value != nil || *envVar.Public != body.Public {
		t.Fatalf("EnvVars.UpdateByRepoSlug returned invalid EnvVar: %v", envVar)
	}

	// Be nice to the API
	time.Sleep(2 * time.Second)

	// Delete
	res, err = integrationClient.EnvVars.DeleteByRepoSlug(context.TODO(), integrationRepoSlug, *envVar.Id)

	if err != nil {
		t.Fatalf("EnvVars.DeleteByRepoSlug returned unexpected error: %s", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("EnvVars.DeleteByRepoSlug returned invalid http status: %s", res.Status)
	}
}
