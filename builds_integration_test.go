// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

package travis

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestBuildService_Integration_Find(t *testing.T) {
	build, res, err := integrationClient.Builds.Find(context.TODO(), integrationBuildId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if build.Id != integrationBuildId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", integrationBuildId, build.Id)
	}
}

func TestBuildsService_Integration_List(t *testing.T) {
	cases := []*BuildsOption{
		{},
		{Limit: 1},
		{SortBy: "id"},
		{Offset: 0},
	}

	for i, opt := range cases {
		builds, res, err := integrationClient.Builds.List(context.TODO(), opt)

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

func TestBuildsService_Integration_ListByRepoId(t *testing.T) {
	cases := []*BuildsByRepositoryOption{
		{},
		{Limit: 1},
		{SortBy: "id"},
		{Offset: 0},
		{State: []string{BuildStateCanceled}},
		{PreviousState: []string{BuildStatePassed}},
		{EventType: []string{BuildEventTypePush}},
		{CreatedBy: []string{"shuheiktgwtest"}},
		{BranchName: "master"},
	}

	for i, opt := range cases {
		builds, res, err := integrationClient.Builds.ListByRepoId(context.TODO(), integrationRepoId, opt)

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

func TestBuildsService_Integration_ListByRepoSlug(t *testing.T) {
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
		builds, res, err := integrationClient.Builds.ListByRepoSlug(context.TODO(), integrationRepoSlug, opt)

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

func TestBuildService_Integration_RestartAndCancel(t *testing.T) {
	build, res, err := integrationClient.Builds.Restart(context.TODO(), integrationBuildId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("#invalid http status: %s", res.Status)
	}

	if build.Id != integrationBuildId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", integrationBuildId, build.Id)
	}

	// Wait till the build has successfully processed
	time.Sleep(2 * time.Second)

	build, res, err = integrationClient.Builds.Cancel(context.TODO(), integrationBuildId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if build.Id != integrationBuildId {
		t.Fatalf("unexpected job returned: want job id %d: got job id %d", integrationBuildId, build.Id)
	}
}
