package main

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// new_deck := hand.toString()
	// fmt.Println(new_deck)
	// hand.print()
	// remainingCards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
