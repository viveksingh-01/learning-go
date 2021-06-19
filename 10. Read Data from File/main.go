package main

func main() {
	cards := newDeckFromFile("My Cards.txt")
	cards.print()
}
