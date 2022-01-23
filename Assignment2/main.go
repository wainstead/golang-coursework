package main

// This is assignment 2, from Section 6: Interfaces

import "fmt"

// Types
type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

// Main

func main() {
	sq := square{10.0}
	tr := triangle{5.0, 6.0}
	printArea(sq)
	printArea(tr)
}

// Functions

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
