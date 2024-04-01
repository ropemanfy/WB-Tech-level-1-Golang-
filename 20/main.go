package main

import "fmt"

func main() {
	str := "sn◯w d◯g sun"
	fmt.Println(ReverseWords(str))
}

func ReverseWords(str string) string {
	var result string
	var count int

	m := make(map[int]string) // создаём мапу, в которую будем помещать слова, а ключом будет его позиция в строке
	runes := []rune(str)      // конвертируем данную строку в слайс рун

	for i := range runes { // проходясь по слайсу
		char := string(runes[i])
		if char == " " { // ищем пробелы
			count++ // если пробел найден, то инкрементим счётчик и переходим к другому слову
		} else {
			m[count] += char // если нет, то добавляем символ к строке под соответсвующим ключом
		}
	}

	for i := 0; i < len(m); i++ { // далее идём по циклу, где len(m) - это количество слов
		index := len(m) - 1 - i
		result += m[index] // складываем слова в обратном порядке в результирующую строку
		if i != len(m)-1 { // и если слово не последнее, то ставим пробел
			result += " "
		}
	}
	return result
}
