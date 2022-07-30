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
	ShowStats() (string, error)
	Version() (string, error)
}

// LocalCommand runs ccache commands in a local shell.
type LocalCommand struct {
	path string
}

func (w *LocalCommand) exec(option string) (string, error) {
	out, err := exec.Command(w.path, option).Output()

	if err != nil {
		return "", err
	}

	return string(out[:]), nil
}

// ShowStats returns the result of ``ccache --show-stats''.
func (w *LocalCommand) ShowStats() (string, error) {
	return w.exec("--show-stats")
}

// Version returns the result of ``ccache --version''.
func (w *LocalCommand) Version() (string, error) {
	return w.exec("--version")
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
