package main

import "fmt"

func main() {
	sl := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	m := make(map[int][]float64) // группировать будем в мапу

	for _, v := range sl { // идём по данному слайсу
		a := (int(v) / 10) * 10 // и ищем ближайшее значение кратное 10
		m[a] = append(m[a], v)  // пишем его в ключ, а под ним же значение float
	}

	fmt.Println(m)
}
