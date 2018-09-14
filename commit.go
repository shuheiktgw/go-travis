package travis

// MinimalCommit is a minimal representation of a GitHub commit as seen by Travis CI
//
// Travis CI API docs: https://developer.travis-ci.com/resource/commit#minimal-representation
type MinimalCommit struct {
	// Value uniquely identifying the commit
	Id uint `json:"id,omitempty"`
	// Checksum the commit has in git and is identified by
	Sha string `json:"sha,omitempty"`
	// Named reference the commit has in git.
	Ref string `json:"ref,omitempty"`
	// Commit mesage
	Message string `json:"message,omitempty"`
	// URL to the commit's diff on GitHub
	CompareUrl string `json:"compare_url,omitempty"`
	// Commit date from git
	CommittedAt string `json:"committed_at,omitempty"`
}
