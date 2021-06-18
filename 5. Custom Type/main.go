package main

import "fmt"

// Using type "deck" to declare the type of parameter
func printCards(cards deck) {
	for _, card := range cards {
		fmt.Println(card)
	}
}

func main() {
	// Using type "deck" to declare the cards
	cards := deck{"Ace of Spades", "Five of Diamonds", "Two of Hearts", "Nine of Clubs"}
	printCards(cards)
}
