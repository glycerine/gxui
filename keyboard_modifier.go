// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gxui

type KeyboardModifier int

const (
	ModNone    KeyboardModifier = 0
	ModShift   KeyboardModifier = 1
	ModControl KeyboardModifier = 2
	ModAlt     KeyboardModifier = 4
	ModSuper   KeyboardModifier = 8
)

func (m KeyboardModifier) Shift() bool {
	return m&ModShift != 0
}

func (m KeyboardModifier) Control() bool {
	return m&ModControl != 0
}

func (m KeyboardModifier) Alt() bool {
	return m&ModAlt != 0
}

func (m KeyboardModifier) Super() bool {
	return m&ModSuper != 0
}

func (m KeyboardModifier) String() string {
	var ev string
	if m.Control() {
		ev = concat(ev, "Ctrl")
	}
	if m.Super() {
		ev = concat(ev, platformSuper)
	}
	if m.Shift() {
		ev = concat(ev, "Shift")
	}
	if m.Alt() {
		ev = concat(ev, "Alt")
	}
	return ev
}

func concat(root, add string) string {
	if root != "" {
		root += "-"
	}
	return root + add
}
