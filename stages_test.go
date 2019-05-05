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

func TestStagesService_ListByBuild(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	var buildId uint = 10
	mux.HandleFunc(fmt.Sprintf("/build/%d/stages", buildId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testFormValues(t, r, values{"include": "stage.jobs"})
		fmt.Fprint(w, `{"stages": [{"id":1,"number":2,"name":"Test"}]}`)
	})

	opt := StagesOption{Include: []string{"stage.jobs"}}
	stages, _, err := client.Stages.ListByBuild(context.Background(), buildId, &opt)

	if err != nil {
		t.Errorf("Repository.List returned error: %v", err)
	}

	want := &Stage{Id: Uint(1), Number: Uint(2), Name: String("Test")}
	if !reflect.DeepEqual(stages[0], want) {
		t.Errorf("Stages.ListByBuild returned %+v, want %+v", stages[0], want)
	}
}
