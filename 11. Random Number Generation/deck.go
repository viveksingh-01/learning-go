package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ", ")
	return deck(s)
}

// Shuffle the deck of cards
func (d deck) shuffle() {
	/*
	* Simply using newPosition := rand.Intn(len(d) - 1) doesn't work here
	* Since rand uses the same source (seed-value), it generates the same sequence of random values everytime
	* So to solve that problem, we need to create a new Rand with different source (seed-value)
	 */
	// Creating a new source with "time.Now().UnixNano()" as the seed value
	source := rand.NewSource(time.Now().UnixNano())
	// Creating a new Rand using the source created above
	r := rand.New(source)

	for i := range d {
		// Using the new Rand to generate a random value
		newPosition := r.Intn(len(d) - 1)
		// Swapping the value at index with that at the newPosition
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
