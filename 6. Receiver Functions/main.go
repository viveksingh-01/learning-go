package main

func main() {
	cards := deck{"Ace of Spades", "Five of Diamonds", "Two of Hearts", "Nine of Clubs"}

	// Invoking the Receiver Function (works similar to how we invoke methods of objects in JS)
	// Since the variable cards is of "deck", it can call the "print()" method on itself.
	cards.print()
}
