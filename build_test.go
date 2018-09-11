package travis

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const testBuildId = 1

func TestBuildService_Find(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/build/%d", testBuildId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1,"number":"1","state":"created","duration":10}`)
	})

	build, _, err := client.Build.Find(context.Background(), testBuildId)

	if err != nil {
		t.Errorf("Build.Find returned error: %v", err)
	}

	want := &Build{Id: testBuildId, Number: "1", State: BuildStateCreated, Duration: 10}
	if !reflect.DeepEqual(build, want) {
		t.Errorf("Build.Find returned %+v, want %+v", build, want)
	}
}

func TestBuildService_Cancel(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/build/%d/cancel", testBuildId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"build":{"id":1,"number":"1","state":"created","duration":10}}`)
	})

	build, _, err := client.Build.Cancel(context.Background(), testBuildId)

	if err != nil {
		t.Errorf("Build.Cancel returned error: %v", err)
	}

	want := &MinimalBuild{Id: testBuildId, Number: "1", State: BuildStateCreated, Duration: 10}
	if !reflect.DeepEqual(build, want) {
		t.Errorf("Build.Cancel returned %+v, want %+v", build, want)
	}
}

func TestBuildService_Restart(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/build/%d/restart", testBuildId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"build":{"id":1,"number":"1","state":"created","duration":10}}`)
	})

	build, _, err := client.Build.Restart(context.Background(), testBuildId)

	if err != nil {
		t.Errorf("Build.Restart returned error: %v", err)
	}

	want := &MinimalBuild{Id: testBuildId, Number: "1", State: BuildStateCreated, Duration: 10}
	if !reflect.DeepEqual(build, want) {
		t.Errorf("Build.Restart returned %+v, want %+v", build, want)
	}
}
