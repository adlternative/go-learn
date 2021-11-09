package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	strs := []string{}
	for input.Scan() {
		s := input.Text()
		strs = append(strs, strings.Replace(s, "\n", " ", -1))
	}
	lens := len(strs)
	for i := 0; i < lens; i++ {
		fmt.Fprint(os.Stdout, strs[i])
		fmt.Fprint(os.Stderr, strs[i])
	}
}
