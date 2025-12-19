package main

import "fmt"

type circle struct {
	radius float64
}
type rectangle struct {
	width  float64
	height float64
}

// Define interface
type shape interface {
	area() float64
}

func main() {
	c := circle{radius: 3.14}
	re := rectangle{width: 4, height: 5}

	printArea(c)
	printArea(re)
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (re rectangle) area() float64 {
	return re.height * re.width
}

func printArea(s shape) {
	fmt.Println("Area: ", s)
}
