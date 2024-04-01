package main

import (
	"fmt"
	"strings"
)

type Set struct {
	s map[rune]struct{}
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aAbcd"

	fmt.Println(Unique(str1))
	fmt.Println(Unique(str2))
	fmt.Println(Unique(str3))
}

func Unique(str string) bool {
	set := NewSet()            // создаем множество
	str = strings.ToLower(str) // приводим все символы к нижнему регистру
	for _, char := range str { // проходим по символам строки
		_, ok := set.s[char] // проверяя на наличие их в множестве
		if ok {
			return false // если находим совпадения, то возвращаем false
		}
		set.Add(char) // если нет, то добавляем в множество
	}
	return true
}

func NewSet() *Set { // новое множество
	return &Set{s: make(map[rune]struct{})}
}

func (s *Set) Add(v rune) { // добавить элемент в множество
	s.s[v] = struct{}{}
}
