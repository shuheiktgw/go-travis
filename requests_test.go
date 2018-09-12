package travis

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRequestsService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/requests", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"limit": "5", "offset": "5"})
		fmt.Fprint(w, `{"requests": [{"id":1,"state":"processed","result":"rejected"}]}`)
	})

	repos, _, err := client.Requests.FindByRepoId(context.Background(), testRepoId, &FindRequestsOption{Limit: 5, Offset: 5})

	if err != nil {
		t.Errorf("RequestsService.FindByRepoId returned error: %v", err)
	}

	want := []Request{{Id: 1, State: "processed", Result: "rejected"}}
	if !reflect.DeepEqual(repos, want) {
		t.Errorf("RequestsService.FindByRepoId returned %+v, want %+v", repos, want)
	}
}

func TestRequestsService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/requests", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"limit": "5", "offset": "5"})
		fmt.Fprint(w, `{"requests": [{"id":1,"state":"processed","result":"rejected"}]}`)
	})

	repos, _, err := client.Requests.FindByRepoSlug(context.Background(), testRepoSlug, &FindRequestsOption{Limit: 5, Offset: 5})

	if err != nil {
		t.Errorf("RequestsService.FindByRepoSlug returned error: %v", err)
	}

	want := []Request{{Id: 1, State: "processed", Result: "rejected"}}
	if !reflect.DeepEqual(repos, want) {
		t.Errorf("RequestsService.FindByRepoSlug returned %+v, want %+v", repos, want)
	}
}

func TestRequestsService_CreateByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/requests", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"config": "testConfig", "message": "testMessage", "branch": "master", "token": "testToken"})
		fmt.Fprint(w, `{"request": {"id":1,"message":"message!"}}`)
	})

	repo, _, err := client.Requests.CreateByRepoId(context.Background(), testRepoId, &CreateRequestOption{Config: "testConfig", Message: "testMessage", Branch: "master", Token: "testToken"})

	if err != nil {
		t.Errorf("RequestsService.CreateByRepoId returned error: %v", err)
	}

	want := &MinimalRequest{Id: 1, Message: "message!"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("RequestsService.CreateByRepoId returned %+v, want %+v", repo, want)
	}
}

func TestRequestsService_CreateByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/requests", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"config": "testConfig", "message": "testMessage", "branch": "master", "token": "testToken"})
		fmt.Fprint(w, `{"request": {"id":1,"message":"message!"}}`)
	})

	repo, _, err := client.Requests.CreateByRepoSlug(context.Background(), testRepoSlug, &CreateRequestOption{Config: "testConfig", Message: "testMessage", Branch: "master", Token: "testToken"})

	if err != nil {
		t.Errorf("RequestsService.CreateByRepoSlug returned error: %v", err)
	}

	want := &MinimalRequest{Id: 1, Message: "message!"}
	if !reflect.DeepEqual(repo, want) {
		t.Errorf("RequestsService.CreateByRepoSlug returned %+v, want %+v", repo, want)
	}
}
