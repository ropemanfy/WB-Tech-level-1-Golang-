package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	increment int
	sync.Mutex
}

func main() {
	c := NewCounter(0)

	var wg sync.WaitGroup
	wg.Add(150)

	for i := 0; i < 150; i++ { // пускаем горутины
		go func() {
			defer wg.Done()
			c.Increase()
		}()
	}

	wg.Wait() // ждём выполнения горутин
	fmt.Println(c.increment)
}

func NewCounter(i int) *Counter { // новый инстанс
	return &Counter{increment: i}
}

func (c *Counter) Increase() { // увеличение счетчика
	c.Lock()         // лочим доступ
	defer c.Unlock() // анлочим после выхода
	c.increment++    // инкрементим
}
