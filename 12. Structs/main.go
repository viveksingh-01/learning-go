package main

import "fmt"

// Structs in Go are similar to Objects in JS
// Defining Struct for "employee" type ~ [Data-structure which represents employee data]
type employee struct {
	id     uint16
	name   string
	salary float64
}

func main() {
	// Creating a new "employee" (employee type struct)

	// FIRST Approach of creating a new struct
	// This approach is not recommended since it will store incorrect data if the order in struct definition changes
	alex := employee{21001, "Alex Anderson", 40000}
	fmt.Println(alex)

	// SECOND Approach of creating struct
	jim := employee{id: 21002, name: "Jim Morrison", salary: 42000}
	fmt.Println(jim)

	// SECOND Approach (Multiline)
	mark := employee{
		id:     21003,
		name:   "Mark Thompson",
		salary: 45000, // NOTE: Don't forget the trailing comma here!
	}
	fmt.Println(mark)

	// THIRD Approach
	var oliver employee
	/*
	* The below will print {id:0 name: salary:0} ~ Similar to an empty Object in JS
	* And that is because we haven't assigned any value to the struct so it's initialized with the ZERO VALUES by default.
	* ZERO VALUES => string -> "", int -> 0, float -> 0, bool -> false
	*/
	fmt.Printf("%+v \n", oliver)
	// Updating the properties/fields on employee struct
	oliver.id = 21004
	oliver.name = "Oliver Clarkson"
	oliver.salary = 36000
	fmt.Println(oliver)
}
