package main

import "fmt"

func main() {
	// Two ways of declaring variables
	// 1. Declaring the data-type explicitly
	var name string = "Joey Tribbiani"
	// 2. Go infers the data-type implicitly
	age := 25
	// Printf - Can be used when we need to print a formatted text with arguments.
	fmt.Printf("Hi! My name is %v and I'm %v years old.", name, age)
	fmt.Printf("\nMy age is %v and it's type is %T", age, age)
}
