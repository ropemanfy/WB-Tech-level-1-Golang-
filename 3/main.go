package main

import (
	"fmt"
	"sync"
)

func main() {
	array := SyncArray{arr: []int{2, 4, 6, 8, 10}}

	ch := make(chan int, len(array.arr)) // создаем канал длиною равной массиву
	array.Add(len(array.arr))            // устанавливаем количество wg

	for _, v := range array.arr { // пускаем горутины, в которые передаем элемент слайса
		go array.Square(ch, v)
	}

	go func() { // ждём выполнения горутин и закрываем канал
		array.Wait()
		close(ch)
	}()

	fmt.Println(Sum(ch)) // считаем и принтим результат
}

type SyncArray struct {
	arr []int
	sync.WaitGroup
}

func (array *SyncArray) Square(ch chan int, v int) { // считаем квадрат
	ch <- v * v
	array.Done()
}

func Sum(ch chan int) (result int) { // суммируем приходящие из горутин квадраты
	for i := range ch {
		result += i
	}
	return
}
