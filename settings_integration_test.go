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
	"time"
)

func TestSettingsService_Integration_FindByRepoId(t *testing.T) {
	setting, res, err := integrationClient.Settings.FindByRepoId(context.TODO(), integrationRepoId, BuildsOnlyWithTravisYmlSetting)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	want := &Setting{
		Name:  String(BuildsOnlyWithTravisYmlSetting),
		Value: false,
		Metadata: &Metadata{
			Type:           String("setting"),
			Href:           String("/repo/20783933/setting/builds_only_with_travis_yml"),
			Representation: String("standard"),
			Permissions:    &Permissions{"read": true, "write": true},
		},
	}

	if !reflect.DeepEqual(setting, want) {
		t.Errorf("Settings.FindByRepoId returned %+v, want %+v", setting, want)
	}
}

func TestSettingsService_Integration_FindByRepoSlug(t *testing.T) {
	setting, res, err := integrationClient.Settings.FindByRepoSlug(context.TODO(), integrationRepoSlug, BuildsOnlyWithTravisYmlSetting)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	want := &Setting{
		Name:  String(BuildsOnlyWithTravisYmlSetting),
		Value: false,
		Metadata: &Metadata{
			Type:           String("setting"),
			Href:           String("/repo/20783933/setting/builds_only_with_travis_yml"),
			Representation: String("standard"),
			Permissions:    &Permissions{"read": true, "write": true},
		},
	}

	if !reflect.DeepEqual(setting, want) {
		t.Errorf("Settings.FindByRepoSlug returned %+v, want %+v", setting, want)
	}
}

func TestSettingsService_Integration_ListByRepoId(t *testing.T) {
	settings, res, err := integrationClient.Settings.ListByRepoId(context.TODO(), integrationRepoId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if len(settings) == 0 {
		t.Fatalf("Settings.ListByRepoId settings cannot be empty")
	}
}

func TestSettingsService_Integration_ListByRepoSlug(t *testing.T) {
	settings, res, err := integrationClient.Settings.ListByRepoSlug(context.TODO(), integrationRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if len(settings) == 0 {
		t.Fatalf("Settings.ListByRepoSlug settings cannot be empty")
	}
}

func TestSettingsService_Integration_UpdateByRepoIdAndSlug(t *testing.T) {
	s := SettingBody{Name: BuildsOnlyWithTravisYmlSetting, Value: true}
	setting, res, err := integrationClient.Settings.UpdateByRepoId(context.TODO(), integrationRepoId, &s)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	want := &Setting{
		Name:  String(BuildsOnlyWithTravisYmlSetting),
		Value: true,
		Metadata: &Metadata{
			Type:           String("setting"),
			Href:           String("/repo/20783933/setting/builds_only_with_travis_yml"),
			Representation: String("standard"),
			Permissions:    &Permissions{"read": true, "write": true},
		},
	}

	if !reflect.DeepEqual(setting, want) {
		t.Errorf("Settings.UpdateByRepoId returned %+v, want %+v", setting, want)
	}

	time.Sleep(2 * time.Second)

	s = SettingBody{Name: BuildsOnlyWithTravisYmlSetting, Value: false}
	setting, res, err = integrationClient.Settings.UpdateByRepoSlug(context.TODO(), integrationRepoSlug, &s)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	want = &Setting{
		Name:  String(BuildsOnlyWithTravisYmlSetting),
		Value: false,
		Metadata: &Metadata{
			Type:           String("setting"),
			Href:           String("/repo/20783933/setting/builds_only_with_travis_yml"),
			Representation: String("standard"),
			Permissions:    &Permissions{"read": true, "write": true},
		},
	}

	if !reflect.DeepEqual(setting, want) {
		t.Errorf("Settings.UpdateByRepoSlug returned %+v, want %+v", setting, want)
	}
}
