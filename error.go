// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgostdio

import (
	"fmt"

	"github.com/chrisfelesoid/cgostdio/internal/stdio"
)

// CcallError is clang's error
type CcallError struct {
	message string
	errno   int
	strerr  string
}

func (e *CcallError) Error() string {
	return fmt.Sprintf("%s: errno=%d, strerr=%s", e.message, e.errno, e.strerr)
}

// NewCcallError creates clang's error
func NewCcallError(message string) error {
	errno := stdio.GetErrno()
	stdio.ClearErrno()

	return &CcallError{
		message: message,
		errno:   errno,
		strerr:  stdio.Strerror(errno),
	}
}
