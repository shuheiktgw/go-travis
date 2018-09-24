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

func TestLintService_Lint(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/lint", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, `{"content":"foo:bar"}`+"\n")
		fmt.Fprint(w, `{"warnings":[{"key":["test"],"message":"test!"}]}`)
	})

	warnings, _, err := client.Lint.Lint(context.Background(), &TravisYml{Content: "foo:bar"})

	if err != nil {
		t.Errorf("Lint.Lint returned error: %v", err)
	}

	want := []Warning{{Key: []string{"test"}, Message: "test!"}}
	if !reflect.DeepEqual(warnings, want) {
		t.Errorf("Lint.Lint returned %+v, want %+v", warnings, want)
	}
}
