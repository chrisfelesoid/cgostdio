# Cgo Stdio

Cgo Stdio is replaced C language stdio.

# Using

```
go get github.com/chrisfelesoid/cgostdio
```

## for Mac OS X

"fmemopen" isn't defined on Max OS X. However, you can use ported library like this.

- https://github.com/NimbusKit/memorymapping


```
git clone git@github.com:NimbusKit/memorymapping.git
cd memorymapping/
clang -c -g0 -Wall src/fmemopen.c
ar rcs libfmemopen.a fmemopen.o
```

copy to library search path.(ex. /usr/local/)
```
cp libfmemopen.a /usr/local/lib/libfmemopen.a
cp src/fmemopen.h /usr/local/include/fmemopen.h
```


## Example

If you pass data to C.stdin

```
package main

/*
#include <stdio.h>

void echo(FILE* in, FILE *out) {
	int c;
	while((c = fgetc(in)) != EOF) {
		fputc(c, out);
	}
}

void echo_from_io() {
	echo(stdin, stdout);
}
*/
import "C"
import (
	"os"

	"github.com/chrisfelesoid/cgostdio"
)

func main() {
	fi, _ := os.Open("stdin.txt")
	// swap C.stdin as "stdin.txt"
	stdin := cgostdio.NewStdin(fi)
	// restore swapped C.stdin
	defer stdin.Close()

    // read "stdin.txt" and write stdout
	C.echo_from_io()
	C.fflush(C.stdout)
}
```


# License

BSD-style license found in the LICENSE file.