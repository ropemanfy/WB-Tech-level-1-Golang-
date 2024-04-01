package main

import (
	"fmt"
)

func main() {
	array := []int{2, 4, 6, 8, 10}

	ch := make(chan struct{}) // WaitGroup можно заменить на канал

	Squares(array, ch)

	for range array { // ждём пока не отработают все горутины
		<-ch
	}
}

func Squares(array []int, ch chan struct{}) {
	for _, v := range array { // перебираем данный массив
		go func(v int) { // пускаем горутины
			result := v * v                                 // считаем квадрат
			fmt.Printf("Квадрат %v равен: %v\n", v, result) // принтим ответ
			ch <- struct{}{}                                // отправляем сингал, что работа выполнена
		}(v)
	}
}
