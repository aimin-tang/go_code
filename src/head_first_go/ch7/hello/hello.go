package main

import (
	"ch7/datafile"
	"fmt"
)

func main() {
	names, err := datafile.GetStrings("names.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(names)
}
