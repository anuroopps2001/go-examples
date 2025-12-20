package main

import "fmt"

// struct named circle
type circle struct {
	radius float64 // radius and variable and it accept the values of type float64
}

// struct named rectangle
type rectangle struct {
	width  float64 // width is an variable and accept the values of type float64
	height float64 // heigh is an variable again in this stuct and accept the values of type float64
}

// Interface called 'shape' which manage the methods having 'area() float64' method definition
type shape interface {
	area() float64 // This says:- 'Any type that has a method area() float64 is automatically a shape.'
}

func main() {
	/*
	   A type does NOT “become” an interface type.
	   It is only represented AS an interface value when passed to something expecting that interface.
	   Here, c will be of type 'circle' and re will be of type 'reactangle'
	   However, when they are passed into an function, which is expecting an interface type, then they become variables of interface type.
	*/

	c := circle{radius: 3.14}            // assigning values into struct
	re := rectangle{width: 4, height: 5} // assigning values into struct

	/*
			When you pass c or re into a function that expects values of type 'shape',
			Go performs implicit interface conversion:
			circle → shape
		    rectangle → shape
	*/
	printArea(c)
	printArea(re)
}

// This method is managed by Interface named 'shape'
func (c circle) area() float64 { // as a receiver of custom type 'circle' and variables calling this method should hold same type data.
	return 3.14 * c.radius * c.radius
}

// This method is managed by Interface named 'shape'
func (re rectangle) area() float64 { // needs to be called by variables holding the values of custom type 'rectangle'
	return re.height * re.width
}

func printArea(s shape) {
	fmt.Println("Area: ", s.area())
}
