package travis

import (
	"testing"
)

func TestRepositoryOption_Identifier_Fail(t *testing.T) {
	op := RepositoryOption{}

	_, err := op.Identifier()

	if err == nil {
		t.Fatalf("error is not supposed to be nil")
	}
}

func TestRepositoryOption_Identifier_Success(t *testing.T) {
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
		op := RepositoryOption{Id: tc.id, Slug: tc.slug}

		id, err := op.Identifier()

		if err != nil {
			t.Fatalf("#%d Identifier fails: unexpected error occured: %s", i, err)
		}

		if id != tc.want {
			t.Fatalf("#%d Identifier fails: invalid identifier: want %s got %s", i, tc.want, id)
		}
	}
}
