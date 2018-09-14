package travis

// MinimalStage is a minimal representation of an individual stage
//
// Travis CI API docs: https://developer.travis-ci.com/resource/stage#minimal-representation
type MinimalStage struct {
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
}
