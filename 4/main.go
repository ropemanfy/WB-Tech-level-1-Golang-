package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

func main() {
	var num int
	fmt.Println("Enter number of workers")
	fmt.Scan(&num) // сканируем ввод с клавиатуры

	shutdown := make(chan os.Signal, 1)   // создаем канал для сигнала остановки
	signal.Notify(shutdown, os.Interrupt) // указываем значение сигнала

	ch := make(chan int) // создаем канал для передачи данных от генератора к воркерам

	for i := 0; i < num; i++ { // пускаем воркеры
		go Worker(ch)
	}

	Generator(ch, shutdown) // пускаем генератор данных
}

func Worker(ch chan int) { // принтит данные из канала, пока тот не закроется
	for range ch {
		fmt.Println(<-ch)
	}
}

func Generator(ch chan int, shutdown chan os.Signal) { // генерит инты пока не получит сигнал остановки
	for {
		select {
		case <-shutdown: // ждём сигнал остановки
			close(ch) // закрываем канал
			fmt.Println("Shutdown")
			return
		default:
			ch <- rand.Int() // генерируем рандомные инты
		}
	}
}
