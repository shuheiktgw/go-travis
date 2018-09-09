// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
)

func TestRepositoryService_Find(t *testing.T) {
	t.Parallel()

	cases := []struct {
		id   uint
		slug string
		want string
	}{
		{id: integrationRepoId, want: integrationRepoSlug},
		{slug: integrationRepoSlug, want: integrationRepoSlug},
		{id: integrationRepoId, slug: integrationRepoSlug, want: integrationRepoSlug},
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

		if got, want := repo.Slug, tc.want; got != want {
			t.Fatalf("#%d unexpected repository returned: want %s: got %s", i, want, got)
		}
	}
}

func TestRepositoryService_Activation(t *testing.T) {
	t.Parallel()

	op := &RepositoryOption{Id: integrationRepoId}

	repo, res, err := integrationClient.Repository.Deactivate(context.TODO(), op)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if repo.Slug != integrationRepoSlug {
		t.Fatalf("unexpected repository returned: want %s: got %s", integrationRepoSlug, repo.Slug)
	}

	repo, res, err = integrationClient.Repository.Activate(context.TODO(), op)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if repo.Slug != integrationRepoSlug {
		t.Fatalf("unexpected repository returned: want %s: got %s", integrationRepoSlug, repo.Slug)
	}
}

func TestRepositoryService_Star(t *testing.T) {
	t.Parallel()

	op := &RepositoryOption{Id: integrationRepoId}

	repo, res, err := integrationClient.Repository.Star(context.TODO(), op)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if repo.Slug != integrationRepoSlug {
		t.Fatalf("unexpected repository returned: want %s: got %s", integrationRepoSlug, repo.Slug)
	}

	repo, res, err = integrationClient.Repository.Unstar(context.TODO(), op)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if repo.Slug != integrationRepoSlug {
		t.Fatalf("unexpected repository returned: want %s: got %s", integrationRepoSlug, repo.Slug)
	}
}
