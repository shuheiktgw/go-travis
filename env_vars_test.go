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

func TestEnvVarsService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/env_var/%s", testRepoId, testEnvVarId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":"test-12345-absde","name":"TEST","value":"test","public":false,"branch":"foo"}`)
	})

	envVar, _, err := client.EnvVars.FindByRepoId(context.Background(), testRepoId, testEnvVarId)

	if err != nil {
		t.Errorf("EnvVars.FindByRepoId returned error: %v", err)
	}

	want := &EnvVar{Id: String(testEnvVarId), Name: String("TEST"), Value: String("test"), Public: Bool(false), Branch: String("foo")}
	if !reflect.DeepEqual(envVar, want) {
		t.Errorf("EnvVars.FindByRepoId returned %+v, want %+v", envVar, want)
	}
}

func TestEnvVarsService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/env_var/%s", testRepoSlug, testEnvVarId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":"test-12345-absde","name":"TEST","value":"test","public":false,"branch":"foo"}`)
	})

	envVar, _, err := client.EnvVars.FindByRepoSlug(context.Background(), testRepoSlug, testEnvVarId)

	if err != nil {
		t.Errorf("EnvVar.FindByRepoSlug returned error: %v", err)
	}

	want := &EnvVar{Id: String(testEnvVarId), Name: String("TEST"), Value: String("test"), Public: Bool(false), Branch: String("foo")}
	if !reflect.DeepEqual(envVar, want) {
		t.Errorf("EnvVars.FindByRepoSlug returned %+v, want %+v", envVar, want)
	}
}

func TestEnvVarsService_ListByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/env_vars", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"env_vars": [{"id":"test-12345-absde","name":"TEST","value":"test","public":false,"branch":"foo"}]}`)
	})

	envVar, _, err := client.EnvVars.ListByRepoId(context.Background(), testRepoId)

	if err != nil {
		t.Errorf("EnvVars.FindByRepoId returned error: %v", err)
	}

	want := []*EnvVar{{Id: String(testEnvVarId), Name: String("TEST"), Value: String("test"), Public: Bool(false), Branch: String("foo")}}
	if !reflect.DeepEqual(envVar, want) {
		t.Errorf("EnvVars.FindByRepoId returned %+v, want %+v", envVar, want)
	}
}

func TestEnvVarsService_ListByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/env_vars", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"env_vars": [{"id":"test-12345-absde","name":"TEST","value":"test","public":false,"branch":"foo"}]}`)
	})

	envVar, _, err := client.EnvVars.ListByRepoSlug(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("EnvVars.FindByRepoSlug returned error: %v", err)
	}

	want := []*EnvVar{{Id: String(testEnvVarId), Name: String("TEST"), Value: String("test"), Public: Bool(false), Branch: String("foo")}}
	if !reflect.DeepEqual(envVar, want) {
		t.Errorf("EnvVars.FindByRepoSlug returned %+v, want %+v", envVar, want)
	}
}

func TestEnvVarsService_CreateByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/env_vars", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"env_var.name":"TEST","env_var.value":"test","env_var.public":false,"env_var.branch":"foo"}`+"\n")
		fmt.Fprint(w, `{"id":"test-12345-absde","name":"TEST","value":"test","public":false,"branch":"foo"}`)
	})

	opt := EnvVarBody{Name: "TEST", Value: "test", Public: false, Branch: "foo"}
	envVar, _, err := client.EnvVars.CreateByRepoId(context.Background(), testRepoId, &opt)

	if err != nil {
		t.Errorf("EnvVars.CreateByRepoId returned error: %v", err)
	}

	want := &EnvVar{Id: String(testEnvVarId), Name: String("TEST"), Value: String("test"), Public: Bool(false), Branch: String("foo")}
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
		fmt.Fprint(w, `{"id":"test-12345-absde","name":"TEST","value":"test","public":false,"branch":"foo"}`)
	})

	opt := EnvVarBody{Name: "TEST", Value: "test", Public: false}
	envVar, _, err := client.EnvVars.CreateByRepoSlug(context.Background(), testRepoSlug, &opt)

	if err != nil {
		t.Errorf("EnvVars.CreateByRepoSlug returned error: %v", err)
	}

	want := &EnvVar{Id: String(testEnvVarId), Name: String("TEST"), Value: String("test"), Public: Bool(false), Branch: String("foo")}
	if !reflect.DeepEqual(envVar, want) {
		t.Errorf("EnvVars.CreateByRepoSlug returned %+v, want %+v", envVar, want)
	}
}

func TestEnvVarsService_UpdateByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/env_var/%s", testRepoId, testEnvVarId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testBody(t, r, `{"env_var.name":"TEST","env_var.value":"test","env_var.public":false,"env_var.branch":"foo"}`+"\n")
		fmt.Fprint(w, `{"id":"test-12345-absde","name":"TEST","value":"test","public":false,"branch":"foo"}`)
	})

	envVar := EnvVarBody{Name: "TEST", Value: "test", Public: false, Branch: "foo"}
	e, _, err := client.EnvVars.UpdateByRepoId(context.Background(), testRepoId, testEnvVarId, &envVar)

	if err != nil {
		t.Errorf("EnvVar.UpdateByRepoId returned error: %v", err)
	}

	want := &EnvVar{Id: String(testEnvVarId), Name: String("TEST"), Value: String("test"), Public: Bool(false), Branch: String("foo")}
	if !reflect.DeepEqual(e, want) {
		t.Errorf("EnvVars.UpdateByRepoId returned %+v, want %+v", envVar, want)
	}
}

func TestEnvVarsService_UpdateByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/env_var/%s", testRepoSlug, testEnvVarId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testBody(t, r, `{"env_var.name":"TEST","env_var.value":"test","env_var.public":false,"env_var.branch":"foo"}`+"\n")
		fmt.Fprint(w, `{"id":"test-12345-absde","name":"TEST","value":"test","public":false,"branch":"foo"}`)
	})

	envVar := EnvVarBody{Name: "TEST", Value: "test", Public: false, Branch: "foo"}
	e, _, err := client.EnvVars.UpdateByRepoSlug(context.Background(), testRepoSlug, testEnvVarId, &envVar)

	if err != nil {
		t.Errorf("EnvVars.UpdateByRepoSlug returned error: %v", err)
	}

	want := &EnvVar{Id: String(testEnvVarId), Name: String("TEST"), Value: String("test"), Public: Bool(false), Branch: String("foo")}
	if !reflect.DeepEqual(e, want) {
		t.Errorf("EnvVars.UpdateByRepoSlug returned %+v, want %+v", e, want)
	}
}

func TestEnvVarsService_DeleteByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/env_var/%s", testRepoId, testEnvVarId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, `{}`)
	})

	_, err := client.EnvVars.DeleteByRepoId(context.Background(), testRepoId, testEnvVarId)

	if err != nil {
		t.Errorf("EnvVars.DeleteByRepoId returned error: %v", err)
	}
}

func TestEnvVarsService_DeleteByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/env_var/%s", testRepoSlug, testEnvVarId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, `{}`)
	})

	_, err := client.EnvVars.DeleteByRepoSlug(context.Background(), testRepoSlug, testEnvVarId)

	if err != nil {
		t.Errorf("EnvVars.DeleteByRepoSlug returned error: %v", err)
	}
}
