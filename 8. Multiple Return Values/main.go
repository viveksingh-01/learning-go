package main

import "fmt"

func main() {
	cards := newDeck()
	// Assigning multiple returned values
	hand, remainingCards := deal(cards, 9)
	fmt.Println("Dealt Cards:")
	hand.print()
	fmt.Println("\nRemaining Cards:")
	remainingCards.print()
}
