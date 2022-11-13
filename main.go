package main

var deckSize int

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
