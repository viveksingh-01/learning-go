package main

import "fmt"

// We can declare functions in Go using the "func" keyword
// We need to declare the data-type for the parameters (if any) as well as for the return-type if the function returns something.
func printCards(cards []string) {
	for _, card := range cards {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Jack of Diamonds"
}

func main() {
	cards := []string{"Ace of Spades", "Five of Diamonds", "Two of Hearts", "Nine of Clubs", newCard()}
	// Calling/invoking printCards function with the cards as the argument
	printCards(cards)
}