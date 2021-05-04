package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Test() error {

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read failed: %v", err)
		}
		fmt.Printf("%q", r)
	}
	return nil
}

func main() {
	if err := Test(); err != nil {
		fmt.Println(err)
	}
}
