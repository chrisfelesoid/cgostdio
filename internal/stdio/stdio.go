// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdio

/*
#include <stdio.h>
*/
import "C"

type CStdio struct {
	pstdio **C.FILE
	stdio  *CFile
	file   *CFile
}

var (
	stdin  *C.FILE
	stdout *C.FILE
	stderr *C.FILE
)

func (s *CStdio) Set(file *CFile) {
	s.file = file
	*(s.pstdio) = file.file
}

func (s *CStdio) Clear() {
	*(s.pstdio) = s.stdio.file
}

func (s *CStdio) Close() int {
	return s.file.Close()
}

func (s *CStdio) Flush() int {
	return s.file.Flush()
}

func NewStdin() *CStdio {
	f := &CFile{stdin, false}
	// TODO: check file descriptor?

	return &CStdio{
		pstdio: &C.stdin,
		stdio:  f,
	}
}

func NewStdout() *CStdio {
	f := &CFile{stdout, false}

	return &CStdio{
		pstdio: &C.stdout,
		stdio:  f,
	}
}

func NewStderr() *CStdio {
	f := &CFile{stderr, false}

	return &CStdio{
		pstdio: &C.stderr,
		stdio:  f,
	}
}

func init() {
	stdin = C.stdin
	stdout = C.stdout
	stderr = C.stdout
}
