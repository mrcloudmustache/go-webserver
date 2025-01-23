package main

import "fmt"

// Create a new type of Deck with a slice a strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
