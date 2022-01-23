// Write a version of 'cat' in Go.
// Take the name of the file as a command line argument.
// Look up os.Args in the standard Go documentation
// Look up the 'Open' function in the 'os' package
// What interfaces does the 'File' type implement?
// Reuse io.Copy if File implements 'Reader'!

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [filename]")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	check(err)
	io.Copy(os.Stdout, file)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
