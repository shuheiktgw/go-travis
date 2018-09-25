// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

// PreferenceService handles communication with the
// preference related methods of the Travis CI API.
type PreferenceService struct {
	client *Client
}

// Preference is a standard representation of an individual preference
//
// Travis CI API docs: https://developer.travis-ci.com/resource/preference#standard-representation
type Preference struct {
	// The preference's name
	Name string `json:"name,omitempty"`
	// The preference's value
	Value bool `json:"value,omitempty"`
}
