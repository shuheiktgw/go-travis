// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
)

func TestBranchService_Integration_FindByRepoId(t *testing.T) {
	t.Parallel()

	branch, res, err := integrationClient.Branch.FindByRepoId(context.TODO(), integrationRepoId, "master")

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if branch.Name != "master" {
		t.Fatalf("unexpected branch returned: want %s: got %s", "master", branch.Name)
	}

	if branch.Repository.Id != integrationRepoId {
		t.Fatalf("unexpected branch returned: want %d: got %d", integrationRepoId, branch.Repository.Id)
	}
}

func TestBranchService_Integration_FindByRepoSlug(t *testing.T) {
	t.Parallel()

	branch, res, err := integrationClient.Branch.FindByRepoSlug(context.TODO(), integrationRepoSlug, "master")

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if branch.Name != "master" {
		t.Fatalf("unexpected branch returned: want %s: got %s", "master", branch.Name)
	}

	if branch.Repository.Slug != integrationRepoSlug {
		t.Fatalf("unexpected branch returned: want %s: got %s", integrationRepoSlug, branch.Repository.Slug)
	}
}
