// Our namespace. There are two kinds of packages: executable and reusable.
// The name of the package you use determines if your thing is executable or reusable.
// "main" has a special magic.
// package blahblah would build as reusable. Any other name builds a reusable.
// And package main must somewhere have a function "main" in it, same as C.
package main

// load a package
// see more at golang.org/pkg
import "fmt"

func main() {
	fmt.Println("Hello, sailor!")
}

/*
Five questions to answer from this simple program:

1. How do we run the code in our project? "go run main.go" (see also "go build main.go")
2. What does 'package main' mean? It's our namespace, and "main" is magical
3. What does 'import "fmt"' mean? It means pull in a function library
4. What's that 'func' thing? A keyword indicating you're declaring a function
5. How is this main.go file organized? Just as we see it: pacakge name, then import statements, then functions

*/
