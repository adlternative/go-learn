package main

import "fmt"
import "sort"

func main() {

	ages := make(map[string]int) // mapping from strings to ints
	fmt.Println(ages == nil)     // "false"
	fmt.Println(len(ages) == 0)  // "true"
	var ages2 map[string]int
	fmt.Println(ages2 == nil)    // "true"
	fmt.Println(len(ages2) == 0) // "true"
	ages3 := map[string]int{}
	fmt.Println(ages3 == nil)    // "false"
	fmt.Println(len(ages3) == 0) // "true"

	_, ok := ages2["bob"]
	if !ok {
		println("!")
	}
	// append(ages2, "sad":1)

	delete(ages, "sad")
	ages["bob"]++
	ages["gg"] = 213
	ages["gr"] = 3
	fmt.Println(ages["bob"])
	// for name, age := range ages {
	// 	fmt.Printf("%s\t%d\n", name, age)
	// }
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
