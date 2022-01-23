// Another student's work, and it's incorrect (see line 12)
package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	if len(os.Args) >= 2 {
		fileName := os.Args[1]
		file, err := os.Open(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		l := logWriter{}

		io.Copy(l, file)
	} else {
		fmt.Printf("Error: filename undefined \n")
		os.Exit(1)
	}

}
