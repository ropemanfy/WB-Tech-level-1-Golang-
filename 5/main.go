package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int
	fmt.Println("Enter the function run time(sec)")
	fmt.Scan(&num)                                    // сканируем ввод с клавиатуры
	ch := make(chan int)                              // создаем канал для передачи данных
	t := time.After(time.Duration(num) * time.Second) // пишем в канал, когда истечет установленное время
	go Reader(ch)                                     // читаем данные
	Writer(ch, t)                                     // генерируем данные
}

func Reader(ch chan int) { // принтит данные из канала, пока тот не закроется
	for range ch {
		fmt.Println(<-ch)
	}
}

func Writer(ch chan int, t <-chan time.Time) { // генерит инты пока не истечет установленное время
	for {
		select {
		case <-t: // ждём
			close(ch) // закрываем канал
			fmt.Println("Timeout")
			return
		default:
			ch <- rand.Int() // генерируем рандомные инты
		}
	}
}
