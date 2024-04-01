package main

import "fmt"

type Set struct {
	s map[string]struct{}
}

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}

	set := NewSet()

	for _, v := range sl { // добавляем элементы данного слайса в множество
		set.Add(v)
	}

	for k := range set.s { // выводим результат
		fmt.Println(k)
	}
}

func NewSet() *Set { // новое множество
	return &Set{s: make(map[string]struct{})}
}

func (s *Set) Add(v string) { // добавить элемент в множество
	s.s[v] = struct{}{}
}
