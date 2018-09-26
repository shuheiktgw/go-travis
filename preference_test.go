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

const testPreferenceName = "builds_email"

func TestPreferenceService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/preference/%s", testPreferenceName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"name":"builds_email","value":true}`)
	})

	preference, _, err := client.Preference.Find(context.Background(), testPreferenceName)

	if err != nil {
		t.Errorf("Preference.Find returned error: %v", err)
	}

	want := &Preference{Name: "builds_email", Value: true}
	if !reflect.DeepEqual(preference, want) {
		t.Errorf("Preference.Find returned %+v, want %+v", preference, want)
	}
}

func TestPreferenceService_Update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/preference/%s", testPreferenceName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testBody(t, r, `{"name":"builds_email","value":false}`+"\n")
		fmt.Fprint(w, `{"name":"builds_email","value":false}`)
	})

	preference, _, err := client.Preference.Update(context.Background(), &Preference{Name: "builds_email", Value: false})

	if err != nil {
		t.Errorf("Preference.Update returned error: %v", err)
	}

	want := &Preference{Name: "builds_email", Value: false}
	if !reflect.DeepEqual(preference, want) {
		t.Errorf("Preference.Update returned %+v, want %+v", preference, want)
	}
}
