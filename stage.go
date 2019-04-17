// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

// Stage is a standard representation of an individual stage
//
// Travis CI API docs: https://developer.travis-ci.com/resource/stage#standard-representation
type Stage struct {
	// Value uniquely identifying the stage
	Id uint `json:"id,omitempty"`
	// Incremental number for a stage
	Number uint `json:"number,omitempty"`
	// The name of the stage
	Name string `json:"name,omitempty"`
	// Current state of the stage
	State string `json:"state,omitempty"`
	// When the stage started
	StartedAt string `json:"started_at,omitempty"`
	// When the stage finished
	FinishedAt string `json:"finished_at,omitempty"`
	// The jobs of a stage.
	Jobs []*Job `json:"jobs,omitempty"`
	*Metadata
}
