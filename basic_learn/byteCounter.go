package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type ByteCount int

func (c ByteCount) String() string { return fmt.Sprintf("[%d]", c) }

// func (c *ByteCount) String() string { return fmt.Sprintf("[%d]", c) }

func (c *ByteCount) Write(p []byte) (int, error) {
	*c += ByteCount(len(p))
	return len(p), nil
}

func main() {
	os.Stdout.Write([]byte("hello")) // OK: *os.File has Write method
	os.Stdout.Close()                // OK: *os.File has Close method

	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello")) // OK: io.Writer has Write method
	// w.Close()

	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	fmt.Println(any)

	var c ByteCount
	c.Write([]byte("helloc"))
	fmt.Println(&c)
	c = 3 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(&c) // "12", = len("hello, Dolly")
	fmt.Println(c)  // "12", = len("hello, Dolly")

}
