// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdio

/*
#include <stdio.h>
#include <stdlib.h>
#include "memory.h"
*/
import "C"
import "unsafe"

type CFile struct {
	file      *C.FILE
	closeable bool
}

const (
	EOF int = C.EOF
)

func (f *CFile) Close() int {
	closeable := f.closeable
	f.closeable = false
	if closeable {
		return int(C.fclose(f.file))
	}
	return 0
}

func (f *CFile) Flush() int {
	return int(C.fflush(f.file))
}

func Fileno(file *CFile) int {
	return int(C.fileno(file.file))
}

func Fdopen(fd uintptr, mode string, closeable bool) *CFile {
	cmode := C.CString(mode)
	defer C.free(unsafe.Pointer(cmode))
	f := C.fdopen((C.int)(fd), cmode)
	return &CFile{f, closeable}
}

func Fmemopen(data []byte, mode string) *CFile {
	cmode := C.CString(mode)
	defer C.free(unsafe.Pointer(cmode))
	f := C.fmemopen(unsafe.Pointer(&data[0]), (C.size_t)(len(data)), cmode)
	return &CFile{f, true}
}
