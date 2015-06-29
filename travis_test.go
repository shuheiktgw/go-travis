// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"net/url"
	"testing"
)

func TestClient_NewRequest(t *testing.T) {
	c := NewClient(TRAVIS_API_DEFAULT_URL, "")

	req, err := c.NewRequest("GET", "/test", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	assert(
		t,
		req.Method == "GET",
		"Wrong Request Method set",
	)

	assert(
		t,
		req.URL.String() == TRAVIS_API_DEFAULT_URL+"test",
		"Wrong Request URL set",
	)

}

func TestClient_NewRequest_with_nil_headers_provided(t *testing.T) {
	baseUrl, _ := url.Parse(TRAVIS_API_DEFAULT_URL)
	c := NewClient(TRAVIS_API_DEFAULT_URL, "")

	req, err := c.NewRequest("GET", "/users", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	assert(
		t,
		req.Header.Get("User-Agent") == TRAVIS_USER_AGENT,
		"Wrong Request User-Agent header set",
	)

	assert(
		t,
		req.Header.Get("Accept") == TRAVIS_REQUEST_ACCEPT_HEADER,
		"Wrong Request Accept header set",
	)

	assert(
		t,
		req.Header.Get("Host") == baseUrl.Host,
		"Wrong Request Host header set",
	)
}

func TestClient_NewRequest_with_non_overriding_headers_provided(t *testing.T) {
	baseUrl, _ := url.Parse(TRAVIS_API_DEFAULT_URL)
	c := NewClient(TRAVIS_API_DEFAULT_URL, "")
	h := map[string]string{
		"Abc": "123",
	}

	req, err := c.NewRequest("GET", "/users", nil, h)
	if err != nil {
		t.Fatal(err)
	}

	assert(
		t,
		req.Header.Get("Abc") == "123",
		"Wrong Request Abc header set",
	)

	assert(
		t,
		req.Header.Get("User-Agent") == TRAVIS_USER_AGENT,
		"Wrong Request User-Agent header set",
	)

	assert(
		t,
		req.Header.Get("Accept") == TRAVIS_REQUEST_ACCEPT_HEADER,
		"Wrong Request Accept header set",
	)

	assert(
		t,
		req.Header.Get("Host") == baseUrl.Host,
		"Wrong Request Host header set",
	)
}

func TestClient_NewRequest_with_overriding_headers_provided(t *testing.T) {
	c := NewClient(TRAVIS_API_DEFAULT_URL, "")
	h := map[string]string{
		"Host": "api.travis-ci.com",
	}

	req, err := c.NewRequest("GET", "/users", nil, h)
	if err != nil {
		t.Fatal(err)
	}

	assert(
		t,
		req.Header.Get("User-Agent") == TRAVIS_USER_AGENT,
		"Wrong Request User-Agent header set",
	)

	assert(
		t,
		req.Header.Get("Accept") == TRAVIS_REQUEST_ACCEPT_HEADER,
		"Wrong Request Accept header set",
	)

	assert(
		t,
		req.Header.Get("Host") == "api.travis-ci.com",
		"Wrong Request Host header set",
	)
}
