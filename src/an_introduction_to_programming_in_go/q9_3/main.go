package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float32
}

type Rectangle struct {
	x1, x2, y1, y2 float32
}

type Shape interface {
	area() float32
	perimeter() float32
}

func (c Circle) area() float32 {
	return math.Pi * c.r * c.r
}

func (c Circle) perimeter() float32 {
	return float32(2) * math.Pi * c.r
}

func (r Rectangle) area() float32 {
	return (r.x2 - r.x1) * (r.y2 - r.y1)
}

func (r Rectangle) perimeter() float32 {
	return float32(2) * (r.x2 - r.x1) * (r.y2 - r.y1)
}

func get_area(shapes ...Shape) float32 {
	area := float32(0)
	for _, shape := range shapes {
		area += shape.area()
	}
	return area
}

func get_perimeter(shapes ...Shape) float32 {
	perimeter := float32(0)
	for _, shape := range shapes {
		perimeter += shape.perimeter()
	}
	return perimeter
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{1, 2, 1, 2}

	fmt.Println(get_area(c, r))
	fmt.Println(get_perimeter(c, r))
}
