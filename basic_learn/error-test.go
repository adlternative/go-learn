package main

import (
	"errors"
	"fmt"
	"syscall"
)

func main() {
	fmt.Println(errors.New("EOF") == errors.New("EOF"))                 // "false"
	fmt.Println(errors.New("EOF").Error() == errors.New("EOF").Error()) // "false"
	var err error = syscall.Errno(2)
	fmt.Println(err.Error()) // "no such file or directory"
	fmt.Println(err)
}
