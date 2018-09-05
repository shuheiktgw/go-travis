package travis

import (
	"net/url"
	"testing"
)

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
		req.URL.String() == defaultBaseURL+"test",
		"Wrong Request URL set",
	)

}

func TestClient_NewRequest_with_nil_headers_provided(t *testing.T) {
	baseUrl, _ := url.Parse(defaultBaseURL)
	c := NewClient(defaultBaseURL, "")

	req, err := c.NewRequest("GET", "/users", nil, nil)
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

	req, err := c.NewRequest("GET", "/users", nil, h)
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
