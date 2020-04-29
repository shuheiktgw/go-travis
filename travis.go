// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	ApiOrgUrl = "https://api.travis-ci.org/"
	ApiComUrl = "https://api.travis-ci.com/"

	userAgent          = "go-travis/" + version
	defaultContentType = "application/json"

	apiVersion3 = "3"
	mediaTypeV2 = "application/vnd.travis-ci.2.1+json"
)

// A Client manages communication with the Travis CI API.
type Client struct {
	// HTTP client used to communicate with the API
	HTTPClient *http.Client

	// Headers to attach to every requests made with the client.
	// As a default, Headers will be used to provide Travis API authentication
	// token and other necessary headers.
	// However these could be updated per-request through a parameters.
	Headers map[string]string

	// Base URL for api requests. Defaults to the public Travis API, but
	// can be set to an alternative endpoint to use with Travis Pro or Enterprise.
	// BaseURL should always be terminated by a slash.
	BaseURL *url.URL

	// User agent used when communicating with the Travis API
	UserAgent string

	// Services used to manipulate API entities
	Active                *ActiveService
	BetaFeatures          *BetaFeaturesService
	BetaMigrationRequests *BetaMigrationRequestsService
	Branches              *BranchesService
	Broadcasts            *BroadcastsService
	Builds                *BuildsService
	Caches                *CachesService
	Crons                 *CronsService
	EmailSubscriptions    *EmailSubscriptionsService
	EnvVars               *EnvVarsService
	GeneratedKeyPair      *GeneratedKeyPairService
	Installations         *InstallationsService
	KeyPair               *KeyPairService
	Jobs                  *JobsService
	Lint                  *LintService
	Logs                  *LogsService
	Messages              *MessagesService
	Organizations         *OrganizationsService
	Owner                 *OwnerService
	Preferences           *PreferencesService
	Repositories          *RepositoriesService
	Requests              *RequestsService
	Settings              *SettingsService
	Stages                *StagesService
	User                  *UserService
}

// NewClient returns a new Travis API client.
// If travisToken is not provided, the client can be authenticated at any time,
// using it's Authentication exposed service.
func NewClient(baseUrl string, travisToken string) *Client {
	bu, _ := url.Parse(baseUrl)
	bh := map[string]string{
		"Content-Type":       defaultContentType,
		"User-Agent":         userAgent,
		"Travis-API-Version": apiVersion3,
		"Host":               bu.Host,
	}

	c := &Client{
		HTTPClient: http.DefaultClient,
		Headers:    bh,
		BaseURL:    bu,
		UserAgent:  userAgent,
	}

	c.Active = &ActiveService{client: c}
	c.BetaFeatures = &BetaFeaturesService{client: c}
	c.BetaMigrationRequests = &BetaMigrationRequestsService{client: c}
	c.Branches = &BranchesService{client: c}
	c.Broadcasts = &BroadcastsService{client: c}
	c.Builds = &BuildsService{client: c}
	c.Caches = &CachesService{client: c}
	c.Crons = &CronsService{client: c}
	c.EmailSubscriptions = &EmailSubscriptionsService{client: c}
	c.EnvVars = &EnvVarsService{client: c}
	c.GeneratedKeyPair = &GeneratedKeyPairService{client: c}
	c.Installations = &InstallationsService{client: c}
	c.Jobs = &JobsService{client: c}
	c.KeyPair = &KeyPairService{client: c}
	c.Lint = &LintService{client: c}
	c.Logs = &LogsService{client: c}
	c.Messages = &MessagesService{client: c}
	c.Organizations = &OrganizationsService{client: c}
	c.Owner = &OwnerService{client: c}
	c.Preferences = &PreferencesService{client: c}
	c.Repositories = &RepositoriesService{client: c}
	c.Requests = &RequestsService{client: c}
	c.Settings = &SettingsService{client: c}
	c.Stages = &StagesService{client: c}
	c.User = &UserService{client: c}

	if travisToken != "" {
		c.SetToken(travisToken)
	}

	return c
}

// NewDefaultClient returns a new Travis API client bound to the public travis API.
// If travisToken is not provided, the client can be authenticated at any time,
// using it's Authentication exposed service.
func NewDefaultClient(travisToken string) *Client {
	return NewClient(ApiOrgUrl, travisToken)
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body. If specified, the map provided by headers will be used to udate
// request headers.
func (c *Client) NewRequest(method, urlStr string, body interface{}, headers map[string]string) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	var h = c.Headers
	if headers != nil {
		for k, v := range headers {
			h[k] = v
		}
	}

	for k, v := range h {
		req.Header.Set(k, v)
	}

	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = withContext(ctx, req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}

// SetToken formats and writes provided
// Travis API token in the client's headers.
func (c *Client) SetToken(token string) {
	c.Headers["Authorization"] = "token " + token
}

// IsAuthenticated indicates if Authorization headers were
// found in Client.Headers mapping.
func (c *Client) IsAuthenticated() bool {
	authHeader, ok := c.Headers["Authorization"]

	if !ok || (ok && authHeader == "token ") {
		return false
	}

	return true
}

// ErrorResponse reports an error caused by an API request.
// ErrorResponse implemented the Error interface.
//
// Travis CI API docs: https://developer.travis-ci.com/resource/error#error
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// The error's type
	ErrorType string `json:"error_type"`

	// The error's message
	ErrorMessage string `json:"error_message"`

	// The error's resource type
	ResourceType string `json:"resource_type"`

	// The error's permission
	Permission string `json:"permission"`
}

func (er *ErrorResponse) Error() string {
	return fmt.Sprintf(
		"%v %v: %d %v %v",
		er.Response.Request.Method,
		er.Response.Request.URL.String(),
		er.Response.StatusCode,
		er.ErrorType,
		er.ErrorMessage,
	)
}

// checkResponse checks the API response for errors; and returns them
// if present.
// A Response is considered an error if it has a status code outside the 2XX range.
func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}

	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		if err := json.Unmarshal(data, errorResponse); err != nil {
			errorResponse.ErrorMessage = string(data)
		}
	}

	return errorResponse
}

func urlWithOptions(s string, opt interface{}) (string, error) {
	rv := reflect.ValueOf(opt)
	if rv.Kind() == reflect.Ptr && rv.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func withContext(ctx context.Context, req *http.Request) *http.Request {
	return req.WithContext(ctx)
}

// Permissions represents permissions of Travis CI API
type Permissions map[string]bool

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Uint is a helper routine that allocates a new Uint value
// to store v and returns a pointer to it.
func Uint(v uint) *uint { return &v }

// Int64 is a helper routine that allocates a new Int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
