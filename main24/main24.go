package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p *Point) SetX(x float64) {
	p.x = x
}

func (p *Point) SetY(y float64) {
	p.y = y
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func (firstPoint *Point) GetDistance(secondPoint *Point) float64 {
	x := math.Pow((secondPoint.x - firstPoint.x), 2)
	y := math.Pow((secondPoint.y - firstPoint.y), 2)
	return math.Sqrt(x + y)
}

func main() {
	var (
		x float64
		y float64
	)

	fmt.Println("1st point")
	fmt.Printf("Enter x: ")
	fmt.Scan(&x)
	fmt.Printf("Enter y: ")
	fmt.Scan(&y)
	first := NewPoint(x, y)

	fmt.Println("2nd point")
	fmt.Printf("Enter x: ")
	fmt.Scan(&x)
	fmt.Printf("Enter y: ")
	fmt.Scan(&y)

	second := NewPoint(x, y)

	fmt.Println(first.GetDistance(&second))
}
