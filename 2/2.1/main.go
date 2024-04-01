package main

import (
	"fmt"
	"sync"
)

func main() {
	array := SyncArray{arr: []int{2, 4, 6, 8, 10}}

	array.Add(len(array.arr)) // устанавливаем количество wg
	Squares(&array)
	array.Wait() // ждём исполнения Squares

}

type SyncArray struct {
	arr []int
	sync.WaitGroup
}

func Squares(array *SyncArray) {
	for _, v := range array.arr { // перебираем данный массив
		go func(v int) { // пускаем горутины
			defer array.Done()                              // декрементим wg после выполнения горутины
			result := v * v                                 // считаем квадрат
			fmt.Printf("Квадрат %v равен: %v\n", v, result) // принтим ответ
		}(v)
	}
}
