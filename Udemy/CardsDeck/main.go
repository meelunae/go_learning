package main

func main() {
	deck := newDeck()
	deck.shuffle()
	deck.print()
	hand, remainingDeck := deal(deck, 5)
	hand.print()
	remainingDeck.print()
	deck.saveToFile("deck.txt")
}
