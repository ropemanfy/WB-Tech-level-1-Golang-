package main

import "fmt"

func main() {
	action := Action{Human: Human{Name: "Ryan", Lastname: "Gosling"}}

	action.Dance()
	action.Sing()
}

type Human struct { // родительская структура
	Name     string
	Lastname string
}

type Action struct { // дочерняя структура
	Human // встраиваем родителя
}

func (a *Action) Dance() {
	fmt.Printf("%v %v Dancing\n", a.Name, a.Lastname)
}

func (a *Action) Sing() {
	fmt.Printf("%v %v Singing\n", a.Name, a.Lastname)
}
