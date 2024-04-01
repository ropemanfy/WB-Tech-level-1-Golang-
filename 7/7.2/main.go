package main

import (
	"fmt"
	"sync"
)

// всё то же, но уже с готовой реализацией из пакета sync

func main() {
	var m sync.Map

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			data := fmt.Sprintf("data %v", i)
			m.Store(i, data)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 5; i++ {
		l, _ := m.Load(i)
		fmt.Println(l)
	}
}
