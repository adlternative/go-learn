package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
func add1(r rune) rune { return r + 1 }
func main() {
	fmt.Println(strings.Map(add1, "hhhheheakldhjakl"))
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000"))
}
