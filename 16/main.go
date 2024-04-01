package main

import "fmt"

func main() {
	var sl = []int{3, 5, 7, 2, 8, 3, 0, 1, 13, 64, 25}
	sl = QuickSort(sl)
	fmt.Println(sl)
}

func QuickSort(sl []int) []int {
	if len(sl) <= 1 { // выход из рекурсии
		return sl
	}

	anchor := sl[0]       // опорный элемент
	var left, right []int // левый слайс для элементов меньше опорного, правый для элементов больше опорного

	for _, v := range sl[1:] { // проходим по слайсу, минуя опорный элемент
		if v <= anchor {
			left = append(left, v) // заполняем левый слайс
		} else {
			right = append(right, v) // заполняем правый слайс
		}
	}
	result := append(QuickSort(left), anchor)    // элементы меньше опроного помещаем слева
	result = append(result, QuickSort(right)...) // больше - справа
	return result
}
