package main

import (
	"fmt"
	"math"
)

func main() {
	point1 := NewPoint(13.6, 15.1)
	point2 := NewPoint(-50.9, 81.4)
	fmt.Println(Distance(point1, point2))
}

type Point struct { // точка в двумерном пространстве
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(p1 *Point, p2 *Point) float64 {
	x := p1.x - p2.x               // x катет
	y := p1.y - p2.y               // y катет
	result := math.Sqrt(x*x + y*y) // по теореме Пифагора находим результат
	return result
}
