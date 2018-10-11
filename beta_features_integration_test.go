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

const integrationBetaFeatureId = 1

func TestBetaFeaturesService_Integration_List(t *testing.T) {
	_, res, err := integrationClient.BetaFeatures.List(context.TODO(), integrationUserId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}

func TestBetaFeaturesService_Integration_Update(t *testing.T) {
	feature, res, err := integrationClient.BetaFeatures.Update(context.TODO(), integrationUserId, integrationBetaFeatureId, true)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if feature.Id != integrationBetaFeatureId || feature.Enabled != true {
		t.Fatalf("unexpected beta feature has returned: %v", feature)
	}

	time.Sleep(2 * time.Second)

	feature, res, err = integrationClient.BetaFeatures.Update(context.TODO(), integrationUserId, integrationBetaFeatureId, false)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if feature.Id != integrationBetaFeatureId || feature.Enabled != false {
		t.Fatalf("unexpected beta feature has returned: %v", feature)
	}
}
