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

func TestSettingsService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/setting/%s", testRepoId, BuildsOnlyWithTravisYmlSetting), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"name":"builds_only_with_travis_yml","value":false}`)
	})

	setting, _, err := client.Settings.FindByRepoId(context.Background(), testRepoId, BuildsOnlyWithTravisYmlSetting)

	if err != nil {
		t.Errorf("Settings.FindByRepoId returned error: %v", err)
	}

	want := &Setting{Name: BuildsOnlyWithTravisYmlSetting, Value: false}
	if !reflect.DeepEqual(setting, want) {
		t.Errorf("Settings.FindByRepoId returned %+v, want %+v", setting, want)
	}
}

func TestSettingsService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/setting/%s", testRepoSlug, BuildsOnlyWithTravisYmlSetting), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"name":"builds_only_with_travis_yml","value":false}`)
	})

	setting, _, err := client.Settings.FindByRepoSlug(context.Background(), testRepoSlug, BuildsOnlyWithTravisYmlSetting)

	if err != nil {
		t.Errorf("Settings.FindByRepoSlug returned error: %v", err)
	}

	want := &Setting{Name: BuildsOnlyWithTravisYmlSetting, Value: false}
	if !reflect.DeepEqual(setting, want) {
		t.Errorf("Settings.FindByRepoSlug returned %+v, want %+v", setting, want)
	}
}

func TestSettingsService_UpdateByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/setting/%s", testRepoId, BuildsOnlyWithTravisYmlSetting), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testBody(t, r, `{"setting.value":true}`+"\n")
		fmt.Fprint(w, `{"name":"builds_only_with_travis_yml","value":true}`)
	})

	s := Setting{Name: BuildsOnlyWithTravisYmlSetting, Value: true}
	setting, _, err := client.Settings.UpdateByRepoId(context.Background(), testRepoId, &s)

	if err != nil {
		t.Errorf("Settings.UpdateByRepoId returned error: %v", err)
	}

	want := &Setting{Name: BuildsOnlyWithTravisYmlSetting, Value: true}
	if !reflect.DeepEqual(setting, want) {
		t.Errorf("Settings.UpdateByRepoId returned %+v, want %+v", setting, want)
	}
}

func TestSettingsService_UpdateByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/setting/%s", testRepoSlug, BuildsOnlyWithTravisYmlSetting), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testBody(t, r, `{"setting.value":true}`+"\n")
		fmt.Fprint(w, `{"name":"builds_only_with_travis_yml","value":true}`)
	})

	s := Setting{Name: BuildsOnlyWithTravisYmlSetting, Value: true}
	setting, _, err := client.Settings.UpdateByRepoSlug(context.Background(), testRepoSlug, &s)

	if err != nil {
		t.Errorf("Settings.UpdateByRepoSlug returned error: %v", err)
	}

	want := &Setting{Name: BuildsOnlyWithTravisYmlSetting, Value: true}
	if !reflect.DeepEqual(setting, want) {
		t.Errorf("Settings.UpdateByRepoSlug returned %+v, want %+v", setting, want)
	}
}
