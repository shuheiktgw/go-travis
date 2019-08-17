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
	"time"
)

const buildId = 420907933

func TestJobsService_Integration_Find(t *testing.T) {
	opt := JobOption{Include: []string{"job.repository", "job.config"}}
	job, res, err := integrationClient.Jobs.Find(context.TODO(), integrationJobId, &opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if *job.Id != integrationJobId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", integrationBuildId, job.Id)
	}

	if job.Repository.IsMinimal() {
		t.Fatal("repository is minimal representation")
	}

	if job.Commit.IsStandard() {
		t.Fatal("commit is standard representation")
	}
}

func TestJobsService_Integration_ListByBuild(t *testing.T) {
	jobs, res, err := integrationClient.Jobs.ListByBuild(context.TODO(), buildId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if len(jobs) == 0 {
		t.Fatalf("returned empty jobs")
	}
}

func TestJobsService_Integration_List(t *testing.T) {
	opt := &JobsOption{}
	jobs, res, err := integrationClient.Jobs.List(context.TODO(), opt)

	// This endpoint returns 500 as of 2019/04/17
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

func TestJobsService_Integration_RestartAndCancel(t *testing.T) {
	// Start a job
	job, res, err := integrationClient.Jobs.Restart(context.TODO(), integrationJobId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if *job.Id != integrationJobId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", integrationJobId, job.Id)
	}

	// Wait till the job has successfully processed
	time.Sleep(2 * time.Second)

	job, res, err = integrationClient.Jobs.Cancel(context.TODO(), integrationJobId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if *job.Id != integrationJobId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", integrationJobId, job.Id)
	}
}
