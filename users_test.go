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

const testUserId = 4321

func TestUserService_Current(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"include": "user.repositories,user.installation,user.emails"})
		fmt.Fprint(w, `{"id":1,"login":"shuheiktgw","name":"shuheiktgw","github_id":1}`)
	})

	opt := UserOption{Include: []string{"user.repositories", "user.installation", "user.emails"}}
	repo, _, err := client.User.Current(context.Background(), &opt)

	if err != nil {
		t.Errorf("UserService.Current returned error: %v", err)
	}

	want := &User{Id: 1, Login: "shuheiktgw", Name: "shuheiktgw", GithubId: 1}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("UserService.Current returned %+v, want %+v", repo, want)
	}
}

func TestUserService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/user/%d", testUserId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"include": "user.repositories,user.installation,user.emails"})
		fmt.Fprint(w, `{"id":1,"login":"shuheiktgw","name":"shuheiktgw","github_id":1}`)
	})

	opt := UserOption{Include: []string{"user.repositories", "user.installation", "user.emails"}}
	repo, _, err := client.User.Find(context.Background(), testUserId, &opt)

	if err != nil {
		t.Errorf("UserService.Find returned error: %v", err)
	}

	want := &User{Id: 1, Login: "shuheiktgw", Name: "shuheiktgw", GithubId: 1}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("UserService.Find returned %+v, want %+v", repo, want)
	}
}

func TestUserService_Sync(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/user/%d/sync", testUserId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1,"login":"shuheiktgw","name":"shuheiktgw","github_id":1}`)
	})

	repo, _, err := client.User.Sync(context.Background(), testUserId)

	if err != nil {
		t.Errorf("UserService.Sync returned error: %v", err)
	}

	want := &User{Id: 1, Login: "shuheiktgw", Name: "shuheiktgw", GithubId: 1}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("UserService.Sync returned %+v, want %+v", repo, want)
	}
}
