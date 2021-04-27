package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func init() {
	var err error
	cwd, err = os.Getwd() // compile error: unused: cwd
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	fmt.Println(cwd)
}

func test(fname string) error {
	f, err := os.Open(fname)
	if err != nil { // compile error: unused: f
		return err
	} else {
		// f.Read()
		f.Close()
	}
	return nil
}

func main() {
	fmt.Println(cwd)

	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
		fmt.Printf("%c", x)
	}
	fmt.Println(x)
}
