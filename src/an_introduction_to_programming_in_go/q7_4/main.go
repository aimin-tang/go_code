package main

import "fmt"

func make_odd_gen() func() int {
	i := 1
	result_func := func() int {
		result := i
		i += 2
		return result
	}
	return result_func
}

func main() {
	f := make_odd_gen()
	fmt.Println(f())
	fmt.Println(f())
}
