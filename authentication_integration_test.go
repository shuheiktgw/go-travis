// +build integration

package travis

import (
	"context"
	"testing"
)

func TestAuthenticationService_UsingGithubToken(t *testing.T) {
	t.Parallel()

	token, res, err := integrationClient.Authentication.UsingGithubToken(context.TODO(), integrationGitHubToken)

	if err != nil {
		t.Fatalf("UsingGithubToken fails: %s", err)
	}

	if got, want := res.StatusCode, 200; got != want {
		t.Fatalf("UsingGithubToken fails: invalid http response %s", res.Status)
	}

	if token == "" {
		t.Fatalf("UsingGithubToken fails: token is empty")
	}
}
