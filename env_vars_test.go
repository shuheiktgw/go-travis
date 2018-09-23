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

func TestEnvVarsService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/env_vars", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"env_vars": [{"id":"test-12345-absde","name":"TEST","value":"test","public":false}]}`)
	})

	envVar, _, err := client.EnvVars.FindByRepoId(context.Background(), testRepoId)

	if err != nil {
		t.Errorf("EnvVars.FindByRepoId returned error: %v", err)
	}

	want := []EnvVar{{Id: testEnvVarId, Name: "TEST", Value: "test", Public: false}}
	if !reflect.DeepEqual(envVar, want) {
		t.Errorf("EnvVars.FindByRepoId returned %+v, want %+v", envVar, want)
	}
}

func TestEnvVarsService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/env_vars", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"env_vars": [{"id":"test-12345-absde","name":"TEST","value":"test","public":false}]}`)
	})

	envVar, _, err := client.EnvVars.FindByRepoSlug(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("EnvVars.FindByRepoSlug returned error: %v", err)
	}

	want := []EnvVar{{Id: testEnvVarId, Name: "TEST", Value: "test", Public: false}}
	if !reflect.DeepEqual(envVar, want) {
		t.Errorf("EnvVars.FindByRepoSlug returned %+v, want %+v", envVar, want)
	}
}

func TestEnvVarsService_CreateByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/env_vars", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"env_var.name":"TEST","env_var.value":"test","env_var.public":false}`+"\n")
		fmt.Fprint(w, `{"id":"test-12345-absde","name":"TEST","value":"test","public":false}`)
	})

	opt := EnvVarBody{Name: "TEST", Value: "test", Public: false}
	envVar, _, err := client.EnvVars.CreateByRepoId(context.Background(), testRepoId, &opt)

	if err != nil {
		t.Errorf("EnvVars.CreateByRepoId returned error: %v", err)
	}

	want := &EnvVar{Id: testEnvVarId, Name: "TEST", Value: "test", Public: false}
	if !reflect.DeepEqual(envVar, want) {
		t.Errorf("EnvVars.CreateByRepoId returned %+v, want %+v", envVar, want)
	}
}

func TestEnvVarsService_CreateByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/env_vars", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"env_var.name":"TEST","env_var.value":"test","env_var.public":false}`+"\n")
		fmt.Fprint(w, `{"id":"test-12345-absde","name":"TEST","value":"test","public":false}`)
	})

	opt := EnvVarBody{Name: "TEST", Value: "test", Public: false}
	envVar, _, err := client.EnvVars.CreateByRepoSlug(context.Background(), testRepoSlug, &opt)

	if err != nil {
		t.Errorf("EnvVars.CreateByRepoSlug returned error: %v", err)
	}

	want := &EnvVar{Id: testEnvVarId, Name: "TEST", Value: "test", Public: false}
	if !reflect.DeepEqual(envVar, want) {
		t.Errorf("EnvVars.CreateByRepoSlug returned %+v, want %+v", envVar, want)
	}
}
