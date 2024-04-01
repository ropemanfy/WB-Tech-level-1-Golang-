package main

import (
	"fmt"
)

//	проблема 1: глобальная justString ссылается на часть строки v
//	после такой операции GC не сможет очистить память v
//	вследствие чего возникает утечка памяти

//	проблема 2: если в v[:100] есть символы unicode
//	то justString получится не такой, какой нам бы хотелось
//	тк итерация происходит по байтам, а не по рунам

var justString string

func someFunc(i int) { // убиваем сразу 2х зайцев
	var count int
	v := createHugeString(1 << 10)
	for _, char := range v { // итерируемся по рунам
		justString += string(char) // конвертируем в string и пишем их в результирующую строку
		count++
		if count == i { // выходим из функции при достижении нужного количества символов
			return
		}
	}
}

func main() {
	someFunc(100)
	fmt.Println(justString)
}

func createHugeString(i int) (s string) {
	for j := 0; j < i; j++ {
		s += "ж"
	}
	return
}
