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

func TestBuildsService_FindByRepoId(t *testing.T) {
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
		builds, res, err := integrationClient.Builds.FindByRepoId(context.TODO(), integrationRepoId, opt)

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

func TestBuildsService_FindByRepoSlug(t *testing.T) {
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
		builds, res, err := integrationClient.Builds.FindByRepoSlug(context.TODO(), integrationRepoSlug, opt)

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
