package main

import (
	"fmt"

	"rsc.io/quote"
)

/*
- go mod init tutorial/get_started (go.mod)
- write hello.go
- go run .
- add rsc.io/quote in code
- go mod tidy (go.mod + go.sum)
*/
func main() {
	fmt.Println(quote.Go())
}
