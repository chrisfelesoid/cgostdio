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

var mstdin *sync.Mutex

// NewStdin creates wrapped stdin
func NewStdin(file *os.File) *Stdio {
	mstdin.Lock()
	stdin := stdio.NewStdin()
	stdin.Set(stdio.Fdopen(file.Fd(), "rb", false))

	return &Stdio{
		file:  file,
		m:     mstdin,
		stdio: stdin,
	}
}

// NewStdinFromBuffer creates wrapped stdin from buffer
func NewStdinFromBuffer(data []byte) *Stdio {
	mstdin.Lock()
	stdin := stdio.NewStdin()
	stdin.Set(stdio.Fmemopen(data, "rb"))

	return &Stdio{
		buf:   bytes.NewBuffer(data),
		m:     mstdin,
		stdio: stdin,
	}
}

func init() {
	mstdin = new(sync.Mutex)
}
