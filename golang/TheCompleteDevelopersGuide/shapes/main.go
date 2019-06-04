package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct{
	sideLength float64
}

type triangle struct{
	base float64
	height float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println("Shape Area:", s.getArea())
}

func main() {
	s := square{sideLength: 10}
	t := triangle{base: 3.0, height: 5.0}
	
	printArea(s)
	printArea(t)
}