package main

import (
	"fmt"
	"math"
)

func greatest(nums ...float64) float64 {
	greatest := math.Inf(-1)
	for _, num := range nums {
		if num > greatest {
			greatest = num
		}
	}
	return greatest
}

func main() {
	s := []float64{1.1, 2.2, 3.3, 2.2}
	fmt.Println(greatest(s...))
}
