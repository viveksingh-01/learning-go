package main

import (
	"fmt"
)

type employee struct {
	id     uint16
	name   string
	salary float64
}

/*
* POINTER -> Stores the memory address of a variable.
* For example:
* 	var fruit string = "apple"
* 	fruitPointer := &fruit // "fruitPointer" will store the address of the "fruit" variable.
*
* Here, fruitPointer ~ &fruit => Memory address of "fruit"
* 		  *fruitPointer ~ fruit => Value stored at that address, i.e. "apple"
 */

// "*employee" in the function parameter just represents the type description which means that we're working with a pointer to a "employee"
func (pointerToEmployee *employee) updateSalary(newSalary float64) {
	// "*pointerToEmployee" represents an actual operator. It provides the value stored in the variable the pointer is referencing
	(*pointerToEmployee).salary = newSalary // Here, we are updating the actual "mark" struct using its pointer.
}

func main() {
	mark := employee{
		id:     21003,
		name:   "Mark Thompson",
		salary: 45000,
	}
	fmt.Println(mark) //  {21003 Mark Thompson 45000}
	// Creating a POINTER which points to the memory address of "mark"
	var markPointer *employee = &mark
	// Invoking the function using the pointer
	markPointer.updateSalary(50000)
	// After updating salary
	fmt.Println(mark) //  {21003 Mark Thompson 50000} -> the salary gets updated successfully!!
}
