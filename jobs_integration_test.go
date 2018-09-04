// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
)

const buildId = 420907933

func TestJobsService_FindByBuild(t *testing.T) {
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

func TestJobsService_Find(t *testing.T) {
	opt := &JobOption{}
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
