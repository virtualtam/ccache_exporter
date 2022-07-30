// Copyright 2018-2022 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ccache

import (
	"errors"
	"regexp"

	"github.com/Masterminds/semver/v3"
)

var (
	ErrVersionMissing error = errors.New("command: missing version")
)

var (
	versionRegex = regexp.MustCompile("ccache version (.+)")
)

// Wrapper provides an abstraction for ccache commands.
type Wrapper struct {
	command Command
	parser  *LegacyParser
}

// NewWrapper initializes and returns a new Wrapper.
func NewWrapper(c Command) *Wrapper {
	p := NewLegacyParser()

	return &Wrapper{
		command: c,
		parser:  p,
	}
}

// Configuration returns the current cache configuration.
func (w *Wrapper) Configuration() (*Configuration, error) {
	out, err := w.command.ShowStats()
	if err != nil {
		return &Configuration{}, err
	}

	config, _, err := w.parser.ParseShowStats(out)
	if err != nil {
		return &Configuration{}, err
	}

	return config, nil
}

// Statistics returns the current ccache statistics.
func (w *Wrapper) Statistics() (*Statistics, error) {
	out, err := w.command.ShowStats()
	if err != nil {
		return &Statistics{}, err
	}

	_, stats, err := w.parser.ParseShowStats(out)
	if err != nil {
		return &Statistics{}, err
	}

	return stats, nil
}

// Version returns the semantic version for ccache.
func (w *Wrapper) Version() (*semver.Version, error) {
	out, err := w.command.Version()
	if err != nil {
		return &semver.Version{}, err
	}

	matches := versionRegex.FindStringSubmatch(out)
	if len(matches) != 2 {
		return &semver.Version{}, ErrVersionMissing
	}

	version, err := semver.NewVersion(matches[1])
	if err != nil {
		return &semver.Version{}, err
	}

	return version, nil
}
