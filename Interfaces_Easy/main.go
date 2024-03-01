package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{
		height: 10.5,
		base:   10,
	}
	s := square{
		sideLength: 4,
	}
	printArea(s)
	printArea(t)

}
