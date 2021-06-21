package main

import (
	"fmt"
)

type address struct {
	street  string
	city    string
	zipCode uint32
}

type contactInfo struct {
	email   string
	phone   string
	address // If the field-name is same as the struct's name, we can just mention the field-name and it's type will be implicit
}

type employee struct {
	id      uint16
	name    string
	salary  float64
	contact contactInfo  // Embedded struct
}

func main() {
	mark := employee{
		id:     21003,
		name:   "Mark Thompson",
		salary: 45000,
		contact: contactInfo{
			email: "markt@gmail.com",
			phone: "+1 234 56789",
			address: address{
				street:  "ABC Street",
				city:    "Los Angeles",
				zipCode: 994525,
			},
		},
	}
	fmt.Println(mark)
}
