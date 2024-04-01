package main

import "fmt"

func main() {
	sl := []int{1, 5, 7, 3, 8, 3, 8}
	sl = RemoveElement(sl, 1)
	fmt.Println(sl)
}

func RemoveElement(sl []int, index int) []int {
	var result []int
	for i, v := range sl { // итерируюясь по слайсу
		if i != index { // проверяем равен ли индекс данному
			result = append(result, v) // если нет, то пишем его в результирующий слайс
		}
	}
	return result
}
