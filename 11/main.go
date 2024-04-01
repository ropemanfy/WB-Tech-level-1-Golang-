package main

import "fmt"

type Set struct {
	s map[string]struct{}
}

func main() {
	set1 := NewSet()
	set2 := NewSet()

	values1 := []string{"three", "hundred", "bucks", "master"} // создаем слайсы строк
	values2 := []string{"hundred", "bucks", "slave"}

	for _, v := range values1 { // заполняем ими множества
		set1.Add(v)
	}
	for _, v := range values2 {
		set2.Add(v)
	}

	set3 := set1.Intersection(set2) // создаем пересечение

	for k := range set3.s { // выводим результат
		fmt.Println(k)
	}
}

func NewSet() *Set { // новое множество
	return &Set{s: make(map[string]struct{})}
}

func (s *Set) Add(v string) { // добавить элемент в множество
	s.s[v] = struct{}{}
}

func (s1 *Set) Intersection(s2 *Set) *Set { // создать пересечение множеств
	intersection := NewSet()
	for k := range s1.s { // итерируемся по одному из множеств
		_, ok := s2.s[k] // делаем проверку на совпадение ключей
		if ok {          // если совпадение есть
			intersection.s[k] = struct{}{} // добавляем значение в новое множество
		}
	}
	return intersection
}
