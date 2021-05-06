package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	r := []rune("H和")
	buf.WriteRune(r[1])
	return buf.String()
}

func main() {
	angel := "Heros never die"
	angleBytes := []byte(angel)
	for i := 5; i <= 10; i++ {
		angleBytes[i] = '#'
	}
	fmt.Println(string(angleBytes))
	fmt.Println(angel)
	fmt.Println(strings.Contains("basic_learn", "bc"))
	fmt.Println(strings.Count("abbccbbcc aa bbcc", "bbcc"))
	strs := strings.Fields("basic_learn def eee")
	for _, str := range strs {
		println(str)
	}
	fmt.Println(strings.HasPrefix("basic_learn def eee", "basic_learn def"))
	fmt.Println(strings.Index("basic_learn def eee", "def"))
	println(strings.Join(strs, "#"))
	println(bytes.Contains([]byte{'a', 'c', 'd'}, []byte{'c', 'd'}))
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
	x := 123
	y := fmt.Sprintf("?%d", x)
	println(y)
	x1, err := strconv.Atoi("123") // x is an int
	if err != nil {
		fmt.Println(err)
	}
	println(x1)

	y1, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	if err != nil {
		fmt.Println(err)
	}
	println(y1)

}
