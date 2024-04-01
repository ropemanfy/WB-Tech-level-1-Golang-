package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	Sleep(5 * time.Second)
	fmt.Println(time.Since(start))
}

func Sleep(t time.Duration) {
	<-time.After(t) // пишем в канал по истечении времени
}
