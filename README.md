# BitReader

[![Build Status](https://app.travis-ci.com/kzeratal/bitreader.svg?branch=main)](https://app.travis-ci.com/kzeratal/bitreader) [![Coverage Status](https://coveralls.io/repos/github/kzeratal/bitreader/badge.svg?branch=main)](https://coveralls.io/github/kzeratal/bitreader?branch=main)

Reads bits from bytes

```go
package main

import (
	"fmt"

	"github.com/kzeratal/bitreader"
)

func main() {
	r := bitreader.NewReader([]byte{255})
	if byte, err := r.Read(7); err == nil {
		fmt.Println(int(byte)) // read 7 bits, output 127 and left 1 bit
	}
	if byte, err := r.Read(7); err == nil {
		fmt.Println(int(byte)) // there is only 1 bit left, so the output is 1
	}
}
```
