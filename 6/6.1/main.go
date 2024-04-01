package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// остановка через context

func main() {
	ctx, cancle := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done(): // в случае отмены контекста завершаем работу
				fmt.Println("shutdown")
				return
			default: // выполняем какую то работу
				fmt.Println("im working")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(7 * time.Second) // ждём 7 секунд
	cancle()                    // отменяем контекст
	wg.Wait()                   // ждём остановки горутины
}
