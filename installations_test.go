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
		testFormValues(t, r, values{"include": "installation.owner"})
		fmt.Fprintf(w, `{"id":%d,"github_id":%d}`, testInstallationId, testGitHubId)
	})

	opt := InstallationOption{Include: []string{"installation.owner"}}
	installation, _, err := client.Installations.Find(context.Background(), testInstallationId, &opt)

	if err != nil {
		t.Errorf("Installation.Find returned error: %v", err)
	}

	want := &Installation{Id: Uint(testInstallationId), GitHubId: Uint(testGitHubId)}
	if !reflect.DeepEqual(installation, want) {
		t.Errorf("Installation.Find returned %+v, want %+v", installation, want)
	}
}
