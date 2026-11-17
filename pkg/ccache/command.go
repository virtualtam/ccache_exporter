// Copyright VirtualTam 2018, 2026
// SPDX-License-Identifier: MIT

package ccache

import (
	"context"
	"os/exec"
)

const (
	DefaultBinaryPath = "/usr/bin/ccache"
)

var _ Command = &LocalCommand{}

// Command exposes supported ccache commands.
type Command interface {
	PrintConfig(ctx context.Context) (string, error)
	ShowConfig(ctx context.Context) (string, error)

	PrintStats(ctx context.Context) (string, error)
	ShowStats(ctx context.Context) (string, error)

	Version(ctx context.Context) (string, error)
}

// LocalCommand runs ccache commands in a local shell.
type LocalCommand struct {
	path string
}

func (c *LocalCommand) exec(ctx context.Context, option string) (string, error) {
	out, err := exec.CommandContext(ctx, c.path, option).Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

// PrintConfig returns the result of `ccache --print-config`.
//
// Available for ccache < 3.7.
func (c *LocalCommand) PrintConfig(ctx context.Context) (string, error) {
	return c.exec(ctx, "--print-config")
}

// ShowConfig returns the result of `ccache --show-config`.
//
// Available since ccache 3.7.
func (c *LocalCommand) ShowConfig(ctx context.Context) (string, error) {
	return c.exec(ctx, "--show-config")
}

// PrintStats returns the result of `ccache --print-stats`.
//
// Available since ccache 3.7.
func (c *LocalCommand) PrintStats(ctx context.Context) (string, error) {
	return c.exec(ctx, "--print-stats")
}

// ShowStats returns the result of “ccache --show-stats”.
func (c *LocalCommand) ShowStats(ctx context.Context) (string, error) {
	return c.exec(ctx, "--show-stats")
}

// Version returns the result of “ccache --version”.
func (c *LocalCommand) Version(ctx context.Context) (string, error) {
	return c.exec(ctx, "--version")
}

// NewLocalCommand ensures the ccache executable exists and can be invoked, and
// returns an initialized BinaryWrapper.
func NewLocalCommand(ctx context.Context, path string) (*LocalCommand, error) {
	err := exec.CommandContext(ctx, path, "-s").Run()
	if err != nil {
		return &LocalCommand{}, err
	}

	return &LocalCommand{path: path}, nil
}
