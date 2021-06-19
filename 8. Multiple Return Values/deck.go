package main

import "fmt"

type deck []string

// Create and return a deck of cards
func newDeck() deck {
	var cards deck
	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "One", "Two", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	// SLICE RANGE: Syntax -> slice[startIndexInclusive : endIndexNotInclusive]
	return d[:handSize], d[handSize:]  // Returning multiple values
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
