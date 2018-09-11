package travis

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

var (
	testRepoId   uint = 12345
	testRepoSlug      = url.QueryEscape("shuheiktgw/go-travis-test")
)

func TestBranchService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/branch/%s", testRepoId, "master"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name":"master","repository":{"id":1,"name":"test","slug":"shuheiktgw/test"},"default_branch":true,"exists_on_github":true}`)
	})

	branch, _, err := client.Branch.FindByRepoId(context.Background(), testRepoId, "master")

	if err != nil {
		t.Errorf("Branch.FindByRepoId returned error: %v", err)
	}

	want := &Branch{Name: "master", Repository: MinimalRepository{Id: 1, Name: "test", Slug: "shuheiktgw/test"}, DefaultBranch: true, ExistsOnGithub: true}
	if !reflect.DeepEqual(branch, want) {
		t.Errorf("Branch.FindByRepoId returned %+v, want %+v", branch, want)
	}
}

func TestBranchService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/branch/%s", testRepoSlug, "master"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name":"master","repository":{"id":1,"name":"test","slug":"shuheiktgw/test"},"default_branch":true,"exists_on_github":true}`)
	})

	branch, _, err := client.Branch.FindByRepoSlug(context.Background(), testRepoSlug, "master")

	if err != nil {
		t.Errorf("Branch.FindByRepoId returned error: %v", err)
	}

	want := &Branch{Name: "master", Repository: MinimalRepository{Id: 1, Name: "test", Slug: "shuheiktgw/test"}, DefaultBranch: true, ExistsOnGithub: true}
	if !reflect.DeepEqual(branch, want) {
		t.Errorf("Branch.FindByRepoId returned %+v, want %+v", branch, want)
	}
}
