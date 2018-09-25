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

func TestPreferencesService_Integration_Find(t *testing.T) {
	preferences, res, err := integrationClient.Preferences.Find(context.TODO())

	if err != nil {
		t.Fatalf("Preferences.Find unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Preferences.Find invalid http status: %s", res.Status)
	}

	want := []Preference{{Name: "build_emails", Value: true}}
	if !reflect.DeepEqual(preferences, want) {
		t.Errorf("Preferences.Find returned %+v, want %+v", preferences, want)
	}
}
