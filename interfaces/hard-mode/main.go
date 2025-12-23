package main

import (
	"fmt"
	"os"
)

/*
You CAN use any type (int, string, slice, map).
But you SHOULD use a struct.
*/
type logWriter struct{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter: please provide a file path")
		return
	}

	// os.Args is an slice of strings, which stores the fileName at first index and that's why we are using os.Args[1] to get the actual file content
	filepath := os.Args[1]
	fmt.Printf("Reading file: %s\n", filepath)
	lw := logWriter{}

	// file is an pointer
	file, err := os.Open(filepath)
	// Error handling
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	buf := make([]byte, 1000)
	n, err := file.Read(buf)
	lw.Write(buf[:n])
	// func (f *File) Read(b []byte) (n int, err error) and Read function a pointer receiver, Only a *File can call .Read()
	// so we can call Read with 'file' variable created above because it;s *File
	/* buf := make([]byte, 1000)
	n, err := file.Read(buf)
	fmt.Println(string(buf[:n])) // converting byte slice into an string


	defer file.Close() */

	//io.Copy(lw, file)

}

func (logWriter) Write(p []byte) (n int, err error) {

	fmt.Println(string(p))
	return len(p), nil
}
