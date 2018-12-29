# getpasswd in Go [![GoDoc](https://godoc.org/github.com/glycerine/gopass?status.svg)](https://godoc.org/github.com/glycerine/gopass) [![Build Status](https://secure.travis-ci.org/glycerine/gopass.png?branch=master)](http://travis-ci.org/glycerine/gopass)

Retrieve password from user terminal or piped input without echo.

Verified on BSD, Linux, and Windows.

Example:
```go
package main

import "fmt"
import "github.com/glycerine/gopass"

func main() {
	fmt.Printf("Password: ")

	// Silent. For printing *'s use gopass.GetPasswdMasked()
	pass, err := gopass.GetPasswd()
	if err != nil {
		// Handle gopass.ErrInterrupted or getch() read error
	}

	// Do something with pass
}
```

Caution: Multi-byte characters not supported!

Forked from https://github.com/howeyc/gopass
