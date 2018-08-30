// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
)

func TestBranchService_Find(t *testing.T) {
	t.Parallel()

	op := BranchOption{Slug: integrationRepo, BranchName: "master"}
	branch, res, err := integrationClient.Branch.Find(context.TODO(), &op)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if branch.Name != "master" {
		t.Fatalf("unexpected branch returned: want %s: got %s", "master", branch.Name)
	}

	if branch.Repository.Slug != integrationRepo {
		t.Fatalf("unexpected branch returned: want %s: got %s", integrationRepo, branch.Repository.Slug)
	}
}
