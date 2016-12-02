# Cgo Stdio

Cgo Stdio is replaced C language stdio.


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