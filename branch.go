package travis

import (
	"fmt"
	"net/url"

	"context"
	"net/http"

	"github.com/pkg/errors"
)

// BranchesService handles communication with the branch
// related methods of the Travis CI API.
type BranchService struct {
	client *Client
}

// Branch represents a branch of a GitHub repository
//
// https://developer.travis-ci.com/resource/branch#standard-representation
type Branch struct {
	Name           string     `json:"name,omitempty"`
	Repository     Repository `json:"repository,omitempty"`
	DefaultBranch  bool       `json:"default_branch,omitempty"`
	ExistsOnGithub bool       `json:"exists_on_github,omitempty"`
	LastBuild      Build      `json:"last_build,omitempty"`
}

// MinimalBranch included when the resource is returned as part of another resource
//
// https://developer.travis-ci.com/resource/branch#minimal-representation
type MinimalBranch struct {
	Name string `json:"name,omitempty"`
}

// BranchOption specifies the optional parameters for the
// BranchService.
type BranchOption struct {
	// Repository Id on Travis.
	// Do not confuse with a Repository Id on GitHub.
	RepositoryId uint `url:"repository_id,omitempty"`

	// GitHub owner name / GitHub repository name.
	// ex. "shuheiktgw/go-travis"
	Slug string `url:"slug,omitempty"`

	// GitHub branch name.
	BranchName string `url:"branch_name,omitempty"`
}

// RepoIdentifier returns repository's identifier, either repository id or slug
func (bo *BranchOption) RepoIdentifier() (string, error) {
	if bo.RepositoryId != 0 {
		return fmt.Sprint(bo.RepositoryId), nil
	}

	if bo.Slug != "" {
		return url.QueryEscape(bo.Slug), nil
	}

	return "", errors.New("missing repository identifier: you need to specify either repository id or slug")
}

// Find fetches a branch based on the provided option
//
// Travis CI API docs: https://developer.travis-ci.com/resource/branch#find
func (bs *BranchService) Find(ctx context.Context, bo *BranchOption) (*Branch, *http.Response, error) {
	ri, err := bo.RepoIdentifier()

	if err != nil {
		return nil, nil, err
	}

	u, err := urlWithOptions(fmt.Sprintf("/repo/%s/branch/%s", ri, bo.BranchName), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := bs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var branch Branch
	resp, err := bs.client.Do(ctx, req, &branch)
	if err != nil {
		return nil, resp, err
	}

	return &branch, resp, err
}
