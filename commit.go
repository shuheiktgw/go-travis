package travis

// MinimalCommit is a minimal representation of a GitHub commit as seen by Travis CI
type MinimalCommit struct {
	Id          uint   `json:"id,omitempty"`
	Sha         string `json:"sha,omitempty"`
	Ref         string `json:"ref,omitempty"`
	Message     string `json:"message,omitempty"`
	CompareUrl  string `json:"compare_url,omitempty"`
	CommittedAt string `json:"committed_at,omitempty"`
}
