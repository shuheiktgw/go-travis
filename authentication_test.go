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

func TestAuthenticationService_UsingGithubToken_without_token(t *testing.T) {
	as := &AuthenticationService{client: NewClient(defaultBaseURL, "")}

	_, _, err := as.UsingGithubToken(context.TODO(), "")

	if err == nil {
		t.Fatal("AuthenticationService.UsingGithubToken with empty token: error is not supposed to be nil")
	}
}

func TestAuthenticationService_UsingGithubToken_with_token(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/auth/github", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeader(t, r, "Accept", "application/vnd.travis-ci.2.1+json")
		fmt.Fprint(w, `{"access_token":"test_access_token"}`)
	})

	repo, _, err := client.Authentication.UsingGithubToken(context.Background(), "test_github_token")

	if err != nil {
		t.Errorf("AuthenticationService.UsingGithubToken returned error: %v", err)
	}

	want := AccessToken("test_access_token")
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("AuthenticationService.UsingGithubToken returned %+v, want %+v", repo, want)
	}
}

func TestAuthenticationService_UsingTravisToken_without_token(t *testing.T) {
	as := &AuthenticationService{client: NewClient(defaultBaseURL, "")}

	err := as.UsingTravisToken("")
	if err == nil {
		t.Fatal("AuthenticationService.UsingTravisToken with empty token: error is not supposed to be nil")
	}
}

func TestAuthenticationService_UsingTravisToken_with_token(t *testing.T) {
	token := "abc123easyasdoremi"
	as := &AuthenticationService{client: NewClient(defaultBaseURL, token)}

	err := as.UsingTravisToken(token)
	if err != nil {
		t.Fatalf("AuthenticationService.UsingTravisToken: unexpected error occurred: %s", err)
	}

	authHeader := as.client.Headers["Authorization"]
	if authHeader != fmt.Sprintf("token %s", token) {
		t.Fatalf("AuthenticationService.Headers: unexpected Authorization %s; expected %s", authHeader, token)
	}
}
