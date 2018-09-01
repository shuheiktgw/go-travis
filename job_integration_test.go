// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
	"time"
)

const jobId = 420908541

func TestJobService_Find(t *testing.T) {
	job, res, err := integrationClient.Job.Find(context.TODO(), jobId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if job.Id != jobId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", jobId, job.Id)
	}
}

func TestJobService_RestartAndCancel(t *testing.T) {
	// Start a job
	job, res, err := integrationClient.Job.Restart(context.TODO(), jobId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if job.Id != jobId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", jobId, job.Id)
	}

	// Wait till the job has successfully processed
	time.Sleep(2 * time.Second)

	job, res, err = integrationClient.Job.Cancel(context.TODO(), jobId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if job.Id != jobId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", jobId, job.Id)
	}
}
