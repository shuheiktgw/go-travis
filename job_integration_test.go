// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestJobService_Integration_Find(t *testing.T) {
	job, res, err := integrationClient.Job.Find(context.TODO(), integrationJobId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if job.Id != integrationJobId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", integrationBuildId, job.Id)
	}
}

func TestJobService_Integration_RestartAndCancel(t *testing.T) {
	// Start a job
	job, res, err := integrationClient.Job.Restart(context.TODO(), integrationJobId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if job.Id != integrationJobId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", integrationJobId, job.Id)
	}

	// Wait till the job has successfully processed
	time.Sleep(2 * time.Second)

	job, res, err = integrationClient.Job.Cancel(context.TODO(), integrationJobId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if job.Id != integrationJobId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", integrationJobId, job.Id)
	}
}
