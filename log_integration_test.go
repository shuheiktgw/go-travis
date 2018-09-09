// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
)

func TestLogService_FindByJob(t *testing.T) {
	_, res, err := integrationClient.Log.FindByJob(context.TODO(), integrationJobId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}
}
