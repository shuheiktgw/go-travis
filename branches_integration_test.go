// +build integration

package travis

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestBranchesService_FindByRepositoryId(t *testing.T) {
	t.Parallel()

	t.Skip("As of 2018/9/7 this endpoints returns 404 and does not seem to work correctly")

	cases := []*BranchesOption{
		{},
		{Limit: 1},
		{SortBy: "id"},
		{Offset: 0},
	}

	for i, opt := range cases {
		branches, res, err := integrationClient.Branches.FindByRepositoryId(context.TODO(), integrationRepoId, opt)

		if err != nil {
			t.Fatalf("#%d unexpected error occured: %s", i, err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("#%d invalid http status: %s", i, res.Status)
		}

		if len(branches) == 0 {
			t.Fatalf("#%d returned empty branches", i)
		}

		fmt.Println(branches)
	}
}

func TestBranchesService_FindByRepositorySlug(t *testing.T) {
	t.Parallel()

	t.Skip("As of 2018/9/7 this endpoints returns 404 and does not seem to work correctly")

	cases := []*BranchesOption{
		{},
		{Limit: 1},
		{SortBy: "id"},
		{Offset: 0},
	}

	for i, opt := range cases {
		branches, res, err := integrationClient.Branches.FindByRepositorySlug(context.TODO(), integrationRepo, opt)

		if err != nil {
			t.Fatalf("#%d unexpected error occured: %s", i, err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("#%d invalid http status: %s", i, res.Status)
		}

		if len(branches) == 0 {
			t.Fatalf("#%d returned empty branches", i)
		}

		fmt.Println(branches)
	}
}
