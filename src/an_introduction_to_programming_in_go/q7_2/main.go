package main

import "fmt"

func halve(x int) (int, bool) {
	if x%2 == 1 {
		return x / 2, false
	} else {
		return x / 2, true
	}
}

func main() {
	fmt.Println(halve(1))
	fmt.Println(halve(2))
}
