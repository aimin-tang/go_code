package main

import (
	"fmt"
	"log"

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
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
