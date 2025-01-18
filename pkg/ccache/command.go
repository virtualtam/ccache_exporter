// Copyright 2018-2022 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ccache

import "os/exec"

const (
	DefaultBinaryPath = "/usr/bin/ccache"
)

var _ Command = &LocalCommand{}

// Command exposes supported ccache commands.
type Command interface {
	PrintConfig() (string, error)
	ShowConfig() (string, error)

	PrintStats() (string, error)
	ShowStats() (string, error)

	Version() (string, error)
}

// LocalCommand runs ccache commands in a local shell.
type LocalCommand struct {
	path string
}

func (c *LocalCommand) exec(option string) (string, error) {
	out, err := exec.Command(c.path, option).Output()

	if err != nil {
		return "", err
	}

	return string(out[:]), nil
}

// PrintConfig returns the result of `ccache --print-config`.
//
// Available for ccache < 3.7
func (c *LocalCommand) PrintConfig() (string, error) {
	return c.exec("--print-config")
}

// ShowConfig returns the result of `ccache --show-config`.
//
// Available since ccache 3.7
func (c *LocalCommand) ShowConfig() (string, error) {
	return c.exec("--show-config")
}

// PrintStats returns the result of `ccache --print-stats`.
//
// Available since ccache 3.7
func (c *LocalCommand) PrintStats() (string, error) {
	return c.exec("--print-stats")
}

// ShowStats returns the result of “ccache --show-stats”.
func (c *LocalCommand) ShowStats() (string, error) {
	return c.exec("--show-stats")
}

// Version returns the result of “ccache --version”.
func (c *LocalCommand) Version() (string, error) {
	return c.exec("--version")
}

// NewLocalCommand ensures the ccache executable exists and can be invoked, and
// returns an initialized BinaryWrapper.
func NewLocalCommand(path string) (*LocalCommand, error) {
	err := exec.Command(path, "-s").Run()
	if err != nil {
		return &LocalCommand{}, err
	}

	return &LocalCommand{path: path}, nil
}
