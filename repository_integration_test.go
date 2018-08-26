// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
)

const GoTravisRepoId = 20762031

func TestRepositoryService_Find_Success(t *testing.T) {
	t.Parallel()

	cases := []struct {
		id   uint
		slug string
		want string
	}{
		{id: GoTravisRepoId, want: "shuheiktgw/go-travis"},
		{slug: "shuheiktgw/go-travis", want: "shuheiktgw/go-travis"},
		{id: GoTravisRepoId, slug: "shuheiktgw/go-travis", want: "shuheiktgw/go-travis"},
	}

	for i, tc := range cases {
		op := &RepositoryOption{Id: tc.id, Slug: tc.slug}

		repo, res, err := integrationClient.Repository.Find(context.TODO(), op)

		if err != nil {
			t.Fatalf("#%d unexpected error occured: %s", i, err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("#%d invalid http status: %s", i, res.Status)
		}

		if got, want := repo.Slug, "shuheiktgw/go-travis"; got != want {
			t.Fatalf("#%d unexpected repository returned: want %s: got %s", i, want, got)
		}
	}
}
