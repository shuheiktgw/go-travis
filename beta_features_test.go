// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBetaFeaturesService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/user/%d/beta_features", testUserId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"beta_features":[{"id":1,"name":"dashboard","description":"Try the new personal Dashboard layout","enabled":true,"feedback_url":"https://github.com/travis-ci/beta-features/issues/5"}]}`)
	})

	features, _, err := client.BetaFeatures.List(context.Background(), testUserId)

	if err != nil {
		t.Errorf("BetaFeatures.List returned error: %v", err)
	}

	want := []BetaFeature{{Id: 1, Name: "dashboard", Description: "Try the new personal Dashboard layout", Enabled: true, FeedbackUrl: "https://github.com/travis-ci/beta-features/issues/5"}}
	if !reflect.DeepEqual(features, want) {
		t.Errorf("BetaFeatures.List returned %+v, want %+v", features, want)
	}
}
