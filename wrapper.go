// Copyright 2018 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ccache

import "os/exec"
const (
	DefaultBinaryPath = "/usr/bin/ccache"
)

// Wrapper exposes supported ccache commands.
type Wrapper interface {
	ShowStats(string) (string, error)
	Version() (string, error)
}

// BinaryWrapper runs ccache commands locally.
type BinaryWrapper struct {
	path string
}

func (w *BinaryWrapper) exec(option string, env []string) (string, error) {
	cmd := exec.Command(w.path, option)
	cmd.Env = env
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out[:]), nil
}

// ShowStats returns the result of ``ccache --show-stats''.
func (w *BinaryWrapper) ShowStats(cacheDir string) (string, error) {

	return w.exec("--show-stats", []string{"CCACHE_DIR="+cacheDir})
}

// Version returns the result of ``ccache --version''.
func (w *BinaryWrapper) Version() (string, error) {
	return w.exec("--version",  []string{})
}

// NewBinaryWrapper ensures the ccache executable exists and can be invoked, and
// returns an initialized BinaryWrapper.
func NewBinaryWrapper(path string) (*BinaryWrapper, error) {
	err := exec.Command(path, "-s").Run()
	if err != nil {
		return &BinaryWrapper{}, err
	}

	return &BinaryWrapper{path: path}, nil
}
