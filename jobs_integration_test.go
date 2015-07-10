// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

package travis

import "testing"

func TestJobsService_Find_without_options(t *testing.T) {
	jobs, _, err := integrationClient.Jobs.Find(nil)
	ok(t, err)

	assert(
		t,
		jobs != nil,
		"JobsService.Find returned nil",
	)

	assert(
		t,
		len(jobs) > 0,
		"JobsService.Find returned no jobs",
	)
}

func TestJobsService_Find_with_options(t *testing.T) {
	// Make sure to fetch jobs, and extract first returned
	// job queue to be able to filter against an existing one
	// later on
	unfilteredJobs, _, err := integrationClient.Jobs.Find(nil)
	jobQueue := unfilteredJobs[0].Queue

	opt := &JobFindOptions{Queue: jobQueue}
	jobs, _, err := integrationClient.Jobs.Find(opt)
	ok(t, err)

	assert(
		t,
		jobs != nil,
		"JobsService.Find returned nil",
	)

	assert(
		t,
		len(jobs) > 0,
		"JobsService.Find returned no jobs",
	)

	for _, j := range jobs {
		assert(
			t,
			j.Queue == jobQueue,
			"JobsService.Find return a job with Queue %s; expected %s", j.Queue, jobQueue,
		)
	}
}

func TestJobsService_ListFromBuild(t *testing.T) {
	// Fetch an existing build id
	builds, _, _, _, err := integrationClient.Builds.List(nil)
	buildId := builds[0].Id

	jobs, _, err := integrationClient.Jobs.ListFromBuild(buildId)
	ok(t, err)

	assert(
		t,
		jobs != nil,
		"JobsService.ListFromBuild returned nil",
	)

	for _, j := range jobs {
		assert(
			t,
			j.BuildId == buildId,
			"JobsService.ListFromBuild return a job with BuildId %d; expected %d", j.BuildId, buildId,
		)
	}
}

func TestJobsService_Get(t *testing.T) {
	jobs, _, err := integrationClient.Jobs.Find(nil)
	jobId := jobs[0].Id

	job, _, err := integrationClient.Jobs.Get(jobId)
	ok(t, err)

	assert(
		t,
		job != nil,
		"JobsService.Get returned a nil job",
	)

	assert(
		t,
		job.Id == jobId,
		"JobsService.Get returned job with id %d; expected %d",
	)
}

func TestJobFindOptions_IsValid(t *testing.T) {
	opt := &JobFindOptions{}
	assert(
		t,
		opt.IsValid() == true,
		"JobFindOptions.IsValid returned %t; expected %t", opt.IsValid(), true,
	)

	opt = &JobFindOptions{Queue: "test"}
	assert(
		t,
		opt.IsValid() == true,
		"JobFindOptions.IsValid returned %t; expected %t", opt.IsValid(), true,
	)

	opt = &JobFindOptions{Queue: "test", State: "a state"}
	assert(
		t,
		opt.IsValid() == false,
		"JobFindOptions.IsValid returned %t; expected %t", opt.IsValid(), false,
	)
}
