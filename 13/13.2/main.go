package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Println(a, b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}
