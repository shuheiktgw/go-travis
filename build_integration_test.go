// +build integration

package travis

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestBuildService_Find(t *testing.T) {
	t.Parallel()

	build, res, err := integrationClient.Build.Find(context.TODO(), integrationBuildId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if build.Id != integrationBuildId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", integrationBuildId, build.Id)
	}

	fmt.Println(build)
}

func TestBuildService_RestartAndCancel(t *testing.T) {
	t.Parallel()

	build, res, err := integrationClient.Build.Restart(context.TODO(), integrationBuildId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if build.Id != integrationBuildId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", integrationBuildId, build.Id)
	}

	// Wait till the build has successfully processed
	time.Sleep(2 * time.Second)

	build, res, err = integrationClient.Build.Cancel(context.TODO(), integrationBuildId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if build.Id != integrationBuildId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", integrationBuildId, build.Id)
	}
}
