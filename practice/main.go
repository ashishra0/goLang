package main

import "fmt"

type shape interface {
	getArea() (string, float64)
}

type triangle struct {
	height float64
	base   float64
	object string
}

type square struct {
	sideLength float64
	object     string
}

func main() {
	t := triangle{2.0, 4.0, "triangle"}
	s := square{5.5, "square"}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() (string, float64) {
	return "the area of" + " " + t.object + " is", 0.5 * t.height * t.base
}

func (s square) getArea() (string, float64) {
	return "the area of" + " " + s.object + " is", s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
