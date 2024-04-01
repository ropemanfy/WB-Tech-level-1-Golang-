package main

import (
	"fmt"
	"sync"
)

type Mapper interface {
	Set(k int, v string)
	Get(k int) string
}

type syncMap struct {
	m map[int]string
	sync.Mutex
}

func main() {
	m := NewMap()

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ { // конкурентно пишем
		go func(i int) {
			defer wg.Done()
			data := fmt.Sprintf("data %v", i)
			m.Set(i, data)
		}(i)
	}

	wg.Wait() // ждём когда отработают горутины

	for i := 0; i < 5; i++ { // проверяем данные
		fmt.Println(m.Get(i))
	}
}

func NewMap() Mapper { // создает инстанс и возвращает интерфейс для работы с ним
	return &syncMap{m: make(map[int]string)}
}

func (s *syncMap) Set(k int, v string) { // устанавливает значение под ключом
	s.Lock()
	defer s.Unlock()
	s.m[k] = v
}

func (s *syncMap) Get(k int) string { // возвращает значение под ключом
	return s.m[k]
}
