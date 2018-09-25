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

func TestPreferencesService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/preferences", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"preferences": [{"name":"builds_email","value":true}]}`)
	})

	preferences, _, err := client.Preferences.Find(context.Background())

	if err != nil {
		t.Errorf("Preferences.Find returned error: %v", err)
	}

	want := []Preference{{Name: "builds_email", Value: true}}
	if !reflect.DeepEqual(preferences, want) {
		t.Errorf("Preferences.Find returned %+v, want %+v", preferences, want)
	}
}
