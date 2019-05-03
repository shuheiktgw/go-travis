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

func TestMessagesService_ListByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/request/%d/messages", testRepoId, testRequestId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"offset": "5", "limit": "5"})
		fmt.Fprint(w, `{"messages":[{"id":1,"level":"info","key":"testKey","code":"testCode","args":{"test":"test"}}]}`)
	})

	opt := MessagesOption{Offset: 5, Limit: 5}
	messages, _, err := client.Messages.ListByRepoId(context.Background(), testRepoId, testRequestId, &opt)

	if err != nil {
		t.Errorf("Messages.ListByRepoId returned error: %v", err)
	}

	want := &Message{Id: 1, Level: "info", Key: "testKey", Code: "testCode", Args: []byte(`{"test":"test"}`)}
	if !reflect.DeepEqual(messages[0], want) {
		t.Errorf("Messages.ListByRepoId returned %+v, want %+v", messages[0], want)
	}
}

func TestMessagesService_ListByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/request/%d/messages", testRepoSlug, testRequestId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"offset": "5", "limit": "5"})
		fmt.Fprint(w, `{"messages":[{"id":1,"level":"info","key":"testKey","code":"testCode","args":{"test":"test"}}]}`)
	})

	opt := MessagesOption{Offset: 5, Limit: 5}
	messages, _, err := client.Messages.ListByRepoSlug(context.Background(), testRepoSlug, testRequestId, &opt)

	if err != nil {
		t.Errorf("Messages.ListByRepoId returned error: %v", err)
	}

	want := &Message{Id: 1, Level: "info", Key: "testKey", Code: "testCode", Args: []byte(`{"test":"test"}`)}
	if !reflect.DeepEqual(messages[0], want) {
		t.Errorf("Messages.ListByRepoId returned %+v, want %+v", messages[0], want)
	}
}
