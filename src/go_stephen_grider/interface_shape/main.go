package main

import "fmt"

func main() {
	s1 := Square{side: 3.0}
	t1 := Triangle{base: 1.0, height: 2.0}

	total := float64(0)
	total += printArea(s1)
	total += printArea(t1)

	fmt.Println(total)
}
