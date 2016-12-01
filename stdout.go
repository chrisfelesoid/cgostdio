// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgostdio

import (
	"os"
	"sync"

	"bytes"

	"github.com/chrisfelesoid/cgostdio/internal/stdio"
)

var mstdout *sync.Mutex

// NewStdout creates wrapped stdout
func NewStdout(file *os.File) *Stdio {
	mstdout.Lock()
	stdout := stdio.NewStdout()
	stdout.Set(stdio.Fdopen(file.Fd(), "wb", false))

	return &Stdio{
		file:  file,
		m:     mstdout,
		stdio: stdout,
	}
}

// NewStdoutFromBuffer creates wrapped stdout from buffer
func NewStdoutFromBuffer(data []byte) *Stdio {
	mstdout.Lock()
	stdout := stdio.NewStdout()
	stdout.Set(stdio.Fmemopen(data, "wb"))

	return &Stdio{
		buf:   bytes.NewBuffer(data),
		m:     mstdout,
		stdio: stdout,
	}
}

func init() {
	mstdout = new(sync.Mutex)
}
