package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// остановка через сигнал

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt) // устанавливаем сигнал остановки (Ctrl+C)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch: // в случае получения сигнала завершаем работу
				fmt.Println("shutdown")
				return
			default: // выполняем какую то работу
				fmt.Println("im working")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	wg.Wait() // ждём остановки горутины
}
