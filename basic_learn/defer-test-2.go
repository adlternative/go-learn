package main
import "fmt"

func main(){
i:=1
defer func(i int){fmt.Print(i)}(i)
defer func(){fmt.Print(i)}()
i=2
}
