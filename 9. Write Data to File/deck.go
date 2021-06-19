package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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
	return d[:handSize], d[handSize:]
}

// Convert "deck" type (slice of string) to a string 
func (d deck) toString() string {
	return strings.Join(d, ", ")
}

// Save (Write) file to disk
func (d deck) saveToFile(filename string) error {
	// We need to convert the data (in this case "deck") to a "byte slice" to write it to the disk.
	// So we convert the "deck" type to a string and then the string to a Byte-Slice.
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) 
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
