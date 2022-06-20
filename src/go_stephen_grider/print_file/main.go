package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong num of args. Got:", os.Args)
		os.Exit(1)
	}

	file_name := os.Args[1]
	fh, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error open file", file_name)
		os.Exit(1)
	}

	io.Copy(os.Stdout, fh)
}
