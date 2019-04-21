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

func TestStagesService_Integration_ListByBuild(t *testing.T) {
	opt := StagesOption{Include: []string{"stage.jobs"}}
	_, res, err := integrationClient.Stages.ListByBuild(context.TODO(), integrationBuildId, &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}
