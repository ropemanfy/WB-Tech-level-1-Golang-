package main

import (
	"fmt"
)

func main() {
	var sl = []int{2, 4, 5, 6, 8, 9}
	up := make(chan int)
	down := make(chan int)

	go Writer(up, sl)   // запускаем
	go Square(up, down) // конвейер

	for i := range down { // выводим результаты
		fmt.Println(i)
	}
}

func Writer(in chan<- int, sl []int) { // пишет значения из слайса в канал
	defer close(in)
	for _, v := range sl {
		in <- v
	}
}

func Square(in <-chan int, out chan<- int) { // принимает значения из одного канала, возводит их в квадрат и пишет в другой канал
	defer close(out)
	for i := range in {
		out <- i * i
	}
}
