// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

// Tag is a standard representation of a Travis CI tag
type Tag struct {
	// Value uniquely identifying a repository of the build belongs to
	RepositoryId *uint `json:"repository_id"`
	// Name of the tag
	Name *string `json:"name,omitempty"`
	// Id of a last build on the branch
	LastBuildId *uint `json:"last_build_id"`
	*Metadata
}
