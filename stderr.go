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

var mstderr *sync.Mutex

// NewStderr creates wrapped stderr
func NewStderr(file *os.File) *Stdio {
	mstderr.Lock()
	stderr := stdio.NewStderr()
	stderr.Set(stdio.Fdopen(file.Fd(), "wb", false))

	return &Stdio{
		file:  file,
		m:     mstderr,
		stdio: stderr,
	}
}

// NewStderrFromBuffer creates wrapped stderr from buffer
func NewStderrFromBuffer(data []byte) *Stdio {
	mstderr.Lock()
	stderr := stdio.NewStderr()
	stderr.Set(stdio.Fmemopen(data, "wb"))

	return &Stdio{
		buf:   bytes.NewBuffer(data),
		m:     mstderr,
		stdio: stderr,
	}
}

func init() {
	mstderr = new(sync.Mutex)
}
