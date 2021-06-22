package main

import (
	"fmt"
)

type employee struct {
	id     uint16
	name   string
	salary float64
}

// Since we have already mentioned the parameter to be of the "employee" Pointer type, Go takes care of it even
// if we directly call this function without a pointer
func (pointerToEmployee *employee) updateSalary(newSalary float64) {
	(*pointerToEmployee).salary = newSalary
}

func main() {
	mark := employee{
		id:     21003,
		name:   "Mark Thompson",
		salary: 45000,
	}
	// Instead of creating a Pointer here and calling the "updateSalary" function using it, we can use a Pointer shortcut that GO provides us
	var markPointer *employee = &mark
	markPointer.updateSalary(50000)
	
	// We can use our "employee" type struct directly to call the function.
	mark.updateSalary(60000)
	fmt.Println(mark)
}
