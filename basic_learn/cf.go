// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	// "abc/popcount"
	"abc/tempconv"
	"flag"
	"fmt"
	// "os"
	// "strconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
	// fmt.Print(popcount.PopCount(2133))
	// for _, arg := range os.Args[1:] {
	// 	t, err := strconv.ParseFloat(arg, 64)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	f := tempconv.Fahrenheit(t)
	// 	c := tempconv.Celsius(t)
	// 	fmt.Printf("%s = %s, %s = %s\n",
	// 		f, tempconv.FToC(f), c, tempconv.CToF(c))
	// }
}
