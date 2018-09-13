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
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1,"login":"shuheiktgw","name":"shuheiktgw","github_id":1}`)
	})

	repo, _, err := client.User.Current(context.Background())

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
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1,"login":"shuheiktgw","name":"shuheiktgw","github_id":1}`)
	})

	repo, _, err := client.User.Find(context.Background(), testUserId)

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
		testMethod(t, r, "POST")
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
