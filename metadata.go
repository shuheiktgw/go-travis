// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

const (
	standardRepresentation = "standard"
	minimalRepresentation  = "minimal"
)

// Metadata is a metadata returned from the Travis CI API
type Metadata struct {
	// The type of data returned from the API
	Type *string `json:"@type,omitempty"`
	// The link for data returned from the API
	Href *string `json:"@href,omitempty"`
	// The representation of data returned from the API, standard or minimal
	Representation *string `json:"@representation,omitempty"`
	// The permissions of data returned from the API
	Permissions *Permissions `json:"@permissions,omitempty"`
}

// IsStandard tells if the struct is in a standard representation
func (m *Metadata) IsStandard() bool {
	return *m.Representation == standardRepresentation
}

// IsStandard tells if the struct is in a minimal representation
func (m *Metadata) IsMinimal() bool {
	return *m.Representation == minimalRepresentation
}
