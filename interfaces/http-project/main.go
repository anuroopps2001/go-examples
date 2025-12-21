package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// http.Get takes url and returns response (resp) and error (err)
	// It returns a struct of type *http.Response. and that struct has filed Body io.ReadCloser
	// resp.Body is a stream — something you have to read from.
	resp, err := http.Get("http://google.com")

	// Error handlig
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1) // If there is a error, we need to perminate the program immediatley.
	}

	// Streams give you bytes, not text.
	// So you must prepare a “bucket” (byte slice) to hold the data:

	// Create a bucket that can hold 99,999 bytes of data.
	bs := make([]byte, 99999) // Initializing an slice of bytes with 99999 empty storage to store data in next steps

	// Take the data from the HTTP response body and fill this bucket with bytes.
	// reading bytes from an HTTP response body into a byte slice.  i.e Take the response body, and read some bytes into the bs buffer.
	// Read(p []byte) (n int, err error)  :- Read function takes byte slice as input and returns n which is an integer and err which is error
	n, err := resp.Body.Read(bs) // write data into byte slice i.e []byte. At this moment, bs now contains HTML (but as raw bytes).
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Now, bs = [ real HTML (8500 bytes) | 91500 zeros ] when n=8500,

	// bs[:n] means: Give me a new slice from index 0 up to index n (not including n).
	fmt.Println(string(bs[:n])) //read slice of bytes in the strings format by making type conversion syntax:- <what_we_want>(what_we_have)

	// Defer function, which runs after the execution of all the function and if there is no error, program exits smoothly without memory leak
	defer func() {
		erro := resp.Body.Close()
		if erro != nil {
			fmt.Println("Error fron closer() function: ", erro)
		}
	}()

	// Copy everything from resp.Body and Write it to Stdout.
	// io.Copy(os.Stdout, resp.Body)
}
