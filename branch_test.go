package travis

import (
	"testing"
)

func TestBranchOption_RepoIdentifier_Fail(t *testing.T) {
	op := BranchOption{}

	_, err := op.RepoIdentifier()

	if err == nil {
		t.Fatalf("error is not supposed to be nil")
	}
}

func TestBranchOption_RepoIdentifier_Success(t *testing.T) {
	cases := []struct {
		id   uint
		slug string
		want string
	}{
		{id: 1, want: "1"},
		{slug: "shuheiktgw/go-travis", want: "shuheiktgw%2Fgo-travis"},
		{id: 1, slug: "shuheiktgw/go-travis", want: "1"},
	}

	for i, tc := range cases {
		op := BranchOption{RepositoryId: tc.id, Slug: tc.slug}

		id, err := op.RepoIdentifier()

		if err != nil {
			t.Fatalf("#%d RepoIdentifier fails: unexpected error occured: %s", i, err)
		}

		if id != tc.want {
			t.Fatalf("#%d RepoIdentifier fails: invalid identifier: want %s got %s", i, tc.want, id)
		}
	}
}
