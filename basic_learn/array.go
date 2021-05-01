package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
)

func zero(ptr *[32]byte) {
	*ptr = [32]byte{0}
	// for i := range ptr {
	// 	ptr[i] = 0
	// }
}
func main() {
	array := [32]byte{}
	for i, _ := range array {
		array[i] = byte(rand.Int())
	}
	zero(&array)
	for _, vv := range array {
		println(vv)
	}

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// d := [3]int{1, 2}
	// fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
}
