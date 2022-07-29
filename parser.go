// Copyright 2018-2022 VirtualTam.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package ccache

type Parser interface {
	// Parse reads machine-readable ccache statistics.
	Parse(text string) (*Statistics, error)
}
