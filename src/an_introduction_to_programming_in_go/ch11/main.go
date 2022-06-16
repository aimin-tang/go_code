package main

import (
	"an_introduction_to_programming_in_go/ch11/math"
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
