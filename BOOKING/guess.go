package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to this CLI guessing game. Goodluck!")
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(11)

	var guess int
	fmt.Println("\nType in your guess between 0 and 10:")

	fmt.Scan(&guess)
	guessCount := 0
	for guess != randomNumber {
		guessCount = guessCount + 1
		if guess < randomNumber {
			fmt.Println("Guess is lower than the number. Try again!")
			fmt.Scan(&guess)
		} else {
			fmt.Println("Guess is higher than the number. Try again!")
			fmt.Scan(&guess)
		}

	}

	fmt.Printf("Congrats, you got it right with %v guesses!", guessCount)

}
