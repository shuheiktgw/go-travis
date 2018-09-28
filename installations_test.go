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

const testInstallationId = 111

func TestInstallationsService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/installation/%d", testInstallationId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"id":%d,"github_id":%d}`, testInstallationId, testGitHubId)
	})

	installation, _, err := client.Installations.Find(context.Background(), testInstallationId)

	if err != nil {
		t.Errorf("Installation.Find returned error: %v", err)
	}

	want := &Installation{Id: testInstallationId, GitHubId: testGitHubId}
	if !reflect.DeepEqual(installation, want) {
		t.Errorf("Installation.Find returned %+v, want %+v", installation, want)
	}
}
