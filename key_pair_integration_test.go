// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build integration

package travis

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"testing"
)

var (
	integrationPrivateGitHubRepoSlug      = "shuheiktgwtest/go-travis-test-private"
	integrationPrivateGitHubRepoId   uint = 9221409
)

func TestKeyPairService_Integration_CRUD_ByRepoSlug(t *testing.T) {
	// Make sure no key is registered
	integrationClient.KeyPair.DeleteByRepoSlug(context.TODO(), integrationPrivateGitHubRepoSlug)

	v, err := generateKey()
	if err != nil {
		t.Fatalf("unexpected error occured while creating a key pair: %s", err)
	}

	// CreateByRepoSlug
	b := KeyPairBody{Description: "test", Value: v}
	k, res, err := integrationClient.KeyPair.CreateByRepoSlug(context.TODO(), integrationPrivateGitHubRepoSlug, &b)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if got, want := *k.Description, "test"; got != want {
		t.Fatalf("invalid description: gpt: %v, want: %v", got, want)
	}

	// FindByRepoSlug
	k, res, err = integrationClient.KeyPair.FindByRepoSlug(context.TODO(), integrationPrivateGitHubRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if got, want := *k.Description, "test"; got != want {
		t.Fatalf("invalid description: gpt: %v, want: %v", got, want)
	}

	// UpdateByRepoSlug
	b = KeyPairBody{Description: "updated"}
	k, res, err = integrationClient.KeyPair.UpdateByRepoSlug(context.TODO(), integrationPrivateGitHubRepoSlug, &b)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if got, want := *k.Description, "updated"; got != want {
		t.Fatalf("invalid description: gpt: %v, want: %v", got, want)
	}

	// DeleteByRepoSlug
	res, err = integrationClient.KeyPair.DeleteByRepoSlug(context.TODO(), integrationPrivateGitHubRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}

func TestKeyPairService_Integration_CRUD_ByRepoId(t *testing.T) {
	// Make sure no key is registered
	integrationClient.KeyPair.DeleteByRepoId(context.TODO(), integrationPrivateGitHubRepoId)

	v, err := generateKey()
	if err != nil {
		t.Fatalf("unexpected error occured while creating a key pair: %s", err)
	}

	// CreateByRepoSlug
	b := KeyPairBody{Description: "test", Value: v}
	k, res, err := integrationClient.KeyPair.CreateByRepoId(context.TODO(), integrationPrivateGitHubRepoId, &b)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if got, want := *k.Description, "test"; got != want {
		t.Fatalf("invalid description: gpt: %v, want: %v", got, want)
	}

	// FindByRepoSlug
	k, res, err = integrationClient.KeyPair.FindByRepoId(context.TODO(), integrationPrivateGitHubRepoId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if got, want := *k.Description, "test"; got != want {
		t.Fatalf("invalid description: gpt: %v, want: %v", got, want)
	}

	// UpdateByRepoSlug
	b = KeyPairBody{Description: "updated"}
	k, res, err = integrationClient.KeyPair.UpdateByRepoId(context.TODO(), integrationPrivateGitHubRepoId, &b)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	if got, want := *k.Description, "updated"; got != want {
		t.Fatalf("invalid description: gpt: %v, want: %v", got, want)
	}

	// DeleteByRepoSlug
	res, err = integrationClient.KeyPair.DeleteByRepoId(context.TODO(), integrationPrivateGitHubRepoId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}

func TestGeneratedKeyPairService_Integration_Create_Read_ByRepoSlug(t *testing.T) {
	// CreateGeneratedByRepoSlug
	_, res, err := integrationClient.GeneratedKeyPair.CreateByRepoSlug(context.TODO(), integrationPrivateGitHubRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	// FindGeneratedByRepoSlug
	_, res, err = integrationClient.GeneratedKeyPair.FindByRepoSlug(context.TODO(), integrationPrivateGitHubRepoSlug)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}

func TestGeneratedKeyPairService_Integration_Create_Read_ByRepoId(t *testing.T) {
	// CreateGeneratedByRepoId
	_, res, err := integrationClient.GeneratedKeyPair.CreateByRepoId(context.TODO(), integrationPrivateGitHubRepoId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("invalid http status: %s", res.Status)
	}

	// FindGeneratedByRepoId
	_, res, err = integrationClient.GeneratedKeyPair.FindByRepoId(context.TODO(), integrationPrivateGitHubRepoId)

	if err != nil {
		t.Fatalf("unexpected error occured: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("invalid http status: %s", res.Status)
	}
}

func generateKey() (string, error) {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		return "", err
	}

	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	return string(pem.EncodeToMemory(block)), nil
}
