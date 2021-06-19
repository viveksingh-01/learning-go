package main

import "log"

func main() {
	cards := newDeck()
	// Since the WriteFile method returns error in case it occurs, we log it.  
	log.Fatal(cards.saveToFile("My Cards.txt"))
}
