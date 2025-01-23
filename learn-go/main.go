package main

import "fmt"

func newCard() string {
	return "King of hearts"
}

func main() {
	cards := deck{"Ace", "Jack", "Ten", newCard()}
	cards = append(cards, "Queen")
	cards.print()

	fmt.Println(newDeck())
}
