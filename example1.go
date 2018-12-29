// +build ignore

package main

import (
	"fmt"

	"github.com/glycerine/gopass"
)

func main() {
	fmt.Println("before")
	c := gopass.NewCleartextReader()
	input, err := c.ReadSlice()
	fmt.Println("after")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("read %q %v\n", string(input), input)
	}
}
