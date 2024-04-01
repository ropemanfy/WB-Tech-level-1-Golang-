package main

import "fmt"

func main() { // не сработает при 0
	a := 1
	b := 2
	fmt.Println(a, b)

	a = a * b
	b = a / b
	a = a / b
	fmt.Println(a, b)
}
