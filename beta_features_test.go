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

	want := []*BetaFeature{{Id: Uint(1), Name: String("dashboard"), Description: String("Try the new personal Dashboard layout"), Enabled: Bool(true), FeedbackUrl: String("https://github.com/travis-ci/beta-features/issues/5")}}
	if !reflect.DeepEqual(features, want) {
		t.Errorf("BetaFeatures.List returned %+v, want %+v", features, want)
	}
}

func TestBetaFeaturesService_Update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/user/%d/beta_feature/1", testUserId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testBody(t, r, `{"enabled":true}`+"\n")
		fmt.Fprint(w, `{"id":1,"name":"dashboard","description":"Try the new personal Dashboard layout","enabled":true,"feedback_url":"https://github.com/travis-ci/beta-features/issues/5"}`)
	})

	features, _, err := client.BetaFeatures.Update(context.Background(), testUserId, 1, true)

	if err != nil {
		t.Errorf("BetaFeatures.Update returned error: %v", err)
	}

	want := &BetaFeature{Id: Uint(1), Name: String("dashboard"), Description: String("Try the new personal Dashboard layout"), Enabled: Bool(true), FeedbackUrl: String("https://github.com/travis-ci/beta-features/issues/5")}
	if !reflect.DeepEqual(features, want) {
		t.Errorf("BetaFeatures.Update returned %+v, want %+v", features, want)
	}
}

func TestBetaFeaturesService_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/user/%d/beta_feature/1", testUserId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, `{"id":1,"name":"dashboard","description":"Try the new personal Dashboard layout","enabled":true,"feedback_url":"https://github.com/travis-ci/beta-features/issues/5"}`)
	})

	features, _, err := client.BetaFeatures.Delete(context.Background(), testUserId, 1)

	if err != nil {
		t.Errorf("BetaFeatures.Delete returned error: %v", err)
	}

	want := &BetaFeature{Id: Uint(1), Name: String("dashboard"), Description: String("Try the new personal Dashboard layout"), Enabled: Bool(true), FeedbackUrl: String("https://github.com/travis-ci/beta-features/issues/5")}
	if !reflect.DeepEqual(features, want) {
		t.Errorf("BetaFeatures.Delete returned %+v, want %+v", features, want)
	}
}
