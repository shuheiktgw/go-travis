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
)

const integrationPreferenceName = "build_emails"

func TestPreferencesService_Integration_Find(t *testing.T) {
	preference, res, err := integrationClient.Preferences.Find(context.TODO(), integrationPreferenceName)

	if err != nil {
		t.Fatalf("Preference.Find unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Preference.Find invalid http status: %s", res.Status)
	}

	want := &Preference{
		Name:     String(integrationPreferenceName),
		Value:    true,
		Metadata: &Metadata{Type: String("preference"), Href: String("/v3/preference/build_emails"), Representation: String("standard")},
	}

	if !reflect.DeepEqual(preference, want) {
		t.Errorf("Preference.Find returned %+v, want %+v", preference, want)
	}
}

func TestPreferencesService_Integration_List(t *testing.T) {
	preferences, res, err := integrationClient.Preferences.List(context.TODO())

	if err != nil {
		t.Fatalf("Preferences.List unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Preferences.List invalid http status: %s", res.Status)
	}

	want := []*Preference{
		{
			Name:     String("build_emails"),
			Value:    true,
			Metadata: &Metadata{Type: String("preference"), Href: String("/v3/preference/build_emails"), Representation: String("standard")},
		},
		{
			Name:     String("private_insights_visibility"),
			Value:    "private",
			Metadata: &Metadata{Type: String("preference"), Href: String("/v3/preference/private_insights_visibility"), Representation: String("standard")},
		},
	}

	if !reflect.DeepEqual(preferences, want) {
		t.Errorf("Preferences.Find returned %+v, want %+v", preferences, want)
	}
}

func TestPreferenceServices_Integration_Update(t *testing.T) {
	// Change build_emails = false
	preference, res, err := integrationClient.Preferences.Update(context.TODO(), &PreferenceBody{Name: integrationPreferenceName, Value: false})

	if err != nil {
		t.Fatalf("Preference.Update unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Preference.Update invalid http status: %s", res.Status)
	}

	want := &Preference{
		Name:     String(integrationPreferenceName),
		Value:    false,
		Metadata: &Metadata{Type: String("preference"), Href: String("/v3/preference/build_emails"), Representation: String("standard")},
	}

	if !reflect.DeepEqual(preference, want) {
		t.Errorf("Preference.Update returned %+v, want %+v", preference, want)
	}

	// Change build_emails = true
	preference, res, err = integrationClient.Preferences.Update(context.TODO(), &PreferenceBody{Name: integrationPreferenceName, Value: true})

	if err != nil {
		t.Fatalf("Preference.Update unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Preference.Update invalid http status: %s", res.Status)
	}

	want = &Preference{
		Name:     String(integrationPreferenceName),
		Value:    true,
		Metadata: &Metadata{Type: String("preference"), Href: String("/v3/preference/build_emails"), Representation: String("standard")},
	}
	if !reflect.DeepEqual(preference, want) {
		t.Errorf("Preference.Update returned %+v, want %+v", preference, want)
	}
}
