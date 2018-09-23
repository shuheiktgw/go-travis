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

func TestCronService_Integration_CreateAndFindAndDeleteCron(t *testing.T) {
	// Create a cron by repository id
	opt := CronBody{Interval: CronIntervalMonthly, DontRunIfRecentBuildExists: true}
	createdCron, res, err := integrationClient.Cron.CreateByRepoId(context.TODO(), integrationRepoId, "master", &opt)

	if err != nil {
		t.Fatalf("Cron.CreateByRepoId unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Cron.CreateByRepoId invalid http status: %s", res.Status)
	}

	if got, want := createdCron.Interval, CronIntervalMonthly; got != want {
		t.Errorf("Cron.CreateByRepoId unexpected cron interval returned: want %s got %s", want, got)
	}

	if got, want := createdCron.DontRunIfRecentBuildExists, true; got != want {
		t.Errorf("Cron.CreateByRepoId unexpected cron DontRunIfRecentBuildExists returned: want %v got %v", want, got)
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

	time.Sleep(2 * time.Second)

	// Create a cron by repository slug
	createdCron, res, err = integrationClient.Cron.CreateByRepoSlug(context.TODO(), integrationRepoSlug, "master", &opt)

	if err != nil {
		t.Fatalf("Cron.CreateByRepoSlug unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Cron.CreateByRepoSlug invalid http status: %s", res.Status)
	}

	if got, want := createdCron.Interval, CronIntervalMonthly; got != want {
		t.Errorf("Cron.CreateByRepoSlug unexpected cron interval returned: want %s got %s", want, got)
	}

	if got, want := createdCron.DontRunIfRecentBuildExists, true; got != want {
		t.Errorf("Cron.CreateByRepoSlug unexpected cron DontRunIfRecentBuildExists returned: want %v got %v", want, got)
	}

	time.Sleep(2 * time.Second)

	// Find a cron
	findCron, res, err := integrationClient.Cron.Find(context.TODO(), createdCron.Id)

	if err != nil {
		t.Fatalf("Cron.Find unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Cron.Find invalid http status: %s", res.Status)
	}

	if got, want := findCron.Id, createdCron.Id; got != want {
		t.Errorf("Cron.Find unexpected cron interval returned: want %d got %d", want, got)
	}

	if got, want := findCron.Interval, CronIntervalMonthly; got != want {
		t.Errorf("Cron.Find unexpected cron interval returned: want %s got %s", want, got)
	}

	if got, want := findCron.DontRunIfRecentBuildExists, true; got != want {
		t.Errorf("Cron.Find unexpected cron DontRunIfRecentBuildExists returned: want %v got %v", want, got)
	}

	time.Sleep(2 * time.Second)

	// Find a cron by repository id
	findCron, res, err = integrationClient.Cron.FindByRepoId(context.TODO(), integrationRepoId, "master")

	if err != nil {
		t.Fatalf("Cron.FindByRepoId unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Cron.FindByRepoId invalid http status: %s", res.Status)
	}

	if got, want := findCron.Id, createdCron.Id; got != want {
		t.Errorf("Cron.FindByRepoId unexpected cron interval returned: want %d got %d", want, got)
	}

	if got, want := findCron.Interval, CronIntervalMonthly; got != want {
		t.Errorf("Cron.FindByRepoId unexpected cron interval returned: want %s got %s", want, got)
	}

	if got, want := findCron.DontRunIfRecentBuildExists, true; got != want {
		t.Errorf("Cron.FindByRepoId unexpected cron DontRunIfRecentBuildExists returned: want %v got %v", want, got)
	}

	time.Sleep(2 * time.Second)

	// Find a cron by repository slug
	findCron, res, err = integrationClient.Cron.FindByRepoSlug(context.TODO(), integrationRepoSlug, "master")

	if err != nil {
		t.Fatalf("Cron.FindByRepoSlug unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Cron.FindByRepoSlug invalid http status: %s", res.Status)
	}

	if got, want := findCron.Id, createdCron.Id; got != want {
		t.Errorf("Cron.FindByRepoSlug unexpected cron interval returned: want %d got %d", want, got)
	}

	if got, want := findCron.Interval, CronIntervalMonthly; got != want {
		t.Errorf("Cron.FindByRepoSlug unexpected cron interval returned: want %s got %s", want, got)
	}

	if got, want := findCron.DontRunIfRecentBuildExists, true; got != want {
		t.Errorf("Cron.FindByRepoSlug unexpected cron DontRunIfRecentBuildExists returned: want %v got %v", want, got)
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
