// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

// setup sets up a test HTTP server along with a travis.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle("/", mux)

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the GitHub client being tested and is
	// configured to use test server.
	client = NewClient("", "")
	u, _ := url.Parse(server.URL + "/")
	client.BaseURL = u

	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

type values map[string]string

func testFormValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Set(k, v)
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func TestClient_NewDefaultClient(t *testing.T) {
	c := NewDefaultClient("")

	assert(
		t,
		c.BaseURL.String() == defaultBaseURL,
		"Client.BaseURL = %s; expected %s", c.BaseURL.String(), defaultBaseURL,
	)
}

func TestClient_NewRequest(t *testing.T) {
	c := NewClient(defaultBaseURL, "")

	req, err := c.NewRequest(http.MethodGet, "/test", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	assert(
		t,
		req.Method == http.MethodGet,
		"Wrong Request Method set",
	)

	assert(
		t,
		req.URL.String() == defaultBaseURL+"test",
		"Wrong Request URL set",
	)

}

func TestClient_NewRequest_with_nil_headers_provided(t *testing.T) {
	baseUrl, _ := url.Parse(defaultBaseURL)
	c := NewClient(defaultBaseURL, "")

	req, err := c.NewRequest(http.MethodGet, "/users", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	assert(
		t,
		req.Header.Get("User-Agent") == userAgent,
		"Wrong Request User-Agent header set",
	)

	assert(
		t,
		req.Header.Get("Travis-API-Version") == apiVersion3,
		"Wrong Request Travis-API-Version header set",
	)

	assert(
		t,
		req.Header.Get("Host") == baseUrl.Host,
		"Wrong Request Host header set",
	)
}

func TestClient_NewRequest_with_non_overriding_headers_provided(t *testing.T) {
	baseUrl, _ := url.Parse(defaultBaseURL)
	c := NewClient(defaultBaseURL, "")
	h := map[string]string{
		"Abc": "123",
	}

	req, err := c.NewRequest(http.MethodGet, "/users", nil, h)
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
		req.Header.Get("User-Agent") == userAgent,
		"Wrong Request User-Agent header set",
	)

	assert(
		t,
		req.Header.Get("Travis-API-Version") == apiVersion3,
		"Wrong Request Travis-API-Version header set",
	)

	assert(
		t,
		req.Header.Get("Host") == baseUrl.Host,
		"Wrong Request Host header set",
	)
}

func TestClient_NewRequest_with_overriding_headers_provided(t *testing.T) {
	c := NewClient(defaultBaseURL, "")
	h := map[string]string{
		"Host": "api.travis-ci.com",
	}

	req, err := c.NewRequest(http.MethodGet, "/users", nil, h)
	if err != nil {
		t.Fatal(err)
	}

	assert(
		t,
		req.Header.Get("User-Agent") == userAgent,
		"Wrong Request User-Agent header set",
	)

	assert(
		t,
		req.Header.Get("Travis-API-Version") == apiVersion3,
		"Wrong Request Travis-API-Version header set",
	)

	assert(
		t,
		req.Header.Get("Host") == "api.travis-ci.com",
		"Wrong Request Host header set",
	)
}
