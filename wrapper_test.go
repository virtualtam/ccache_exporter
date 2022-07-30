package ccache

import (
	"testing"

	"github.com/Masterminds/semver/v3"
)

type fakeCommand struct {
	version string
}

func (c *fakeCommand) ShowStats() (string, error) {
	return "", nil
}

func (c *fakeCommand) Version() (string, error) {
	return c.version, nil
}

func TestWrapperVersion(t *testing.T) {
	cases := []struct {
		tname      string
		cmdVersion string
		want       string
	}{
		{
			tname: "ccache 3.3.4",
			cmdVersion: `
ccache version 3.3.4

Copyright (C) 2002-2007 Andrew Tridgell
Copyright (C) 2009-2017 Joel Rosdahl

This program is free software; you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free Software
Foundation; either version 3 of the License, or (at your option) any later
version.
`,
			want: "3.3.4",
		},
		{
			tname: "ccache 4.6.1",
			cmdVersion: `ccache version 4.6.1
Features: file-storage http-storage redis-storage

Copyright (C) 2002-2007 Andrew Tridgell
Copyright (C) 2009-2022 Joel Rosdahl and other contributors

See <https://ccache.dev/credits.html> for a complete list of contributors.

This program is free software; you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free Software
Foundation; either version 3 of the License, or (at your option) any later
version.
`,
			want: "4.6.1",
		},
	}

	for _, tc := range cases {
		t.Run(tc.tname, func(t *testing.T) {
			cmd := &fakeCommand{
				version: tc.cmdVersion,
			}
			wrapper := NewWrapper(cmd)

			got, err := wrapper.Version()
			if err != nil {
				t.Fatalf("want no error, got %q", err)
			}

			wantVersion, err := semver.NewVersion(tc.want)
			if err != nil {
				t.Fatalf("failed to instantiate wantVersion: %q", err)
			}

			if !got.Equal(wantVersion) {
				t.Errorf("want version %q, got %q", wantVersion.String(), got.String())
			}
		})
	}
}