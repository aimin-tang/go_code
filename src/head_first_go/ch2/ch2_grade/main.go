package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter your grade: ")

	// get grade
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed in ReadString")
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal("Failed in ParseFloat")
		log.Fatal(err)
	}

	// test grade
	var status bool
	if grade >= 60 {
		status = true
	} else {
		status = false
	}

	fmt.Println(status)
}
