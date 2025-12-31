package main

import "fmt"

type printer interface {
	printDetails()
}
type book struct {
	title  string
	author string
}

type pen struct {
	brand string
	color string
}

func main() {
	bookVar1 := book{title: "The Go Programming Language", author: "Alan A. Donovan"}
	penVar1 := pen{brand: "Reynolds", color: "(Blue)"}
	showDetails(bookVar1)
	showDetails(penVar1)

}

func (b book) printDetails() {
	fmt.Println("Book:", b.title, "by", b.author)
}

func (p pen) printDetails() {
	fmt.Println("Pen:", p.brand, "of", p.color)

}

// This specific function expects the values of type 'printer' which is an Interface
func showDetails(p printer) {
	p.printDetails()
}
