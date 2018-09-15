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

const buildId = 420907933

func TestJobsService_Integration_FindByBuild(t *testing.T) {
	jobs, res, err := integrationClient.Jobs.FindByBuild(context.TODO(), buildId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if len(jobs) == 0 {
		t.Fatalf("returned empty jobs")
	}
}

func TestJobsService_Integration_Find(t *testing.T) {
	opt := &JobsOption{}
	jobs, res, err := integrationClient.Jobs.Find(context.TODO(), opt)

	// This endpoint returns 500 as of 2019/09/04
	t.Skip()

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if len(jobs) == 0 {
		t.Fatalf("returned empty jobs")
	}
}
