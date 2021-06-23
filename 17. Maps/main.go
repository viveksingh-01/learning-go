package main

import "fmt"

func main() {
	// Declaring a Map
	var colors map[string]string
	fmt.Println(colors) // map[]
	// Maps in Go are similar to dictionaries in Python
	colors = map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colors) // map[blue:#0000ff green:#00ff00 red:#ff0000]
	// NOTE: Map items can only be manipulated or accessed using square-brackets and NOT the dot operator.
	colors["black"] = "#000000"
	colors["white"] = "#ffffff"
	fmt.Println(colors) // map[black:#000000 blue:#0000ff green:#00ff00 red:#ff0000 white:#ffffff]

	// Map with key as integers
	numInWords := map[int]string{
		0: "Zero",
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
	}
	fmt.Println(numInWords[3]) // "Three"
	fmt.Println(numInWords)    // map[0:Zero 1:One 2:Two 3:Three 4:Four 5:Five]
	delete(numInWords, 4)      // We can delete a specific item using "delete" function
	fmt.Println(numInWords)    // map[0:Zero 1:One 2:Two 3:Three 5:Five]

	// We can also create Maps using "make" function
	menu := make(map[string]float64)
	menu = map[string]float64{
		"sandwich": 3.99,
		"cake":     4.99,
		"juice":    1.50,
	}
	// Iterating over map items
	for key, value := range menu {
		fmt.Printf("%-10v $%0.2f \n", key, value)
	}
}
