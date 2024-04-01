package main

import (
	"fmt"
	"time"
)

// работает чуть точнее

func main() {
	start := time.Now()
	Sleep(5 * time.Second)
	fmt.Println(time.Since(start))
}

func Sleep(t time.Duration) {
	start := time.Now()          // засекаем время
	end := start.Add(t)          // добавляем к нему заданное
	for time.Now().Before(end) { // выходим из цикла по его достижении
	}
}
