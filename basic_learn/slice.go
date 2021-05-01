package main

import "fmt"

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
func appendInt2(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	// ...expand z to at least zlen...
	copy(z[len(x):], y)
	return z
}
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
func main() {
	sss := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(sss, 2)) // "[5 6 8 9]"
	fmt.Println(sss)            // "[5 6 8 8 9]"
	data := []string{"one", "", "three"}
	data = nonempty(data)
	fmt.Printf("%q\n", data) //
	var ints []int
	ints = appendInt(ints, 1)
	ints = appendInt(ints, 2)
	ints = appendInt2(ints, 2, 3)
	fmt.Println(ints)
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	fmt.Println(x)

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"	fmt.Println(runes)

	strs := []string{"Sad", "gg", "am", "ca", "Sad", "gg"}
	fmt.Println(equal(strs[0:2], strs[4:6]))
	s := [...]int{1, 2, 3, 5, 6, 7, 8}
	reverse(s[:2])
	fmt.Println(s)
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May"}
	Q2 := months[1:3]
	summer := months[2:4]
	fmt.Println(Q2)     //
	fmt.Println(summer) //
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
	fmt.Println(summer[:4])
}
