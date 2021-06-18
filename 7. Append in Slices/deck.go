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
			// The "append" function appends the new item/element at the end of the slice and returns the modified slice.
			// NOTE: Append function DOES NOT modify the original slice but creates a copy of it and then appends.

			// append(cards, value+" of "+suit)  -->  Doesn't append to the original slice
			// fmt.Println(cards)  --> Same as before.
			// So, it's important that we assign the modified slice (returned by the "append" function) to our slice.
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
