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
)

func TestLogsService_Integration_FindByJobId(t *testing.T) {
	_, res, err := integrationClient.Logs.FindByJobId(context.TODO(), integrationJobId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}

func TestLogsService_Integration_DeleteByJobId(t *testing.T) {
	var id uint = 420907934

	_, res, _ := integrationClient.Logs.DeleteByJobId(context.TODO(), id)

	if res.StatusCode != http.StatusConflict {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}
