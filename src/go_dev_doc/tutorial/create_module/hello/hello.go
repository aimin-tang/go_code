package main

import (
    "fmt"

    "tutorial/greetings"
)

/*
- cd tutorial
- mkdir greetings; cd greetings
- go mod init tutorial/greetings
- code greetings
- cd ..; mkdir hello; cd hello
- go mod init tutorial/hello
- code hello.go
- go mod edit -replace tutorial/greetings=../greetings
- go mod tidy
*/

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
