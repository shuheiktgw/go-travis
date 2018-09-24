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

func TestLintService_Integration_Lint(t *testing.T) {
	warnings, res, err := integrationClient.Lint.Lint(context.TODO(), &TravisYml{Content: "foo: bar"})

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if len(warnings) == 0 {
		t.Fatal("Lint.Lint returned empty warning")
	}
}
