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

func TestLogService_FindByJob(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/job/%d/log", testJobId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1,"content":"test"}`)
	})

	log, _, err := client.Log.FindByJob(context.Background(), testJobId)

	if err != nil {
		t.Errorf("Log.FindByJob returned error: %v", err)
	}

	want := &Log{Id: 1, Content: "test"}
	if !reflect.DeepEqual(log, want) {
		t.Errorf("Log.FindByJob returned %+v, want %+v", log, want)
	}
}
