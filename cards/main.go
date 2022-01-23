package main

// Declare and assign a variable. Basic types include bool, string, int, float64.
// Long form:
// var card string = "Ace of Spaces"
// Short form:
// Only use := when defining a new variable, natch.
// card := "Queen of Hearts"
// illegal
// card := "Jack of Clubs"
// but we can assign a new value with plain old '='
//card = "King of Diamonds"
// card := newCard()
// fmt.Println(card)
// printState()
//cards := []string{"Ace of Diamonds", newCard()}
//cards := deck{"Ace of Spaces", newCard()}
//cards = append(cards, "Queen of Hearts")

// We use := syntax here because we get a new index and card on every iteration
// for i, card := range cards {
// 	fmt.Println(i, card)
// }
func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// fmt.Println(cards.toString())
	//cards.saveToFile("/tmp/my_cards.csv")
	// cards := newDeckFromFile("/tmp/my_cards.csv")
	// cards.print()
	// greeting := "Hello sailor!"
	// fmt.Println([]byte(greeting))
	//cards.print()
	// legal but not so readable
	//fmt.Println(cards)
}

// Go requires us to declare the return type
func newCard() string {
	return "Five of Diamonds"
}
