package main

import (
	"fmt"
	"io"
	"net/http" // http is part of the net package
	"os"
)

// for the fourth iteration: we define our own thing that conforms to the Writer inferface
// but doesn't really do anything. His point is it's not like the type enforcement you get
// in bondage-and-discipline languages like Java, where you are placed in a straight jacket
// and your eyes are propped open like in A Clockwork Orange
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// first iteration: just print the body. We get a reference.
	// fmt.Println(resp.Body)

	// second iteration:
	// After long explanations of the Reader interface, we switch to declaring a byte slice
	// to pass to the read function to be filled up. We're using the builtin make() function
	// to do this.
	// bs := make([]byte, 99999)
	// // The response body implements the Reader interface
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// third iteration: just use io.Copy
	// this is lesson 65, "The Writer Interface"
	// io.Copy(os.Stdout, resp.Body)

	// fourth iteration: we kludge something that conforms to the interface, but does nothing
	lw := logWriter{}
	io.Copy(lw, resp.Body)

	fmt.Println("End of line.")
}

// fourth iteration: writing our own thing that conforms to the Writer interface
// func (logWriter) Write(bs []byte) (int, error) {
// 	return 1, nil
// }
// fourth iteration part two: do something
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
