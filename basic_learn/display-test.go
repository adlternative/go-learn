package main

import (
	. "basic_learn/display"
	"basic_learn/sexpr"
	"fmt"
	"os"
	"reflect"
	// "os"
)

type student struct {
	no   int
	name string
}
type Movie struct {
	Title, Subtitle string
	Year            int
	// Color           bool
	Actor  map[string]string
	Oscars []string
	Sequel *string
}

func main() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		// Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	// s := student{1, "hehe"}
	// Display("stu", s)
	// Display("stdout", os.Stdout)
	// var i interface{} = 3
	// Display("i", i)
	// Display("&i", &i)
	// Display("stringelove", strangelove)
	message, err := sexpr.Marshal(strangelove)
	if err != nil {
		fmt.Printf("Marshal%v", err)
	}
	fmt.Print(string(message))
	// type Cycle struct {
	// 	Value int
	// 	Tail  *Cycle
	// }
	// var c Cycle
	// c = Cycle{42, &c}
	// Display("c", c)
	Display("2", reflect.ValueOf(2))
	x := 2
	Display("x", reflect.ValueOf(x))
	c := reflect.ValueOf(&x)
	Display("&x", reflect.ValueOf(c))
	d := c.Elem()
	Display("&x.Elem", reflect.ValueOf(d))
	fmt.Println(reflect.ValueOf(2).CanAddr()) // "false"
	fmt.Println(reflect.ValueOf(x).CanAddr()) // "false"
	fmt.Println(c.CanAddr())                  // "false"
	fmt.Println(d.CanAddr())                  // "true"
	// x := 2
	e := reflect.ValueOf(&x).Elem()   // e refers to the variable x
	px := e.Addr().Interface().(*int) // px := &x
	*px = 3                           // x = 3
	fmt.Println(x)
	e.Set(reflect.ValueOf(4))
	// e.Set(reflect.ValueOf(int64(5))) // panic: int64 is not assignable to int
	fmt.Println(x)
	f := reflect.ValueOf(&x).Elem()
	f.SetInt(-1)
	fmt.Println(x) // "3"

	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	// ry.SetInt(2)                     // panic: SetInt called on interface Value
	ry.Set(reflect.ValueOf(3)) // OK, y = int(3)
	fmt.Println(y)
	// ry.SetString("hello")            // panic: SetString called on interface Value
	ry.Set(reflect.ValueOf("hello")) // OK, y = "hello"
	fmt.Println(y)

	stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
	fmt.Println(stdout.Type())                  // "os.File"
	// fd := stdout.FieldByName("fd")
	// fmt.Println(fd) // "1"
	Display("stdout", os.Stdout)
	name := stdout.FieldByName("file").Elem().FieldByName("name")
	fmt.Println(name)

}
