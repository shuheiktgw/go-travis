// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
)

func TestBuildsService_Find_WithEmptyOption(t *testing.T) {
	opt := &BuildsOption{}
	builds, res, err := integrationClient.Builds.Find(context.TODO(), opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if len(builds) == 0 {
		t.Fatalf("returned empty builds")
	}
}

func TestBuildsService_Find_WithOption(t *testing.T) {
	opt := &BuildsOption{Limit: 1}
	builds, res, err := integrationClient.Builds.Find(context.TODO(), opt)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if len(builds) != 1 {
		t.Fatalf("limit 1 does not seem to work correctly")
	}
}
