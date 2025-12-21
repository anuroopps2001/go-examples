package main

import "fmt"

type shape interface {
	getArea() float64
}
type square struct { // has function of with def: = getArea() float64, so it becomes type shape
	sideLengrh float64
}

type triangle struct { // has function of with def: = getArea() float64, so it comes type shape
	height float64
	base   float64
}

func main() {
	sq_val1 := square{sideLengrh: 3.2}
	tri_var1 := triangle{height: 2.4, base: 3.7}
	printArea(sq_val1)  // because square has func getArea() float64, it can be passed into printArea() which accepts the values of type shape
	printArea(tri_var1) // // because triangle has func getArea() float64, it can be passed into printArea() which accepts the values of type shape

}
func (sq square) getArea() float64 {
	return sq.sideLengrh * sq.sideLengrh // for sq_var1 struct variable
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println("Area: ", s.getArea())
}
