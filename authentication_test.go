package travis

import (
	"context"
	"fmt"
	"testing"
)

func TestAuthenticate_UsingGithubToken_with_empty_token(t *testing.T) {
	as := &AuthenticationService{client: NewClient(defaultBaseURL, "")}

	_, _, err := as.UsingGithubToken(context.TODO(), "")
	notOk(t, err)
}

func TestAuthenticate_UsingTravisToken_with_empty_token(t *testing.T) {
	as := &AuthenticationService{client: NewClient(defaultBaseURL, "")}

	err := as.UsingTravisToken("")
	notOk(t, err)
}

func TestAuthenticate_UsingTravisToken_with_string_token(t *testing.T) {
	token := "abc123easyasdoremi"
	as := &AuthenticationService{client: NewClient(defaultBaseURL, token)}

	err := as.UsingTravisToken(token)
	ok(t, err)

	authHeader := as.client.Headers["Authorization"]
	assert(
		t,
		authHeader == fmt.Sprintf("token %s", token),
		fmt.Sprintf("travis token found in AuthenticationService.Headers: %s; expected %s", authHeader, token),
	)
}
