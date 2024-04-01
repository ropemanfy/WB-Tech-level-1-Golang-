package main

import "fmt"

func main() {
	var a interface{}

	a = 3
	PrintType(a)

	a = "bimbimbambam"
	PrintType(a)

	a = false
	PrintType(a)

	a = make(chan int)
	PrintType(a)
}

func PrintType(v interface{}) {
	fmt.Printf("%T\n", v)
}
