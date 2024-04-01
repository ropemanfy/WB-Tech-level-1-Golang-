package main

import "fmt"

func main() {
	str := "глав℟ыба"
	fmt.Println(ReverseStr(str))
}

func ReverseStr(str string) (result string) {
	s := []rune(str)   // конвертируем данную строку в слайс рун
	for i := range s { // перебираем его
		index := len(s) - 1 - i    // элементы по порядку, начиная с последнего
		result += string(s[index]) // конвертируем обратно в строку и добавляем к результирующей строке
	}
	return
}
