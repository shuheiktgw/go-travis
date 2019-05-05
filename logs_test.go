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

func TestLogsService_FindByJobId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/job/%d/log", testJobId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1,"content":"test"}`)
	})

	log, _, err := client.Logs.FindByJobId(context.Background(), testJobId)

	if err != nil {
		t.Errorf("Log.FindByJobId returned error: %v", err)
	}

	want := &Log{Id: Uint(1), Content: String("test")}
	if !reflect.DeepEqual(log, want) {
		t.Errorf("Log.FindByJobId returned %+v, want %+v", log, want)
	}
}

func TestLogsService_DeleteByJobId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/job/%d/log", testJobId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, `{"id":1,"content":"Log removed by XXX at 2017-02-13 16:00:00 UTC"}`)
	})

	log, _, err := client.Logs.DeleteByJobId(context.Background(), testJobId)

	if err != nil {
		t.Errorf("Log.DeleteByJobId returned error: %v", err)
	}

	want := &Log{Id: Uint(1), Content: String("Log removed by XXX at 2017-02-13 16:00:00 UTC")}
	if !reflect.DeepEqual(log, want) {
		t.Errorf("Log.DeleteByJobId returned %+v, want %+v", log, want)
	}
}
