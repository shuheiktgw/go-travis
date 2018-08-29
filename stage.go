package travis

// MinimalStage is a minimal representation of an individual stage
type MinimalStage struct {
	Id         uint   `json:"id,omitempty"`
	Number     uint   `json:"number,omitempty"`
	Name       string `json:"name,omitempty"`
	State      string `json:"state,omitempty"`
	StartedAt  string `json:"started_at,omitempty"`
	FinishedAt string `json:"finished_at,omitempty"`
}
