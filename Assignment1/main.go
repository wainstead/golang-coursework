package main

import "fmt"

func main() {
	fmt.Println("Hello, sailor! This is assignment 1.")
	listy := []int{31337, 13, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, element := range listy {
		if element%2 == 0 {
			fmt.Println(element, " is even ")
		} else {
			fmt.Println(element, " is odd")
		}
	}

	/**
	Some additional hacking fun. Use printf style string formatting.
	*/
	for counter, element := range listy {
		fmt.Printf("%d: %d\n", counter, element)
	}
}
