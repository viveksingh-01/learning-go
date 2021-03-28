package main

import "fmt"

func main() {
	// Declaring a Slice (A slice in Go is an array with infinite length)
	cards := []string{"Ace of Spades", "Five of Diamonds", "Two of Hearts", "Nine of Clubs"}

	fmt.Println("--- With index ---")
	for index, card := range cards {
		fmt.Println(index, card)
	}

	fmt.Println("\n--- Without index ---")
	// If we want to skip the index part, we can do that by replacing it with an "_" (since Go doesn't allow to declare any variable without use)
	for _, card := range cards {
		fmt.Println(card)
	}
}
