package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func getTarget() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3)
}

func getResult(round int) int {
	fmt.Print("Round #")
	fmt.Print(round)
	fmt.Print(" - Enter your guess: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed in ReadString")
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Failed in ParseFload")
		log.Fatal(err)
	}

	return num
}

func main() {
	target := getTarget()
	var i int
	for i = 1; i < 10; i++ {
		result := getResult(i)
		if result == target {
			fmt.Println("Good Guess!")
			break
		} else if result < target {
			fmt.Println("Too low!")
		} else {
			fmt.Println("Too high!")
		}
	}

	if i == 10 {
		fmt.Println("Failed: too many guesses!")
	}
}
