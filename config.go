package travis

type Config struct {
	Os            string            `json:"os,omitempty"`
	Language      string            `json:"language,omitempty"`
	BeforeScript  []string          `json:"before_script,omitempty"`
	Script        string            `json:"script,omitempty"`
	AfterScript   []string          `json:"after_script,omitempty"`
	BeforeInstall []string          `json:"before_install,omitempty"`
	Install       []string          `json:"install,omitempty"`
	AfterSuccess  []string          `json:"after_success,omitempty"`
	AfterFailure  []string          `json:"after_failure,omitempty"`
	Addons        map[string]string `json:"addons,omitempty"`
	Notifications map[string]string `json:"notifications,omitempty"`
}
