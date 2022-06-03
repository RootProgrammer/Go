/*
	The io package specifies the io.Reader interface, which represents the read end of a stream of data.

	The Go standard library contains many implementations of this interface,
	including files, network connections, compressors, ciphers, and others.

	The io.Reader interface has a Read method:

	func (T) Read(b []byte) (n int, err error)
	Read populates the given byte slice with data and returns the number of bytes populated
	and an error value. It returns an io.EOF error when the stream ends.

	The example code creates a strings.Reader and consumes its output 8 bytes at a time.
*/

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// for exercise, go to https://tour.golang.org/methods/22 , https://tour.golang.org/methods/23
