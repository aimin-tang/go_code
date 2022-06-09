package main

import "fmt"

func sum(nums []float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	nums := []float64{1, 2, 3}
	fmt.Println(sum(nums))
}
