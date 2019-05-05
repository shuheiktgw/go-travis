// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

// Commit is a standard representation of a GitHub commit as seen by Travis CI
//
// Travis CI API docs: https://developer.travis-ci.com/resource/commit#standard-representation
type Commit struct {
	// Value uniquely identifying the commit
	Id *uint `json:"id,omitempty"`
	// Checksum the commit has in git and is identified by
	Sha *string `json:"sha,omitempty"`
	// Named reference the commit has in git.
	Ref *string `json:"ref,omitempty"`
	// Commit message
	Message *string `json:"message,omitempty"`
	// URL to the commit's diff on GitHub
	CompareUrl *string `json:"compare_url,omitempty"`
	// Commit date from git
	CommittedAt *string `json:"committed_at,omitempty"`
	// Committer of the commit
	Committer *Committer `json:"committer,omitempty"`
	// Author of the commit
	Author *Author `json:"author,omitempty"`
	*Metadata
}

// Committer is a committer of the commit
type Committer struct {
	Name      string `json:"name,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

// Author is an author of the commit
type Author Committer
