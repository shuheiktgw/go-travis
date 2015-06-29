package travis

import (
	"fmt"
	"net/http"
)

type AuthenticationService struct {
	client *Client
}

type AccessToken string
type AccessTokenResponse struct {
	Token AccessToken `json:"access_token"`
}

func (as *AuthenticationService) UsingGithubToken(githubToken string) (AccessToken, *http.Response, error) {
	var u string = "/auth/github"
	var b map[string]string = map[string]string{"github_token": githubToken}

	req, err := as.client.NewRequest("POST", u, b, nil)
	if err != nil {
		return "", nil, err
	}

	atr := &AccessTokenResponse{}
	resp, err := as.client.Do(req, atr)
	if err != nil {
		return "", nil, err
	}

	as.UsingTravisToken(string(atr.Token))

	return atr.Token, resp, err
}

func (as *AuthenticationService) UsingTravisToken(travisToken string) error {
	if travisToken == "" {
		fmt.Errorf("unable to authenticate client; empty token provided")
	}

	as.client.Headers["Authorization"] = "token " + travisToken

	return nil
}
