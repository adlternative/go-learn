package main

import (
	"fmt"
	// "unicode"
	"unicode/utf8"
)

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // "255 0 1"

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	ascii := 'a'
	unicode := "国家真大"
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)        // "97 a 'a'"
	fmt.Printf("%[1]s %[1]q\n", unicode)         // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)            // "10 '\n'"
	fmt.Println(len(unicode))                    // "10 '\n'"
	fmt.Println(utf8.RuneCountInString(unicode)) // UTF8 字符串长度

	for i := 0; i < len(unicode); {
		/* 解格式成UTF8 */
		r, size := utf8.DecodeRuneInString(unicode[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	n := 0
	// for _, _ = range unicode {
	// 	n++
	// }
	// fmt.Println(n)
	// n := 0
	for range unicode {
		n++
	}
	fmt.Println(n)
	s := "プログラム"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))       // "プログラム"
	fmt.Println(string(65))      // "A", not "65"
	fmt.Println(string(0x4eac))  // "京"
	fmt.Println(string(1234567)) // "?"
	// fmt.Println(basename("a/b/c.go")) // "c"
	// fmt.Println(basename("c.d.go"))   // "c.d"
	// fmt.Println(basename("basic_learn"))      // "basic_learn"
	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)
}
