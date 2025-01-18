// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

package metrics

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/virtualtam/ccache_exporter/internal/version"
)

type versionCollector struct {
	committedAtEpoch string
	isDirty          string
	revision         string
	version          string

	versionDesc *prometheus.Desc
}

func newVersionCollector(metricsPrefix string, versionDetails *version.Details) prometheus.Collector {
	var committedAtEpoch string

	if versionDetails.CommittedAt != nil && !versionDetails.CommittedAt.IsZero() {
		committedAtEpoch = strconv.FormatInt(versionDetails.CommittedAt.Unix(), 10)
	}

	return &versionCollector{
		committedAtEpoch: committedAtEpoch,
		isDirty:          strconv.FormatBool(versionDetails.DirtyBuild),
		revision:         versionDetails.Revision,
		version:          versionDetails.Short,

		versionDesc: prometheus.NewDesc(
			prometheus.BuildFQName(metricsPrefix, "", "version"),
			"Build Version",
			[]string{"committed_at_seconds", "is_dirty", "revision", "version"},
			nil,
		),
	}
}

// Describe publishes the description of each version metric to a metrics
// channel.
func (vc *versionCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- vc.versionDesc
}

// Collect returns version metrics.
func (vc *versionCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		vc.versionDesc,
		prometheus.UntypedValue,
		1,
		vc.committedAtEpoch,
		vc.isDirty,
		vc.revision,
		vc.version,
	)
}
