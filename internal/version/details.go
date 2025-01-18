// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package version

import (
	"time"

	"github.com/earthboundkid/versioninfo/v2"
)

// Details provides detailed build information about the application.
type Details struct {
	Short         string     `json:"short"`
	Revision      string     `json:"revision"`
	CommittedAt   *time.Time `json:"committed_at,omitempty"`
	DirtyBuild    bool       `json:"dirty_build"`
	CcacheVersion string     `json:"ccache_version"`
}

// NewDetails retrieves and returns build version Details.
func NewDetails(ccacheVersion string) *Details {
	v := &Details{
		Short:         versioninfo.Short(),
		Revision:      versioninfo.Revision,
		DirtyBuild:    versioninfo.DirtyBuild,
		CcacheVersion: ccacheVersion,
	}

	if !versioninfo.LastCommit.IsZero() {
		v.CommittedAt = &versioninfo.LastCommit
	}

	return v
}
