package main

import (
	"fmt"
	"sync"
	"time"
)

// остановка через закрытие канала

func main() {
	ch := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch: // завершаем работу, если канал закрылся
				fmt.Println("shutdown")
				return
			default: // выполняем какую то работу
				fmt.Println("im working")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(7 * time.Second) // ждём 7 секунд
	close(ch)                   // закрываем канал
	wg.Wait()                   // ждём остановки горутины
}
