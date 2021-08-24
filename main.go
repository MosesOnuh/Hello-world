package main
import (
	"fmt"
	"math"
)


func main() {
	fmt.Println("Hello world")
	work()
	anything(  "Moses" , 34)

	t, h  := fruit("orange", 500)
	fmt.Println(t,h)
}
func anything (name string, age int)( string, int){
	fmt.Println("name is", name)
	fmt.Println("age is", age)

	return name, age
}
func fruit (name string, number int)(string, int){
   number = number + 3
	return name, number
}


func work() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(144))
}
