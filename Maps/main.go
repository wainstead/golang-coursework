package main

import "fmt"

func main() {
	fmt.Println("Hello, sailor!")
	fmt.Println("Need a map?!")
	// map[KeyType]ValueType
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#745745",
		"blue":  "#00ff00",
	}
	printMap(colors)
	fanciermaps()
}

func printMap(c map[string]string) {
	fmt.Println("Printing map:")
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}
}
func fanciermaps() {
	var colorsv map[string]string
	fmt.Println(colorsv) // prints an empty map

	colors := make(map[string]string)
	colors["dog"] = "Fido"
	colors["cat"] = "Fluffy"
	colors["parrot"] = "Polly"
	fmt.Println(colors)

	delete(colors, "dog")
	fmt.Println(colors)
}
