// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const (
	testKeyPairDescription = "My key pair"
	testKeyPairPublicKey   = "-----BEGIN PUBLIC KEY----- abcdefg -----END PUBLIC KEY-----"
	testKeyPairPrivateKey  = "-----BEGIN RSA PRIVATE KEY----- hijklmnop -----END RSA PRIVATE KEY-----"
	testKeyPairFingerprint = "a1:b2:c3:d4:e5:d6:f7:g8:h9:i0:j1:k2:l3:m4:n5:o6"
)

func TestKeyPairService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/key_pair", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"description":"%s","public_key":"%s","fingerprint":"%s"}`,
			testKeyPairDescription, testKeyPairPublicKey, testKeyPairFingerprint)
	})

	keyPair, _, err := client.KeyPair.FindByRepoId(context.Background(), testRepoId)

	if err != nil {
		t.Errorf("KeyPair.FindByRepoId returned error: %v", err)
	}

	want := &KeyPair{
		Description: String(testKeyPairDescription),
		PublicKey:   String(testKeyPairPublicKey),
		Fingerprint: String(testKeyPairFingerprint),
	}
	if !reflect.DeepEqual(keyPair, want) {
		t.Errorf("KeyPair.FindByRepoId returned %+v, want %+v", keyPair, want)
	}
}

func TestKeyPairService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/key_pair", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"description":"%s","public_key":"%s","fingerprint":"%s"}`,
			testKeyPairDescription, testKeyPairPublicKey, testKeyPairFingerprint)
	})

	keyPair, _, err := client.KeyPair.FindByRepoSlug(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("KeyPair.FindByRepoSlug returned error: %v", err)
	}

	want := &KeyPair{
		Description: String(testKeyPairDescription),
		PublicKey:   String(testKeyPairPublicKey),
		Fingerprint: String(testKeyPairFingerprint),
	}
	if !reflect.DeepEqual(keyPair, want) {
		t.Errorf("KeyPair.FindByRepoSlug returned %+v, want %+v", keyPair, want)
	}
}

func TestKeyPairService_CreateByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/key_pair", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, fmt.Sprintf(`{"key_pair.description":"%s","key_pair.value":"%s"}`+"\n",
			testKeyPairDescription, testKeyPairPrivateKey))
		fmt.Fprintf(w, `{"description":"%s","public_key":"%s","fingerprint":"%s"}`,
			testKeyPairDescription, testKeyPairPublicKey, testKeyPairFingerprint)
	})

	opt := KeyPairBody{Description: testKeyPairDescription, Value: testKeyPairPrivateKey}
	keyPair, _, err := client.KeyPair.CreateByRepoId(context.Background(), testRepoId, &opt)

	if err != nil {
		t.Errorf("KeyPair.CreateByRepoId returned error: %v", err)
	}

	want := &KeyPair{
		Description: String(testKeyPairDescription),
		PublicKey:   String(testKeyPairPublicKey),
		Fingerprint: String(testKeyPairFingerprint),
	}
	if !reflect.DeepEqual(keyPair, want) {
		t.Errorf("KeyPair.CreateByRepoId returned %+v, want %+v", keyPair, want)
	}
}

func TestKeyPairService_CreateByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/key_pair", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBody(t, r, fmt.Sprintf(`{"key_pair.description":"%s","key_pair.value":"%s"}`+"\n",
			testKeyPairDescription, testKeyPairPrivateKey))
		fmt.Fprintf(w, `{"description":"%s","public_key":"%s","fingerprint":"%s"}`,
			testKeyPairDescription, testKeyPairPublicKey, testKeyPairFingerprint)
	})

	opt := KeyPairBody{Description: testKeyPairDescription, Value: testKeyPairPrivateKey}
	keyPair, _, err := client.KeyPair.CreateByRepoSlug(context.Background(), testRepoSlug, &opt)

	if err != nil {
		t.Errorf("KeyPair.CreateByRepoSlug returned error: %v", err)
	}

	want := &KeyPair{
		Description: String(testKeyPairDescription),
		PublicKey:   String(testKeyPairPublicKey),
		Fingerprint: String(testKeyPairFingerprint),
	}
	if !reflect.DeepEqual(keyPair, want) {
		t.Errorf("KeyPair.CreateByRepoSlug returned %+v, want %+v", keyPair, want)
	}
}

func TestKeyPairService_UpdateByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/key_pair", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testBody(t, r, fmt.Sprintf(`{"key_pair.description":"%s","key_pair.value":"%s"}`+"\n",
			testKeyPairDescription, testKeyPairPrivateKey))
		fmt.Fprintf(w, `{"description":"%s","public_key":"%s","fingerprint":"%s"}`,
			testKeyPairDescription, testKeyPairPublicKey, testKeyPairFingerprint)
	})

	opt := KeyPairBody{Description: testKeyPairDescription, Value: testKeyPairPrivateKey}
	keyPair, _, err := client.KeyPair.UpdateByRepoId(context.Background(), testRepoId, &opt)

	if err != nil {
		t.Errorf("KeyPair.UpdateByRepoId returned error: %v", err)
	}

	want := &KeyPair{
		Description: String(testKeyPairDescription),
		PublicKey:   String(testKeyPairPublicKey),
		Fingerprint: String(testKeyPairFingerprint),
	}
	if !reflect.DeepEqual(keyPair, want) {
		t.Errorf("KeyPair.UpdateByRepoId returned %+v, want %+v", keyPair, want)
	}
}

func TestKeyPairService_UpdateByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/key_pair", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testBody(t, r, fmt.Sprintf(`{"key_pair.description":"%s","key_pair.value":"%s"}`+"\n",
			testKeyPairDescription, testKeyPairPrivateKey))
		fmt.Fprintf(w, `{"description":"%s","public_key":"%s","fingerprint":"%s"}`,
			testKeyPairDescription, testKeyPairPublicKey, testKeyPairFingerprint)
	})

	opt := KeyPairBody{Description: testKeyPairDescription, Value: testKeyPairPrivateKey}
	keyPair, _, err := client.KeyPair.UpdateByRepoSlug(context.Background(), testRepoSlug, &opt)

	if err != nil {
		t.Errorf("KeyPair.UpdateByRepoSlug returned error: %v", err)
	}

	want := &KeyPair{
		Description: String(testKeyPairDescription),
		PublicKey:   String(testKeyPairPublicKey),
		Fingerprint: String(testKeyPairFingerprint),
	}
	if !reflect.DeepEqual(keyPair, want) {
		t.Errorf("KeyPair.UpdateByRepoSlug returned %+v, want %+v", keyPair, want)
	}
}

func TestKeyPairService_DeleteByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/key_pair", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, `{}`)
	})

	_, err := client.KeyPair.DeleteByRepoId(context.Background(), testRepoId)

	if err != nil {
		t.Errorf("KeyPair.DeleteByRepoId returned error: %v", err)
	}
}

func TestKeyPairService_DeleteByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/key_pair", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, `{}`)
	})

	_, err := client.KeyPair.DeleteByRepoSlug(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("KeyPair.DeleteByRepoSlug returned error: %v", err)
	}
}

func TestGeneratedKeyPairService_FindByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/key_pair/generated", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"description":"%s","public_key":"%s","fingerprint":"%s"}`,
			testKeyPairDescription, testKeyPairPublicKey, testKeyPairFingerprint)
	})

	keyPair, _, err := client.GeneratedKeyPair.FindByRepoId(context.Background(), testRepoId)

	if err != nil {
		t.Errorf("KeyPair.FindGeneratedByRepoId returned error: %v", err)
	}

	want := &KeyPair{
		Description: String(testKeyPairDescription),
		PublicKey:   String(testKeyPairPublicKey),
		Fingerprint: String(testKeyPairFingerprint),
	}
	if !reflect.DeepEqual(keyPair, want) {
		t.Errorf("KeyPair.FindGeneratedByRepoId returned %+v, want %+v", keyPair, want)
	}
}

func TestGeneratedKeyPairService_FindByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/key_pair/generated", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"description":"%s","public_key":"%s","fingerprint":"%s"}`,
			testKeyPairDescription, testKeyPairPublicKey, testKeyPairFingerprint)
	})

	keyPair, _, err := client.GeneratedKeyPair.FindByRepoSlug(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("KeyPair.FindGeneratedByRepoSlug returned error: %v", err)
	}

	want := &KeyPair{
		Description: String(testKeyPairDescription),
		PublicKey:   String(testKeyPairPublicKey),
		Fingerprint: String(testKeyPairFingerprint),
	}
	if !reflect.DeepEqual(keyPair, want) {
		t.Errorf("KeyPair.FindGeneratedByRepoSlug returned %+v, want %+v", keyPair, want)
	}
}

func TestGeneratedKeyPair_CreateByRepoId(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%d/key_pair/generated", testRepoId), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `{"description":"%s","public_key":"%s","fingerprint":"%s"}`,
			testKeyPairDescription, testKeyPairPublicKey, testKeyPairFingerprint)
	})

	keyPair, _, err := client.GeneratedKeyPair.CreateByRepoId(context.Background(), testRepoId)

	if err != nil {
		t.Errorf("KeyPair.CreateGeneratedByRepoId returned error: %v", err)
	}

	want := &KeyPair{
		Description: String(testKeyPairDescription),
		PublicKey:   String(testKeyPairPublicKey),
		Fingerprint: String(testKeyPairFingerprint),
	}
	if !reflect.DeepEqual(keyPair, want) {
		t.Errorf("KeyPair.CreateGeneratedByRepoId returned %+v, want %+v", keyPair, want)
	}
}

func TestGeneratedKeyPairService_CreateByRepoSlug(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/repo/%s/key_pair/generated", testRepoSlug), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `{"description":"%s","public_key":"%s","fingerprint":"%s"}`,
			testKeyPairDescription, testKeyPairPublicKey, testKeyPairFingerprint)
	})

	keyPair, _, err := client.GeneratedKeyPair.CreateByRepoSlug(context.Background(), testRepoSlug)

	if err != nil {
		t.Errorf("KeyPair.CreateByRepoSlug returned error: %v", err)
	}

	want := &KeyPair{
		Description: String(testKeyPairDescription),
		PublicKey:   String(testKeyPairPublicKey),
		Fingerprint: String(testKeyPairFingerprint),
	}
	if !reflect.DeepEqual(keyPair, want) {
		t.Errorf("KeyPair.CreateByRepoSlug returned %+v, want %+v", keyPair, want)
	}
}
