// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
)

func TestInstallations_Integration_Find(t *testing.T) {
	opt := InstallationOption{Include: []string{"installation.owner"}}
	_, res, err := integrationClient.Installations.Find(context.TODO(), 111, &opt)

	if err == nil {
		t.Error("error is not supposed to be nil")
	}

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("invalid http status: %s", res.Status)
	}
}
