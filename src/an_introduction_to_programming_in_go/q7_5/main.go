package main

import "fmt"

func fib(n uint) uint {
	if n < 0 {
		panic("n must not be negative!")
	} else if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	fmt.Println(fib(10))
}
