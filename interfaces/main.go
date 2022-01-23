package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
	greeting string
}
type spanishBot struct {
	greeting string
}

type hungarianBot struct {
	greeting string
}

func main() {
	eb := englishBot{"hello sailor"}
	sb := spanishBot{"hola señor"}
	hb := hungarianBot{"szia!"}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(hb)
}

// Takes anything conforming to the 'bot' interface
// i.e. "To whom it may concern, if the thing you received
// conforms to my interface it walks and quacks just like me"
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// Note he removed the 'eb' in the first part of the declaration
// because we don't use it, and it's permissible under Go to
// do so.
func (eb englishBot) getGreeting() string {
	// VERY custom logic here
	return eb.greeting
}

func (sb spanishBot) getGreeting() string {
	// MUY custom logic here
	return sb.greeting
}

// Jan 17 2022: catching up again
func (sb hungarianBot) getGreeting() string {
	// NAGYON custom logic here
	return sb.greeting
}
