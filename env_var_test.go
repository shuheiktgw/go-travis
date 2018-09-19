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

const testEnvVarId = "test-12345-absde"

func TestEnvVarService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/env_var/%s", testRepoId, testEnvVarId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":"test-12345-absde","name":"TEST","value":"test","public":false}`)
	})

	envVar, _, err := client.EnvVar.FindByRepoId(context.Background(), testRepoId, testEnvVarId)

	if err != nil {
		t.Errorf("EnvVar.FindByRepoId returned error: %v", err)
	}

	want := &EnvVar{Id: testEnvVarId, Name: "TEST", Value: "test", Public: false}
	if !reflect.DeepEqual(envVar, want) {
		t.Errorf("EnvVar.FindByRepoId returned %+v, want %+v", envVar, want)
	}
}

func TestEnvVarService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/env_var/%s", testRepoSlug, testEnvVarId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":"test-12345-absde","name":"TEST","value":"test","public":false}`)
	})

	envVar, _, err := client.EnvVar.FindByRepoSlug(context.Background(), testRepoSlug, testEnvVarId)

	if err != nil {
		t.Errorf("EnvVar.FindByRepoSlug returned error: %v", err)
	}

	want := &EnvVar{Id: testEnvVarId, Name: "TEST", Value: "test", Public: false}
	if !reflect.DeepEqual(envVar, want) {
		t.Errorf("EnvVar.FindByRepoSlug returned %+v, want %+v", envVar, want)
	}
}

func TestEnvVarService_UpdateByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/env_var/%s", testRepoId, testEnvVarId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testFormValues(t, r, values{"env_var.name": "TEST", "env_var.value": "test", "env_var.public": "false"})
		fmt.Fprint(w, `{"id":"test-12345-absde","name":"TEST","value":"test","public":false}`)
	})

	opt := EnvVarOption{Name: "TEST", Value: "test", Public: false}
	envVar, _, err := client.EnvVar.UpdateByRepoId(context.Background(), testRepoId, testEnvVarId, &opt)

	if err != nil {
		t.Errorf("EnvVar.UpdateByRepoId returned error: %v", err)
	}

	want := &EnvVar{Id: testEnvVarId, Name: "TEST", Value: "test", Public: false}
	if !reflect.DeepEqual(envVar, want) {
		t.Errorf("EnvVar.UpdateByRepoId returned %+v, want %+v", envVar, want)
	}
}

func TestEnvVarService_UpdateByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/env_var/%s", testRepoSlug, testEnvVarId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testFormValues(t, r, values{"env_var.name": "TEST", "env_var.value": "test", "env_var.public": "false"})
		fmt.Fprint(w, `{"id":"test-12345-absde","name":"TEST","value":"test","public":false}`)
	})

	opt := EnvVarOption{Name: "TEST", Value: "test", Public: false}
	envVar, _, err := client.EnvVar.UpdateByRepoSlug(context.Background(), testRepoSlug, testEnvVarId, &opt)

	if err != nil {
		t.Errorf("EnvVar.UpdateByRepoSlug returned error: %v", err)
	}

	want := &EnvVar{Id: testEnvVarId, Name: "TEST", Value: "test", Public: false}
	if !reflect.DeepEqual(envVar, want) {
		t.Errorf("EnvVar.UpdateByRepoSlug returned %+v, want %+v", envVar, want)
	}
}

func TestEnvVarService_DeleteByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/env_var/%s", testRepoId, testEnvVarId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		fmt.Fprint(w, `{}`)
	})

	_, err := client.EnvVar.DeleteByRepoId(context.Background(), testRepoId, testEnvVarId)

	if err != nil {
		t.Errorf("EnvVar.DeleteByRepoId returned error: %v", err)
	}
}

func TestEnvVarService_DeleteByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/env_var/%s", testRepoSlug, testEnvVarId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		fmt.Fprint(w, `{}`)
	})

	_, err := client.EnvVar.DeleteByRepoSlug(context.Background(), testRepoSlug, testEnvVarId)

	if err != nil {
		t.Errorf("EnvVar.DeleteByRepoSlug returned error: %v", err)
	}
}
