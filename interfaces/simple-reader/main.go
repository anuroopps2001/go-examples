package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing file path:")
		os.Exit(1)
	}
	fmt.Println("Currently reading file:", os.Args[1])
	filepath := os.Args[1]

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	bs := make([]byte, 1000)
	// func (f *File) Read(b []byte) (n int, err error) i.e Read needs to be called with vars of type *File and file variable is of same type
	// Read reads up to len(b) bytes from the File and stores them in b
	n, err := file.Read(bs)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(bs[:n]))
}
