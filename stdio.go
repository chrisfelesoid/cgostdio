// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgostdio

import (
	"bytes"
	"os"
	"sync"

	"github.com/chrisfelesoid/cgostdio/internal/stdio"
)

// Stdio is wrapped clang stdin/stdout/stderr
type Stdio struct {
	file  *os.File
	buf   *bytes.Buffer
	m     *sync.Mutex
	stdio *stdio.CStdio
}

func (s *Stdio) Close() error {
	defer s.m.Unlock()
	s.stdio.Clear()

	if ret := s.stdio.Close(); ret == stdio.EOF {
		return NewCcallError("close error")
	}
	if s.file != nil {
		return s.file.Close()
	}
	return nil
}

func (s *Stdio) Read(p []byte) (n int, err error) {
	if s.file != nil {
		return s.file.Read(p)
	}
	return s.buf.Read(p)
}

func (s *Stdio) Write(p []byte) (n int, err error) {
	if s.file != nil {
		return s.file.Write(p)
	}
	return s.buf.Write(p)
}

// Flush writes any buffered data
func (s *Stdio) Flush() error {
	if ret := s.stdio.Flush(); ret == stdio.EOF {
		return NewCcallError("flush error")
	}
	return nil
}
