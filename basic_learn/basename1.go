package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' { /* 找到最后一个/ */
			s = s[i+1:] /* 只要文件部分 */
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func basename2(s string) string {
	if slash := strings.LastIndex(s, "/"); slash >= 0 {
		s = s[slash+1:]
		if dot := strings.LastIndex(s, "."); dot >= 0 {
			s = s[:dot]
		}
	}
	return s
}
func main() {
	s := "basic_learn"
	b := []byte(s)
	s2 := string(b)
	println(s2)
	paths := "/home/adl/gitTest4/a.txt"
	fmt.Println(basename2(paths))
	fmt.Println(comma("1234532"))
}
