package main

import "fmt"

func make_odd_gen() func() int {
	i := 1
	odds := func() (next_odd int) {
		next_odd = i
		i += 2
		return
	}
	return odds
}

func main() {
	odds := make_odd_gen()
	fmt.Println(odds())
	fmt.Println(odds())
}
