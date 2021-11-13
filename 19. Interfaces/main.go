package main

import (
	"fmt"
	"math"
)

// INTERFACES:
// Interfaces in Go work a little different than most of the other languages.
// In Go, Interfaces are not implemented explicitly.

// The idea behind Interfaces in Go is duck-typing. 
// Which simply translates into: If something looks like a duck and quacks like a duck then it must be a duck. 
// Meaning that if our object implements all duck’s features then there should be no problem using it as a duck.

// Declaring the Interface Shape
type Shape interface {
	area() float64
}

// By declaring the above Interface "Shape", we're telling Go that our program has a new type called 'Shape'
// and our 'Shape' Interface defines the following behaviour/method => "area() float64"
// and if any type in the entire program implements exactly the same behavior as defined by the Interface
// meaning, if any type has an associated function called "area()" and it returns a "float64" value
// then it will be considered an honorary member of "Shape".

type Rectangle struct {
	length  float64
	breadth float64
}

type Triangle struct {
	base   float64
	height float64
}

type Circle struct {
	radius float64
}

// type Rectangle's implementation of "area() float64"
func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

// type Triangle's implementation of "area() float64"
func (t Triangle) area() float64 {
	return 0.5 * t.base * t.height
}

// type Circle's implementation of "area() float64"
func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func printArea(s Shape) {
	fmt.Println(s.area())
}

func main() {
	r := Rectangle{length: 25, breadth: 10}
	t := Triangle{base: 10, height: 20}
	c := Circle{radius: 8}

	// Since all the types namely, Rectange, Triangle and Circle have the function "area() float64" associated with them
	// They can call (use) the "printArea()" function.
	printArea(r)
	printArea(t)
	printArea(c)
}
