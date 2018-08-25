package travis

import (
	"context"
	"testing"
)

func TestJobFindOptions_IsValid_with_only_one_value_set(t *testing.T) {
	jfo := JobFindOptions{Ids: []uint{1234, 5678}}
	assert(
		t,
		jfo.IsValid() == true,
		"JobFindOptions.IsValid returned false; expected it to be true",
	)
}

func TestJobFindOptions_IsValid_with_multiple_values_set(t *testing.T) {
	jfo := JobFindOptions{Ids: []uint{1234, 5678}, State: "datqueue"}
	assert(
		t,
		jfo.IsValid() == false,
		"JobFindOptions.IsValid returned true; expected it to be false",
	)
}

func TestJobsService_Find_fails_with_invalid_opt(t *testing.T) {
	// No need to instantiate client. Should fail fast
	js := &JobsService{}
	opt := &JobFindOptions{Ids: []uint{1234, 5678}, State: "datqueue"}
	_, _, err := js.Find(context.TODO(), opt)
	notOk(t, err)
}
