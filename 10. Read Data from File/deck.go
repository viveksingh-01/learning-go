package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Read File from disk
func newDeckFromFile(filename string) deck {
	// ReadFile returns a byte-slice or an error (if any)
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// Convert byte-slice to slice of strings
	s := strings.Split(string(bs), ", ")
	return deck(s)
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
