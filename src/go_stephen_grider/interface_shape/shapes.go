package main

import "fmt"

// shape interface
type Shape interface {
	getArea() float64
}

func printArea(s Shape) float64 {
	fmt.Println(s.getArea())
	return s.getArea()
}

// square
type Square struct {
	side float64
}

func (s Square) getArea() float64 {
	return s.side * s.side
}

// triangle
type Triangle struct {
	base, height float64
}

func (t Triangle) getArea() float64 {
	return t.base * t.height / 2
}
