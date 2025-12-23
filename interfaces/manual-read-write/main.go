package main

import (
	"fmt"
	"io"
)

type logWriter struct{}

type myReader struct {
	data     string
	position int
}

func main() {
	lw := logWriter{}

	r := myReader{data: "Hello Anuroop, this is custom reader", position: 9}
	rr := &r // rr is an pointer which stores memory address of r and rr is pointer of type myReader

	myCopy(lw, rr)

	io.Copy(lw, rr)
}

/*
	✔ myCopy DOES:

it takes a Reader (any type that has Read)
it takes a Writer (any type that has Write)
it moves bytes from Reader → Writer
*/
func myCopy(dst io.Writer, src io.Reader) error {

	buf := make([]byte, 60) // buf = [0,0,0,0,0,0,0, ... 60 times]

	for {
		// Read data into 'buf' which is an slice of bytes
		n, err := src.Read(buf) // Ex:= buf = ['H','e','l','l','o', 0,0,0,0, ...]  and n=5

		if n > 0 { // Ex :=  n=5
			_, writeErr := dst.Write(buf[:n]) // buf[:n]  →  buf[0:5]  →  ['H','e','l','l','o']

			// Error Handling
			if writeErr != nil {
				return writeErr
			}
		}

		// Stop reading, the file/stream is over. i,e n=0
		if err == io.EOF {
			// end of file → stop the loop
			break
		}

		// actual error → stop and return error
		if err != nil {
			return err
		}

	}
	return nil

}

func (logWriter) Write(p []byte) (int, error) {

	fmt.Println("Custom Writer:", string(p)) // Type conversion from byte slice into string
	return len(p), nil
}

/*
**Since the Read method has a pointer receiver (*myReader),

calling Read on a pointer means:
→ the method works on the ORIGINAL struct i,e myReader struct,
→ not on a copy,
→ and any changes are applied directly to the real object.**
*/
func (r *myReader) Read(p []byte) (int, error) {
	/* 	// “If my reading cursor reaches the end of the string, tell the caller that there is no more data.”
		Assume:=
		data = "Hello Anuroop, this is custom reader"
		len(data) = 36
		position = 0


		FIRST READ:=
		position >= len(data)? 0 >= 36 → NO
	    n = copy(p, data[0:]) → n=36
	    position = position + n → position = 36
	    return (36, nil)

	*/
	if r.position >= len(r.data) {
		return 0, io.EOF
	}

	//
	n := copy(p, r.data[r.position:])
	r.position += n
	return n, nil
}
