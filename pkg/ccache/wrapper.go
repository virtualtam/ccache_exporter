// Copyright (c) VirtualTam
// SPDX-License-Identifier: MIT

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
	versionRegex                    = regexp.MustCompile("ccache version (.+)")
	useLegacyParserForVersionsBelow = semver.MustParse("3.7")
)

// Wrapper provides an abstraction for ccache commands.
type Wrapper struct {
	command    Command
	version    semver.Version
	versionStr string
}

// NewWrapper initializes and returns a new Wrapper.
func NewWrapper(c Command) *Wrapper {
	w := &Wrapper{
		command: c,
	}

	v, err := w.ParseVersion()
	if err != nil {
		panic(err)
	}

	w.version = *v
	w.versionStr = v.Original()

	return w
}

// Configuration returns the current ccache configuration.
func (w *Wrapper) Configuration() (*Configuration, error) {
	var out string
	var err error

	if w.version.LessThan(useLegacyParserForVersionsBelow) {
		out, err = w.command.PrintConfig()
	} else {
		out, err = w.command.ShowConfig()
	}

	if err != nil {
		return &Configuration{}, err
	}

	return ParseConfiguration(out)
}

// Statistics returns the current ccache statistics.
func (w *Wrapper) Statistics() (*Statistics, error) {
	if w.version.LessThan(useLegacyParserForVersionsBelow) {
		return w.legacyStatistics()
	}

	return w.tsvStatistics()
}

func (w *Wrapper) legacyStatistics() (*Statistics, error) {
	out, err := w.command.ShowStats()
	if err != nil {
		return &Statistics{}, err
	}

	stats, err := ParsePre37Statistics(out)

	return stats, err
}

func (w *Wrapper) tsvStatistics() (*Statistics, error) {
	out, err := w.command.PrintStats()
	if err != nil {
		return &Statistics{}, err
	}

	return ParseTSVStatistics(out)
}

// ParseVersion parses the semantic version for ccache.
func (w *Wrapper) ParseVersion() (*semver.Version, error) {
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

// Version returns the parsed version for ccache.
func (w *Wrapper) Version() string {
	return w.versionStr
}
