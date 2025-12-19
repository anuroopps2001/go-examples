package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// http.Get takes url and returns response (resp) and error (err)
	// It returns a struct of type *http.Response. and that struct has filed Body io.ReadCloser
	resp, err := http.Get("http://google.com")

	// Error handlig
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1) // If there is a error, we need to perminate the program immediatley.
	}

	bs := make([]byte, 99999) // Initializing an slice of bytes with 99999 empty storage to store data in next steps

	// reading bytes from an HTTP response body into a byte slice.  i.e Take the response body, and read some bytes into the bs buffer.
	resp.Body.Read(bs)      // write data into byte slice i.e []byte
	fmt.Println(string(bs)) //read slice of bytes in the strings format by making type conversion syntax:- <what_we_want>(what_we_have)

}
