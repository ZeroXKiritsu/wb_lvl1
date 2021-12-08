package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func (p *Point) New(x, y int) {
	p.x = x
	p.y = y
}

func getDistance(p1, p2 *Point) float64 {
	x := math.Abs(float64(p1.x - p2.x))
	y := math.Abs(float64(p1.y - p2.y))
	distance := math.Sqrt(x*x + y*y)
	return distance
}

func main() {
	p1 := new(Point)
	p1.New(1, 2)
	p2 := new(Point)
	p2.New(3, 4)
	result := getDistance(p1, p2)
	fmt.Printf("%f",result)
}
