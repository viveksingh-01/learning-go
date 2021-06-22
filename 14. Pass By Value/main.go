package main

import (
	"fmt"
)

type employee struct {
	id     uint16
	name   string
	salary float64
}

// Receiver function to update salary of employee
func (e employee) updateSalary(newSalary float64) {
	e.salary = newSalary
}

func main() {
	mark := employee{
		id:     21003,
		name:   "Mark Thompson",
		salary: 45000,
	}
	fmt.Println(mark) //  {21003 Mark Thompson 45000}
	mark.updateSalary(50000)
	// After updating salary
	fmt.Println(mark) //  {21003 Mark Thompson 45000} -> the salary remains the same
}

/*
* The receiver function "updateSalary" doesn't actually work as expected.
* Meaning, it doesn't update the salary of the "employee" (on which it's called upon, "mark" in this case) to a new value.

* REASON:
* It's because Go is considered to be a "Pass By Value" language, which means that when we pass a parameter to a function, Go
* creates a copy of that variable to use it inside the function.

* In the above function "updateSalary", the variable "e" is the copy of "mark" and both of them have different memory addresses.
* mark -> 0xc000006030 | e -> 0xc000006040

* So when we update the salary: e.salary = newSalary
* It's actually updating the salary of the variable "e" and not "mark" 
* In other words, since they both don't share the same memory address, the updated field value doesn't reflect in "mark".

* In Go, we can broadly categorize the data-types into two categories:
* 	(1) Value Types (Pass By Value) => int, float, string, bool, structs
* 	(2) Reference Types (Pass By Reference) => slices, maps, channels, pointers, functions

* Due to the above categorization, "slices" worked perfectly alright for us but "structs" did not.
* So we need to handle the "Value Types" differently to make them work.
* And that's when the "POINTERS" come to our rescue!
*/
