// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
)

func TestBranchesService_FindByRepoId(t *testing.T) {
	t.Parallel()

	t.Skip("As of 2018/9/7 this endpoints returns 404 and does not seem to work correctly")

	cases := []*BranchesOption{
		{},
		{Limit: 1},
		{SortBy: "id"},
		{Offset: 0},
	}

	for i, opt := range cases {
		branches, res, err := integrationClient.Branches.FindByRepoId(context.TODO(), integrationRepoId, opt)

		if err != nil {
			t.Fatalf("#%d unexpected error occured: %s", i, err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("#%d invalid http status: %s", i, res.Status)
		}

		if len(branches) == 0 {
			t.Fatalf("#%d returned empty branches", i)
		}
	}
}

func TestBranchesService_FindByRepoSlug(t *testing.T) {
	t.Parallel()

	t.Skip("As of 2018/9/7 this endpoints returns 404 and does not seem to work correctly")

	cases := []*BranchesOption{
		{},
		{Limit: 1},
		{SortBy: "id"},
		{Offset: 0},
	}

	for i, opt := range cases {
		branches, res, err := integrationClient.Branches.FindByRepoSlug(context.TODO(), integrationRepoSlug, opt)

		if err != nil {
			t.Fatalf("#%d unexpected error occured: %s", i, err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("#%d invalid http status: %s", i, res.Status)
		}

		if len(branches) == 0 {
			t.Fatalf("#%d returned empty branches", i)
		}
	}
}
