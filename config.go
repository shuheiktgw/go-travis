// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

// Config represents Travis CI job's configuration
type Config struct {
	Os            *string            `json:"os,omitempty"`
	Env           *string            `json:"env,omitempty"`
	Rvm           *string            `json:"rvm,omitempty"`
	Dist          *string            `json:"dist,omitempty"`
	Sudo          *bool              `json:"sudo,omitempty"`
	Cache         *map[string]bool   `json:"cache,omitempty"`
	Group         *string            `json:"group,omitempty"`
	Addons        *map[string]string `json:"addons,omitempty"`
	Script        *[]string          `json:"script,omitempty"`
	Result        *string            `json:".result,omitempty"`
	Language      *string            `json:"language,omitempty"`
	Services      *[]string          `json:"services,omitempty"`
	GlobalEnv     *string            `json:"global_env,omitempty"`
	BeforeScript  *[]string          `json:"before_script,omitempty"`
	BeforeInstall *[]string          `json:"before_install,omitempty"`
}
