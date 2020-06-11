package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func basicReaderTest() {
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

// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
// import "golang.org/x/tour/reader"
// MyReader struct
type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}
func readerExercise() {
	// reader.Validate(MyReader{})
}

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, e := rot.r.Read(b)

	for i := 0; i < n; i++ {
		if (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n') {
			b[i] += 13
		} else if (b[i] > 'M' && b[i] <= 'Z') || (b[i] > 'm' && b[i] <= 'z') {
			b[i] -= 13
		}
	}
	return n, e
}

func rot13ReaderTest() {
	// Build string reader
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// Build rot13 reader
	r := rot13Reader{s}
	// Call reader.read
	io.Copy(os.Stdout, &r)
}

func main() {
	basicReaderTest()
	rot13ReaderTest()
}
