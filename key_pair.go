// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// KeyPairService handles communication with the key pair endpoints
// of Travis CI API
type KeyPairService struct {
	client *Client
}

// GeneratedKeyPairService handles communication with the key pair (generated) endpoints
// of Travis CI API
type GeneratedKeyPairService struct {
	client *Client
}

// KeyPairBody specifies options for
// creating and updating key pair.
type KeyPairBody struct {
	// A text description.
	Description string `json:"key_pair.description,omitempty"`
	// The private key.
	Value string `json:"key_pair.value,omitempty"`
}

// KeyPair is a standard representation of a public/private RSA key pair on Travis CI
//
// Travis CI API docs: https://developer.travis-ci.com/resource/key_pair#standard-representation
type KeyPair struct {
	// A text description.
	Description *string `json:"description,omitempty"`
	// The public key.
	PublicKey *string `json:"public_key,omitempty"`
	// The fingerprint.
	Fingerprint *string `json:"fingerprint,omitempty"`
	*Metadata
}

// FindByRepoId fetches the key pair based on the given repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/key_pair#find
func (ks *KeyPairService) FindByRepoId(ctx context.Context, repoId uint) (*KeyPair, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/key_pair", repoId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ks.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var keyPair KeyPair
	resp, err := ks.client.Do(ctx, req, &keyPair)
	if err != nil {
		return nil, resp, err
	}

	return &keyPair, resp, err
}

// FindByRepoSlug fetches the key pair based on the given repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/key_pair#find
func (ks *KeyPairService) FindByRepoSlug(ctx context.Context, repoSlug string) (*KeyPair, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/key_pair", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ks.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var keyPair KeyPair
	resp, err := ks.client.Do(ctx, req, &keyPair)
	if err != nil {
		return nil, resp, err
	}

	return &keyPair, resp, err
}

// CreateByRepoId creates the key pair based on the given repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/key_pair#create
func (ks *KeyPairService) CreateByRepoId(ctx context.Context, repoId uint, keyPair *KeyPairBody) (*KeyPair, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/key_pair", repoId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ks.client.NewRequest(http.MethodPost, u, keyPair, nil)
	if err != nil {
		return nil, nil, err
	}

	var k KeyPair
	resp, err := ks.client.Do(ctx, req, &k)
	if err != nil {
		return nil, resp, err
	}

	return &k, resp, err
}

// CreateByRepoSlug creates a new key pair based on the given repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/key_pair#create
func (ks *KeyPairService) CreateByRepoSlug(ctx context.Context, repoSlug string, keyPair *KeyPairBody) (*KeyPair, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/key_pair", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ks.client.NewRequest(http.MethodPost, u, keyPair, nil)
	if err != nil {
		return nil, nil, err
	}

	var k KeyPair
	resp, err := ks.client.Do(ctx, req, &k)
	if err != nil {
		return nil, resp, err
	}

	return &k, resp, err
}

// UpdateByRepoId updates the key pair variable based on the given option
//
// Travis CI API docs: https://developer.travis-ci.com/resource/key_pair#update
func (ks *KeyPairService) UpdateByRepoId(ctx context.Context, repoId uint, keyPair *KeyPairBody) (*KeyPair, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/key_pair", repoId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ks.client.NewRequest(http.MethodPatch, u, keyPair, nil)
	if err != nil {
		return nil, nil, err
	}

	var k KeyPair
	resp, err := ks.client.Do(ctx, req, &k)
	if err != nil {
		return nil, resp, err
	}

	return &k, resp, err
}

// UpdateByRepoSlug updates the key pair variable based on the given option
//
// Travis CI API docs: https://developer.travis-ci.com/resource/key_pair#update
func (ks *KeyPairService) UpdateByRepoSlug(ctx context.Context, repoSlug string, keyPair *KeyPairBody) (*KeyPair, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/key_pair", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ks.client.NewRequest(http.MethodPatch, u, keyPair, nil)
	if err != nil {
		return nil, nil, err
	}

	var k KeyPair
	resp, err := ks.client.Do(ctx, req, &k)
	if err != nil {
		return nil, resp, err
	}

	return &k, resp, err
}

// DeleteByRepoId deletes the key pair based on the given repository id and
//
// Travis CI API docs: https://developer.travis-ci.com/resource/key_pair#delete
func (ks *KeyPairService) DeleteByRepoId(ctx context.Context, repoId uint) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/key_pair", repoId), nil)
	if err != nil {
		return nil, err
	}

	req, err := ks.client.NewRequest(http.MethodDelete, u, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := ks.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// DeleteByRepoSlug deletes the key pair based on the given repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/key_pair#delete
func (ks *KeyPairService) DeleteByRepoSlug(ctx context.Context, repoSlug string) (*http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/key_pair", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, err
	}

	req, err := ks.client.NewRequest(http.MethodDelete, u, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := ks.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// FindByRepoId fetches the default key pair based on the given repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/key_pair_generated#find
func (ks *GeneratedKeyPairService) FindByRepoId(ctx context.Context, repoId uint) (*KeyPair, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/key_pair/generated", repoId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ks.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var keyPair KeyPair
	resp, err := ks.client.Do(ctx, req, &keyPair)
	if err != nil {
		return nil, resp, err
	}

	return &keyPair, resp, err
}

// FindByRepoSlug fetches the default key pair based on the given repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/key_pair_generated#find
func (ks *GeneratedKeyPairService) FindByRepoSlug(ctx context.Context, repoSlug string) (*KeyPair, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/key_pair/generated", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ks.client.NewRequest(http.MethodGet, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var keyPair KeyPair
	resp, err := ks.client.Do(ctx, req, &keyPair)
	if err != nil {
		return nil, resp, err
	}

	return &keyPair, resp, err
}

// CreateByRepoId creates the new default key pair based on the given repository id
//
// Travis CI API docs: https://developer.travis-ci.com/resource/key_pair_generated#create
func (ks *GeneratedKeyPairService) CreateByRepoId(ctx context.Context, repoId uint) (*KeyPair, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%d/key_pair/generated", repoId), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ks.client.NewRequest(http.MethodPost, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var k KeyPair
	resp, err := ks.client.Do(ctx, req, &k)
	if err != nil {
		return nil, resp, err
	}

	return &k, resp, err
}

// CreateByRepoSlug creates the new default key pair based on the given repository slug
//
// Travis CI API docs: https://developer.travis-ci.com/resource/key_pair_generated#create
func (ks *GeneratedKeyPairService) CreateByRepoSlug(ctx context.Context, repoSlug string) (*KeyPair, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("repo/%s/key_pair/generated", url.QueryEscape(repoSlug)), nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := ks.client.NewRequest(http.MethodPost, u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var k KeyPair
	resp, err := ks.client.Do(ctx, req, &k)
	if err != nil {
		return nil, resp, err
	}

	return &k, resp, err
}
