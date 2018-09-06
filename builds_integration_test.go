// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
)

func TestBuildsService_Find(t *testing.T) {
	cases := []*BuildsOption{
		{},
		{Limit: 1},
		{SortBy: "id"},
		{Offset: 0},
	}

	for i, opt := range cases {
		builds, res, err := integrationClient.Builds.Find(context.TODO(), opt)

		if err != nil {
			t.Fatalf("#%d unexpected error occured: %s", i, err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("#%d invalid http status: %s", i, res.Status)
		}

		if len(builds) == 0 {
			t.Fatalf("#%d returned empty builds", i)
		}
	}
}

func TestBuildsService_FindByRepositoryId(t *testing.T) {
	cases := []*BuildsByRepositoryOption{
		{},
		{Limit: 1},
		{SortBy: "id"},
		{Offset: 0},
		{State: []string{BuildStateCanceled}},
		{PreviousState: []string{BuildStatePassed}},
		{EventType: []string{BuildEventTypePush}},
		{CreatedBy: []string{"shuheiktgwtest"}},
	}

	for i, opt := range cases {
		builds, res, err := integrationClient.Builds.FindByRepositoryId(context.TODO(), integrationRepoId, opt)

		if err != nil {
			t.Fatalf("#%d unexpected error occured: %s", i, err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("#%d invalid http status: %s", i, res.Status)
		}

		if len(builds) == 0 {
			t.Fatalf("#%d returned empty builds", i)
		}
	}
}

func TestBuildsService_FindByRepositorySlug(t *testing.T) {
	cases := []*BuildsByRepositoryOption{
		{},
		{Limit: 1},
		{SortBy: "id"},
		{Offset: 0},
		{State: []string{BuildStateCanceled}},
		{PreviousState: []string{BuildStatePassed}},
		{EventType: []string{BuildEventTypePush}},
		{CreatedBy: []string{"shuheiktgwtest"}},
	}

	for i, opt := range cases {
		builds, res, err := integrationClient.Builds.FindByRepositorySlug(context.TODO(), integrationRepo, opt)

		if err != nil {
			t.Fatalf("#%d unexpected error occured: %s", i, err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("#%d invalid http status: %s", i, res.Status)
		}

		if len(builds) == 0 {
			t.Fatalf("#%d returned empty builds", i)
		}
	}
}
