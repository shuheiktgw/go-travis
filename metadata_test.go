// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

import (
	"testing"
)

func TestMetadata_isStandard(t *testing.T) {
	cases := []struct {
		representation string
		want           bool
	}{
		{"standard", true},
		{"minimal", false},
	}

	for i, c := range cases {
		m := Metadata{Representation: c.representation}

		if got := m.IsStandard(); got != c.want {
			t.Fatalf("#%d invalid: got %v, want: %v", i, got, c.want)
		}
	}
}

func TestMetadata_isMinimal(t *testing.T) {
	cases := []struct {
		representation string
		want           bool
	}{
		{"standard", false},
		{"minimal", true},
	}

	for i, c := range cases {
		m := Metadata{Representation: c.representation}

		if got := m.IsMinimal(); got != c.want {
			t.Fatalf("#%d invalid: got %v, want: %v", i, got, c.want)
		}
	}
}
