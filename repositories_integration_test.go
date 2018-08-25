// +build integration

package travis

import (
	"context"
	"testing"
)

func TestRepositoriesService_Find_without_options(t *testing.T) {
	t.Parallel()

	_, _, err := integrationClient.Repositories.Find(context.TODO(), nil)
	ok(t, err)
}

func TestRepositoriesService_Find_with_options(t *testing.T) {
	t.Parallel()

	opt := &RepositoryListOptions{Slug: integrationRepo}
	repositories, _, err := integrationClient.Repositories.Find(context.TODO(), opt)
	ok(t, err)

	assert(
		t,
		len(repositories) == 1,
		"Repositories.Find returned no repositories; expected 1",
	)

	assert(
		t,
		repositories[0].Slug == integrationRepo,
		"Repositories.Find returned a repository with slug %s; expected %s", repositories[0].Slug, integrationRepo,
	)
}

func TestRepositoriesService_GetFromSlug(t *testing.T) {
	t.Parallel()

	repository, _, err := integrationClient.Repositories.GetFromSlug(context.TODO(), integrationRepo)
	ok(t, err)

	assert(
		t,
		repository != nil,
		"Repositories.GetFromSlug returned nil repository",
	)

	assert(
		t,
		repository.Slug == integrationRepo,
		"Repositories.GetFromSlug returned a repository with slug %s; expected %s", repository.Slug, integrationRepo,
	)
}

func TestRepositoriesService_Get(t *testing.T) {
	t.Parallel()

	repoFromSlug, _, err := integrationClient.Repositories.GetFromSlug(context.TODO(), integrationRepo)
	repositoryId := repoFromSlug.Id

	repository, _, err := integrationClient.Repositories.Get(context.TODO(), repositoryId)
	ok(t, err)

	assert(
		t,
		repository != nil,
		"Repositories.Get returned nil repository",
	)

	assert(
		t,
		repository.Id == repositoryId,
		"Repositories.Get returned a repository with Id %d; expected %d", repository.Id, repositoryId,
	)

	assert(
		t,
		repository.Slug == integrationRepo,
		"Repositories.Get returned a repository with slug %s; expected %s", repository.Slug, integrationRepo,
	)
}
