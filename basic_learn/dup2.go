// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		/* 如果没有输入行 */
		countLines(os.Stdin, counts)
	} else {
		/* 说明传入了参数 */
		for _, arg := range files {
			/* 打开文件 */
			f, err := os.Open(arg)
			/* 打不开文件，打印错误类型 */
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			/* 对文件的每一行计数 */
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

/* Open返回值也就是一个 *os.File */
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f) /* 传入一个文件 */
	for input.Scan() {           /* 对每一行计数 */
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
