package main

func main() {
	// cards := newDeck()
	// cards.toFile("my_cards")
	cards := fromFile("my_cards")
	cards.shuffle()
	// cards.print()
	hand, _ := deal(cards, 5)
	hand.print()
	// remainingCards.print()
}
