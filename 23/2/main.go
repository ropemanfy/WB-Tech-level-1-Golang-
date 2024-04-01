package main

import "fmt"

func main() {
	sl := []int{1, 5, 7, 3, 8, 3, 8}
	sl = RemoveElement(sl, 6)
	fmt.Println(sl)
}

func RemoveElement(sl []int, i int) []int { // делим слайс на две части
	sl1 := sl[:i]            // с элементами до данного индекса
	sl2 := sl[i+1:]          // и после него
	sl = append(sl1, sl2...) // объединяем получившиеся слайсы
	return sl
}
