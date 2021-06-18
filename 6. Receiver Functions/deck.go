package main

import "fmt"

// Create a custom-type "deck" which is a slice of string
type deck []string

// Receiver function (Any variable of type "deck" will now have access to the "print()" method and can call it on itself.)
// Here, the variable "d" (passed to the function) is the actual copy of the deck (cards array) we're working with.
// It works similar to the "this" keyword in JS.
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
