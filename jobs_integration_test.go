// +build integration

package travis

import (
	"context"
	"testing"
)

func TestJobsService_Find_without_options(t *testing.T) {
	_, _, err := integrationClient.Jobs.Find(context.TODO(), nil)
	ok(t, err) // As jobs could be nil, that's unfortunately the only assertion we can make
}

func TestJobsService_Find_with_options(t *testing.T) {
	// Make sure to fetch jobs, and extract first returned
	// job queue to be able to filter against an existing one
	// later on
	unfilteredJobs, _, err := integrationClient.Jobs.Find(context.TODO(), nil)
	jobQueue := unfilteredJobs[0].Queue

	opt := &JobFindOptions{Queue: jobQueue}
	jobs, _, err := integrationClient.Jobs.Find(context.TODO(), opt)
	ok(t, err)

	if jobs != nil {
		for _, j := range jobs {
			assert(
				t,
				j.Queue == jobQueue,
				"JobsService.FindByID return a job with Queue %s; expected %s", j.Queue, jobQueue,
			)
		}
	}
}

func TestJobsService_ListFromBuild(t *testing.T) {
	// Fetch an existing build id
	builds, _, _, _, err := integrationClient.Builds.List(context.TODO(), nil)
	buildId := builds[0].Id

	jobs, _, err := integrationClient.Jobs.ListFromBuild(context.TODO(), buildId)
	ok(t, err)

	if jobs != nil {
		for _, j := range jobs {
			assert(
				t,
				j.Build.Id == buildId,
				"JobsService.ListFromBuild return a job with BuildId %d; expected %d", j.Build.Id, buildId,
			)
		}
	}
}

func TestJobsService_Get(t *testing.T) {
	jobs, _, err := integrationClient.Jobs.Find(context.TODO(), nil)
	if jobs == nil || len(jobs) == 0 {
		t.Skip("No jobs found for the provided integration repo. skipping test")
	}
	jobId := jobs[0].Id

	job, _, err := integrationClient.Jobs.Get(context.TODO(), jobId)
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
