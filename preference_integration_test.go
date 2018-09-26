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

func TestPreferenceService_Integration_Find(t *testing.T) {
	preference, res, err := integrationClient.Preference.Find(context.TODO(), integrationPreferenceName)

	if err != nil {
		t.Fatalf("Preference.Find unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Preference.Find invalid http status: %s", res.Status)
	}

	want := &Preference{Name: integrationPreferenceName, Value: true}
	if !reflect.DeepEqual(preference, want) {
		t.Errorf("Preference.Find returned %+v, want %+v", preference, want)
	}
}

func TestPreferenceService_Integration_Update(t *testing.T) {
	// Change build_emails = false
	preference, res, err := integrationClient.Preference.Update(context.TODO(), &Preference{Name: integrationPreferenceName, Value: false})

	if err != nil {
		t.Fatalf("Preference.Update unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Preference.Update invalid http status: %s", res.Status)
	}

	want := &Preference{Name: integrationPreferenceName, Value: false}
	if !reflect.DeepEqual(preference, want) {
		t.Errorf("Preference.Update returned %+v, want %+v", preference, want)
	}

	// Change build_emails = true
	preference, res, err = integrationClient.Preference.Update(context.TODO(), &Preference{Name: integrationPreferenceName, Value: true})

	if err != nil {
		t.Fatalf("Preference.Update unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Preference.Update invalid http status: %s", res.Status)
	}

	want = &Preference{Name: integrationPreferenceName, Value: true}
	if !reflect.DeepEqual(preference, want) {
		t.Errorf("Preference.Update returned %+v, want %+v", preference, want)
	}
}
