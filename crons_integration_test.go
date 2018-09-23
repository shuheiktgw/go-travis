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

func TestCronService_Integration_FindByRepoId(t *testing.T) {
	// Create a cron by repository id
	createdCron, res, err := integrationClient.Cron.CreateByRepoId(
		context.TODO(),
		integrationRepoId,
		"master",
		&CronBody{Interval: CronIntervalMonthly, DontRunIfRecentBuildExists: true},
	)

	if err != nil {
		t.Fatalf("Cron.CreateByRepoId unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Cron.CreateByRepoId invalid http status: %s", res.Status)
	}

	time.Sleep(2 * time.Second)

	// Find crons
	opt := CronsOption{Limit: 5}
	crons, res, err := integrationClient.Crons.FindByRepoId(context.TODO(), integrationRepoId, &opt)

	if err != nil {
		t.Fatalf("Cron.FindByRepoId unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Cron.FindByRepoId invalid http status: %s", res.Status)
	}

	if got, want := len(crons), 1; got != want {
		t.Fatalf("Cron.FindByRepoId returns invalid number of items: want %d, got %d", want, got)
	}

	if got, want := crons[0].Id, createdCron.Id; got != want {
		t.Fatalf("Cron.FindByRepoId returns invalid item id: want %d, got %d", want, got)
	}

	time.Sleep(2 * time.Second)

	// Delete a cron
	res, err = integrationClient.Cron.Delete(context.TODO(), createdCron.Id)

	if err != nil {
		t.Fatalf("Cron.Delete unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("Cron.Delete invalid http status: %s", res.Status)
	}
}

func TestCronService_Integration_FindByRepoSlug(t *testing.T) {
	// Create a cron by repository id
	createdCron, res, err := integrationClient.Cron.CreateByRepoId(
		context.TODO(),
		integrationRepoId,
		"master",
		&CronBody{Interval: CronIntervalMonthly, DontRunIfRecentBuildExists: true},
	)

	if err != nil {
		t.Fatalf("Cron.CreateByRepoId unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Cron.CreateByRepoId invalid http status: %s", res.Status)
	}

	time.Sleep(2 * time.Second)

	// Find crons
	opt := CronsOption{Limit: 5}
	crons, res, err := integrationClient.Crons.FindByRepoSlug(context.TODO(), integrationRepoSlug, &opt)

	if err != nil {
		t.Fatalf("Cron.FindByRepoSlug unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Cron.FindByRepoSlug invalid http status: %s", res.Status)
	}

	if got, want := len(crons), 1; got != want {
		t.Fatalf("Cron.FindByRepoSlug returns invalid number of items: want %d, got %d", want, got)
	}

	if got, want := crons[0].Id, createdCron.Id; got != want {
		t.Fatalf("Cron.FindByRepoSlug returns invalid item id: want %d, got %d", want, got)
	}

	time.Sleep(2 * time.Second)

	// Delete a cron
	res, err = integrationClient.Cron.Delete(context.TODO(), createdCron.Id)

	if err != nil {
		t.Fatalf("Cron.Delete unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("Cron.Delete invalid http status: %s", res.Status)
	}
}
