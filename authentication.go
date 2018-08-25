package travis

import (
	"fmt"
	"net/http"
)

// BuildsService handles communication with the builds
// related methods of the Travis CI API.
type AuthenticationService struct {
	client *Client
}

type AccessToken string
type accessTokenResponse struct {
	Token AccessToken `json:"access_token"`
}

// UsingGithubToken will generate a Travis CI API authentication
// token and call the UsingTravisToken method with it, leaving your
// client authenticated and ready to use.
func (as *AuthenticationService) UsingGithubToken(githubToken string) (AccessToken, *http.Response, error) {
	if githubToken == "" {
		return "", nil, fmt.Errorf("unable to authenticate client; empty github token provided")
	}

	b := map[string]string{"github_token": githubToken}
	h := map[string]string{"Accept": V2_HEADER}

	req, err := as.client.NewRequest("POST", "/auth/github", b, h)
	if err != nil {
		return "", nil, err
	}

	atr := &accessTokenResponse{}
	resp, err := as.client.Do(req, atr)
	if err != nil {
		return "", nil, err
	}

	as.UsingTravisToken(string(atr.Token))

	return atr.Token, resp, err
}

// UsingTravisToken will format and write provided
// travisToken in the AuthenticationService client's headers.
func (as *AuthenticationService) UsingTravisToken(travisToken string) error {
	if travisToken == "" {
		return fmt.Errorf("unable to authenticate client; empty travis token provided")
	}

	as.client.Headers["Authorization"] = "token " + travisToken

	return nil
}
