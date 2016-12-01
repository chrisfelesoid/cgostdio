// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdio

/*
#include <stdio.h>
#include <errno.h>
#include <string.h>

void clear_errno() {
	errno = 0;
}
int get_errno() {
	return errno;
}
*/
import "C"

func ClearErrno() {
	C.clear_errno()
}

func GetErrno() int {
	return int(C.get_errno())
}

func Strerror(errno int) string {
	return C.GoString(C.strerror((C.int)(errno)))
}
