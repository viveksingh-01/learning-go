package main

import (
	"fmt"
	"math"
)

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

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

func (t Triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// NOTE:
// The above receiver functions associated with Rectangle,Triangle and Circle have very different logic implementations.
// But the functions below which are used to print the respective areas of Rectangle,Triangle and Circle contain
// pretty much the same logic only with different types.

// So, we need to find a way to generalize the below functions into one which will contain the
// same logic and will have a generic type.
// That's when our hero "The INTERFACES" comes into play.

// Function to print area of Rectangle
func printAreaOfRectangle(r Rectangle) {
	fmt.Println(r.area())
}

// Another similar function to print area of Triangle
func printAreaOfTriangle(t Triangle) {
	fmt.Println(t.area())
}

// Yet another similar function to print area of Circle
func printAreaOfCircle(c Circle) {
	fmt.Println(c.area())
}

func main() {
	r := Rectangle{length: 25, breadth: 10}
	t := Triangle{base: 10, height: 20}
	c := Circle{radius: 8}
	printAreaOfRectangle(r)
	printAreaOfTriangle(t)
	printAreaOfCircle(c)
}
