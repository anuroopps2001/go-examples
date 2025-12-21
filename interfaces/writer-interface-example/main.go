package main

import (
	"fmt"
	"io"
	"net/http"
)

// custom type to associate the 'Write' function
type logWriter struct{}

func main() {
	resp, _ := http.Get("http://googl.com")

	lw := logWriter{} // defining variable to hold values of type struct
	/* bs := make([]byte, 1000)
	n, err := resp.Body.Read(bs)
	lw.Write(bs[:n]) // You will only get the first 1000 bytes. The rest of the data will be ignored. ❌ Not acceptable for real-world networking.
	lw.Write(resp.Body.Read(bs)) */ // Read function expects byte slice and return int and erro but Write function is expecting an byte slice but Write gets bytes only from your slice, not from Read’s return.
	io.Copy(lw, resp.Body)          // io.Copy will reuses buffer and handles all the bytes and does reading and writing and it;s never grows byte slice size
}

// "call Write function with an variable holding the values of type logWriter (which is struct in this case) and pass arguments of type byte slice"
func (logWriter) Write(p []byte) (n int, err error) { // Any value of type logWriter automatically satisfies the io.Writer interface. Because logWriter has a Write method with the correct signature
	/*
	   Go’s standard library already defines:
	   type Writer interface {
	       Write(p []byte) (n int, err error)
	   }

	   This lives in the io package.

	   So, we never define Writer interface again, as below
	   type Writer interface { ... }   // ❌ Do NOT do this
	*/
	fmt.Println(string(p))
	fmt.Println("Just wrote this many bytes:", len(p))
	return len(p), nil // It returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the write to stop early.
}
