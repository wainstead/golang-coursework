package main

import (
	"fmt"
)

// INITIALIZING: If we define a struct with two fields...
type foo struct {
	thing string
	ding  string
}

type bar struct{} // This is just an empty struct, which is legal

type barBlee interface {
	barSay() string
	laughDamnit() string
}

func (b bar) barSay() string {
	return "ha ha only bar-serious!"
}

func (b bar) laughDamnit() string {
	return "well I say YA HA HA HA HA HA HA HA"
}

func yammer(b barBlee) {
	fmt.Println(b.barSay())
}

func main() {
	fmt.Println("Hello sailor!")
	// INITIALIZING: ...we initialize it this way; to do anything else is a compile error
	var f = foo{"hippy", "is there a better way?"}
	fmt.Println(f)

	// Arrays are low level and you should prefer slices to arrays, as a rule of thumb
	licks := [3]string{"one", "two", "three"}
	fmt.Println(licks)
	licks[0] = "uno"
	fmt.Println(licks)

	var claps = []string{"one", "two", "three", "four"}
	fmt.Println(claps)

	b := bar{}
	bb := bar{}
	yammer(b)
	yammer(bb)

	xy := X{}
	yy := Y{}
	Gogo(xy)
	Gogo(yy)
}
